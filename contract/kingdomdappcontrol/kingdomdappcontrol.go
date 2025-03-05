// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package kingdomdappcontrol

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// CallConfig is an auto generated low-level Go binding around an user-defined struct.
type CallConfig struct {
	UserNoncesSequential      bool
	DappNoncesSequential      bool
	RequirePreOps             bool
	TrackPreOpsReturnData     bool
	TrackUserReturnData       bool
	DelegateUser              bool
	RequirePreSolver          bool
	RequirePostSolver         bool
	RequirePostOps            bool
	ZeroSolvers               bool
	ReuseUserOp               bool
	UserAuctioneer            bool
	SolverAuctioneer          bool
	UnknownAuctioneer         bool
	VerifyCallChainHash       bool
	ForwardReturnData         bool
	RequireFulfillment        bool
	TrustedOpHash             bool
	InvertBidValue            bool
	ExPostBids                bool
	AllowAllocateValueFailure bool
}

// DAppConfig is an auto generated low-level Go binding around an user-defined struct.
type DAppConfig struct {
	To             common.Address
	CallConfig     uint32
	BidToken       common.Address
	SolverGasLimit uint32
}

// FeeSetterIntent is an auto generated low-level Go binding around an user-defined struct.
type FeeSetterIntent struct {
	PoolAddress         common.Address
	BidToken            common.Address
	Fee                 *big.Int
	Token0BalanceBefore *big.Int
	Token1BalanceBefore *big.Int
}

// SolverOperation is an auto generated low-level Go binding around an user-defined struct.
type SolverOperation struct {
	From         common.Address
	To           common.Address
	Value        *big.Int
	Gas          *big.Int
	MaxFeePerGas *big.Int
	Deadline     *big.Int
	Solver       common.Address
	Control      common.Address
	UserOpHash   [32]byte
	BidToken     common.Address
	BidAmount    *big.Int
	Data         []byte
	Signature    []byte
}

// UserOperation is an auto generated low-level Go binding around an user-defined struct.
type UserOperation struct {
	From         common.Address
	To           common.Address
	Value        *big.Int
	Gas          *big.Int
	MaxFeePerGas *big.Int
	Nonce        *big.Int
	Deadline     *big.Int
	Dapp         common.Address
	Control      common.Address
	CallConfig   uint32
	SessionKey   common.Address
	Data         []byte
	Signature    []byte
}

// KingdomDAppControlMetaData contains all meta data concerning the KingdomDAppControl contract.
var KingdomDAppControlMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_atlas\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_govSharePercent\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_govPayoutAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_fastlane\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_factory\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"ATLAS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ATLAS_VERIFICATION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"CALL_CONFIG\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"CONTROL\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PERCENTAGE_DENOMINATOR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"SOURCE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"acceptGovernance\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addAuthorizedUser\",\"inputs\":[{\"name\":\"_user\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"allocateValueCall\",\"inputs\":[{\"name\":\"bidToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"bidAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"executionEnvironmentWhitelist\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"factory\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"fastlane\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"feeSetterIntent\",\"inputs\":[{\"name\":\"_feeSetterIntent\",\"type\":\"tuple\",\"internalType\":\"structFeeSetterIntent\",\"components\":[{\"name\":\"poolAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"bidToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"fee\",\"type\":\"uint24\",\"internalType\":\"uint24\"},{\"name\":\"token0BalanceBefore\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"token1BalanceBefore\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structFeeSetterIntent\",\"components\":[{\"name\":\"poolAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"bidToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"fee\",\"type\":\"uint24\",\"internalType\":\"uint24\"},{\"name\":\"token0BalanceBefore\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"token1BalanceBefore\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getBidFormat\",\"inputs\":[{\"name\":\"userOp\",\"type\":\"tuple\",\"internalType\":\"structUserOperation\",\"components\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"gas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxFeePerGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"deadline\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"dapp\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"control\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callConfig\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"sessionKey\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getBidValue\",\"inputs\":[{\"name\":\"solverOp\",\"type\":\"tuple\",\"internalType\":\"structSolverOperation\",\"components\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"gas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxFeePerGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"deadline\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"solver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"control\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"userOpHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"bidToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"bidAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCallConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structCallConfig\",\"components\":[{\"name\":\"userNoncesSequential\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"dappNoncesSequential\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"requirePreOps\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"trackPreOpsReturnData\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"trackUserReturnData\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"delegateUser\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"requirePreSolver\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"requirePostSolver\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"requirePostOps\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"zeroSolvers\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"reuseUserOp\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"userAuctioneer\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"solverAuctioneer\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"unknownAuctioneer\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"verifyCallChainHash\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"forwardReturnData\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"requireFulfillment\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"trustedOpHash\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"invertBidValue\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"exPostBids\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"allowAllocateValueFailure\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDAppConfig\",\"inputs\":[{\"name\":\"userOp\",\"type\":\"tuple\",\"internalType\":\"structUserOperation\",\"components\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"gas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxFeePerGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"deadline\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"dapp\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"control\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callConfig\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"sessionKey\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"dConfig\",\"type\":\"tuple\",\"internalType\":\"structDAppConfig\",\"components\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callConfig\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"bidToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"solverGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDAppSignatory\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFactory\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPayoutData\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSolverGasLimit\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"govPayoutAddr\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"govSharePercent\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"governance\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pendingGovernance\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"postOpsCall\",\"inputs\":[{\"name\":\"solved\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"postSolverCall\",\"inputs\":[{\"name\":\"solverOp\",\"type\":\"tuple\",\"internalType\":\"structSolverOperation\",\"components\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"gas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxFeePerGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"deadline\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"solver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"control\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"userOpHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"bidToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"bidAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"returnData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"preOpsCall\",\"inputs\":[{\"name\":\"userOp\",\"type\":\"tuple\",\"internalType\":\"structUserOperation\",\"components\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"gas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxFeePerGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"deadline\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"dapp\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"control\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callConfig\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"sessionKey\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"preSolverCall\",\"inputs\":[{\"name\":\"solverOp\",\"type\":\"tuple\",\"internalType\":\"structSolverOperation\",\"components\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"gas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxFeePerGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"deadline\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"solver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"control\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"userOpHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"bidToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"bidAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"returnData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"removeAuthorizedUser\",\"inputs\":[{\"name\":\"_user\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"requireSequentialDAppNonces\",\"inputs\":[],\"outputs\":[{\"name\":\"isSequential\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"requireSequentialUserNonces\",\"inputs\":[],\"outputs\":[{\"name\":\"isSequential\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"resetFee\",\"inputs\":[{\"name\":\"_poolAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_fee\",\"type\":\"uint24\",\"internalType\":\"uint24\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setFastlanePayoutAddr\",\"inputs\":[{\"name\":\"_fastlanePayoutAddr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setFeeSetter\",\"inputs\":[{\"name\":\"_feeSetter\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setGovPayoutAddr\",\"inputs\":[{\"name\":\"_govPayoutAddr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setGovSharePercent\",\"inputs\":[{\"name\":\"_govSharePercent\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setSolverGasLimit\",\"inputs\":[{\"name\":\"_solverGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"solverGasLimit\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferGovernance\",\"inputs\":[{\"name\":\"newGovernance\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"userDelegated\",\"inputs\":[],\"outputs\":[{\"name\":\"delegated\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"FastlanePayoutAddressUpdated\",\"inputs\":[{\"name\":\"oldFastlanePayoutAddr\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newFastlanePayoutAddr\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"GovernancePayoutAddressUpdated\",\"inputs\":[{\"name\":\"oldGovPayoutAddr\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newGovPayoutAddr\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"GovernanceSharePercentUpdated\",\"inputs\":[{\"name\":\"oldGovSharePercent\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"newGovSharePercent\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"GovernanceTransferStarted\",\"inputs\":[{\"name\":\"previousGovernance\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newGovernance\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"GovernanceTransferred\",\"inputs\":[{\"name\":\"previousGovernance\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newGovernance\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AlteredControl\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"BothPreOpsAndUserReturnDataCannotBeTracked\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"BothUserAndDAppNoncesCannotBeSequential\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientBalance\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidControl\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidDelegateCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidFeeSetter\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidPayoutData\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidSolver\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvertBidValueCannotBeExPostBids\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustBeDelegatecalled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NoDelegatecall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotAuthorized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotImplemented\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyAtlas\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyGovernance\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyGovernance\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PoolFeeAlreadyZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PoolFeeNotSetToZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ResetFeeFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Token0BalanceDidNotChange\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Token1BalanceDidNotChange\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Unauthorized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"WrongPhase\",\"inputs\":[]}]",
}

// KingdomDAppControlABI is the input ABI used to generate the binding from.
// Deprecated: Use KingdomDAppControlMetaData.ABI instead.
var KingdomDAppControlABI = KingdomDAppControlMetaData.ABI

// KingdomDAppControl is an auto generated Go binding around an Ethereum contract.
type KingdomDAppControl struct {
	KingdomDAppControlCaller     // Read-only binding to the contract
	KingdomDAppControlTransactor // Write-only binding to the contract
	KingdomDAppControlFilterer   // Log filterer for contract events
}

// KingdomDAppControlCaller is an auto generated read-only Go binding around an Ethereum contract.
type KingdomDAppControlCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KingdomDAppControlTransactor is an auto generated write-only Go binding around an Ethereum contract.
type KingdomDAppControlTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KingdomDAppControlFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type KingdomDAppControlFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KingdomDAppControlSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type KingdomDAppControlSession struct {
	Contract     *KingdomDAppControl // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// KingdomDAppControlCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type KingdomDAppControlCallerSession struct {
	Contract *KingdomDAppControlCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// KingdomDAppControlTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type KingdomDAppControlTransactorSession struct {
	Contract     *KingdomDAppControlTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// KingdomDAppControlRaw is an auto generated low-level Go binding around an Ethereum contract.
type KingdomDAppControlRaw struct {
	Contract *KingdomDAppControl // Generic contract binding to access the raw methods on
}

// KingdomDAppControlCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type KingdomDAppControlCallerRaw struct {
	Contract *KingdomDAppControlCaller // Generic read-only contract binding to access the raw methods on
}

// KingdomDAppControlTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type KingdomDAppControlTransactorRaw struct {
	Contract *KingdomDAppControlTransactor // Generic write-only contract binding to access the raw methods on
}

