package main

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"flag"
	"fmt"
	"math/big"
	"os"
	"os/signal"
	"time"

	"github.com/FastLane-Labs/atlas-sdk-go/config"
	"github.com/FastLane-Labs/atlas-sdk-go/types"
	"github.com/FastLane-Labs/atlas-sdk-go/utils"
	"github.com/FastLane-Labs/kingdom-solver/contract"
	"github.com/FastLane-Labs/kingdom-solver/contract/kingdomdappcontrol"
	"github.com/FastLane-Labs/kingdom-solver/contract/ramsesv2pool"
	"github.com/FastLane-Labs/kingdom-solver/log"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/gorilla/websocket"
	"github.com/mitchellh/mapstructure"
	"golang.org/x/sync/errgroup"
)

const (
	permissionlessAuctionMethod = "atlas_permissionlessAuction"
	submitSolverOperationMethod = "solver_submitSolverOperation"
	solverNamespace             = "solver"
	userOperationsSubscription  = "userOperations"
	feeSetterIntentFunction     = "feeSetterIntent"
	demoSwapFunction            = "demoSwap"
)

var (
	networkTimeout    = 3 * time.Second
	networkRetryDelay = 1 * time.Second
)

type AuctionRequest struct {
	ChainId            *hexutil.Big
	DAppControlAddress common.Address
	Data               hexutil.Bytes
}

type UserOperationNotification struct {
	AuctionId            string                         `json:"auction_id"`
	PartialUserOperation *types.UserOperationPartialRaw `json:"partial_user_operation"`
}

type Solution struct {
	AuctionId       string                    `json:"auction_id"`
	AuctionSolution *types.SolverOperationRaw `json:"auction_solution"`
}

func main() {
	log.InitLogger("debug")

	ctx, cancel := getContextWithTimeout()
	ethClient, err := ethclient.DialContext(ctx, os.Getenv("RPC_URL"))
	cancel()

	if err != nil {
		log.Error("failed to connect to the RPC", "error", err)
		os.Exit(1)
	}

	ctx, cancel = getContextWithTimeout()
	chainId, err := ethClient.ChainID(ctx)
	cancel()

	if err != nil {
		log.Error("failed to get the chain ID", "error", err)
		os.Exit(1)
	}

	flag.Parse()

	switch flag.Arg(0) {
	case "solver":
		runSolver(chainId, ethClient)

	case "trigger-auction":
		triggerAuction(chainId)

	case "print-bonded-atleth":
		printBondedAtleth(chainId, ethClient)

	case "deposit-and-bond-atleth":
		depositAndBondAtleth(chainId, ethClient)

	default:
		log.Error("invalid command", "command", flag.Arg(0))
		os.Exit(1)
	}
}

