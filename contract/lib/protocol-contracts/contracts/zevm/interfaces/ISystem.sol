// SPDX-License-Identifier: MIT
pragma solidity 0.8.7;

/**
 * @dev Interfaces of SystemContract and ZRC20 to make easier to import.
 */
interface ISystem {
    function FUNGIBLE_MODULE_ADDRESS() external view returns (address);

    function wZetaContractAddress() external view returns (address);

    function uniswapv2FactoryAddress() external view returns (address);

    function gasPriceByChainId(uint256 chainID) external view returns (uint256);

    function gasCoinZRC20ByChainId(uint256 chainID) external view returns (address);

    function gasZetaPoolByChainId(uint256 chainID) external view returns (address);
}
