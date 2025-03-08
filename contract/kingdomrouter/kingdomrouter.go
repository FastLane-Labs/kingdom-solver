// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package kingdomrouter

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

// IUniswapV3RouterExactInputSingleParams is an auto generated low-level Go binding around an user-defined struct.
type IUniswapV3RouterExactInputSingleParams struct {
	TokenIn           common.Address
	TokenOut          common.Address
	Fee               *big.Int
	Recipient         common.Address
	Deadline          *big.Int
	AmountIn          *big.Int
	AmountOutMinimum  *big.Int
	SqrtPriceLimitX96 *big.Int
}

// KingdomRouterMetaData contains all meta data concerning the KingdomRouter contract.
var KingdomRouterMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"exactInputSingle\",\"inputs\":[{\"name\":\"params\",\"type\":\"tuple\",\"internalType\":\"structIUniswapV3Router.ExactInputSingleParams\",\"components\":[{\"name\":\"tokenIn\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenOut\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"fee\",\"type\":\"uint24\",\"internalType\":\"uint24\"},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"deadline\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"amountIn\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"amountOutMinimum\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"sqrtPriceLimitX96\",\"type\":\"uint160\",\"internalType\":\"uint160\"}]}],\"outputs\":[{\"name\":\"amountOut\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"payable\"}]",
}

// KingdomRouterABI is the input ABI used to generate the binding from.
// Deprecated: Use KingdomRouterMetaData.ABI instead.
var KingdomRouterABI = KingdomRouterMetaData.ABI

// KingdomRouter is an auto generated Go binding around an Ethereum contract.
type KingdomRouter struct {
	KingdomRouterCaller     // Read-only binding to the contract
	KingdomRouterTransactor // Write-only binding to the contract
	KingdomRouterFilterer   // Log filterer for contract events
}

// KingdomRouterCaller is an auto generated read-only Go binding around an Ethereum contract.
type KingdomRouterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KingdomRouterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type KingdomRouterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KingdomRouterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type KingdomRouterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KingdomRouterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type KingdomRouterSession struct {
	Contract     *KingdomRouter    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// KingdomRouterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type KingdomRouterCallerSession struct {
	Contract *KingdomRouterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// KingdomRouterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type KingdomRouterTransactorSession struct {
	Contract     *KingdomRouterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// KingdomRouterRaw is an auto generated low-level Go binding around an Ethereum contract.
type KingdomRouterRaw struct {
	Contract *KingdomRouter // Generic contract binding to access the raw methods on
}

// KingdomRouterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type KingdomRouterCallerRaw struct {
	Contract *KingdomRouterCaller // Generic read-only contract binding to access the raw methods on
}

// KingdomRouterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type KingdomRouterTransactorRaw struct {
	Contract *KingdomRouterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewKingdomRouter creates a new instance of KingdomRouter, bound to a specific deployed contract.
func NewKingdomRouter(address common.Address, backend bind.ContractBackend) (*KingdomRouter, error) {
	contract, err := bindKingdomRouter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &KingdomRouter{KingdomRouterCaller: KingdomRouterCaller{contract: contract}, KingdomRouterTransactor: KingdomRouterTransactor{contract: contract}, KingdomRouterFilterer: KingdomRouterFilterer{contract: contract}}, nil
}

// NewKingdomRouterCaller creates a new read-only instance of KingdomRouter, bound to a specific deployed contract.
func NewKingdomRouterCaller(address common.Address, caller bind.ContractCaller) (*KingdomRouterCaller, error) {
	contract, err := bindKingdomRouter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &KingdomRouterCaller{contract: contract}, nil
}

// NewKingdomRouterTransactor creates a new write-only instance of KingdomRouter, bound to a specific deployed contract.
func NewKingdomRouterTransactor(address common.Address, transactor bind.ContractTransactor) (*KingdomRouterTransactor, error) {
	contract, err := bindKingdomRouter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &KingdomRouterTransactor{contract: contract}, nil
}

// NewKingdomRouterFilterer creates a new log filterer instance of KingdomRouter, bound to a specific deployed contract.
func NewKingdomRouterFilterer(address common.Address, filterer bind.ContractFilterer) (*KingdomRouterFilterer, error) {
	contract, err := bindKingdomRouter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &KingdomRouterFilterer{contract: contract}, nil
}

// bindKingdomRouter binds a generic wrapper to an already deployed contract.
func bindKingdomRouter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(KingdomRouterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KingdomRouter *KingdomRouterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _KingdomRouter.Contract.KingdomRouterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KingdomRouter *KingdomRouterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KingdomRouter.Contract.KingdomRouterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KingdomRouter *KingdomRouterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KingdomRouter.Contract.KingdomRouterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KingdomRouter *KingdomRouterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _KingdomRouter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KingdomRouter *KingdomRouterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KingdomRouter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KingdomRouter *KingdomRouterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KingdomRouter.Contract.contract.Transact(opts, method, params...)
}

// ExactInputSingle is a paid mutator transaction binding the contract method 0x414bf389.
//
// Solidity: function exactInputSingle((address,address,uint24,address,uint256,uint256,uint256,uint160) params) payable returns(uint256 amountOut)
func (_KingdomRouter *KingdomRouterTransactor) ExactInputSingle(opts *bind.TransactOpts, params IUniswapV3RouterExactInputSingleParams) (*types.Transaction, error) {
	return _KingdomRouter.contract.Transact(opts, "exactInputSingle", params)
}

// ExactInputSingle is a paid mutator transaction binding the contract method 0x414bf389.
//
// Solidity: function exactInputSingle((address,address,uint24,address,uint256,uint256,uint256,uint160) params) payable returns(uint256 amountOut)
func (_KingdomRouter *KingdomRouterSession) ExactInputSingle(params IUniswapV3RouterExactInputSingleParams) (*types.Transaction, error) {
	return _KingdomRouter.Contract.ExactInputSingle(&_KingdomRouter.TransactOpts, params)
}

// ExactInputSingle is a paid mutator transaction binding the contract method 0x414bf389.
//
// Solidity: function exactInputSingle((address,address,uint24,address,uint256,uint256,uint256,uint160) params) payable returns(uint256 amountOut)
func (_KingdomRouter *KingdomRouterTransactorSession) ExactInputSingle(params IUniswapV3RouterExactInputSingleParams) (*types.Transaction, error) {
	return _KingdomRouter.Contract.ExactInputSingle(&_KingdomRouter.TransactOpts, params)
}