// NewKingdomDAppControl creates a new instance of KingdomDAppControl, bound to a specific deployed contract.
func NewKingdomDAppControl(address common.Address, backend bind.ContractBackend) (*KingdomDAppControl, error) {
	contract, err := bindKingdomDAppControl(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &KingdomDAppControl{KingdomDAppControlCaller: KingdomDAppControlCaller{contract: contract}, KingdomDAppControlTransactor: KingdomDAppControlTransactor{contract: contract}, KingdomDAppControlFilterer: KingdomDAppControlFilterer{contract: contract}}, nil
}

// NewKingdomDAppControlCaller creates a new read-only instance of KingdomDAppControl, bound to a specific deployed contract.
func NewKingdomDAppControlCaller(address common.Address, caller bind.ContractCaller) (*KingdomDAppControlCaller, error) {
	contract, err := bindKingdomDAppControl(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &KingdomDAppControlCaller{contract: contract}, nil
}

// NewKingdomDAppControlTransactor creates a new write-only instance of KingdomDAppControl, bound to a specific deployed contract.
func NewKingdomDAppControlTransactor(address common.Address, transactor bind.ContractTransactor) (*KingdomDAppControlTransactor, error) {
	contract, err := bindKingdomDAppControl(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &KingdomDAppControlTransactor{contract: contract}, nil
}

// NewKingdomDAppControlFilterer creates a new log filterer instance of KingdomDAppControl, bound to a specific deployed contract.
func NewKingdomDAppControlFilterer(address common.Address, filterer bind.ContractFilterer) (*KingdomDAppControlFilterer, error) {
	contract, err := bindKingdomDAppControl(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &KingdomDAppControlFilterer{contract: contract}, nil
}

// bindKingdomDAppControl binds a generic wrapper to an already deployed contract.
func bindKingdomDAppControl(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(KingdomDAppControlABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KingdomDAppControl *KingdomDAppControlRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _KingdomDAppControl.Contract.KingdomDAppControlCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KingdomDAppControl *KingdomDAppControlRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KingdomDAppControl.Contract.KingdomDAppControlTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KingdomDAppControl *KingdomDAppControlRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KingdomDAppControl.Contract.KingdomDAppControlTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KingdomDAppControl *KingdomDAppControlCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _KingdomDAppControl.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KingdomDAppControl *KingdomDAppControlTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KingdomDAppControl.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KingdomDAppControl *KingdomDAppControlTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KingdomDAppControl.Contract.contract.Transact(opts, method, params...)
}

// ATLAS is a free data retrieval call binding the contract method 0xe982ae92.
//
// Solidity: function ATLAS() view returns(address)
func (_KingdomDAppControl *KingdomDAppControlCaller) ATLAS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _KingdomDAppControl.contract.Call(opts, &out, "ATLAS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ATLAS is a free data retrieval call binding the contract method 0xe982ae92.
//
// Solidity: function ATLAS() view returns(address)
func (_KingdomDAppControl *KingdomDAppControlSession) ATLAS() (common.Address, error) {
	return _KingdomDAppControl.Contract.ATLAS(&_KingdomDAppControl.CallOpts)
}

// ATLAS is a free data retrieval call binding the contract method 0xe982ae92.
//
// Solidity: function ATLAS() view returns(address)
func (_KingdomDAppControl *KingdomDAppControlCallerSession) ATLAS() (common.Address, error) {
	return _KingdomDAppControl.Contract.ATLAS(&_KingdomDAppControl.CallOpts)
}

// ATLASVERIFICATION is a free data retrieval call binding the contract method 0xbf230cfb.
//
// Solidity: function ATLAS_VERIFICATION() view returns(address)
func (_KingdomDAppControl *KingdomDAppControlCaller) ATLASVERIFICATION(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _KingdomDAppControl.contract.Call(opts, &out, "ATLAS_VERIFICATION")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ATLASVERIFICATION is a free data retrieval call binding the contract method 0xbf230cfb.
//
// Solidity: function ATLAS_VERIFICATION() view returns(address)
func (_KingdomDAppControl *KingdomDAppControlSession) ATLASVERIFICATION() (common.Address, error) {
	return _KingdomDAppControl.Contract.ATLASVERIFICATION(&_KingdomDAppControl.CallOpts)
}

// ATLASVERIFICATION is a free data retrieval call binding the contract method 0xbf230cfb.
//
// Solidity: function ATLAS_VERIFICATION() view returns(address)
func (_KingdomDAppControl *KingdomDAppControlCallerSession) ATLASVERIFICATION() (common.Address, error) {
	return _KingdomDAppControl.Contract.ATLASVERIFICATION(&_KingdomDAppControl.CallOpts)
}

// CALLCONFIG is a free data retrieval call binding the contract method 0x8d212978.
//
// Solidity: function CALL_CONFIG() view returns(uint32)
func (_KingdomDAppControl *KingdomDAppControlCaller) CALLCONFIG(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _KingdomDAppControl.contract.Call(opts, &out, "CALL_CONFIG")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// CALLCONFIG is a free data retrieval call binding the contract method 0x8d212978.
//
// Solidity: function CALL_CONFIG() view returns(uint32)
func (_KingdomDAppControl *KingdomDAppControlSession) CALLCONFIG() (uint32, error) {
	return _KingdomDAppControl.Contract.CALLCONFIG(&_KingdomDAppControl.CallOpts)
}

// CALLCONFIG is a free data retrieval call binding the contract method 0x8d212978.
//
// Solidity: function CALL_CONFIG() view returns(uint32)
func (_KingdomDAppControl *KingdomDAppControlCallerSession) CALLCONFIG() (uint32, error) {
	return _KingdomDAppControl.Contract.CALLCONFIG(&_KingdomDAppControl.CallOpts)
}

// CONTROL is a free data retrieval call binding the contract method 0x4953ecc7.
//
// Solidity: function CONTROL() view returns(address)
func (_KingdomDAppControl *KingdomDAppControlCaller) CONTROL(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _KingdomDAppControl.contract.Call(opts, &out, "CONTROL")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CONTROL is a free data retrieval call binding the contract method 0x4953ecc7.
//
// Solidity: function CONTROL() view returns(address)
func (_KingdomDAppControl *KingdomDAppControlSession) CONTROL() (common.Address, error) {
	return _KingdomDAppControl.Contract.CONTROL(&_KingdomDAppControl.CallOpts)
}

// CONTROL is a free data retrieval call binding the contract method 0x4953ecc7.
//
// Solidity: function CONTROL() view returns(address)
func (_KingdomDAppControl *KingdomDAppControlCallerSession) CONTROL() (common.Address, error) {
	return _KingdomDAppControl.Contract.CONTROL(&_KingdomDAppControl.CallOpts)
}

// PERCENTAGEDENOMINATOR is a free data retrieval call binding the contract method 0xb3cd4254.
//
// Solidity: function PERCENTAGE_DENOMINATOR() view returns(uint256)
func (_KingdomDAppControl *KingdomDAppControlCaller) PERCENTAGEDENOMINATOR(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _KingdomDAppControl.contract.Call(opts, &out, "PERCENTAGE_DENOMINATOR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PERCENTAGEDENOMINATOR is a free data retrieval call binding the contract method 0xb3cd4254.
//
// Solidity: function PERCENTAGE_DENOMINATOR() view returns(uint256)
func (_KingdomDAppControl *KingdomDAppControlSession) PERCENTAGEDENOMINATOR() (*big.Int, error) {
	return _KingdomDAppControl.Contract.PERCENTAGEDENOMINATOR(&_KingdomDAppControl.CallOpts)
}

// PERCENTAGEDENOMINATOR is a free data retrieval call binding the contract method 0xb3cd4254.
//
// Solidity: function PERCENTAGE_DENOMINATOR() view returns(uint256)
func (_KingdomDAppControl *KingdomDAppControlCallerSession) PERCENTAGEDENOMINATOR() (*big.Int, error) {
	return _KingdomDAppControl.Contract.PERCENTAGEDENOMINATOR(&_KingdomDAppControl.CallOpts)
}

// SOURCE is a free data retrieval call binding the contract method 0xf230b4c2.
//
// Solidity: function SOURCE() view returns(address)
func (_KingdomDAppControl *KingdomDAppControlCaller) SOURCE(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _KingdomDAppControl.contract.Call(opts, &out, "SOURCE")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SOURCE is a free data retrieval call binding the contract method 0xf230b4c2.
//
// Solidity: function SOURCE() view returns(address)
func (_KingdomDAppControl *KingdomDAppControlSession) SOURCE() (common.Address, error) {
	return _KingdomDAppControl.Contract.SOURCE(&_KingdomDAppControl.CallOpts)
}

// SOURCE is a free data retrieval call binding the contract method 0xf230b4c2.
//
// Solidity: function SOURCE() view returns(address)
func (_KingdomDAppControl *KingdomDAppControlCallerSession) SOURCE() (common.Address, error) {
	return _KingdomDAppControl.Contract.SOURCE(&_KingdomDAppControl.CallOpts)
}

// ExecutionEnvironmentWhitelist is a free data retrieval call binding the contract method 0xd02a1984.
//
// Solidity: function executionEnvironmentWhitelist(address ) view returns(bool)
func (_KingdomDAppControl *KingdomDAppControlCaller) ExecutionEnvironmentWhitelist(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _KingdomDAppControl.contract.Call(opts, &out, "executionEnvironmentWhitelist", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ExecutionEnvironmentWhitelist is a free data retrieval call binding the contract method 0xd02a1984.
//
// Solidity: function executionEnvironmentWhitelist(address ) view returns(bool)
func (_KingdomDAppControl *KingdomDAppControlSession) ExecutionEnvironmentWhitelist(arg0 common.Address) (bool, error) {
	return _KingdomDAppControl.Contract.ExecutionEnvironmentWhitelist(&_KingdomDAppControl.CallOpts, arg0)
}

// ExecutionEnvironmentWhitelist is a free data retrieval call binding the contract method 0xd02a1984.
//
// Solidity: function executionEnvironmentWhitelist(address ) view returns(bool)
func (_KingdomDAppControl *KingdomDAppControlCallerSession) ExecutionEnvironmentWhitelist(arg0 common.Address) (bool, error) {
	return _KingdomDAppControl.Contract.ExecutionEnvironmentWhitelist(&_KingdomDAppControl.CallOpts, arg0)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_KingdomDAppControl *KingdomDAppControlCaller) Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _KingdomDAppControl.contract.Call(opts, &out, "factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_KingdomDAppControl *KingdomDAppControlSession) Factory() (common.Address, error) {
	return _KingdomDAppControl.Contract.Factory(&_KingdomDAppControl.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_KingdomDAppControl *KingdomDAppControlCallerSession) Factory() (common.Address, error) {
	return _KingdomDAppControl.Contract.Factory(&_KingdomDAppControl.CallOpts)
}

// Fastlane is a free data retrieval call binding the contract method 0x0e292044.
//
// Solidity: function fastlane() view returns(address)
func (_KingdomDAppControl *KingdomDAppControlCaller) Fastlane(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _KingdomDAppControl.contract.Call(opts, &out, "fastlane")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Fastlane is a free data retrieval call binding the contract method 0x0e292044.
//
// Solidity: function fastlane() view returns(address)
func (_KingdomDAppControl *KingdomDAppControlSession) Fastlane() (common.Address, error) {
	return _KingdomDAppControl.Contract.Fastlane(&_KingdomDAppControl.CallOpts)
}

// Fastlane is a free data retrieval call binding the contract method 0x0e292044.
//
// Solidity: function fastlane() view returns(address)
func (_KingdomDAppControl *KingdomDAppControlCallerSession) Fastlane() (common.Address, error) {
	return _KingdomDAppControl.Contract.Fastlane(&_KingdomDAppControl.CallOpts)
}

// GetBidFormat is a free data retrieval call binding the contract method 0x8831b924.
//
// Solidity: function getBidFormat((address,address,uint256,uint256,uint256,uint256,uint256,address,address,uint32,address,bytes,bytes) userOp) view returns(address)
func (_KingdomDAppControl *KingdomDAppControlCaller) GetBidFormat(opts *bind.CallOpts, userOp UserOperation) (common.Address, error) {
	var out []interface{}
	err := _KingdomDAppControl.contract.Call(opts, &out, "getBidFormat", userOp)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetBidFormat is a free data retrieval call binding the contract method 0x8831b924.
//
// Solidity: function getBidFormat((address,address,uint256,uint256,uint256,uint256,uint256,address,address,uint32,address,bytes,bytes) userOp) view returns(address)
func (_KingdomDAppControl *KingdomDAppControlSession) GetBidFormat(userOp UserOperation) (common.Address, error) {
	return _KingdomDAppControl.Contract.GetBidFormat(&_KingdomDAppControl.CallOpts, userOp)
}

// GetBidFormat is a free data retrieval call binding the contract method 0x8831b924.
//
// Solidity: function getBidFormat((address,address,uint256,uint256,uint256,uint256,uint256,address,address,uint32,address,bytes,bytes) userOp) view returns(address)
func (_KingdomDAppControl *KingdomDAppControlCallerSession) GetBidFormat(userOp UserOperation) (common.Address, error) {
	return _KingdomDAppControl.Contract.GetBidFormat(&_KingdomDAppControl.CallOpts, userOp)
}

// GetBidValue is a free data retrieval call binding the contract method 0x6d25fc9a.
//
// Solidity: function getBidValue((address,address,uint256,uint256,uint256,uint256,address,address,bytes32,address,uint256,bytes,bytes) solverOp) view returns(uint256)
func (_KingdomDAppControl *KingdomDAppControlCaller) GetBidValue(opts *bind.CallOpts, solverOp SolverOperation) (*big.Int, error) {
	var out []interface{}
	err := _KingdomDAppControl.contract.Call(opts, &out, "getBidValue", solverOp)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBidValue is a free data retrieval call binding the contract method 0x6d25fc9a.
//
// Solidity: function getBidValue((address,address,uint256,uint256,uint256,uint256,address,address,bytes32,address,uint256,bytes,bytes) solverOp) view returns(uint256)
func (_KingdomDAppControl *KingdomDAppControlSession) GetBidValue(solverOp SolverOperation) (*big.Int, error) {
	return _KingdomDAppControl.Contract.GetBidValue(&_KingdomDAppControl.CallOpts, solverOp)
}

// GetBidValue is a free data retrieval call binding the contract method 0x6d25fc9a.
//
// Solidity: function getBidValue((address,address,uint256,uint256,uint256,uint256,address,address,bytes32,address,uint256,bytes,bytes) solverOp) view returns(uint256)
func (_KingdomDAppControl *KingdomDAppControlCallerSession) GetBidValue(solverOp SolverOperation) (*big.Int, error) {
	return _KingdomDAppControl.Contract.GetBidValue(&_KingdomDAppControl.CallOpts, solverOp)
}

// GetCallConfig is a free data retrieval call binding the contract method 0x6e1ccfc4.
//
// Solidity: function getCallConfig() view returns((bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool))
func (_KingdomDAppControl *KingdomDAppControlCaller) GetCallConfig(opts *bind.CallOpts) (CallConfig, error) {
	var out []interface{}
	err := _KingdomDAppControl.contract.Call(opts, &out, "getCallConfig")

	if err != nil {
		return *new(CallConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(CallConfig)).(*CallConfig)

	return out0, err

}

// GetCallConfig is a free data retrieval call binding the contract method 0x6e1ccfc4.
//
// Solidity: function getCallConfig() view returns((bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool))
func (_KingdomDAppControl *KingdomDAppControlSession) GetCallConfig() (CallConfig, error) {
	return _KingdomDAppControl.Contract.GetCallConfig(&_KingdomDAppControl.CallOpts)
}

// GetCallConfig is a free data retrieval call binding the contract method 0x6e1ccfc4.
//
// Solidity: function getCallConfig() view returns((bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool))
func (_KingdomDAppControl *KingdomDAppControlCallerSession) GetCallConfig() (CallConfig, error) {
	return _KingdomDAppControl.Contract.GetCallConfig(&_KingdomDAppControl.CallOpts)
}

// GetDAppConfig is a free data retrieval call binding the contract method 0x44912b6e.
//
// Solidity: function getDAppConfig((address,address,uint256,uint256,uint256,uint256,uint256,address,address,uint32,address,bytes,bytes) userOp) view returns((address,uint32,address,uint32) dConfig)
func (_KingdomDAppControl *KingdomDAppControlCaller) GetDAppConfig(opts *bind.CallOpts, userOp UserOperation) (DAppConfig, error) {
	var out []interface{}
	err := _KingdomDAppControl.contract.Call(opts, &out, "getDAppConfig", userOp)

	if err != nil {
		return *new(DAppConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(DAppConfig)).(*DAppConfig)

	return out0, err

}

// GetDAppConfig is a free data retrieval call binding the contract method 0x44912b6e.
//
// Solidity: function getDAppConfig((address,address,uint256,uint256,uint256,uint256,uint256,address,address,uint32,address,bytes,bytes) userOp) view returns((address,uint32,address,uint32) dConfig)
func (_KingdomDAppControl *KingdomDAppControlSession) GetDAppConfig(userOp UserOperation) (DAppConfig, error) {
	return _KingdomDAppControl.Contract.GetDAppConfig(&_KingdomDAppControl.CallOpts, userOp)
}

// GetDAppConfig is a free data retrieval call binding the contract method 0x44912b6e.
//
// Solidity: function getDAppConfig((address,address,uint256,uint256,uint256,uint256,uint256,address,address,uint32,address,bytes,bytes) userOp) view returns((address,uint32,address,uint32) dConfig)
func (_KingdomDAppControl *KingdomDAppControlCallerSession) GetDAppConfig(userOp UserOperation) (DAppConfig, error) {
	return _KingdomDAppControl.Contract.GetDAppConfig(&_KingdomDAppControl.CallOpts, userOp)
}

// GetDAppSignatory is a free data retrieval call binding the contract method 0xee418488.
//
// Solidity: function getDAppSignatory() view returns(address)
func (_KingdomDAppControl *KingdomDAppControlCaller) GetDAppSignatory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _KingdomDAppControl.contract.Call(opts, &out, "getDAppSignatory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetDAppSignatory is a free data retrieval call binding the contract method 0xee418488.
//
// Solidity: function getDAppSignatory() view returns(address)
func (_KingdomDAppControl *KingdomDAppControlSession) GetDAppSignatory() (common.Address, error) {
	return _KingdomDAppControl.Contract.GetDAppSignatory(&_KingdomDAppControl.CallOpts)
}

// GetDAppSignatory is a free data retrieval call binding the contract method 0xee418488.
//
// Solidity: function getDAppSignatory() view returns(address)
func (_KingdomDAppControl *KingdomDAppControlCallerSession) GetDAppSignatory() (common.Address, error) {
	return _KingdomDAppControl.Contract.GetDAppSignatory(&_KingdomDAppControl.CallOpts)
}

// GetFactory is a free data retrieval call binding the contract method 0x88cc58e4.
//
// Solidity: function getFactory() view returns(address)
func (_KingdomDAppControl *KingdomDAppControlCaller) GetFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _KingdomDAppControl.contract.Call(opts, &out, "getFactory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetFactory is a free data retrieval call binding the contract method 0x88cc58e4.
//
// Solidity: function getFactory() view returns(address)
func (_KingdomDAppControl *KingdomDAppControlSession) GetFactory() (common.Address, error) {
	return _KingdomDAppControl.Contract.GetFactory(&_KingdomDAppControl.CallOpts)
}

// GetFactory is a free data retrieval call binding the contract method 0x88cc58e4.
//
// Solidity: function getFactory() view returns(address)
func (_KingdomDAppControl *KingdomDAppControlCallerSession) GetFactory() (common.Address, error) {
	return _KingdomDAppControl.Contract.GetFactory(&_KingdomDAppControl.CallOpts)
}

// GetPayoutData is a free data retrieval call binding the contract method 0xa6013654.
//
// Solidity: function getPayoutData() view returns(uint256, address, address)
func (_KingdomDAppControl *KingdomDAppControlCaller) GetPayoutData(opts *bind.CallOpts) (*big.Int, common.Address, common.Address, error) {
	var out []interface{}
	err := _KingdomDAppControl.contract.Call(opts, &out, "getPayoutData")

	if err != nil {
		return *new(*big.Int), *new(common.Address), *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	out2 := *abi.ConvertType(out[2], new(common.Address)).(*common.Address)

	return out0, out1, out2, err

}

// GetPayoutData is a free data retrieval call binding the contract method 0xa6013654.
//
// Solidity: function getPayoutData() view returns(uint256, address, address)
func (_KingdomDAppControl *KingdomDAppControlSession) GetPayoutData() (*big.Int, common.Address, common.Address, error) {
	return _KingdomDAppControl.Contract.GetPayoutData(&_KingdomDAppControl.CallOpts)
}

// GetPayoutData is a free data retrieval call binding the contract method 0xa6013654.
//
// Solidity: function getPayoutData() view returns(uint256, address, address)
func (_KingdomDAppControl *KingdomDAppControlCallerSession) GetPayoutData() (*big.Int, common.Address, common.Address, error) {
	return _KingdomDAppControl.Contract.GetPayoutData(&_KingdomDAppControl.CallOpts)
}

// GetSolverGasLimit is a free data retrieval call binding the contract method 0x99218be5.
//
// Solidity: function getSolverGasLimit() view returns(uint32)
func (_KingdomDAppControl *KingdomDAppControlCaller) GetSolverGasLimit(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _KingdomDAppControl.contract.Call(opts, &out, "getSolverGasLimit")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetSolverGasLimit is a free data retrieval call binding the contract method 0x99218be5.
//
// Solidity: function getSolverGasLimit() view returns(uint32)
func (_KingdomDAppControl *KingdomDAppControlSession) GetSolverGasLimit() (uint32, error) {
	return _KingdomDAppControl.Contract.GetSolverGasLimit(&_KingdomDAppControl.CallOpts)
}

// GetSolverGasLimit is a free data retrieval call binding the contract method 0x99218be5.
//
// Solidity: function getSolverGasLimit() view returns(uint32)
func (_KingdomDAppControl *KingdomDAppControlCallerSession) GetSolverGasLimit() (uint32, error) {
	return _KingdomDAppControl.Contract.GetSolverGasLimit(&_KingdomDAppControl.CallOpts)
}

// GovPayoutAddr is a free data retrieval call binding the contract method 0xedb51dae.
//
// Solidity: function govPayoutAddr() view returns(address)
func (_KingdomDAppControl *KingdomDAppControlCaller) GovPayoutAddr(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _KingdomDAppControl.contract.Call(opts, &out, "govPayoutAddr")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GovPayoutAddr is a free data retrieval call binding the contract method 0xedb51dae.
//
// Solidity: function govPayoutAddr() view returns(address)
func (_KingdomDAppControl *KingdomDAppControlSession) GovPayoutAddr() (common.Address, error) {
	return _KingdomDAppControl.Contract.GovPayoutAddr(&_KingdomDAppControl.CallOpts)
}

// GovPayoutAddr is a free data retrieval call binding the contract method 0xedb51dae.
//
// Solidity: function govPayoutAddr() view returns(address)
func (_KingdomDAppControl *KingdomDAppControlCallerSession) GovPayoutAddr() (common.Address, error) {
	return _KingdomDAppControl.Contract.GovPayoutAddr(&_KingdomDAppControl.CallOpts)
}

// GovSharePercent is a free data retrieval call binding the contract method 0x414a25f3.
//
// Solidity: function govSharePercent() view returns(uint256)
func (_KingdomDAppControl *KingdomDAppControlCaller) GovSharePercent(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _KingdomDAppControl.contract.Call(opts, &out, "govSharePercent")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GovSharePercent is a free data retrieval call binding the contract method 0x414a25f3.
//
// Solidity: function govSharePercent() view returns(uint256)
func (_KingdomDAppControl *KingdomDAppControlSession) GovSharePercent() (*big.Int, error) {
	return _KingdomDAppControl.Contract.GovSharePercent(&_KingdomDAppControl.CallOpts)
}

// GovSharePercent is a free data retrieval call binding the contract method 0x414a25f3.
//
// Solidity: function govSharePercent() view returns(uint256)
func (_KingdomDAppControl *KingdomDAppControlCallerSession) GovSharePercent() (*big.Int, error) {
	return _KingdomDAppControl.Contract.GovSharePercent(&_KingdomDAppControl.CallOpts)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_KingdomDAppControl *KingdomDAppControlCaller) Governance(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _KingdomDAppControl.contract.Call(opts, &out, "governance")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_KingdomDAppControl *KingdomDAppControlSession) Governance() (common.Address, error) {
	return _KingdomDAppControl.Contract.Governance(&_KingdomDAppControl.CallOpts)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_KingdomDAppControl *KingdomDAppControlCallerSession) Governance() (common.Address, error) {
	return _KingdomDAppControl.Contract.Governance(&_KingdomDAppControl.CallOpts)
}

// PendingGovernance is a free data retrieval call binding the contract method 0xf39c38a0.
//
// Solidity: function pendingGovernance() view returns(address)
func (_KingdomDAppControl *KingdomDAppControlCaller) PendingGovernance(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _KingdomDAppControl.contract.Call(opts, &out, "pendingGovernance")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingGovernance is a free data retrieval call binding the contract method 0xf39c38a0.
//
// Solidity: function pendingGovernance() view returns(address)
func (_KingdomDAppControl *KingdomDAppControlSession) PendingGovernance() (common.Address, error) {
	return _KingdomDAppControl.Contract.PendingGovernance(&_KingdomDAppControl.CallOpts)
}

// PendingGovernance is a free data retrieval call binding the contract method 0xf39c38a0.
//
// Solidity: function pendingGovernance() view returns(address)
func (_KingdomDAppControl *KingdomDAppControlCallerSession) PendingGovernance() (common.Address, error) {
	return _KingdomDAppControl.Contract.PendingGovernance(&_KingdomDAppControl.CallOpts)
}

// RequireSequentialDAppNonces is a free data retrieval call binding the contract method 0x72d91684.
//
// Solidity: function requireSequentialDAppNonces() view returns(bool isSequential)
func (_KingdomDAppControl *KingdomDAppControlCaller) RequireSequentialDAppNonces(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _KingdomDAppControl.contract.Call(opts, &out, "requireSequentialDAppNonces")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// RequireSequentialDAppNonces is a free data retrieval call binding the contract method 0x72d91684.
//
// Solidity: function requireSequentialDAppNonces() view returns(bool isSequential)
func (_KingdomDAppControl *KingdomDAppControlSession) RequireSequentialDAppNonces() (bool, error) {
	return _KingdomDAppControl.Contract.RequireSequentialDAppNonces(&_KingdomDAppControl.CallOpts)
}

// RequireSequentialDAppNonces is a free data retrieval call binding the contract method 0x72d91684.
//
// Solidity: function requireSequentialDAppNonces() view returns(bool isSequential)
func (_KingdomDAppControl *KingdomDAppControlCallerSession) RequireSequentialDAppNonces() (bool, error) {
	return _KingdomDAppControl.Contract.RequireSequentialDAppNonces(&_KingdomDAppControl.CallOpts)
}

// RequireSequentialUserNonces is a free data retrieval call binding the contract method 0xe2c0c30f.
//
// Solidity: function requireSequentialUserNonces() view returns(bool isSequential)
func (_KingdomDAppControl *KingdomDAppControlCaller) RequireSequentialUserNonces(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _KingdomDAppControl.contract.Call(opts, &out, "requireSequentialUserNonces")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// RequireSequentialUserNonces is a free data retrieval call binding the contract method 0xe2c0c30f.
//
// Solidity: function requireSequentialUserNonces() view returns(bool isSequential)
func (_KingdomDAppControl *KingdomDAppControlSession) RequireSequentialUserNonces() (bool, error) {
	return _KingdomDAppControl.Contract.RequireSequentialUserNonces(&_KingdomDAppControl.CallOpts)
}

// RequireSequentialUserNonces is a free data retrieval call binding the contract method 0xe2c0c30f.
//
// Solidity: function requireSequentialUserNonces() view returns(bool isSequential)
func (_KingdomDAppControl *KingdomDAppControlCallerSession) RequireSequentialUserNonces() (bool, error) {
	return _KingdomDAppControl.Contract.RequireSequentialUserNonces(&_KingdomDAppControl.CallOpts)
}

// SolverGasLimit is a free data retrieval call binding the contract method 0xd787510d.
//
// Solidity: function solverGasLimit() view returns(uint32)
func (_KingdomDAppControl *KingdomDAppControlCaller) SolverGasLimit(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _KingdomDAppControl.contract.Call(opts, &out, "solverGasLimit")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// SolverGasLimit is a free data retrieval call binding the contract method 0xd787510d.
//
// Solidity: function solverGasLimit() view returns(uint32)
func (_KingdomDAppControl *KingdomDAppControlSession) SolverGasLimit() (uint32, error) {
	return _KingdomDAppControl.Contract.SolverGasLimit(&_KingdomDAppControl.CallOpts)
}

// SolverGasLimit is a free data retrieval call binding the contract method 0xd787510d.
//
// Solidity: function solverGasLimit() view returns(uint32)
func (_KingdomDAppControl *KingdomDAppControlCallerSession) SolverGasLimit() (uint32, error) {
	return _KingdomDAppControl.Contract.SolverGasLimit(&_KingdomDAppControl.CallOpts)
}

// UserDelegated is a free data retrieval call binding the contract method 0x1e151167.
//
// Solidity: function userDelegated() view returns(bool delegated)
func (_KingdomDAppControl *KingdomDAppControlCaller) UserDelegated(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _KingdomDAppControl.contract.Call(opts, &out, "userDelegated")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// UserDelegated is a free data retrieval call binding the contract method 0x1e151167.
//
// Solidity: function userDelegated() view returns(bool delegated)
func (_KingdomDAppControl *KingdomDAppControlSession) UserDelegated() (bool, error) {
	return _KingdomDAppControl.Contract.UserDelegated(&_KingdomDAppControl.CallOpts)
}

// UserDelegated is a free data retrieval call binding the contract method 0x1e151167.
//
// Solidity: function userDelegated() view returns(bool delegated)
func (_KingdomDAppControl *KingdomDAppControlCallerSession) UserDelegated() (bool, error) {
	return _KingdomDAppControl.Contract.UserDelegated(&_KingdomDAppControl.CallOpts)
}

// AcceptGovernance is a paid mutator transaction binding the contract method 0x238efcbc.
//
// Solidity: function acceptGovernance() returns()
func (_KingdomDAppControl *KingdomDAppControlTransactor) AcceptGovernance(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KingdomDAppControl.contract.Transact(opts, "acceptGovernance")
}

// AcceptGovernance is a paid mutator transaction binding the contract method 0x238efcbc.
//
// Solidity: function acceptGovernance() returns()
func (_KingdomDAppControl *KingdomDAppControlSession) AcceptGovernance() (*types.Transaction, error) {
	return _KingdomDAppControl.Contract.AcceptGovernance(&_KingdomDAppControl.TransactOpts)
}

// AcceptGovernance is a paid mutator transaction binding the contract method 0x238efcbc.
//
// Solidity: function acceptGovernance() returns()
func (_KingdomDAppControl *KingdomDAppControlTransactorSession) AcceptGovernance() (*types.Transaction, error) {
	return _KingdomDAppControl.Contract.AcceptGovernance(&_KingdomDAppControl.TransactOpts)
}

// AddAuthorizedUser is a paid mutator transaction binding the contract method 0x177d2a74.
//
// Solidity: function addAuthorizedUser(address _user) returns()
func (_KingdomDAppControl *KingdomDAppControlTransactor) AddAuthorizedUser(opts *bind.TransactOpts, _user common.Address) (*types.Transaction, error) {
	return _KingdomDAppControl.contract.Transact(opts, "addAuthorizedUser", _user)
}

// AddAuthorizedUser is a paid mutator transaction binding the contract method 0x177d2a74.
//
// Solidity: function addAuthorizedUser(address _user) returns()
func (_KingdomDAppControl *KingdomDAppControlSession) AddAuthorizedUser(_user common.Address) (*types.Transaction, error) {
	return _KingdomDAppControl.Contract.AddAuthorizedUser(&_KingdomDAppControl.TransactOpts, _user)
}

// AddAuthorizedUser is a paid mutator transaction binding the contract method 0x177d2a74.
//
// Solidity: function addAuthorizedUser(address _user) returns()
func (_KingdomDAppControl *KingdomDAppControlTransactorSession) AddAuthorizedUser(_user common.Address) (*types.Transaction, error) {
	return _KingdomDAppControl.Contract.AddAuthorizedUser(&_KingdomDAppControl.TransactOpts, _user)
}

// AllocateValueCall is a paid mutator transaction binding the contract method 0x2f5e0d16.
//
// Solidity: function allocateValueCall(address bidToken, uint256 bidAmount, bytes data) returns()
func (_KingdomDAppControl *KingdomDAppControlTransactor) AllocateValueCall(opts *bind.TransactOpts, bidToken common.Address, bidAmount *big.Int, data []byte) (*types.Transaction, error) {
	return _KingdomDAppControl.contract.Transact(opts, "allocateValueCall", bidToken, bidAmount, data)
}

// AllocateValueCall is a paid mutator transaction binding the contract method 0x2f5e0d16.
//
// Solidity: function allocateValueCall(address bidToken, uint256 bidAmount, bytes data) returns()
func (_KingdomDAppControl *KingdomDAppControlSession) AllocateValueCall(bidToken common.Address, bidAmount *big.Int, data []byte) (*types.Transaction, error) {
	return _KingdomDAppControl.Contract.AllocateValueCall(&_KingdomDAppControl.TransactOpts, bidToken, bidAmount, data)
}

// AllocateValueCall is a paid mutator transaction binding the contract method 0x2f5e0d16.
//
// Solidity: function allocateValueCall(address bidToken, uint256 bidAmount, bytes data) returns()
func (_KingdomDAppControl *KingdomDAppControlTransactorSession) AllocateValueCall(bidToken common.Address, bidAmount *big.Int, data []byte) (*types.Transaction, error) {
	return _KingdomDAppControl.Contract.AllocateValueCall(&_KingdomDAppControl.TransactOpts, bidToken, bidAmount, data)
}

// FeeSetterIntent is a paid mutator transaction binding the contract method 0x19f147f4.
//
// Solidity: function feeSetterIntent((address,address,uint24,uint256,uint256) _feeSetterIntent) returns((address,address,uint24,uint256,uint256))
func (_KingdomDAppControl *KingdomDAppControlTransactor) FeeSetterIntent(opts *bind.TransactOpts, _feeSetterIntent FeeSetterIntent) (*types.Transaction, error) {
	return _KingdomDAppControl.contract.Transact(opts, "feeSetterIntent", _feeSetterIntent)
}

// FeeSetterIntent is a paid mutator transaction binding the contract method 0x19f147f4.
//
// Solidity: function feeSetterIntent((address,address,uint24,uint256,uint256) _feeSetterIntent) returns((address,address,uint24,uint256,uint256))
func (_KingdomDAppControl *KingdomDAppControlSession) FeeSetterIntent(_feeSetterIntent FeeSetterIntent) (*types.Transaction, error) {
	return _KingdomDAppControl.Contract.FeeSetterIntent(&_KingdomDAppControl.TransactOpts, _feeSetterIntent)
}

// FeeSetterIntent is a paid mutator transaction binding the contract method 0x19f147f4.
//
// Solidity: function feeSetterIntent((address,address,uint24,uint256,uint256) _feeSetterIntent) returns((address,address,uint24,uint256,uint256))
func (_KingdomDAppControl *KingdomDAppControlTransactorSession) FeeSetterIntent(_feeSetterIntent FeeSetterIntent) (*types.Transaction, error) {
	return _KingdomDAppControl.Contract.FeeSetterIntent(&_KingdomDAppControl.TransactOpts, _feeSetterIntent)
}

// PostOpsCall is a paid mutator transaction binding the contract method 0x836a611b.
//
// Solidity: function postOpsCall(bool solved, bytes data) payable returns()
func (_KingdomDAppControl *KingdomDAppControlTransactor) PostOpsCall(opts *bind.TransactOpts, solved bool, data []byte) (*types.Transaction, error) {
	return _KingdomDAppControl.contract.Transact(opts, "postOpsCall", solved, data)
}

// PostOpsCall is a paid mutator transaction binding the contract method 0x836a611b.
//
// Solidity: function postOpsCall(bool solved, bytes data) payable returns()
func (_KingdomDAppControl *KingdomDAppControlSession) PostOpsCall(solved bool, data []byte) (*types.Transaction, error) {
	return _KingdomDAppControl.Contract.PostOpsCall(&_KingdomDAppControl.TransactOpts, solved, data)
}

// PostOpsCall is a paid mutator transaction binding the contract method 0x836a611b.
//
// Solidity: function postOpsCall(bool solved, bytes data) payable returns()
func (_KingdomDAppControl *KingdomDAppControlTransactorSession) PostOpsCall(solved bool, data []byte) (*types.Transaction, error) {
	return _KingdomDAppControl.Contract.PostOpsCall(&_KingdomDAppControl.TransactOpts, solved, data)
}

// PostSolverCall is a paid mutator transaction binding the contract method 0x6d4d6b2e.
//
// Solidity: function postSolverCall((address,address,uint256,uint256,uint256,uint256,address,address,bytes32,address,uint256,bytes,bytes) solverOp, bytes returnData) payable returns()
func (_KingdomDAppControl *KingdomDAppControlTransactor) PostSolverCall(opts *bind.TransactOpts, solverOp SolverOperation, returnData []byte) (*types.Transaction, error) {
	return _KingdomDAppControl.contract.Transact(opts, "postSolverCall", solverOp, returnData)
}

// PostSolverCall is a paid mutator transaction binding the contract method 0x6d4d6b2e.
//
// Solidity: function postSolverCall((address,address,uint256,uint256,uint256,uint256,address,address,bytes32,address,uint256,bytes,bytes) solverOp, bytes returnData) payable returns()
func (_KingdomDAppControl *KingdomDAppControlSession) PostSolverCall(solverOp SolverOperation, returnData []byte) (*types.Transaction, error) {
	return _KingdomDAppControl.Contract.PostSolverCall(&_KingdomDAppControl.TransactOpts, solverOp, returnData)
}

// PostSolverCall is a paid mutator transaction binding the contract method 0x6d4d6b2e.
//
// Solidity: function postSolverCall((address,address,uint256,uint256,uint256,uint256,address,address,bytes32,address,uint256,bytes,bytes) solverOp, bytes returnData) payable returns()
func (_KingdomDAppControl *KingdomDAppControlTransactorSession) PostSolverCall(solverOp SolverOperation, returnData []byte) (*types.Transaction, error) {
	return _KingdomDAppControl.Contract.PostSolverCall(&_KingdomDAppControl.TransactOpts, solverOp, returnData)
}

// PreOpsCall is a paid mutator transaction binding the contract method 0x77bceb1b.
//
// Solidity: function preOpsCall((address,address,uint256,uint256,uint256,uint256,uint256,address,address,uint32,address,bytes,bytes) userOp) payable returns(bytes)
func (_KingdomDAppControl *KingdomDAppControlTransactor) PreOpsCall(opts *bind.TransactOpts, userOp UserOperation) (*types.Transaction, error) {
	return _KingdomDAppControl.contract.Transact(opts, "preOpsCall", userOp)
}

// PreOpsCall is a paid mutator transaction binding the contract method 0x77bceb1b.
//
// Solidity: function preOpsCall((address,address,uint256,uint256,uint256,uint256,uint256,address,address,uint32,address,bytes,bytes) userOp) payable returns(bytes)
func (_KingdomDAppControl *KingdomDAppControlSession) PreOpsCall(userOp UserOperation) (*types.Transaction, error) {
	return _KingdomDAppControl.Contract.PreOpsCall(&_KingdomDAppControl.TransactOpts, userOp)
}

// PreOpsCall is a paid mutator transaction binding the contract method 0x77bceb1b.
//
// Solidity: function preOpsCall((address,address,uint256,uint256,uint256,uint256,uint256,address,address,uint32,address,bytes,bytes) userOp) payable returns(bytes)
func (_KingdomDAppControl *KingdomDAppControlTransactorSession) PreOpsCall(userOp UserOperation) (*types.Transaction, error) {
	return _KingdomDAppControl.Contract.PreOpsCall(&_KingdomDAppControl.TransactOpts, userOp)
}

// PreSolverCall is a paid mutator transaction binding the contract method 0x283ee1cf.
//
// Solidity: function preSolverCall((address,address,uint256,uint256,uint256,uint256,address,address,bytes32,address,uint256,bytes,bytes) solverOp, bytes returnData) payable returns()
func (_KingdomDAppControl *KingdomDAppControlTransactor) PreSolverCall(opts *bind.TransactOpts, solverOp SolverOperation, returnData []byte) (*types.Transaction, error) {
	return _KingdomDAppControl.contract.Transact(opts, "preSolverCall", solverOp, returnData)
}

// PreSolverCall is a paid mutator transaction binding the contract method 0x283ee1cf.
//
// Solidity: function preSolverCall((address,address,uint256,uint256,uint256,uint256,address,address,bytes32,address,uint256,bytes,bytes) solverOp, bytes returnData) payable returns()
func (_KingdomDAppControl *KingdomDAppControlSession) PreSolverCall(solverOp SolverOperation, returnData []byte) (*types.Transaction, error) {
	return _KingdomDAppControl.Contract.PreSolverCall(&_KingdomDAppControl.TransactOpts, solverOp, returnData)
}

// PreSolverCall is a paid mutator transaction binding the contract method 0x283ee1cf.
//
// Solidity: function preSolverCall((address,address,uint256,uint256,uint256,uint256,address,address,bytes32,address,uint256,bytes,bytes) solverOp, bytes returnData) payable returns()
func (_KingdomDAppControl *KingdomDAppControlTransactorSession) PreSolverCall(solverOp SolverOperation, returnData []byte) (*types.Transaction, error) {
	return _KingdomDAppControl.Contract.PreSolverCall(&_KingdomDAppControl.TransactOpts, solverOp, returnData)
}

// RemoveAuthorizedUser is a paid mutator transaction binding the contract method 0x89fabc80.
//
// Solidity: function removeAuthorizedUser(address _user) returns()
func (_KingdomDAppControl *KingdomDAppControlTransactor) RemoveAuthorizedUser(opts *bind.TransactOpts, _user common.Address) (*types.Transaction, error) {
	return _KingdomDAppControl.contract.Transact(opts, "removeAuthorizedUser", _user)
}

// RemoveAuthorizedUser is a paid mutator transaction binding the contract method 0x89fabc80.
//
// Solidity: function removeAuthorizedUser(address _user) returns()
func (_KingdomDAppControl *KingdomDAppControlSession) RemoveAuthorizedUser(_user common.Address) (*types.Transaction, error) {
	return _KingdomDAppControl.Contract.RemoveAuthorizedUser(&_KingdomDAppControl.TransactOpts, _user)
}

// RemoveAuthorizedUser is a paid mutator transaction binding the contract method 0x89fabc80.
//
// Solidity: function removeAuthorizedUser(address _user) returns()
func (_KingdomDAppControl *KingdomDAppControlTransactorSession) RemoveAuthorizedUser(_user common.Address) (*types.Transaction, error) {
	return _KingdomDAppControl.Contract.RemoveAuthorizedUser(&_KingdomDAppControl.TransactOpts, _user)
}

// ResetFee is a paid mutator transaction binding the contract method 0x56e2793d.
//
// Solidity: function resetFee(address _poolAddress, uint24 _fee) returns()
func (_KingdomDAppControl *KingdomDAppControlTransactor) ResetFee(opts *bind.TransactOpts, _poolAddress common.Address, _fee *big.Int) (*types.Transaction, error) {
	return _KingdomDAppControl.contract.Transact(opts, "resetFee", _poolAddress, _fee)
}

// ResetFee is a paid mutator transaction binding the contract method 0x56e2793d.
//
// Solidity: function resetFee(address _poolAddress, uint24 _fee) returns()
func (_KingdomDAppControl *KingdomDAppControlSession) ResetFee(_poolAddress common.Address, _fee *big.Int) (*types.Transaction, error) {
	return _KingdomDAppControl.Contract.ResetFee(&_KingdomDAppControl.TransactOpts, _poolAddress, _fee)
}

// ResetFee is a paid mutator transaction binding the contract method 0x56e2793d.
//
// Solidity: function resetFee(address _poolAddress, uint24 _fee) returns()
func (_KingdomDAppControl *KingdomDAppControlTransactorSession) ResetFee(_poolAddress common.Address, _fee *big.Int) (*types.Transaction, error) {
	return _KingdomDAppControl.Contract.ResetFee(&_KingdomDAppControl.TransactOpts, _poolAddress, _fee)
}

// SetFastlanePayoutAddr is a paid mutator transaction binding the contract method 0x2e16db0b.
//
// Solidity: function setFastlanePayoutAddr(address _fastlanePayoutAddr) returns()
func (_KingdomDAppControl *KingdomDAppControlTransactor) SetFastlanePayoutAddr(opts *bind.TransactOpts, _fastlanePayoutAddr common.Address) (*types.Transaction, error) {
	return _KingdomDAppControl.contract.Transact(opts, "setFastlanePayoutAddr", _fastlanePayoutAddr)
}

// SetFastlanePayoutAddr is a paid mutator transaction binding the contract method 0x2e16db0b.
//
// Solidity: function setFastlanePayoutAddr(address _fastlanePayoutAddr) returns()
func (_KingdomDAppControl *KingdomDAppControlSession) SetFastlanePayoutAddr(_fastlanePayoutAddr common.Address) (*types.Transaction, error) {
	return _KingdomDAppControl.Contract.SetFastlanePayoutAddr(&_KingdomDAppControl.TransactOpts, _fastlanePayoutAddr)
}

// SetFastlanePayoutAddr is a paid mutator transaction binding the contract method 0x2e16db0b.
//
// Solidity: function setFastlanePayoutAddr(address _fastlanePayoutAddr) returns()
func (_KingdomDAppControl *KingdomDAppControlTransactorSession) SetFastlanePayoutAddr(_fastlanePayoutAddr common.Address) (*types.Transaction, error) {
	return _KingdomDAppControl.Contract.SetFastlanePayoutAddr(&_KingdomDAppControl.TransactOpts, _fastlanePayoutAddr)
}

// SetFeeSetter is a paid mutator transaction binding the contract method 0xb19805af.
//
// Solidity: function setFeeSetter(address _feeSetter) returns()
func (_KingdomDAppControl *KingdomDAppControlTransactor) SetFeeSetter(opts *bind.TransactOpts, _feeSetter common.Address) (*types.Transaction, error) {
	return _KingdomDAppControl.contract.Transact(opts, "setFeeSetter", _feeSetter)
}

// SetFeeSetter is a paid mutator transaction binding the contract method 0xb19805af.
//
// Solidity: function setFeeSetter(address _feeSetter) returns()
func (_KingdomDAppControl *KingdomDAppControlSession) SetFeeSetter(_feeSetter common.Address) (*types.Transaction, error) {
	return _KingdomDAppControl.Contract.SetFeeSetter(&_KingdomDAppControl.TransactOpts, _feeSetter)
}

// SetFeeSetter is a paid mutator transaction binding the contract method 0xb19805af.
//
// Solidity: function setFeeSetter(address _feeSetter) returns()
func (_KingdomDAppControl *KingdomDAppControlTransactorSession) SetFeeSetter(_feeSetter common.Address) (*types.Transaction, error) {
	return _KingdomDAppControl.Contract.SetFeeSetter(&_KingdomDAppControl.TransactOpts, _feeSetter)
}

// SetGovPayoutAddr is a paid mutator transaction binding the contract method 0xf9f9d127.
//
// Solidity: function setGovPayoutAddr(address _govPayoutAddr) returns()
func (_KingdomDAppControl *KingdomDAppControlTransactor) SetGovPayoutAddr(opts *bind.TransactOpts, _govPayoutAddr common.Address) (*types.Transaction, error) {
	return _KingdomDAppControl.contract.Transact(opts, "setGovPayoutAddr", _govPayoutAddr)
}

// SetGovPayoutAddr is a paid mutator transaction binding the contract method 0xf9f9d127.
//
// Solidity: function setGovPayoutAddr(address _govPayoutAddr) returns()
func (_KingdomDAppControl *KingdomDAppControlSession) SetGovPayoutAddr(_govPayoutAddr common.Address) (*types.Transaction, error) {
	return _KingdomDAppControl.Contract.SetGovPayoutAddr(&_KingdomDAppControl.TransactOpts, _govPayoutAddr)
}

// SetGovPayoutAddr is a paid mutator transaction binding the contract method 0xf9f9d127.
//
// Solidity: function setGovPayoutAddr(address _govPayoutAddr) returns()
func (_KingdomDAppControl *KingdomDAppControlTransactorSession) SetGovPayoutAddr(_govPayoutAddr common.Address) (*types.Transaction, error) {
	return _KingdomDAppControl.Contract.SetGovPayoutAddr(&_KingdomDAppControl.TransactOpts, _govPayoutAddr)
}

// SetGovSharePercent is a paid mutator transaction binding the contract method 0xe4030bf0.
//
// Solidity: function setGovSharePercent(uint256 _govSharePercent) returns()
func (_KingdomDAppControl *KingdomDAppControlTransactor) SetGovSharePercent(opts *bind.TransactOpts, _govSharePercent *big.Int) (*types.Transaction, error) {
	return _KingdomDAppControl.contract.Transact(opts, "setGovSharePercent", _govSharePercent)
}

// SetGovSharePercent is a paid mutator transaction binding the contract method 0xe4030bf0.
//
// Solidity: function setGovSharePercent(uint256 _govSharePercent) returns()
func (_KingdomDAppControl *KingdomDAppControlSession) SetGovSharePercent(_govSharePercent *big.Int) (*types.Transaction, error) {
	return _KingdomDAppControl.Contract.SetGovSharePercent(&_KingdomDAppControl.TransactOpts, _govSharePercent)
}

// SetGovSharePercent is a paid mutator transaction binding the contract method 0xe4030bf0.
//
// Solidity: function setGovSharePercent(uint256 _govSharePercent) returns()
func (_KingdomDAppControl *KingdomDAppControlTransactorSession) SetGovSharePercent(_govSharePercent *big.Int) (*types.Transaction, error) {
	return _KingdomDAppControl.Contract.SetGovSharePercent(&_KingdomDAppControl.TransactOpts, _govSharePercent)
}

// SetSolverGasLimit is a paid mutator transaction binding the contract method 0xe2c04d0f.
//
// Solidity: function setSolverGasLimit(uint32 _solverGasLimit) returns()
func (_KingdomDAppControl *KingdomDAppControlTransactor) SetSolverGasLimit(opts *bind.TransactOpts, _solverGasLimit uint32) (*types.Transaction, error) {
	return _KingdomDAppControl.contract.Transact(opts, "setSolverGasLimit", _solverGasLimit)
}

// SetSolverGasLimit is a paid mutator transaction binding the contract method 0xe2c04d0f.
//
// Solidity: function setSolverGasLimit(uint32 _solverGasLimit) returns()
func (_KingdomDAppControl *KingdomDAppControlSession) SetSolverGasLimit(_solverGasLimit uint32) (*types.Transaction, error) {
	return _KingdomDAppControl.Contract.SetSolverGasLimit(&_KingdomDAppControl.TransactOpts, _solverGasLimit)
}

// SetSolverGasLimit is a paid mutator transaction binding the contract method 0xe2c04d0f.
//
// Solidity: function setSolverGasLimit(uint32 _solverGasLimit) returns()
func (_KingdomDAppControl *KingdomDAppControlTransactorSession) SetSolverGasLimit(_solverGasLimit uint32) (*types.Transaction, error) {
	return _KingdomDAppControl.Contract.SetSolverGasLimit(&_KingdomDAppControl.TransactOpts, _solverGasLimit)
}

// TransferGovernance is a paid mutator transaction binding the contract method 0xd38bfff4.
//
// Solidity: function transferGovernance(address newGovernance) returns()
func (_KingdomDAppControl *KingdomDAppControlTransactor) TransferGovernance(opts *bind.TransactOpts, newGovernance common.Address) (*types.Transaction, error) {
	return _KingdomDAppControl.contract.Transact(opts, "transferGovernance", newGovernance)
}

// TransferGovernance is a paid mutator transaction binding the contract method 0xd38bfff4.
//
// Solidity: function transferGovernance(address newGovernance) returns()
func (_KingdomDAppControl *KingdomDAppControlSession) TransferGovernance(newGovernance common.Address) (*types.Transaction, error) {
	return _KingdomDAppControl.Contract.TransferGovernance(&_KingdomDAppControl.TransactOpts, newGovernance)
}

// TransferGovernance is a paid mutator transaction binding the contract method 0xd38bfff4.
//
// Solidity: function transferGovernance(address newGovernance) returns()
func (_KingdomDAppControl *KingdomDAppControlTransactorSession) TransferGovernance(newGovernance common.Address) (*types.Transaction, error) {
	return _KingdomDAppControl.Contract.TransferGovernance(&_KingdomDAppControl.TransactOpts, newGovernance)
}

// KingdomDAppControlFastlanePayoutAddressUpdatedIterator is returned from FilterFastlanePayoutAddressUpdated and is used to iterate over the raw logs and unpacked data for FastlanePayoutAddressUpdated events raised by the KingdomDAppControl contract.
type KingdomDAppControlFastlanePayoutAddressUpdatedIterator struct {
	Event *KingdomDAppControlFastlanePayoutAddressUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *KingdomDAppControlFastlanePayoutAddressUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KingdomDAppControlFastlanePayoutAddressUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(KingdomDAppControlFastlanePayoutAddressUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *KingdomDAppControlFastlanePayoutAddressUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KingdomDAppControlFastlanePayoutAddressUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KingdomDAppControlFastlanePayoutAddressUpdated represents a FastlanePayoutAddressUpdated event raised by the KingdomDAppControl contract.
type KingdomDAppControlFastlanePayoutAddressUpdated struct {
	OldFastlanePayoutAddr common.Address
	NewFastlanePayoutAddr common.Address
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterFastlanePayoutAddressUpdated is a free log retrieval operation binding the contract event 0x4dd1cd506d86060529dd6dab83df9a91d2852ea4f47cec0e340f4bb4f8a48c07.
//
// Solidity: event FastlanePayoutAddressUpdated(address indexed oldFastlanePayoutAddr, address indexed newFastlanePayoutAddr)
func (_KingdomDAppControl *KingdomDAppControlFilterer) FilterFastlanePayoutAddressUpdated(opts *bind.FilterOpts, oldFastlanePayoutAddr []common.Address, newFastlanePayoutAddr []common.Address) (*KingdomDAppControlFastlanePayoutAddressUpdatedIterator, error) {

	var oldFastlanePayoutAddrRule []interface{}
	for _, oldFastlanePayoutAddrItem := range oldFastlanePayoutAddr {
		oldFastlanePayoutAddrRule = append(oldFastlanePayoutAddrRule, oldFastlanePayoutAddrItem)
	}
	var newFastlanePayoutAddrRule []interface{}
	for _, newFastlanePayoutAddrItem := range newFastlanePayoutAddr {
		newFastlanePayoutAddrRule = append(newFastlanePayoutAddrRule, newFastlanePayoutAddrItem)
	}

	logs, sub, err := _KingdomDAppControl.contract.FilterLogs(opts, "FastlanePayoutAddressUpdated", oldFastlanePayoutAddrRule, newFastlanePayoutAddrRule)
	if err != nil {
		return nil, err
	}
	return &KingdomDAppControlFastlanePayoutAddressUpdatedIterator{contract: _KingdomDAppControl.contract, event: "FastlanePayoutAddressUpdated", logs: logs, sub: sub}, nil
}

// WatchFastlanePayoutAddressUpdated is a free log subscription operation binding the contract event 0x4dd1cd506d86060529dd6dab83df9a91d2852ea4f47cec0e340f4bb4f8a48c07.
//
// Solidity: event FastlanePayoutAddressUpdated(address indexed oldFastlanePayoutAddr, address indexed newFastlanePayoutAddr)
func (_KingdomDAppControl *KingdomDAppControlFilterer) WatchFastlanePayoutAddressUpdated(opts *bind.WatchOpts, sink chan<- *KingdomDAppControlFastlanePayoutAddressUpdated, oldFastlanePayoutAddr []common.Address, newFastlanePayoutAddr []common.Address) (event.Subscription, error) {

	var oldFastlanePayoutAddrRule []interface{}
	for _, oldFastlanePayoutAddrItem := range oldFastlanePayoutAddr {
		oldFastlanePayoutAddrRule = append(oldFastlanePayoutAddrRule, oldFastlanePayoutAddrItem)
	}
	var newFastlanePayoutAddrRule []interface{}
	for _, newFastlanePayoutAddrItem := range newFastlanePayoutAddr {
		newFastlanePayoutAddrRule = append(newFastlanePayoutAddrRule, newFastlanePayoutAddrItem)
	}

	logs, sub, err := _KingdomDAppControl.contract.WatchLogs(opts, "FastlanePayoutAddressUpdated", oldFastlanePayoutAddrRule, newFastlanePayoutAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KingdomDAppControlFastlanePayoutAddressUpdated)
				if err := _KingdomDAppControl.contract.UnpackLog(event, "FastlanePayoutAddressUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseFastlanePayoutAddressUpdated is a log parse operation binding the contract event 0x4dd1cd506d86060529dd6dab83df9a91d2852ea4f47cec0e340f4bb4f8a48c07.
//
// Solidity: event FastlanePayoutAddressUpdated(address indexed oldFastlanePayoutAddr, address indexed newFastlanePayoutAddr)
func (_KingdomDAppControl *KingdomDAppControlFilterer) ParseFastlanePayoutAddressUpdated(log types.Log) (*KingdomDAppControlFastlanePayoutAddressUpdated, error) {
	event := new(KingdomDAppControlFastlanePayoutAddressUpdated)
	if err := _KingdomDAppControl.contract.UnpackLog(event, "FastlanePayoutAddressUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KingdomDAppControlGovernancePayoutAddressUpdatedIterator is returned from FilterGovernancePayoutAddressUpdated and is used to iterate over the raw logs and unpacked data for GovernancePayoutAddressUpdated events raised by the KingdomDAppControl contract.
type KingdomDAppControlGovernancePayoutAddressUpdatedIterator struct {
	Event *KingdomDAppControlGovernancePayoutAddressUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *KingdomDAppControlGovernancePayoutAddressUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KingdomDAppControlGovernancePayoutAddressUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(KingdomDAppControlGovernancePayoutAddressUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *KingdomDAppControlGovernancePayoutAddressUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KingdomDAppControlGovernancePayoutAddressUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KingdomDAppControlGovernancePayoutAddressUpdated represents a GovernancePayoutAddressUpdated event raised by the KingdomDAppControl contract.
type KingdomDAppControlGovernancePayoutAddressUpdated struct {
	OldGovPayoutAddr common.Address
	NewGovPayoutAddr common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterGovernancePayoutAddressUpdated is a free log retrieval operation binding the contract event 0x82b0603c5577211b760b9d8a7d9927b8203bbe0a38f10bff438ded510f8faf80.
//
// Solidity: event GovernancePayoutAddressUpdated(address indexed oldGovPayoutAddr, address indexed newGovPayoutAddr)
func (_KingdomDAppControl *KingdomDAppControlFilterer) FilterGovernancePayoutAddressUpdated(opts *bind.FilterOpts, oldGovPayoutAddr []common.Address, newGovPayoutAddr []common.Address) (*KingdomDAppControlGovernancePayoutAddressUpdatedIterator, error) {

	var oldGovPayoutAddrRule []interface{}
	for _, oldGovPayoutAddrItem := range oldGovPayoutAddr {
		oldGovPayoutAddrRule = append(oldGovPayoutAddrRule, oldGovPayoutAddrItem)
	}
	var newGovPayoutAddrRule []interface{}
	for _, newGovPayoutAddrItem := range newGovPayoutAddr {
		newGovPayoutAddrRule = append(newGovPayoutAddrRule, newGovPayoutAddrItem)
	}

	logs, sub, err := _KingdomDAppControl.contract.FilterLogs(opts, "GovernancePayoutAddressUpdated", oldGovPayoutAddrRule, newGovPayoutAddrRule)
	if err != nil {
		return nil, err
	}
	return &KingdomDAppControlGovernancePayoutAddressUpdatedIterator{contract: _KingdomDAppControl.contract, event: "GovernancePayoutAddressUpdated", logs: logs, sub: sub}, nil
}

// WatchGovernancePayoutAddressUpdated is a free log subscription operation binding the contract event 0x82b0603c5577211b760b9d8a7d9927b8203bbe0a38f10bff438ded510f8faf80.
//
// Solidity: event GovernancePayoutAddressUpdated(address indexed oldGovPayoutAddr, address indexed newGovPayoutAddr)
func (_KingdomDAppControl *KingdomDAppControlFilterer) WatchGovernancePayoutAddressUpdated(opts *bind.WatchOpts, sink chan<- *KingdomDAppControlGovernancePayoutAddressUpdated, oldGovPayoutAddr []common.Address, newGovPayoutAddr []common.Address) (event.Subscription, error) {

	var oldGovPayoutAddrRule []interface{}
	for _, oldGovPayoutAddrItem := range oldGovPayoutAddr {
		oldGovPayoutAddrRule = append(oldGovPayoutAddrRule, oldGovPayoutAddrItem)
	}
	var newGovPayoutAddrRule []interface{}
	for _, newGovPayoutAddrItem := range newGovPayoutAddr {
		newGovPayoutAddrRule = append(newGovPayoutAddrRule, newGovPayoutAddrItem)
	}

	logs, sub, err := _KingdomDAppControl.contract.WatchLogs(opts, "GovernancePayoutAddressUpdated", oldGovPayoutAddrRule, newGovPayoutAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KingdomDAppControlGovernancePayoutAddressUpdated)
				if err := _KingdomDAppControl.contract.UnpackLog(event, "GovernancePayoutAddressUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseGovernancePayoutAddressUpdated is a log parse operation binding the contract event 0x82b0603c5577211b760b9d8a7d9927b8203bbe0a38f10bff438ded510f8faf80.
//
// Solidity: event GovernancePayoutAddressUpdated(address indexed oldGovPayoutAddr, address indexed newGovPayoutAddr)
func (_KingdomDAppControl *KingdomDAppControlFilterer) ParseGovernancePayoutAddressUpdated(log types.Log) (*KingdomDAppControlGovernancePayoutAddressUpdated, error) {
	event := new(KingdomDAppControlGovernancePayoutAddressUpdated)
	if err := _KingdomDAppControl.contract.UnpackLog(event, "GovernancePayoutAddressUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KingdomDAppControlGovernanceSharePercentUpdatedIterator is returned from FilterGovernanceSharePercentUpdated and is used to iterate over the raw logs and unpacked data for GovernanceSharePercentUpdated events raised by the KingdomDAppControl contract.
type KingdomDAppControlGovernanceSharePercentUpdatedIterator struct {
	Event *KingdomDAppControlGovernanceSharePercentUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *KingdomDAppControlGovernanceSharePercentUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KingdomDAppControlGovernanceSharePercentUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(KingdomDAppControlGovernanceSharePercentUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *KingdomDAppControlGovernanceSharePercentUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KingdomDAppControlGovernanceSharePercentUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KingdomDAppControlGovernanceSharePercentUpdated represents a GovernanceSharePercentUpdated event raised by the KingdomDAppControl contract.
type KingdomDAppControlGovernanceSharePercentUpdated struct {
	OldGovSharePercent *big.Int
	NewGovSharePercent *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterGovernanceSharePercentUpdated is a free log retrieval operation binding the contract event 0x58a11e6c79f86c7892e8b01697c6d8ff8d8db69f155c88199ab8f5ec9621b53a.
//
// Solidity: event GovernanceSharePercentUpdated(uint256 oldGovSharePercent, uint256 newGovSharePercent)
func (_KingdomDAppControl *KingdomDAppControlFilterer) FilterGovernanceSharePercentUpdated(opts *bind.FilterOpts) (*KingdomDAppControlGovernanceSharePercentUpdatedIterator, error) {

	logs, sub, err := _KingdomDAppControl.contract.FilterLogs(opts, "GovernanceSharePercentUpdated")
	if err != nil {
		return nil, err
	}
	return &KingdomDAppControlGovernanceSharePercentUpdatedIterator{contract: _KingdomDAppControl.contract, event: "GovernanceSharePercentUpdated", logs: logs, sub: sub}, nil
}

// WatchGovernanceSharePercentUpdated is a free log subscription operation binding the contract event 0x58a11e6c79f86c7892e8b01697c6d8ff8d8db69f155c88199ab8f5ec9621b53a.
//
// Solidity: event GovernanceSharePercentUpdated(uint256 oldGovSharePercent, uint256 newGovSharePercent)
func (_KingdomDAppControl *KingdomDAppControlFilterer) WatchGovernanceSharePercentUpdated(opts *bind.WatchOpts, sink chan<- *KingdomDAppControlGovernanceSharePercentUpdated) (event.Subscription, error) {

	logs, sub, err := _KingdomDAppControl.contract.WatchLogs(opts, "GovernanceSharePercentUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KingdomDAppControlGovernanceSharePercentUpdated)
				if err := _KingdomDAppControl.contract.UnpackLog(event, "GovernanceSharePercentUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseGovernanceSharePercentUpdated is a log parse operation binding the contract event 0x58a11e6c79f86c7892e8b01697c6d8ff8d8db69f155c88199ab8f5ec9621b53a.
//
// Solidity: event GovernanceSharePercentUpdated(uint256 oldGovSharePercent, uint256 newGovSharePercent)
func (_KingdomDAppControl *KingdomDAppControlFilterer) ParseGovernanceSharePercentUpdated(log types.Log) (*KingdomDAppControlGovernanceSharePercentUpdated, error) {
	event := new(KingdomDAppControlGovernanceSharePercentUpdated)
	if err := _KingdomDAppControl.contract.UnpackLog(event, "GovernanceSharePercentUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KingdomDAppControlGovernanceTransferStartedIterator is returned from FilterGovernanceTransferStarted and is used to iterate over the raw logs and unpacked data for GovernanceTransferStarted events raised by the KingdomDAppControl contract.
type KingdomDAppControlGovernanceTransferStartedIterator struct {
	Event *KingdomDAppControlGovernanceTransferStarted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *KingdomDAppControlGovernanceTransferStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KingdomDAppControlGovernanceTransferStarted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(KingdomDAppControlGovernanceTransferStarted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *KingdomDAppControlGovernanceTransferStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KingdomDAppControlGovernanceTransferStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KingdomDAppControlGovernanceTransferStarted represents a GovernanceTransferStarted event raised by the KingdomDAppControl contract.
type KingdomDAppControlGovernanceTransferStarted struct {
	PreviousGovernance common.Address
	NewGovernance      common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterGovernanceTransferStarted is a free log retrieval operation binding the contract event 0x1c4bb4e3cab7b72da7eb9f0ae62554dda85dc7fb907c946ad2776095b95ac1ad.
//
// Solidity: event GovernanceTransferStarted(address indexed previousGovernance, address indexed newGovernance)
func (_KingdomDAppControl *KingdomDAppControlFilterer) FilterGovernanceTransferStarted(opts *bind.FilterOpts, previousGovernance []common.Address, newGovernance []common.Address) (*KingdomDAppControlGovernanceTransferStartedIterator, error) {

	var previousGovernanceRule []interface{}
	for _, previousGovernanceItem := range previousGovernance {
		previousGovernanceRule = append(previousGovernanceRule, previousGovernanceItem)
	}
	var newGovernanceRule []interface{}
	for _, newGovernanceItem := range newGovernance {
		newGovernanceRule = append(newGovernanceRule, newGovernanceItem)
	}

	logs, sub, err := _KingdomDAppControl.contract.FilterLogs(opts, "GovernanceTransferStarted", previousGovernanceRule, newGovernanceRule)
	if err != nil {
		return nil, err
	}
	return &KingdomDAppControlGovernanceTransferStartedIterator{contract: _KingdomDAppControl.contract, event: "GovernanceTransferStarted", logs: logs, sub: sub}, nil
}

// WatchGovernanceTransferStarted is a free log subscription operation binding the contract event 0x1c4bb4e3cab7b72da7eb9f0ae62554dda85dc7fb907c946ad2776095b95ac1ad.
//
// Solidity: event GovernanceTransferStarted(address indexed previousGovernance, address indexed newGovernance)
func (_KingdomDAppControl *KingdomDAppControlFilterer) WatchGovernanceTransferStarted(opts *bind.WatchOpts, sink chan<- *KingdomDAppControlGovernanceTransferStarted, previousGovernance []common.Address, newGovernance []common.Address) (event.Subscription, error) {

	var previousGovernanceRule []interface{}
	for _, previousGovernanceItem := range previousGovernance {
		previousGovernanceRule = append(previousGovernanceRule, previousGovernanceItem)
	}
	var newGovernanceRule []interface{}
	for _, newGovernanceItem := range newGovernance {
		newGovernanceRule = append(newGovernanceRule, newGovernanceItem)
	}

	logs, sub, err := _KingdomDAppControl.contract.WatchLogs(opts, "GovernanceTransferStarted", previousGovernanceRule, newGovernanceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KingdomDAppControlGovernanceTransferStarted)
				if err := _KingdomDAppControl.contract.UnpackLog(event, "GovernanceTransferStarted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseGovernanceTransferStarted is a log parse operation binding the contract event 0x1c4bb4e3cab7b72da7eb9f0ae62554dda85dc7fb907c946ad2776095b95ac1ad.
//
// Solidity: event GovernanceTransferStarted(address indexed previousGovernance, address indexed newGovernance)
func (_KingdomDAppControl *KingdomDAppControlFilterer) ParseGovernanceTransferStarted(log types.Log) (*KingdomDAppControlGovernanceTransferStarted, error) {
	event := new(KingdomDAppControlGovernanceTransferStarted)
	if err := _KingdomDAppControl.contract.UnpackLog(event, "GovernanceTransferStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KingdomDAppControlGovernanceTransferredIterator is returned from FilterGovernanceTransferred and is used to iterate over the raw logs and unpacked data for GovernanceTransferred events raised by the KingdomDAppControl contract.
type KingdomDAppControlGovernanceTransferredIterator struct {
	Event *KingdomDAppControlGovernanceTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *KingdomDAppControlGovernanceTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KingdomDAppControlGovernanceTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(KingdomDAppControlGovernanceTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *KingdomDAppControlGovernanceTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KingdomDAppControlGovernanceTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KingdomDAppControlGovernanceTransferred represents a GovernanceTransferred event raised by the KingdomDAppControl contract.
type KingdomDAppControlGovernanceTransferred struct {
	PreviousGovernance common.Address
	NewGovernance      common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterGovernanceTransferred is a free log retrieval operation binding the contract event 0x5f56bee8cffbe9a78652a74a60705edede02af10b0bbb888ca44b79a0d42ce80.
//
// Solidity: event GovernanceTransferred(address indexed previousGovernance, address indexed newGovernance)
func (_KingdomDAppControl *KingdomDAppControlFilterer) FilterGovernanceTransferred(opts *bind.FilterOpts, previousGovernance []common.Address, newGovernance []common.Address) (*KingdomDAppControlGovernanceTransferredIterator, error) {

	var previousGovernanceRule []interface{}
	for _, previousGovernanceItem := range previousGovernance {
		previousGovernanceRule = append(previousGovernanceRule, previousGovernanceItem)
	}
	var newGovernanceRule []interface{}
	for _, newGovernanceItem := range newGovernance {
		newGovernanceRule = append(newGovernanceRule, newGovernanceItem)
	}

	logs, sub, err := _KingdomDAppControl.contract.FilterLogs(opts, "GovernanceTransferred", previousGovernanceRule, newGovernanceRule)
	if err != nil {
		return nil, err
	}
	return &KingdomDAppControlGovernanceTransferredIterator{contract: _KingdomDAppControl.contract, event: "GovernanceTransferred", logs: logs, sub: sub}, nil
}

// WatchGovernanceTransferred is a free log subscription operation binding the contract event 0x5f56bee8cffbe9a78652a74a60705edede02af10b0bbb888ca44b79a0d42ce80.
//
// Solidity: event GovernanceTransferred(address indexed previousGovernance, address indexed newGovernance)
func (_KingdomDAppControl *KingdomDAppControlFilterer) WatchGovernanceTransferred(opts *bind.WatchOpts, sink chan<- *KingdomDAppControlGovernanceTransferred, previousGovernance []common.Address, newGovernance []common.Address) (event.Subscription, error) {

	var previousGovernanceRule []interface{}
	for _, previousGovernanceItem := range previousGovernance {
		previousGovernanceRule = append(previousGovernanceRule, previousGovernanceItem)
	}
	var newGovernanceRule []interface{}
	for _, newGovernanceItem := range newGovernance {
		newGovernanceRule = append(newGovernanceRule, newGovernanceItem)
	}

	logs, sub, err := _KingdomDAppControl.contract.WatchLogs(opts, "GovernanceTransferred", previousGovernanceRule, newGovernanceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KingdomDAppControlGovernanceTransferred)
				if err := _KingdomDAppControl.contract.UnpackLog(event, "GovernanceTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseGovernanceTransferred is a log parse operation binding the contract event 0x5f56bee8cffbe9a78652a74a60705edede02af10b0bbb888ca44b79a0d42ce80.
//
// Solidity: event GovernanceTransferred(address indexed previousGovernance, address indexed newGovernance)
func (_KingdomDAppControl *KingdomDAppControlFilterer) ParseGovernanceTransferred(log types.Log) (*KingdomDAppControlGovernanceTransferred, error) {
	event := new(KingdomDAppControlGovernanceTransferred)
	if err := _KingdomDAppControl.contract.UnpackLog(event, "GovernanceTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
