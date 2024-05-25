require("dotenv").config();
const { ethers, Contract } = require('ethers');
const DepositorABI = require("../../contracts/abi/Depositor.json");
const MainRouterABI = require("../../contracts/abi/MainRouter.json");
const NetworkInfomation = require("../src/NetworkInfomation.json");
const {
    currentChainID,
    getWallet,
    getCurrentChainId,
} = require("./helper")

async function getDepositedAmount(chainId, tokenSymbol) {
    const wallet = getWallet(chainId);

    const DEPOSITOR_ADDRESS = NetworkInfomation[chainId].DEPOSITOR_ADDRESS;
    const depositorContract = new Contract(DEPOSITOR_ADDRESS, DepositorABI, wallet);

    const tokenAddress = NetworkInfomation[chainId]["TOKEN"][tokenSymbol].address;
    const walletAddress = "0xB1A296a720D9AAF5c5e9F805d8095e6d94882eE1";
    const deposited = await depositorContract.getDeposited(walletAddress, tokenAddress);

    const depositedValueFormat = ethers.utils.formatUnits(deposited, "ether");
    console.log(`Deposited in ${chainId} with ${tokenSymbol}: ${deposited.toString()}`);
    return depositedValueFormat;
}

async function getDepositedValue(chainId, tokenSymbol) {
    const avalancheFujiChainId = 43113;
    const wallet = getWallet(avalancheFujiChainId);

    const MAIN_ROUTER_ADDRESS = NetworkInfomation[chainId].MAIN_ROUTER_ADDRESS;
    const mainRouterContract = new Contract(MAIN_ROUTER_ADDRESS, MainRouterABI, wallet);
    const CHAIN_SELECTOR = NetworkInfomation[chainId].CHAIN_SELECTOR;
    const walletAddress = "0xB1A296a720D9AAF5c5e9F805d8095e6d94882eE1";
    const tokenAddress = NetworkInfomation[chainId]["TOKEN"][tokenSymbol].address;

    const depositedAmount = await mainRouterContract._getUserDepositedAmount(walletAddress, CHAIN_SELECTOR, tokenAddress);
    const depositedValue = await mainRouterContract.getUserCollateralValue(walletAddress, CHAIN_SELECTOR, tokenAddress);


    console.log(`Deposited in ${chainId} with ${tokenSymbol} Amount: ${depositedAmount.toString()}`);
    console.log(`Deposited in ${chainId} with ${tokenSymbol} Value: ${depositedValue.toString()}`);
    return depositedValue;
}


async function getTotalDepositedValueOnChain(chainId) {
    const avalancheFujiChainId = 43113;
    const wallet = getWallet(avalancheFujiChainId);

    const MAIN_ROUTER_ADDRESS = NetworkInfomation[chainId].MAIN_ROUTER_ADDRESS;
    const mainRouterContract = new Contract(MAIN_ROUTER_ADDRESS, MainRouterABI, wallet);
    const CHAIN_SELECTOR = NetworkInfomation[chainId].CHAIN_SELECTOR;
    const walletAddress = "0xB1A296a720D9AAF5c5e9F805d8095e6d94882eE1";

    const { totalCollateral, totalMinted } = await mainRouterContract.getUserOnChainInformation(walletAddress, CHAIN_SELECTOR);
    console.log(`Total Collateral in chain ${chainId}: ${totalCollateral}`);
    return totalCollateral.toString();
}

async function getTotalDepositedValueOverallChain() {
    const avalancheFujiChainId = 43113;
    const wallet = getWallet(avalancheFujiChainId);

    const MAIN_ROUTER_ADDRESS = NetworkInfomation[avalancheFujiChainId].MAIN_ROUTER_ADDRESS;
    const mainRouterContract = new Contract(MAIN_ROUTER_ADDRESS, MainRouterABI, wallet);
    const walletAddress = "0xB1A296a720D9AAF5c5e9F805d8095e6d94882eE1";

    const { totalCollateral, totalMinted } = await mainRouterContract.getUserOverallInformation(walletAddress);
    console.log(`Total Collateral overall chain: ${totalCollateral.toString()}`);
    return totalCollateral.toString();
}

module.exports = {
    getDepositedAmount,
    getDepositedValue,
    getTotalDepositedValueOnChain,
    getTotalDepositedValueOverallChain,
}