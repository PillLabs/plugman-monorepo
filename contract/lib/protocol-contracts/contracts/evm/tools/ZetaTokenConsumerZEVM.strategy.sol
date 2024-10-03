// SPDX-License-Identifier: MIT
pragma solidity 0.8.7;

import "@openzeppelin/contracts/interfaces/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import "@uniswap/v2-periphery/contracts/interfaces/IUniswapV2Router02.sol";

import "../interfaces/ZetaInterfaces.sol";
import "../../zevm/interfaces/IWZETA.sol";

interface ZetaTokenConsumerZEVMErrors {
    error InputCantBeZero();

    error ErrorSendingETH();

    error ReentrancyError();

    error NotEnoughValue();

    error InputCantBeZeta();

    error OutputCantBeZeta();

    error OnlyWZETAAllowed();

    error InvalidForZEVM();
}

/**
 * @dev ZetaTokenConsumer for ZEVM
 */
contract ZetaTokenConsumerZEVM is ZetaTokenConsumer, ZetaTokenConsumerZEVMErrors {
    using SafeERC20 for IERC20;
    uint256 internal constant MAX_DEADLINE = 200;

    address public immutable WETH9Address;

    IUniswapV2Router02 internal immutable uniswapV2Router;

    constructor(address WETH9Address_, address uniswapV2Router_) {
        if (WETH9Address_ == address(0) || uniswapV2Router_ == address(0)) revert ZetaCommonErrors.InvalidAddress();

        WETH9Address = WETH9Address_;
        uniswapV2Router = IUniswapV2Router02(uniswapV2Router_);
    }

    function getZetaFromEth(
        address destinationAddress,
        uint256 minAmountOut
    ) external payable override returns (uint256) {
        if (destinationAddress == address(0)) revert ZetaCommonErrors.InvalidAddress();
        if (msg.value == 0) revert InputCantBeZero();
        if (msg.value < minAmountOut) revert NotEnoughValue();

        IWETH9(WETH9Address).deposit{value: msg.value}();
        IERC20(WETH9Address).safeTransfer(destinationAddress, msg.value);
        emit EthExchangedForZeta(msg.value, msg.value);
        return msg.value;
    }

    function getZetaFromToken(
        address destinationAddress,
        uint256 minAmountOut,
        address inputToken,
        uint256 inputTokenAmount
    ) external override returns (uint256) {
        if (destinationAddress == address(0) || inputToken == address(0)) revert ZetaCommonErrors.InvalidAddress();
        if (inputTokenAmount == 0) revert InputCantBeZero();
        if (inputToken == WETH9Address) revert InputCantBeZeta();

        IERC20(inputToken).safeTransferFrom(msg.sender, address(this), inputTokenAmount);
        IERC20(inputToken).safeApprove(address(uniswapV2Router), inputTokenAmount);

        address[] memory path = new address[](2);
        path[0] = inputToken;
        path[1] = WETH9Address;

        uint256[] memory amounts = uniswapV2Router.swapExactTokensForTokens(
            inputTokenAmount,
            minAmountOut,
            path,
            destinationAddress,
            block.timestamp + MAX_DEADLINE
        );
        uint256 amountOut = amounts[path.length - 1];

        emit TokenExchangedForZeta(inputToken, inputTokenAmount, amountOut);
        return amountOut;
    }

    function getEthFromZeta(
        address destinationAddress,
        uint256 minAmountOut,
        uint256 zetaTokenAmount
    ) external override returns (uint256) {
        if (destinationAddress == address(0)) revert ZetaCommonErrors.InvalidAddress();
        if (zetaTokenAmount == 0) revert InputCantBeZero();
        if (zetaTokenAmount < minAmountOut) revert NotEnoughValue();

        IERC20(WETH9Address).safeTransferFrom(msg.sender, address(this), zetaTokenAmount);
        IWETH9(WETH9Address).withdraw(zetaTokenAmount);

        emit ZetaExchangedForEth(zetaTokenAmount, zetaTokenAmount);

        (bool sent, ) = destinationAddress.call{value: zetaTokenAmount}("");
        if (!sent) revert ErrorSendingETH();

        return zetaTokenAmount;
    }

    function getTokenFromZeta(
        address destinationAddress,
        uint256 minAmountOut,
        address outputToken,
        uint256 zetaTokenAmount
    ) external override returns (uint256) {
        if (destinationAddress == address(0) || outputToken == address(0)) revert ZetaCommonErrors.InvalidAddress();
        if (zetaTokenAmount == 0) revert InputCantBeZero();
        if (outputToken == WETH9Address) revert OutputCantBeZeta();

        IERC20(WETH9Address).safeTransferFrom(msg.sender, address(this), zetaTokenAmount);
        IERC20(WETH9Address).safeApprove(address(uniswapV2Router), zetaTokenAmount);

        address[] memory path = new address[](2);
        path[0] = WETH9Address;
        path[1] = outputToken;

        uint256[] memory amounts = uniswapV2Router.swapExactTokensForTokens(
            zetaTokenAmount,
            minAmountOut,
            path,
            destinationAddress,
            block.timestamp + MAX_DEADLINE
        );

        uint256 amountOut = amounts[path.length - 1];

        emit ZetaExchangedForToken(outputToken, zetaTokenAmount, amountOut);
        return amountOut;
    }

    function hasZetaLiquidity() external view override returns (bool) {
        revert InvalidForZEVM();
    }

    receive() external payable {
        if (msg.sender != WETH9Address) revert OnlyWZETAAllowed();
    }
}
