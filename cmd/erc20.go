package main

import (
	"context"
	"math/big"
	"os"
	"time"

	"github.com/FastLane-Labs/kingdom-solver/contract/erc20"
	"github.com/FastLane-Labs/kingdom-solver/contract/kingdomrouter"
	"github.com/FastLane-Labs/kingdom-solver/contract/weth9"
	"github.com/FastLane-Labs/kingdom-solver/log"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	ethSwapAmount = big.NewInt(1e15)
)

func printSolverContractErc20Balance(ethClient *ethclient.Client, erc20Address string) {
	solverContractAddress, err := getSolverContract()
	if err != nil {
		log.Error("failed to get solver contract address", "error", err)
		os.Exit(1)
	}

	if !common.IsHexAddress(erc20Address) {
		log.Error("invalid erc20 address", "address", erc20Address)
		os.Exit(1)
	}

	erc20Contract, err := erc20.NewErc20(common.HexToAddress(erc20Address), ethClient)
	if err != nil {
		log.Error("failed to get erc20 contract", "error", err)
		os.Exit(1)
	}

	opts, cancel := getOptsWithTimeout()
	symbol, err := erc20Contract.Symbol(opts)
	cancel()

	if err != nil {
		log.Error("failed to get erc20 symbol", "error", err)
		os.Exit(1)
	}

	opts, cancel = getOptsWithTimeout()
	balance, err := erc20Contract.BalanceOf(opts, solverContractAddress)
	cancel()

	if err != nil {
		log.Error("failed to get erc20 balance", "error", err)
		os.Exit(1)
	}

	log.Info("erc20 balance", "token", symbol, "balance", balance)
}

func wrapEthToSolverContract(chainId *big.Int, ethClient *ethclient.Client) {
	wethAddress := os.Getenv("WETH_ADDRESS")
	if !common.IsHexAddress(wethAddress) {
		log.Error("invalid weth address", "address", wethAddress)
		os.Exit(1)
	}

	wethContract, err := weth9.NewWeth9(common.HexToAddress(wethAddress), ethClient)
	if err != nil {
		log.Error("failed to get weth contract", "error", err)
		os.Exit(1)
	}

	solverPk, solverAddress, err := getSolverPk()
	if err != nil {
		log.Error("failed to get solver pk", "error", err)
		os.Exit(1)
	}

	solverContractAddress, err := getSolverContract()
	if err != nil {
		log.Error("failed to get solver contract address", "error", err)
		os.Exit(1)
	}

	opts, cancel := getOptsWithTimeout()
	currentBalance, err := wethContract.BalanceOf(opts, solverContractAddress)
	cancel()

	if err != nil {
		log.Error("failed to get weth balance", "error", err)
		os.Exit(1)
	}

	log.Info("current balance", "token", "weth", "balance", currentBalance)

	transactOpts, err := bind.NewKeyedTransactorWithChainID(solverPk, chainId)
	if err != nil {
		log.Error("failed to create transactor", "error", err)
		os.Exit(1)
	}

	transactOpts.From = solverAddress
	transactOpts.Value = ethSwapAmount

	ctx, cancel := getContextWithTimeout()
	transactOpts.Context = ctx

	tx, err := wethContract.Deposit(transactOpts)
	cancel()

	if err != nil {
		log.Error("failed to wrap eth to weth", "error", err)
		os.Exit(1)
	}

	log.Info("wrap eth to weth sent, waiting for inclusion...", "tx", tx.Hash().Hex())

	ctx, cancel = context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	receipt, err := bind.WaitMined(ctx, ethClient, tx)
	if err != nil {
		log.Error("failed to wait for tx inclusion", "error", err)
		os.Exit(1)
	}

	if receipt.Status == types.ReceiptStatusFailed {
		log.Error("transaction reverted")
		os.Exit(1)
	}

	log.Info("wrap eth to weth done")

	transactOpts.Value = common.Big0

	ctx, cancel = getContextWithTimeout()
	transactOpts.Context = ctx

	tx, err = wethContract.Transfer(transactOpts, solverContractAddress, ethSwapAmount)
	cancel()

	if err != nil {
		log.Error("failed to transfer weth to solver contract", "error", err)
		os.Exit(1)
	}

	log.Info("transfer weth to solver contract sent, waiting for inclusion...", "tx", tx.Hash().Hex())

	ctx, cancel = context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	receipt, err = bind.WaitMined(ctx, ethClient, tx)
	if err != nil {
		log.Error("failed to wait for tx inclusion", "error", err)
		os.Exit(1)
	}

	if receipt.Status == types.ReceiptStatusFailed {
		log.Error("transaction reverted")
		os.Exit(1)
	}

	log.Info("transfer weth to solver contract done")

	opts, cancel = getOptsWithTimeout()
	newBalance, err := wethContract.BalanceOf(opts, solverContractAddress)
	cancel()

	if err != nil {
		log.Error("failed to get weth balance", "error", err)
		os.Exit(1)
	}

	log.Info("new balance", "token", "weth", "balance", newBalance)
}

