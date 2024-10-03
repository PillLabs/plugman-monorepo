import { expect } from "chai";
import { Contract } from "ethers";
import { ethers, upgrades } from "hardhat";
import { UniswapV2Deployer } from "uniswap-v2-deploy-plugin";

import {
  ERC20,
  ERC20CustodyNew,
  GatewayEVM,
  Receiver,
  TestERC20,
  UniswapV2Factory,
  UniswapV2Pair,
  UniswapV2Router02,
} from "../../typechain-types";

describe("Uniswap Integration with GatewayEVM", function () {
  let tokenA: TestERC20;
  let tokenB: TestERC20;
  let factory: UniswapV2Factory;
  let router: UniswapV2Router02;
  let pair: UniswapV2Pair;
  let custody: ERC20CustodyNew;
  let gateway: GatewayEVM;
  let owner, addr1, addr2;
  let tssAddress;

  beforeEach(async function () {
    [owner, addr1, addr2, tssAddress] = await ethers.getSigners();

    // Deploy TestERC20 tokens
    const TestERC20 = await ethers.getContractFactory("TestERC20");
    tokenA = (await TestERC20.deploy("Token A", "TKA")) as TestERC20;
    tokenB = (await TestERC20.deploy("Token B", "TKB")) as TestERC20;
    await tokenA.mint(owner.address, ethers.utils.parseEther("1000"));
    await tokenB.mint(owner.address, ethers.utils.parseEther("1000"));

    const { factory: newFactory, router: newRouter, weth9: weth } = await UniswapV2Deployer.deploy(owner);

    factory = newFactory;
    router = newRouter;

    // Approve Router to move tokens
    await tokenA.approve(router.address, ethers.utils.parseEther("1000"));
    await tokenB.approve(router.address, ethers.utils.parseEther("1000"));

    // Add Liquidity
    await router.addLiquidity(
      tokenA.address,
      tokenB.address,
      ethers.utils.parseEther("500"),
      ethers.utils.parseEther("500"),
      0,
      0,
      owner.address,
      Math.floor(Date.now() / 1000) + 60 * 20
    );

    // Deploy contracts
    const Gateway = await ethers.getContractFactory("GatewayEVM");
    const ERC20CustodyNew = await ethers.getContractFactory("ERC20CustodyNew");
    const ZetaConnector = await ethers.getContractFactory("ZetaConnectorNonNative");
    const zeta = await TestERC20.deploy("Zeta", "ZETA");
    gateway = (await upgrades.deployProxy(Gateway, [tssAddress.address, zeta.address], {
      initializer: "initialize",
      kind: "uups",
    })) as GatewayEVM;
    custody = (await ERC20CustodyNew.deploy(gateway.address, tssAddress.address)) as ERC20CustodyNew;
    gateway.connect(tssAddress).setCustody(custody.address);
    const zetaConnector = await ZetaConnector.deploy(gateway.address, zeta.address, tssAddress.address);
    gateway.connect(tssAddress).setConnector(zetaConnector.address);

    // Transfer some tokens to the custody contract
    await tokenA.transfer(custody.address, ethers.utils.parseEther("100"));
    await tokenB.transfer(custody.address, ethers.utils.parseEther("100"));
  });

  it("should perform a token swap on Uniswap and transfer the output tokens to a destination address", async function () {
    const amountIn = ethers.utils.parseEther("50");

    const data = router.interface.encodeFunctionData("swapExactTokensForTokens", [
      amountIn,
      0,
      [tokenA.address, tokenB.address],
      addr2.address,
      Math.floor(Date.now() / 1000) + 60 * 20,
    ]);

    // Withdraw and call
    await custody.connect(tssAddress).withdrawAndCall(tokenA.address, router.address, amountIn, data);

    // Verify the destination address received the tokens
    const destBalance = await tokenB.balanceOf(addr2.address);
    expect(destBalance).to.be.gt(0);

    // Verify the remaining tokens are refunded to the Custody contract
    const remainingBalance = await tokenA.balanceOf(custody.address);
    expect(remainingBalance).to.equal(ethers.utils.parseEther("50"));

    // Verify the approval was reset
    const allowance = await tokenA.allowance(gateway.address, router.address);
    expect(allowance).to.equal(0);
  });
});