func runSolver(chainId *big.Int, ethClient *ethclient.Client) {
	kingdomDappControlAddress, err := getKingdomDappControlAddress()
	if err != nil {
		log.Error("failed to get kingdom dapp control address", "error", err)
		os.Exit(1)
	}

	dialerOption := rpc.WithWebsocketDialer(websocket.Dialer{
		ReadBufferSize:  1024 * 1024,
		WriteBufferSize: 1024 * 1024,
	})

	ctx, cancel := getContextWithTimeout()
	operationsRelayClient, err := rpc.DialOptions(ctx, os.Getenv("OPERATIONS_RELAY_URL"), dialerOption)
	cancel()

	if err != nil {
		log.Error("failed to connect to the operations relay", "error", err)
		os.Exit(1)
	}

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	var (
		uoChan chan *UserOperationNotification
		sub    *rpc.ClientSubscription
	)

	subscribe := func() {
		for {
			uoChan = make(chan *UserOperationNotification, 32)

			ctx, cancel := getContextWithTimeout()
			sub, err = operationsRelayClient.Subscribe(ctx, solverNamespace, uoChan, userOperationsSubscription)
			cancel()

			if err != nil {
				log.Error("failed to subscribe to the operations relay", "retry", networkRetryDelay, "error", err)

				select {
				case <-time.After(networkRetryDelay):
				case <-interrupt:
					os.Exit(0)
				}

				continue
			}

			break
		}
	}

	subscribe()

	log.Info("solver started")

	for {
		select {
		case <-interrupt:
			log.Info("shutting down...")
			return

		case err := <-sub.Err():
			log.Error("subscription error, resubscribing...", "error", err)
			subscribe()

		case n := <-uoChan:
			if n.PartialUserOperation.ChainId.ToInt().Cmp(chainId) != 0 {
				// Skipping other chains
				continue
			}

			if n.PartialUserOperation.Control != kingdomDappControlAddress {
				// Skipping other dapps
				continue
			}

			log.Info("received user operation", "auctionId", n.AuctionId, "userOperation", n.PartialUserOperation.UserOpHash.Hex())

			solution, err := buildSolution(ethClient, n)
			if err != nil {
				log.Error("failed to build solution", "auctionId", n.AuctionId, "error", err)
				continue
			}

			log.Info("sending solution")

			ctx, cancel := getContextWithTimeout()
			err = operationsRelayClient.CallContext(ctx, nil, submitSolverOperationMethod, solution)
			cancel()

			if err != nil {
				log.Error("failed to submit solver operation", "auctionId", n.AuctionId, "error", err)
				continue
			}

			log.Info("solver operation successfully submitted", "auctionId", n.AuctionId)
		}
	}
}

func getSolverPk() (*ecdsa.PrivateKey, common.Address, error) {
	pk, err := crypto.HexToECDSA(os.Getenv("SOLVER_PK"))
	if err != nil {
		return nil, common.Address{}, fmt.Errorf("failed to get solver pk, %w", err)
	}

	return pk, crypto.PubkeyToAddress(pk.PublicKey), nil
}

func getSolverContract() (common.Address, error) {
	solverContractAddress := os.Getenv("SOLVER_CONTRACT_ADDRESS")
	if !common.IsHexAddress(solverContractAddress) {
		return common.Address{}, fmt.Errorf("invalid solver contract address, %s", solverContractAddress)
	}
	return common.HexToAddress(solverContractAddress), nil
}

func getKingdomDappControlAddress() (common.Address, error) {
	kingdomDappControlAddress := os.Getenv("KINGDOM_DAPP_CONTROL")
	if !common.IsHexAddress(kingdomDappControlAddress) {
		return common.Address{}, fmt.Errorf("invalid kingdom dapp control address, %s", kingdomDappControlAddress)
	}
	return common.HexToAddress(kingdomDappControlAddress), nil
}

func getPoolTokens(ethClient *ethclient.Client, poolAddress common.Address) ([]common.Address, error) {
	// return []common.Address{
	// 	common.HexToAddress("0x82aF49447D8a07e3bd95BD0d56f35241523fBab1"),
	// 	common.HexToAddress("0xaf88d065e77c8cC2239327C5EDb3A432268e5831"),
	// }, nil

	pool, err := ramsesv2pool.NewRamsesV2Pool(poolAddress, ethClient)
	if err != nil {
		return nil, fmt.Errorf("failed to get pool, %w", err)
	}

	var (
		g, _   = errgroup.WithContext(context.Background())
		token0 common.Address
		token1 common.Address
	)

	g.Go(func() error {
		opts, cancel := getOptsWithTimeout()
		defer cancel()

		token0_, err := pool.Token0(opts)
		if err != nil {
			return err
		}
		token0 = token0_
		return nil
	})

	g.Go(func() error {
		opts, cancel := getOptsWithTimeout()
		defer cancel()

		token1_, err := pool.Token1(opts)
		if err != nil {
			return err
		}
		token1 = token1_
		return nil
	})

	if err := g.Wait(); err != nil {
		return nil, err
	}

	return []common.Address{token0, token1}, nil
}