func swapErc20ToSolverContract(chainId *big.Int, ethClient *ethclient.Client, erc20Address string) {
	wethAddress := os.Getenv("WETH_ADDRESS")
	if !common.IsHexAddress(wethAddress) {
		log.Error("invalid weth address", "address", wethAddress)
		os.Exit(1)
	}

	routerAddress := os.Getenv("ROUTER_ADDRESS")
	if !common.IsHexAddress(routerAddress) {
		log.Error("invalid router address", "address", routerAddress)
		os.Exit(1)
	}

	routerContract, err := kingdomrouter.NewKingdomRouter(common.HexToAddress(routerAddress), ethClient)
	if err != nil {
		log.Error("failed to get router contract", "error", err)
		os.Exit(1)
	}

	solverPk, solverAddress, err := getSolverPk()
	if err != nil {
		log.Error("failed to get solver pk", "error", err)
		os.Exit(1)
	}

	solverContractAddress, err := getSolverContract()
	if err != nil {
		log.Error("failed to get solver contract address", "error", err)
		os.Exit(1)
	}

	if !common.IsHexAddress(erc20Address) {
		log.Error("invalid erc20 address", "address", erc20Address)
		os.Exit(1)
	}

	erc20Contract, err := erc20.NewErc20(common.HexToAddress(erc20Address), ethClient)
	if err != nil {
		log.Error("failed to get erc20 contract", "error", err)
		os.Exit(1)
	}

	opts, cancel := getOptsWithTimeout()
	symbol, err := erc20Contract.Symbol(opts)
	cancel()

	if err != nil {
		log.Error("failed to get erc20 symbol", "error", err)
		os.Exit(1)
	}

	opts, cancel = getOptsWithTimeout()
	currentBalance, err := erc20Contract.BalanceOf(opts, solverContractAddress)
	cancel()

	if err != nil {
		log.Error("failed to get erc20 balance", "error", err)
		os.Exit(1)
	}

	log.Info("current balance", "token", symbol, "balance", currentBalance)

	transactOpts, err := bind.NewKeyedTransactorWithChainID(solverPk, chainId)
	if err != nil {
		log.Error("failed to create transactor", "error", err)
		os.Exit(1)
	}

	transactOpts.From = solverAddress
	transactOpts.Value = ethSwapAmount

	ctx, cancel := getContextWithTimeout()
	transactOpts.Context = ctx

	tx, err := routerContract.ExactInputSingle(transactOpts, kingdomrouter.IUniswapV3RouterExactInputSingleParams{
		TokenIn:           common.HexToAddress(wethAddress),
		TokenOut:          common.HexToAddress(erc20Address),
		Fee:               big.NewInt(500),
		Recipient:         solverContractAddress,
		Deadline:          big.NewInt(time.Now().Add(1 * time.Minute).Unix()),
		AmountIn:          ethSwapAmount,
		AmountOutMinimum:  common.Big0,
		SqrtPriceLimitX96: common.Big0,
	})
	cancel()

	if err != nil {
		log.Error("failed to swap erc20 to solver contract", "error", err)
		os.Exit(1)
	}

	log.Info("swap erc20 to solver contract sent, waiting for inclusion...", "tx", tx.Hash().Hex())

	ctx, cancel = context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	receipt, err := bind.WaitMined(ctx, ethClient, tx)
	if err != nil {
		log.Error("failed to wait for tx inclusion", "error", err)
		os.Exit(1)
	}

	if receipt.Status == types.ReceiptStatusFailed {
		log.Error("transaction reverted")
		os.Exit(1)
	}

	log.Info("swap erc20 to solver contract done")

	opts, cancel = getOptsWithTimeout()
	newBalance, err := erc20Contract.BalanceOf(opts, solverContractAddress)
	cancel()

	if err != nil {
		log.Error("failed to get erc20 balance", "error", err)
		os.Exit(1)
	}

	log.Info("new balance", "token", symbol, "balance", newBalance)
}
