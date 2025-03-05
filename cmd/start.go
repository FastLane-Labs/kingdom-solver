package main

import (
	"bytes"
	"context"
	"flag"
	"math/big"
	"os"
	"os/signal"
	"time"

	"github.com/FastLane-Labs/atlas-sdk-go/types"
	"github.com/FastLane-Labs/kingdom-solver/log"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

const (
	permissionlessAuctionMethod = "atlas_permissionlessAuction"
	solverNamespace             = "solver"
)

type AuctionRequest struct {
	ChainId            *hexutil.Big
	DAppControlAddress common.Address
	Data               hexutil.Bytes
}

func main() {
	log.InitLogger("debug")

	ethClient, err := ethclient.Dial(os.Getenv("RPC_URL"))
	if err != nil {
		log.Error("failed to connect to the RPC", "error", err)
		os.Exit(1)
	}

	chainId, err := ethClient.ChainID(context.Background())
	if err != nil {
		log.Error("failed to get the chain ID", "error", err)
		os.Exit(1)
	}

	kingdomDappControlAddress := os.Getenv("KINGDOM_DAPP_CONTROL")
	if !common.IsHexAddress(kingdomDappControlAddress) {
		log.Error("invalid kingdom dapp control address", "address", kingdomDappControlAddress)
		os.Exit(1)
	}

	flag.Parse()

	if flag.Arg(0) == "solver" {
		runSolver(chainId, ethClient, common.HexToAddress(kingdomDappControlAddress))
	} else if flag.Arg(0) == "trigger-auction" {
		triggerAuction(chainId, common.HexToAddress(kingdomDappControlAddress))
	} else {
		log.Error("invalid command", "command", flag.Arg(0))
		os.Exit(1)
	}
}

func runSolver(
	chainId *big.Int,
	ethClient *ethclient.Client,
	kingdomDappControlAddress common.Address,
) {
	operationsRelayClient, err := rpc.Dial(os.Getenv("OPERATIONS_RELAY_URL"))
	if err != nil {
		log.Error("failed to connect to the operations relay", "error", err)
		os.Exit(1)
	}

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	var (
		uoChan chan *types.UserOperationPartialRaw
		sub    *rpc.ClientSubscription
	)

	subscribe := func() {
		for {
			uoChan = make(chan *types.UserOperationPartialRaw, 32)

			sub, err = operationsRelayClient.Subscribe(context.Background(), solverNamespace, uoChan)
			if err != nil {
				log.Error("failed to subscribe to the operations relay", "retry", "1s", "error", err)

				select {
				case <-time.After(1 * time.Second):
				case <-interrupt:
					os.Exit(0)
				}

				continue
			}

			break
		}
	}

	subscribe()

	for {
		select {
		case <-interrupt:
			log.Info("shutting down...")
			return

		case err := <-sub.Err():
			log.Error("subscription error, resubscribing...", "error", err)
			subscribe()

		case uo := <-uoChan:
			log.Info("received user operation", "userOperation", uo)
		}
	}
}

func triggerAuction(chainId *big.Int, kingdomDappControlAddress common.Address) {
	auctioneerClient, err := rpc.Dial(os.Getenv("AUCTIONEER_URL"))
	if err != nil {
		log.Error("failed to connect to the auctioneer", "error", err)
		os.Exit(1)
	}

	poolAddressString := os.Getenv("POOL_ADDRESS")
	if !common.IsHexAddress(poolAddressString) {
		log.Error("invalid pool address", "address", poolAddressString)
		os.Exit(1)
	}

	tokenOutAddressString := os.Getenv("TOKEN_OUT_ADDRESS")
	if !common.IsHexAddress(tokenOutAddressString) {
		log.Error("invalid token out address", "address", tokenOutAddressString)
		os.Exit(1)
	}

	poolAddress := common.HexToAddress(poolAddressString)
	tokenOutAddress := common.HexToAddress(tokenOutAddressString)

	request := &AuctionRequest{
		ChainId:            (*hexutil.Big)(chainId),
		DAppControlAddress: kingdomDappControlAddress,
		Data:               bytes.Join([][]byte{poolAddress.Bytes(), tokenOutAddress.Bytes()}, []byte{}),
	}

	if err := auctioneerClient.CallContext(context.Background(), nil, permissionlessAuctionMethod, request); err != nil {
		log.Error("failed to call atlas_permissionlessAuction", "error", err)
	}
}