func buildSolution(ethClient *ethclient.Client, n *UserOperationNotification) (*Solution, error) {
	solverPk, solverAddress, err := getSolverPk()
	if err != nil {
		return nil, fmt.Errorf("failed to get solver pk, %w", err)
	}

	solverContractAddress, err := getSolverContract()
	if err != nil {
		return nil, fmt.Errorf("failed to get solver contract address, %w", err)
	}

	feeSetterIntentMethod, ok := contract.KingdomDAppControlAbi.Methods[feeSetterIntentFunction]
	if !ok {
		return nil, fmt.Errorf("method %s not found", feeSetterIntentFunction)
	}

	inputs, err := feeSetterIntentMethod.Inputs.Unpack(n.PartialUserOperation.Data[4:])
	if err != nil {
		return nil, fmt.Errorf("failed to unpack inputs, %w", err)
	}

	feeIntent := &kingdomdappcontrol.FeeSetterIntent{}
	if err := mapstructure.Decode(inputs[0], feeIntent); err != nil {
		return nil, fmt.Errorf("failed to decode fee intent, %w", err)
	}

	if _, ok := contract.DemoSolverAbi.Methods[demoSwapFunction]; !ok {
		return nil, fmt.Errorf("method %s not found", demoSwapFunction)
	}

	poolTokens, err := getPoolTokens(ethClient, feeIntent.PoolAddress)
	if err != nil {
		return nil, fmt.Errorf("failed to get pool tokens, %w", err)
	}

	var (
		tokenIn  common.Address
		tokenOut common.Address
	)

	if poolTokens[0] == feeIntent.BidToken {
		tokenIn = poolTokens[1]
		tokenOut = poolTokens[0]
	} else {
		tokenIn = poolTokens[0]
		tokenOut = poolTokens[1]
	}

	data, err := contract.DemoSolverAbi.Pack(
		demoSwapFunction,
		feeIntent.PoolAddress,
		tokenIn,
		tokenOut,
		big.NewInt(1e9),
		feeIntent.Fee,
		true,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to pack inputs, %w", err)
	}

	solverOperation := &types.SolverOperation{
		From:         solverAddress,
		To:           n.PartialUserOperation.To,
		Gas:          big.NewInt(500_000),
		MaxFeePerGas: new(big.Int).Set(n.PartialUserOperation.MaxFeePerGas.ToInt()),
		Deadline:     new(big.Int).Set(n.PartialUserOperation.Deadline.ToInt()),
		Solver:       solverContractAddress,
		Control:      n.PartialUserOperation.Control,
		UserOpHash:   n.PartialUserOperation.UserOpHash,
		BidToken:     feeIntent.BidToken,
		BidAmount:    big.NewInt(1e9),
		Data:         data,
	}

	chainId := n.PartialUserOperation.ChainId.ToInt().Uint64()

	atlasVersion, err := config.GetVersionFromAtlasAddress(chainId, n.PartialUserOperation.To)
	if err != nil {
		return nil, fmt.Errorf("failed to get atlas version, %w", err)
	}

	hash, err := solverOperation.Hash(chainId, &atlasVersion)
	if err != nil {
		return nil, fmt.Errorf("failed to get hash, %w", err)
	}

	signature, err := utils.SignMessage(hash.Bytes(), solverPk)
	if err != nil {
		return nil, fmt.Errorf("failed to sign message, %w", err)
	}

	solverOperation.Signature = signature

	return &Solution{
		AuctionId:       n.AuctionId,
		AuctionSolution: solverOperation.EncodeToRaw(),
	}, nil
}

func triggerAuction(chainId *big.Int) {
	kingdomDappControlAddress, err := getKingdomDappControlAddress()
	if err != nil {
		log.Error("failed to get kingdom dapp control address", "error", err)
		os.Exit(1)
	}

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

	ctx, cancel := getContextWithTimeout()
	defer cancel()

	if err := auctioneerClient.CallContext(ctx, nil, permissionlessAuctionMethod, request); err != nil {
		log.Error("failed to call atlas_permissionlessAuction", "error", err)
	}
}

func getContextWithTimeout() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), networkTimeout)
}

func getOptsWithTimeout() (*bind.CallOpts, context.CancelFunc) {
	ctx, cancel := getContextWithTimeout()
	return &bind.CallOpts{Context: ctx}, cancel
}
