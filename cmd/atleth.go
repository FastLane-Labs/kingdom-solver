package main

import (
	"context"
	"fmt"
	"math/big"
	"os"
	"time"

	"github.com/FastLane-Labs/atlas-sdk-go/config"
	atlas_1_1 "github.com/FastLane-Labs/atlas-sdk-go/contract/atlas/1.1"
	"github.com/FastLane-Labs/kingdom-solver/log"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	atlasVersion     = config.AtlasV_1_1
	atlEthBondAmount = big.NewInt(1e15)
)

func printBondedAtleth(chainId *big.Int, ethClient *ethclient.Client) {
	_, solverAddress, err := getSolverPk()
	if err != nil {
		log.Error("failed to get solver pk", "error", err)
		os.Exit(1)
	}

	ctx, cancel := getContextWithTimeout()
	defer cancel()

	balance, err := ethClient.BalanceAt(ctx, solverAddress, nil)
	if err != nil {
		log.Error("failed to get solver balance", "error", err)
		os.Exit(1)
	}

	atlas, err := getAtlasContract(chainId, ethClient)
	if err != nil {
		log.Error("failed to get atlas contract", "error", err)
		os.Exit(1)
	}

	opts, cancel := getOptsWithTimeout()
	defer cancel()

	bondedBalance, err := atlas.BalanceOfBonded(opts, solverAddress)
	if err != nil {
		log.Error("failed to get bonded balance", "error", err)
		os.Exit(1)
	}

	log.Info("balances", "eth", balance, "bondedAtleth", bondedBalance)
}

func depositAndBondAtleth(chainId *big.Int, ethClient *ethclient.Client) {
	solverPk, solverAddress, err := getSolverPk()
	if err != nil {
		log.Error("failed to get solver pk", "error", err)
		os.Exit(1)
	}

	atlas, err := getAtlasContract(chainId, ethClient)
	if err != nil {
		log.Error("failed to get atlas contract", "error", err)
		os.Exit(1)
	}

	opts, err := bind.NewKeyedTransactorWithChainID(solverPk, chainId)
	if err != nil {
		log.Error("failed to create transactor", "error", err)
		os.Exit(1)
	}

	opts.From = solverAddress
	opts.Value = atlEthBondAmount

	ctx, cancel := getContextWithTimeout()
	opts.Context = ctx

	tx, err := atlas.DepositAndBond(opts, atlEthBondAmount)
	cancel()

	if err != nil {
		log.Error("failed to deposit and bond", "error", err)
		os.Exit(1)
	}

	log.Info("deposit and bond sent, waiting for inclusion...", "tx", tx.Hash().Hex())

	ctx, cancel = context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	_, err = bind.WaitMined(ctx, ethClient, tx)
	if err != nil {
		log.Error("failed to wait for tx inclusion", "error", err)
		os.Exit(1)
	}

	log.Info("deposit and bond done")
}

func getAtlasContract(chainId *big.Int, ethClient *ethclient.Client) (*atlas_1_1.Atlas, error) {
	atlasAddress, err := config.GetAtlasAddress(chainId.Uint64(), &atlasVersion)
	if err != nil {
		return nil, fmt.Errorf("failed to get atlas address, %w", err)
	}

	return atlas_1_1.NewAtlas(atlasAddress, ethClient)
}
