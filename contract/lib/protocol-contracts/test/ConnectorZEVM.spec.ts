import { SignerWithAddress } from "@nomiclabs/hardhat-ethers/signers";
import { WETH9, ZetaConnectorZEVM, ZetaReceiverMock } from "@typechain-types";
import { expect } from "chai";
import { parseEther } from "ethers/lib/utils";
import { ethers } from "hardhat";

import { FUNGIBLE_MODULE_ADDRESS } from "./test.helpers";

const hre = require("hardhat");

describe("ConnectorZEVM tests", () => {
  let zetaTokenContract: WETH9;
  let zetaConnectorZEVM: ZetaConnectorZEVM;
  let zetaReceiverMockContract: ZetaReceiverMock;

  let owner: SignerWithAddress;
  let fungibleModuleSigner: SignerWithAddress;
  let addrs: SignerWithAddress[];
  let randomSigner: SignerWithAddress;

  beforeEach(async () => {
    [owner, randomSigner, ...addrs] = await ethers.getSigners();

    // Impersonate the fungible module account
    await hre.network.provider.request({
      method: "hardhat_impersonateAccount",
      params: [FUNGIBLE_MODULE_ADDRESS],
    });

    // Get a signer for the fungible module account
    fungibleModuleSigner = await ethers.getSigner(FUNGIBLE_MODULE_ADDRESS);
    hre.network.provider.send("hardhat_setBalance", [FUNGIBLE_MODULE_ADDRESS, parseEther("1000000").toHexString()]);

    const WZETAFactory = await ethers.getContractFactory("contracts/zevm/WZETA.sol:WETH9");
    zetaTokenContract = (await WZETAFactory.deploy()) as WETH9;

    const ZetaConnectorZEVMFactory = await ethers.getContractFactory("ZetaConnectorZEVM");
    zetaConnectorZEVM = (await ZetaConnectorZEVMFactory.connect(owner).deploy(
      zetaTokenContract.address
    )) as ZetaConnectorZEVM;

    const ZetaReceiverMockContractactory = await ethers.getContractFactory("ZetaReceiverMock");
    zetaReceiverMockContract = (await ZetaReceiverMockContractactory.connect(owner).deploy()) as ZetaReceiverMock;
  });

  describe("ZetaConnectorZEVM", () => {
    it("Should revert if the zetaTxSender has no enough zeta", async () => {
      await zetaTokenContract.connect(randomSigner).approve(zetaConnectorZEVM.address, 100_000);

      const tx = zetaConnectorZEVM.connect(randomSigner).send({
        destinationAddress: randomSigner.address,
        destinationChainId: 1,
        destinationGasLimit: 2500000,
        message: new ethers.utils.AbiCoder().encode(["string"], ["hello"]),
        zetaParams: new ethers.utils.AbiCoder().encode(["string"], ["hello"]),
        zetaValueAndGas: 1000,
      });

      // @dev: As we use the standard WETH contract, there's no error message for not enough balance
      await expect(tx).to.be.reverted;
    });

    it("Should revert if the zetaTxSender didn't allow ZetaConnector to spend Zeta token", async () => {
      await zetaTokenContract.deposit({ value: 100_000 });
      await zetaTokenContract.transfer(randomSigner.address, 100_000);

      const balance = await zetaTokenContract.balanceOf(randomSigner.address);
      expect(balance.toString()).to.equal("100000");

      const tx = zetaConnectorZEVM.send({
        destinationAddress: randomSigner.address,
        destinationChainId: 1,
        destinationGasLimit: 2500000,
        message: new ethers.utils.AbiCoder().encode(["string"], ["hello"]),
        zetaParams: new ethers.utils.AbiCoder().encode(["string"], ["hello"]),
        zetaValueAndGas: 1000,
      });

      // @dev: As we use the standard WETH contract, there's no error message for not enough balance
      await expect(tx).to.be.reverted;

      await zetaTokenContract.connect(randomSigner).transfer(owner.address, balance);
    });

    it("Should emit `ZetaSent` on success", async () => {
      const tx = await zetaConnectorZEVM.send({
        destinationAddress: randomSigner.address,
        destinationChainId: 1,
        destinationGasLimit: 2500000,
        message: new ethers.utils.AbiCoder().encode(["string"], ["hello"]),
        zetaParams: new ethers.utils.AbiCoder().encode(["string"], ["hello"]),
        zetaValueAndGas: 0,
      });

      expect(tx)
        .to.emit(zetaConnectorZEVM, "ZetaSent")
        .withArgs(owner.address, owner.address, 1, randomSigner.address, 0, 2500000, "hello", "hello");
    });

    it("Should transfer value and gas to fungible address", async () => {
      const zetaValueAndGas = 1000;
      await zetaTokenContract.approve(zetaConnectorZEVM.address, zetaValueAndGas);
      await zetaTokenContract.deposit({ value: zetaValueAndGas });

      const balanceBefore = await ethers.provider.getBalance(fungibleModuleSigner.address);

      await zetaConnectorZEVM.send({
        destinationAddress: randomSigner.address,
        destinationChainId: 1,
        destinationGasLimit: 2500000,
        message: new ethers.utils.AbiCoder().encode(["string"], ["hello"]),
        zetaParams: new ethers.utils.AbiCoder().encode(["string"], ["hello"]),
        zetaValueAndGas,
      });

      const balanceAfter = await ethers.provider.getBalance(fungibleModuleSigner.address);
      expect(balanceAfter.sub(balanceBefore)).to.equal(zetaValueAndGas);
    });

    it("Should update wzeta address if is call from fungible address", async () => {
      const WZETAFactory = await ethers.getContractFactory("contracts/zevm/WZETA.sol:WETH9");
      const newZetaTokenContract = (await WZETAFactory.deploy()) as WETH9;

      const tx = zetaConnectorZEVM.connect(fungibleModuleSigner).setWzetaAddress(newZetaTokenContract.address);
      await expect(tx).to.emit(zetaConnectorZEVM, "SetWZETA").withArgs(newZetaTokenContract.address);

      expect(await zetaConnectorZEVM.wzeta()).to.equal(newZetaTokenContract.address);
    });

    it("Should revert if try to update wzeta address from other address", async () => {
      const WZETAFactory = await ethers.getContractFactory("contracts/zevm/WZETA.sol:WETH9");
      const newZetaTokenContract = (await WZETAFactory.deploy()) as WETH9;

      const tx = zetaConnectorZEVM.setWzetaAddress(newZetaTokenContract.address);
      await expect(tx).to.be.revertedWith("OnlyFungibleModule");
    });

    it("Should transfer to the receiver address", async () => {
      // read eth balance of zetaConnectorEVM
      const initialZetaBalanceConnector = await ethers.provider.getBalance(zetaConnectorZEVM.address);
      const initialBalanceConnector = await zetaTokenContract.balanceOf(zetaConnectorZEVM.address);
      const initialBalanceReceiver = await zetaTokenContract.balanceOf(zetaReceiverMockContract.address);
      expect(initialZetaBalanceConnector.toString()).to.equal("0");
      expect(initialBalanceConnector.toString()).to.equal("0");
      expect(initialBalanceReceiver.toString()).to.equal("0");

      const tx = await zetaConnectorZEVM
        .connect(fungibleModuleSigner)
        .onReceive(
          randomSigner.address,
          1,
          zetaReceiverMockContract.address,
          1000,
          new ethers.utils.AbiCoder().encode(["string"], ["hello"]),
          ethers.constants.HashZero,
          { value: 1000 }
        );

      await expect(tx)
        .to.emit(zetaReceiverMockContract, "MockOnZetaMessage")
        .withArgs(zetaReceiverMockContract.address);

      const finalZetaBalanceConnector = await ethers.provider.getBalance(zetaConnectorZEVM.address);
      const finalBalanceConnector = await zetaTokenContract.balanceOf(zetaConnectorZEVM.address);
      const finalBalanceReceiver = await zetaTokenContract.balanceOf(zetaReceiverMockContract.address);

      expect(finalZetaBalanceConnector.toString()).to.equal("0");
      expect(finalBalanceConnector.toString()).to.equal("0");
      expect(finalBalanceReceiver.toString()).to.equal("1000");
    });

    it("Should call onRevert to the original address", async () => {
      // read eth balance of zetaConnectorEVM
      const initialZetaBalanceConnector = await ethers.provider.getBalance(zetaConnectorZEVM.address);
      const initialBalanceConnector = await zetaTokenContract.balanceOf(zetaConnectorZEVM.address);
      const initialBalanceReceiver = await zetaTokenContract.balanceOf(zetaReceiverMockContract.address);
      expect(initialZetaBalanceConnector.toString()).to.equal("0");
      expect(initialBalanceConnector.toString()).to.equal("0");
      expect(initialBalanceReceiver.toString()).to.equal("0");

      const tx = await zetaConnectorZEVM
        .connect(fungibleModuleSigner)
        .onRevert(
          zetaReceiverMockContract.address,
          1,
          randomSigner.address,
          5,
          1000,
          new ethers.utils.AbiCoder().encode(["string"], ["hello"]),
          ethers.constants.HashZero,
          { value: 1000 }
        );

      await expect(tx).to.emit(zetaReceiverMockContract, "MockOnZetaRevert").withArgs(zetaReceiverMockContract.address);

      const finalZetaBalanceConnector = await ethers.provider.getBalance(zetaConnectorZEVM.address);
      const finalBalanceConnector = await zetaTokenContract.balanceOf(zetaConnectorZEVM.address);
      const finalBalanceReceiver = await zetaTokenContract.balanceOf(zetaReceiverMockContract.address);

      expect(finalZetaBalanceConnector.toString()).to.equal("0");
      expect(finalBalanceConnector.toString()).to.equal("0");
      expect(finalBalanceReceiver.toString()).to.equal("1000");
    });

    it("Should accept ZETA from fungible module", async () => {
      const initialZetaBalanceConnector = await ethers.provider.getBalance(zetaConnectorZEVM.address);
      expect(initialZetaBalanceConnector.toString()).to.equal("0");

      await fungibleModuleSigner.sendTransaction({
        to: zetaConnectorZEVM.address,
        value: 1000,
      });

      // read eth balance of zetaConnectorEVM
      const finalZetaBalanceConnector = await ethers.provider.getBalance(zetaConnectorZEVM.address);
      expect(finalZetaBalanceConnector.toString()).to.equal("1000");
    });

    it("Should reject ZETA from other address", async () => {
      await fungibleModuleSigner.sendTransaction({
        to: randomSigner.address,
        value: 1000,
      });

      const tx = randomSigner.sendTransaction({
        to: zetaConnectorZEVM.address,
        value: 1000,
      });

      await expect(tx).to.be.revertedWith("OnlyWZETAOrFungible");
    });

    it("Should reject if call onReceive from other than fungible module", async () => {
      const tx = zetaConnectorZEVM.onReceive(
        randomSigner.address,
        1,
        zetaReceiverMockContract.address,
        1000,
        new ethers.utils.AbiCoder().encode(["string"], ["hello"]),
        ethers.constants.HashZero,
        { value: 1000 }
      );

      await expect(tx).to.be.revertedWith("OnlyFungibleModule");
    });

    it("Should reject if call onRevert from other than fungible module", async () => {
      const tx = zetaConnectorZEVM.onRevert(
        zetaReceiverMockContract.address,
        1,
        randomSigner.address,
        5,
        1000,
        new ethers.utils.AbiCoder().encode(["string"], ["hello"]),
        ethers.constants.HashZero,
        { value: 1000 }
      );

      await expect(tx).to.be.revertedWith("OnlyFungibleModule");
    });

    it("Should revert if value is not the same as declared in onReceive args", async () => {
      const tx = zetaConnectorZEVM
        .connect(fungibleModuleSigner)
        .onReceive(
          randomSigner.address,
          1,
          zetaReceiverMockContract.address,
          1000,
          new ethers.utils.AbiCoder().encode(["string"], ["hello"]),
          ethers.constants.HashZero,
          { value: 900 }
        );

      await expect(tx).to.be.revertedWith("WrongValue");
    });

    it("Should revert if value is not the same as declared in onRevert args", async () => {
      const tx = zetaConnectorZEVM
        .connect(fungibleModuleSigner)
        .onRevert(
          zetaReceiverMockContract.address,
          1,
          randomSigner.address,
          5,
          1000,
          new ethers.utils.AbiCoder().encode(["string"], ["hello"]),
          ethers.constants.HashZero,
          { value: 900 }
        );

      await expect(tx).to.be.revertedWith("WrongValue");
    });
  });
});
