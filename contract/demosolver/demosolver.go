// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package demosolver

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

// DemoSolverMetaData contains all meta data concerning the DemoSolver contract.
var DemoSolverMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"weth_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"atlas_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"router_\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"WETH_ADDRESS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"atlasSolverCall\",\"inputs\":[{\"name\":\"solverOpFrom\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"executionEnvironment\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"bidToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"bidAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"solverOpData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"forwardedData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"demoSwap\",\"inputs\":[{\"name\":\"pool\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenIn\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenOut\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amountIn\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"originalFee\",\"type\":\"uint24\",\"internalType\":\"uint24\"},{\"name\":\"checkCurrentFee\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"router\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"error\",\"name\":\"InvalidCaller\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidEntry\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SolverCallUnsuccessful\",\"inputs\":[]}]",
}

// DemoSolverABI is the input ABI used to generate the binding from.
// Deprecated: Use DemoSolverMetaData.ABI instead.
var DemoSolverABI = DemoSolverMetaData.ABI

// DemoSolver is an auto generated Go binding around an Ethereum contract.
type DemoSolver struct {
	DemoSolverCaller     // Read-only binding to the contract
	DemoSolverTransactor // Write-only binding to the contract
	DemoSolverFilterer   // Log filterer for contract events
}

// DemoSolverCaller is an auto generated read-only Go binding around an Ethereum contract.
type DemoSolverCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DemoSolverTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DemoSolverTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DemoSolverFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DemoSolverFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DemoSolverSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DemoSolverSession struct {
	Contract     *DemoSolver       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DemoSolverCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DemoSolverCallerSession struct {
	Contract *DemoSolverCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// DemoSolverTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DemoSolverTransactorSession struct {
	Contract     *DemoSolverTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// DemoSolverRaw is an auto generated low-level Go binding around an Ethereum contract.
type DemoSolverRaw struct {
	Contract *DemoSolver // Generic contract binding to access the raw methods on
}

// DemoSolverCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DemoSolverCallerRaw struct {
	Contract *DemoSolverCaller // Generic read-only contract binding to access the raw methods on
}

// DemoSolverTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DemoSolverTransactorRaw struct {
	Contract *DemoSolverTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDemoSolver creates a new instance of DemoSolver, bound to a specific deployed contract.
func NewDemoSolver(address common.Address, backend bind.ContractBackend) (*DemoSolver, error) {
	contract, err := bindDemoSolver(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DemoSolver{DemoSolverCaller: DemoSolverCaller{contract: contract}, DemoSolverTransactor: DemoSolverTransactor{contract: contract}, DemoSolverFilterer: DemoSolverFilterer{contract: contract}}, nil
}

// NewDemoSolverCaller creates a new read-only instance of DemoSolver, bound to a specific deployed contract.
func NewDemoSolverCaller(address common.Address, caller bind.ContractCaller) (*DemoSolverCaller, error) {
	contract, err := bindDemoSolver(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DemoSolverCaller{contract: contract}, nil
}

// NewDemoSolverTransactor creates a new write-only instance of DemoSolver, bound to a specific deployed contract.
func NewDemoSolverTransactor(address common.Address, transactor bind.ContractTransactor) (*DemoSolverTransactor, error) {
	contract, err := bindDemoSolver(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DemoSolverTransactor{contract: contract}, nil
}

// NewDemoSolverFilterer creates a new log filterer instance of DemoSolver, bound to a specific deployed contract.
func NewDemoSolverFilterer(address common.Address, filterer bind.ContractFilterer) (*DemoSolverFilterer, error) {
	contract, err := bindDemoSolver(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DemoSolverFilterer{contract: contract}, nil
}

// bindDemoSolver binds a generic wrapper to an already deployed contract.
func bindDemoSolver(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DemoSolverABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DemoSolver *DemoSolverRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DemoSolver.Contract.DemoSolverCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DemoSolver *DemoSolverRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DemoSolver.Contract.DemoSolverTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DemoSolver *DemoSolverRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DemoSolver.Contract.DemoSolverTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DemoSolver *DemoSolverCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DemoSolver.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DemoSolver *DemoSolverTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DemoSolver.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DemoSolver *DemoSolverTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DemoSolver.Contract.contract.Transact(opts, method, params...)
}

// WETHADDRESS is a free data retrieval call binding the contract method 0x040141e5.
//
// Solidity: function WETH_ADDRESS() view returns(address)
func (_DemoSolver *DemoSolverCaller) WETHADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DemoSolver.contract.Call(opts, &out, "WETH_ADDRESS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WETHADDRESS is a free data retrieval call binding the contract method 0x040141e5.
//
// Solidity: function WETH_ADDRESS() view returns(address)
func (_DemoSolver *DemoSolverSession) WETHADDRESS() (common.Address, error) {
	return _DemoSolver.Contract.WETHADDRESS(&_DemoSolver.CallOpts)
}

// WETHADDRESS is a free data retrieval call binding the contract method 0x040141e5.
//
// Solidity: function WETH_ADDRESS() view returns(address)
func (_DemoSolver *DemoSolverCallerSession) WETHADDRESS() (common.Address, error) {
	return _DemoSolver.Contract.WETHADDRESS(&_DemoSolver.CallOpts)
}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_DemoSolver *DemoSolverCaller) Router(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DemoSolver.contract.Call(opts, &out, "router")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_DemoSolver *DemoSolverSession) Router() (common.Address, error) {
	return _DemoSolver.Contract.Router(&_DemoSolver.CallOpts)
}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_DemoSolver *DemoSolverCallerSession) Router() (common.Address, error) {
	return _DemoSolver.Contract.Router(&_DemoSolver.CallOpts)
}

// AtlasSolverCall is a paid mutator transaction binding the contract method 0x024181a6.
//
// Solidity: function atlasSolverCall(address solverOpFrom, address executionEnvironment, address bidToken, uint256 bidAmount, bytes solverOpData, bytes forwardedData) payable returns()
func (_DemoSolver *DemoSolverTransactor) AtlasSolverCall(opts *bind.TransactOpts, solverOpFrom common.Address, executionEnvironment common.Address, bidToken common.Address, bidAmount *big.Int, solverOpData []byte, forwardedData []byte) (*types.Transaction, error) {
	return _DemoSolver.contract.Transact(opts, "atlasSolverCall", solverOpFrom, executionEnvironment, bidToken, bidAmount, solverOpData, forwardedData)
}

// AtlasSolverCall is a paid mutator transaction binding the contract method 0x024181a6.
//
// Solidity: function atlasSolverCall(address solverOpFrom, address executionEnvironment, address bidToken, uint256 bidAmount, bytes solverOpData, bytes forwardedData) payable returns()
func (_DemoSolver *DemoSolverSession) AtlasSolverCall(solverOpFrom common.Address, executionEnvironment common.Address, bidToken common.Address, bidAmount *big.Int, solverOpData []byte, forwardedData []byte) (*types.Transaction, error) {
	return _DemoSolver.Contract.AtlasSolverCall(&_DemoSolver.TransactOpts, solverOpFrom, executionEnvironment, bidToken, bidAmount, solverOpData, forwardedData)
}

// AtlasSolverCall is a paid mutator transaction binding the contract method 0x024181a6.
//
// Solidity: function atlasSolverCall(address solverOpFrom, address executionEnvironment, address bidToken, uint256 bidAmount, bytes solverOpData, bytes forwardedData) payable returns()
func (_DemoSolver *DemoSolverTransactorSession) AtlasSolverCall(solverOpFrom common.Address, executionEnvironment common.Address, bidToken common.Address, bidAmount *big.Int, solverOpData []byte, forwardedData []byte) (*types.Transaction, error) {
	return _DemoSolver.Contract.AtlasSolverCall(&_DemoSolver.TransactOpts, solverOpFrom, executionEnvironment, bidToken, bidAmount, solverOpData, forwardedData)
}

// DemoSwap is a paid mutator transaction binding the contract method 0x42290975.
//
// Solidity: function demoSwap(address pool, address tokenIn, address tokenOut, uint256 amountIn, uint24 originalFee, bool checkCurrentFee) returns()
func (_DemoSolver *DemoSolverTransactor) DemoSwap(opts *bind.TransactOpts, pool common.Address, tokenIn common.Address, tokenOut common.Address, amountIn *big.Int, originalFee *big.Int, checkCurrentFee bool) (*types.Transaction, error) {
	return _DemoSolver.contract.Transact(opts, "demoSwap", pool, tokenIn, tokenOut, amountIn, originalFee, checkCurrentFee)
}

// DemoSwap is a paid mutator transaction binding the contract method 0x42290975.
//
// Solidity: function demoSwap(address pool, address tokenIn, address tokenOut, uint256 amountIn, uint24 originalFee, bool checkCurrentFee) returns()
func (_DemoSolver *DemoSolverSession) DemoSwap(pool common.Address, tokenIn common.Address, tokenOut common.Address, amountIn *big.Int, originalFee *big.Int, checkCurrentFee bool) (*types.Transaction, error) {
	return _DemoSolver.Contract.DemoSwap(&_DemoSolver.TransactOpts, pool, tokenIn, tokenOut, amountIn, originalFee, checkCurrentFee)
}

// DemoSwap is a paid mutator transaction binding the contract method 0x42290975.
//
// Solidity: function demoSwap(address pool, address tokenIn, address tokenOut, uint256 amountIn, uint24 originalFee, bool checkCurrentFee) returns()
func (_DemoSolver *DemoSolverTransactorSession) DemoSwap(pool common.Address, tokenIn common.Address, tokenOut common.Address, amountIn *big.Int, originalFee *big.Int, checkCurrentFee bool) (*types.Transaction, error) {
	return _DemoSolver.Contract.DemoSwap(&_DemoSolver.TransactOpts, pool, tokenIn, tokenOut, amountIn, originalFee, checkCurrentFee)
}
