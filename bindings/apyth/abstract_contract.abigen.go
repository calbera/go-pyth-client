// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package apyth

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
	_ = abi.ConvertType
)

// PythStructsPrice is an auto generated low-level Go binding around an user-defined struct.
type PythStructsPrice struct {
	Price       int64
	Conf        uint64
	Expo        int32
	PublishTime *big.Int
}

// PythStructsPriceFeed is an auto generated low-level Go binding around an user-defined struct.
type PythStructsPriceFeed struct {
	Id       [32]byte
	Price    PythStructsPrice
	EmaPrice PythStructsPrice
}

// ContractMetaData contains all meta data concerning the Contract contract.
var ContractMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"getEmaPrice\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"price\",\"type\":\"tuple\",\"internalType\":\"structPythStructs.Price\",\"components\":[{\"name\":\"price\",\"type\":\"int64\",\"internalType\":\"int64\"},{\"name\":\"conf\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"expo\",\"type\":\"int32\",\"internalType\":\"int32\"},{\"name\":\"publishTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getEmaPriceNoOlderThan\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"age\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"price\",\"type\":\"tuple\",\"internalType\":\"structPythStructs.Price\",\"components\":[{\"name\":\"price\",\"type\":\"int64\",\"internalType\":\"int64\"},{\"name\":\"conf\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"expo\",\"type\":\"int32\",\"internalType\":\"int32\"},{\"name\":\"publishTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getEmaPriceUnsafe\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"price\",\"type\":\"tuple\",\"internalType\":\"structPythStructs.Price\",\"components\":[{\"name\":\"price\",\"type\":\"int64\",\"internalType\":\"int64\"},{\"name\":\"conf\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"expo\",\"type\":\"int32\",\"internalType\":\"int32\"},{\"name\":\"publishTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPrice\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"price\",\"type\":\"tuple\",\"internalType\":\"structPythStructs.Price\",\"components\":[{\"name\":\"price\",\"type\":\"int64\",\"internalType\":\"int64\"},{\"name\":\"conf\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"expo\",\"type\":\"int32\",\"internalType\":\"int32\"},{\"name\":\"publishTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPriceNoOlderThan\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"age\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"price\",\"type\":\"tuple\",\"internalType\":\"structPythStructs.Price\",\"components\":[{\"name\":\"price\",\"type\":\"int64\",\"internalType\":\"int64\"},{\"name\":\"conf\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"expo\",\"type\":\"int32\",\"internalType\":\"int32\"},{\"name\":\"publishTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPriceUnsafe\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"price\",\"type\":\"tuple\",\"internalType\":\"structPythStructs.Price\",\"components\":[{\"name\":\"price\",\"type\":\"int64\",\"internalType\":\"int64\"},{\"name\":\"conf\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"expo\",\"type\":\"int32\",\"internalType\":\"int32\"},{\"name\":\"publishTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getUpdateFee\",\"inputs\":[{\"name\":\"updateData\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"outputs\":[{\"name\":\"feeAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getValidTimePeriod\",\"inputs\":[],\"outputs\":[{\"name\":\"validTimePeriod\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"parsePriceFeedUpdates\",\"inputs\":[{\"name\":\"updateData\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"priceIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"minPublishTime\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"maxPublishTime\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"priceFeeds\",\"type\":\"tuple[]\",\"internalType\":\"structPythStructs.PriceFeed[]\",\"components\":[{\"name\":\"id\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"price\",\"type\":\"tuple\",\"internalType\":\"structPythStructs.Price\",\"components\":[{\"name\":\"price\",\"type\":\"int64\",\"internalType\":\"int64\"},{\"name\":\"conf\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"expo\",\"type\":\"int32\",\"internalType\":\"int32\"},{\"name\":\"publishTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"emaPrice\",\"type\":\"tuple\",\"internalType\":\"structPythStructs.Price\",\"components\":[{\"name\":\"price\",\"type\":\"int64\",\"internalType\":\"int64\"},{\"name\":\"conf\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"expo\",\"type\":\"int32\",\"internalType\":\"int32\"},{\"name\":\"publishTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"parsePriceFeedUpdatesUnique\",\"inputs\":[{\"name\":\"updateData\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"priceIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"minPublishTime\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"maxPublishTime\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"priceFeeds\",\"type\":\"tuple[]\",\"internalType\":\"structPythStructs.PriceFeed[]\",\"components\":[{\"name\":\"id\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"price\",\"type\":\"tuple\",\"internalType\":\"structPythStructs.Price\",\"components\":[{\"name\":\"price\",\"type\":\"int64\",\"internalType\":\"int64\"},{\"name\":\"conf\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"expo\",\"type\":\"int32\",\"internalType\":\"int32\"},{\"name\":\"publishTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"emaPrice\",\"type\":\"tuple\",\"internalType\":\"structPythStructs.Price\",\"components\":[{\"name\":\"price\",\"type\":\"int64\",\"internalType\":\"int64\"},{\"name\":\"conf\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"expo\",\"type\":\"int32\",\"internalType\":\"int32\"},{\"name\":\"publishTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"priceFeedExists\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"exists\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"queryPriceFeed\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"priceFeed\",\"type\":\"tuple\",\"internalType\":\"structPythStructs.PriceFeed\",\"components\":[{\"name\":\"id\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"price\",\"type\":\"tuple\",\"internalType\":\"structPythStructs.Price\",\"components\":[{\"name\":\"price\",\"type\":\"int64\",\"internalType\":\"int64\"},{\"name\":\"conf\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"expo\",\"type\":\"int32\",\"internalType\":\"int32\"},{\"name\":\"publishTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"emaPrice\",\"type\":\"tuple\",\"internalType\":\"structPythStructs.Price\",\"components\":[{\"name\":\"price\",\"type\":\"int64\",\"internalType\":\"int64\"},{\"name\":\"conf\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"expo\",\"type\":\"int32\",\"internalType\":\"int32\"},{\"name\":\"publishTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"updatePriceFeeds\",\"inputs\":[{\"name\":\"updateData\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"updatePriceFeedsIfNecessary\",\"inputs\":[{\"name\":\"updateData\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"priceIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"publishTimes\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"event\",\"name\":\"BatchPriceFeedUpdate\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PriceFeedUpdate\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"publishTime\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"price\",\"type\":\"int64\",\"indexed\":false,\"internalType\":\"int64\"},{\"name\":\"conf\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"InvalidArgument\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NoFreshUpdate\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StalePrice\",\"inputs\":[]}]",
}

// ContractABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractMetaData.ABI instead.
var ContractABI = ContractMetaData.ABI

// Contract is an auto generated Go binding around an Ethereum contract.
type Contract struct {
	ContractCaller     // Read-only binding to the contract
	ContractTransactor // Write-only binding to the contract
	ContractFilterer   // Log filterer for contract events
}

// ContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractSession struct {
	Contract     *Contract         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractCallerSession struct {
	Contract *ContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractTransactorSession struct {
	Contract     *ContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractRaw struct {
	Contract *Contract // Generic contract binding to access the raw methods on
}

// ContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractCallerRaw struct {
	Contract *ContractCaller // Generic read-only contract binding to access the raw methods on
}

// ContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractTransactorRaw struct {
	Contract *ContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContract creates a new instance of Contract, bound to a specific deployed contract.
func NewContract(address common.Address, backend bind.ContractBackend) (*Contract, error) {
	contract, err := bindContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

// NewContractCaller creates a new read-only instance of Contract, bound to a specific deployed contract.
func NewContractCaller(address common.Address, caller bind.ContractCaller) (*ContractCaller, error) {
	contract, err := bindContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractCaller{contract: contract}, nil
}

// NewContractTransactor creates a new write-only instance of Contract, bound to a specific deployed contract.
func NewContractTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractTransactor, error) {
	contract, err := bindContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractTransactor{contract: contract}, nil
}

// NewContractFilterer creates a new log filterer instance of Contract, bound to a specific deployed contract.
func NewContractFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractFilterer, error) {
	contract, err := bindContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractFilterer{contract: contract}, nil
}

// bindContract binds a generic wrapper to an already deployed contract.
func bindContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.ContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transact(opts, method, params...)
}

// GetEmaPrice is a free data retrieval call binding the contract method 0xb5dcc911.
//
// Solidity: function getEmaPrice(bytes32 id) view returns((int64,uint64,int32,uint256) price)
func (_Contract *ContractCaller) GetEmaPrice(opts *bind.CallOpts, id [32]byte) (PythStructsPrice, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getEmaPrice", id)

	if err != nil {
		return *new(PythStructsPrice), err
	}

	out0 := *abi.ConvertType(out[0], new(PythStructsPrice)).(*PythStructsPrice)

	return out0, err

}

// GetEmaPrice is a free data retrieval call binding the contract method 0xb5dcc911.
//
// Solidity: function getEmaPrice(bytes32 id) view returns((int64,uint64,int32,uint256) price)
func (_Contract *ContractSession) GetEmaPrice(id [32]byte) (PythStructsPrice, error) {
	return _Contract.Contract.GetEmaPrice(&_Contract.CallOpts, id)
}

// GetEmaPrice is a free data retrieval call binding the contract method 0xb5dcc911.
//
// Solidity: function getEmaPrice(bytes32 id) view returns((int64,uint64,int32,uint256) price)
func (_Contract *ContractCallerSession) GetEmaPrice(id [32]byte) (PythStructsPrice, error) {
	return _Contract.Contract.GetEmaPrice(&_Contract.CallOpts, id)
}

// GetEmaPriceNoOlderThan is a free data retrieval call binding the contract method 0x711a2e28.
//
// Solidity: function getEmaPriceNoOlderThan(bytes32 id, uint256 age) view returns((int64,uint64,int32,uint256) price)
func (_Contract *ContractCaller) GetEmaPriceNoOlderThan(opts *bind.CallOpts, id [32]byte, age *big.Int) (PythStructsPrice, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getEmaPriceNoOlderThan", id, age)

	if err != nil {
		return *new(PythStructsPrice), err
	}

	out0 := *abi.ConvertType(out[0], new(PythStructsPrice)).(*PythStructsPrice)

	return out0, err

}

// GetEmaPriceNoOlderThan is a free data retrieval call binding the contract method 0x711a2e28.
//
// Solidity: function getEmaPriceNoOlderThan(bytes32 id, uint256 age) view returns((int64,uint64,int32,uint256) price)
func (_Contract *ContractSession) GetEmaPriceNoOlderThan(id [32]byte, age *big.Int) (PythStructsPrice, error) {
	return _Contract.Contract.GetEmaPriceNoOlderThan(&_Contract.CallOpts, id, age)
}

// GetEmaPriceNoOlderThan is a free data retrieval call binding the contract method 0x711a2e28.
//
// Solidity: function getEmaPriceNoOlderThan(bytes32 id, uint256 age) view returns((int64,uint64,int32,uint256) price)
func (_Contract *ContractCallerSession) GetEmaPriceNoOlderThan(id [32]byte, age *big.Int) (PythStructsPrice, error) {
	return _Contract.Contract.GetEmaPriceNoOlderThan(&_Contract.CallOpts, id, age)
}

// GetEmaPriceUnsafe is a free data retrieval call binding the contract method 0x9474f45b.
//
// Solidity: function getEmaPriceUnsafe(bytes32 id) view returns((int64,uint64,int32,uint256) price)
func (_Contract *ContractCaller) GetEmaPriceUnsafe(opts *bind.CallOpts, id [32]byte) (PythStructsPrice, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getEmaPriceUnsafe", id)

	if err != nil {
		return *new(PythStructsPrice), err
	}

	out0 := *abi.ConvertType(out[0], new(PythStructsPrice)).(*PythStructsPrice)

	return out0, err

}

// GetEmaPriceUnsafe is a free data retrieval call binding the contract method 0x9474f45b.
//
// Solidity: function getEmaPriceUnsafe(bytes32 id) view returns((int64,uint64,int32,uint256) price)
func (_Contract *ContractSession) GetEmaPriceUnsafe(id [32]byte) (PythStructsPrice, error) {
	return _Contract.Contract.GetEmaPriceUnsafe(&_Contract.CallOpts, id)
}

// GetEmaPriceUnsafe is a free data retrieval call binding the contract method 0x9474f45b.
//
// Solidity: function getEmaPriceUnsafe(bytes32 id) view returns((int64,uint64,int32,uint256) price)
func (_Contract *ContractCallerSession) GetEmaPriceUnsafe(id [32]byte) (PythStructsPrice, error) {
	return _Contract.Contract.GetEmaPriceUnsafe(&_Contract.CallOpts, id)
}

// GetPrice is a free data retrieval call binding the contract method 0x31d98b3f.
//
// Solidity: function getPrice(bytes32 id) view returns((int64,uint64,int32,uint256) price)
func (_Contract *ContractCaller) GetPrice(opts *bind.CallOpts, id [32]byte) (PythStructsPrice, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getPrice", id)

	if err != nil {
		return *new(PythStructsPrice), err
	}

	out0 := *abi.ConvertType(out[0], new(PythStructsPrice)).(*PythStructsPrice)

	return out0, err

}

// GetPrice is a free data retrieval call binding the contract method 0x31d98b3f.
//
// Solidity: function getPrice(bytes32 id) view returns((int64,uint64,int32,uint256) price)
func (_Contract *ContractSession) GetPrice(id [32]byte) (PythStructsPrice, error) {
	return _Contract.Contract.GetPrice(&_Contract.CallOpts, id)
}

// GetPrice is a free data retrieval call binding the contract method 0x31d98b3f.
//
// Solidity: function getPrice(bytes32 id) view returns((int64,uint64,int32,uint256) price)
func (_Contract *ContractCallerSession) GetPrice(id [32]byte) (PythStructsPrice, error) {
	return _Contract.Contract.GetPrice(&_Contract.CallOpts, id)
}

// GetPriceNoOlderThan is a free data retrieval call binding the contract method 0xa4ae35e0.
//
// Solidity: function getPriceNoOlderThan(bytes32 id, uint256 age) view returns((int64,uint64,int32,uint256) price)
func (_Contract *ContractCaller) GetPriceNoOlderThan(opts *bind.CallOpts, id [32]byte, age *big.Int) (PythStructsPrice, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getPriceNoOlderThan", id, age)

	if err != nil {
		return *new(PythStructsPrice), err
	}

	out0 := *abi.ConvertType(out[0], new(PythStructsPrice)).(*PythStructsPrice)

	return out0, err

}

// GetPriceNoOlderThan is a free data retrieval call binding the contract method 0xa4ae35e0.
//
// Solidity: function getPriceNoOlderThan(bytes32 id, uint256 age) view returns((int64,uint64,int32,uint256) price)
func (_Contract *ContractSession) GetPriceNoOlderThan(id [32]byte, age *big.Int) (PythStructsPrice, error) {
	return _Contract.Contract.GetPriceNoOlderThan(&_Contract.CallOpts, id, age)
}

// GetPriceNoOlderThan is a free data retrieval call binding the contract method 0xa4ae35e0.
//
// Solidity: function getPriceNoOlderThan(bytes32 id, uint256 age) view returns((int64,uint64,int32,uint256) price)
func (_Contract *ContractCallerSession) GetPriceNoOlderThan(id [32]byte, age *big.Int) (PythStructsPrice, error) {
	return _Contract.Contract.GetPriceNoOlderThan(&_Contract.CallOpts, id, age)
}

// GetPriceUnsafe is a free data retrieval call binding the contract method 0x96834ad3.
//
// Solidity: function getPriceUnsafe(bytes32 id) view returns((int64,uint64,int32,uint256) price)
func (_Contract *ContractCaller) GetPriceUnsafe(opts *bind.CallOpts, id [32]byte) (PythStructsPrice, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getPriceUnsafe", id)

	if err != nil {
		return *new(PythStructsPrice), err
	}

	out0 := *abi.ConvertType(out[0], new(PythStructsPrice)).(*PythStructsPrice)

	return out0, err

}

// GetPriceUnsafe is a free data retrieval call binding the contract method 0x96834ad3.
//
// Solidity: function getPriceUnsafe(bytes32 id) view returns((int64,uint64,int32,uint256) price)
func (_Contract *ContractSession) GetPriceUnsafe(id [32]byte) (PythStructsPrice, error) {
	return _Contract.Contract.GetPriceUnsafe(&_Contract.CallOpts, id)
}

// GetPriceUnsafe is a free data retrieval call binding the contract method 0x96834ad3.
//
// Solidity: function getPriceUnsafe(bytes32 id) view returns((int64,uint64,int32,uint256) price)
func (_Contract *ContractCallerSession) GetPriceUnsafe(id [32]byte) (PythStructsPrice, error) {
	return _Contract.Contract.GetPriceUnsafe(&_Contract.CallOpts, id)
}

// GetUpdateFee is a free data retrieval call binding the contract method 0xd47eed45.
//
// Solidity: function getUpdateFee(bytes[] updateData) view returns(uint256 feeAmount)
func (_Contract *ContractCaller) GetUpdateFee(opts *bind.CallOpts, updateData [][]byte) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getUpdateFee", updateData)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUpdateFee is a free data retrieval call binding the contract method 0xd47eed45.
//
// Solidity: function getUpdateFee(bytes[] updateData) view returns(uint256 feeAmount)
func (_Contract *ContractSession) GetUpdateFee(updateData [][]byte) (*big.Int, error) {
	return _Contract.Contract.GetUpdateFee(&_Contract.CallOpts, updateData)
}

// GetUpdateFee is a free data retrieval call binding the contract method 0xd47eed45.
//
// Solidity: function getUpdateFee(bytes[] updateData) view returns(uint256 feeAmount)
func (_Contract *ContractCallerSession) GetUpdateFee(updateData [][]byte) (*big.Int, error) {
	return _Contract.Contract.GetUpdateFee(&_Contract.CallOpts, updateData)
}

// GetValidTimePeriod is a free data retrieval call binding the contract method 0xe18910a3.
//
// Solidity: function getValidTimePeriod() view returns(uint256 validTimePeriod)
func (_Contract *ContractCaller) GetValidTimePeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getValidTimePeriod")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetValidTimePeriod is a free data retrieval call binding the contract method 0xe18910a3.
//
// Solidity: function getValidTimePeriod() view returns(uint256 validTimePeriod)
func (_Contract *ContractSession) GetValidTimePeriod() (*big.Int, error) {
	return _Contract.Contract.GetValidTimePeriod(&_Contract.CallOpts)
}

// GetValidTimePeriod is a free data retrieval call binding the contract method 0xe18910a3.
//
// Solidity: function getValidTimePeriod() view returns(uint256 validTimePeriod)
func (_Contract *ContractCallerSession) GetValidTimePeriod() (*big.Int, error) {
	return _Contract.Contract.GetValidTimePeriod(&_Contract.CallOpts)
}

// PriceFeedExists is a free data retrieval call binding the contract method 0xb5ec0261.
//
// Solidity: function priceFeedExists(bytes32 id) view returns(bool exists)
func (_Contract *ContractCaller) PriceFeedExists(opts *bind.CallOpts, id [32]byte) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "priceFeedExists", id)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// PriceFeedExists is a free data retrieval call binding the contract method 0xb5ec0261.
//
// Solidity: function priceFeedExists(bytes32 id) view returns(bool exists)
func (_Contract *ContractSession) PriceFeedExists(id [32]byte) (bool, error) {
	return _Contract.Contract.PriceFeedExists(&_Contract.CallOpts, id)
}

// PriceFeedExists is a free data retrieval call binding the contract method 0xb5ec0261.
//
// Solidity: function priceFeedExists(bytes32 id) view returns(bool exists)
func (_Contract *ContractCallerSession) PriceFeedExists(id [32]byte) (bool, error) {
	return _Contract.Contract.PriceFeedExists(&_Contract.CallOpts, id)
}

// QueryPriceFeed is a free data retrieval call binding the contract method 0xcaaf43f1.
//
// Solidity: function queryPriceFeed(bytes32 id) view returns((bytes32,(int64,uint64,int32,uint256),(int64,uint64,int32,uint256)) priceFeed)
func (_Contract *ContractCaller) QueryPriceFeed(opts *bind.CallOpts, id [32]byte) (PythStructsPriceFeed, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "queryPriceFeed", id)

	if err != nil {
		return *new(PythStructsPriceFeed), err
	}

	out0 := *abi.ConvertType(out[0], new(PythStructsPriceFeed)).(*PythStructsPriceFeed)

	return out0, err

}

// QueryPriceFeed is a free data retrieval call binding the contract method 0xcaaf43f1.
//
// Solidity: function queryPriceFeed(bytes32 id) view returns((bytes32,(int64,uint64,int32,uint256),(int64,uint64,int32,uint256)) priceFeed)
func (_Contract *ContractSession) QueryPriceFeed(id [32]byte) (PythStructsPriceFeed, error) {
	return _Contract.Contract.QueryPriceFeed(&_Contract.CallOpts, id)
}

// QueryPriceFeed is a free data retrieval call binding the contract method 0xcaaf43f1.
//
// Solidity: function queryPriceFeed(bytes32 id) view returns((bytes32,(int64,uint64,int32,uint256),(int64,uint64,int32,uint256)) priceFeed)
func (_Contract *ContractCallerSession) QueryPriceFeed(id [32]byte) (PythStructsPriceFeed, error) {
	return _Contract.Contract.QueryPriceFeed(&_Contract.CallOpts, id)
}

// ParsePriceFeedUpdates is a paid mutator transaction binding the contract method 0x4716e9c5.
//
// Solidity: function parsePriceFeedUpdates(bytes[] updateData, bytes32[] priceIds, uint64 minPublishTime, uint64 maxPublishTime) payable returns((bytes32,(int64,uint64,int32,uint256),(int64,uint64,int32,uint256))[] priceFeeds)
func (_Contract *ContractTransactor) ParsePriceFeedUpdates(opts *bind.TransactOpts, updateData [][]byte, priceIds [][32]byte, minPublishTime uint64, maxPublishTime uint64) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "parsePriceFeedUpdates", updateData, priceIds, minPublishTime, maxPublishTime)
}

// ParsePriceFeedUpdates is a paid mutator transaction binding the contract method 0x4716e9c5.
//
// Solidity: function parsePriceFeedUpdates(bytes[] updateData, bytes32[] priceIds, uint64 minPublishTime, uint64 maxPublishTime) payable returns((bytes32,(int64,uint64,int32,uint256),(int64,uint64,int32,uint256))[] priceFeeds)
func (_Contract *ContractSession) ParsePriceFeedUpdates(updateData [][]byte, priceIds [][32]byte, minPublishTime uint64, maxPublishTime uint64) (*types.Transaction, error) {
	return _Contract.Contract.ParsePriceFeedUpdates(&_Contract.TransactOpts, updateData, priceIds, minPublishTime, maxPublishTime)
}

// ParsePriceFeedUpdates is a paid mutator transaction binding the contract method 0x4716e9c5.
//
// Solidity: function parsePriceFeedUpdates(bytes[] updateData, bytes32[] priceIds, uint64 minPublishTime, uint64 maxPublishTime) payable returns((bytes32,(int64,uint64,int32,uint256),(int64,uint64,int32,uint256))[] priceFeeds)
func (_Contract *ContractTransactorSession) ParsePriceFeedUpdates(updateData [][]byte, priceIds [][32]byte, minPublishTime uint64, maxPublishTime uint64) (*types.Transaction, error) {
	return _Contract.Contract.ParsePriceFeedUpdates(&_Contract.TransactOpts, updateData, priceIds, minPublishTime, maxPublishTime)
}

// ParsePriceFeedUpdatesUnique is a paid mutator transaction binding the contract method 0xaccca7f9.
//
// Solidity: function parsePriceFeedUpdatesUnique(bytes[] updateData, bytes32[] priceIds, uint64 minPublishTime, uint64 maxPublishTime) payable returns((bytes32,(int64,uint64,int32,uint256),(int64,uint64,int32,uint256))[] priceFeeds)
func (_Contract *ContractTransactor) ParsePriceFeedUpdatesUnique(opts *bind.TransactOpts, updateData [][]byte, priceIds [][32]byte, minPublishTime uint64, maxPublishTime uint64) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "parsePriceFeedUpdatesUnique", updateData, priceIds, minPublishTime, maxPublishTime)
}

// ParsePriceFeedUpdatesUnique is a paid mutator transaction binding the contract method 0xaccca7f9.
//
// Solidity: function parsePriceFeedUpdatesUnique(bytes[] updateData, bytes32[] priceIds, uint64 minPublishTime, uint64 maxPublishTime) payable returns((bytes32,(int64,uint64,int32,uint256),(int64,uint64,int32,uint256))[] priceFeeds)
func (_Contract *ContractSession) ParsePriceFeedUpdatesUnique(updateData [][]byte, priceIds [][32]byte, minPublishTime uint64, maxPublishTime uint64) (*types.Transaction, error) {
	return _Contract.Contract.ParsePriceFeedUpdatesUnique(&_Contract.TransactOpts, updateData, priceIds, minPublishTime, maxPublishTime)
}

// ParsePriceFeedUpdatesUnique is a paid mutator transaction binding the contract method 0xaccca7f9.
//
// Solidity: function parsePriceFeedUpdatesUnique(bytes[] updateData, bytes32[] priceIds, uint64 minPublishTime, uint64 maxPublishTime) payable returns((bytes32,(int64,uint64,int32,uint256),(int64,uint64,int32,uint256))[] priceFeeds)
func (_Contract *ContractTransactorSession) ParsePriceFeedUpdatesUnique(updateData [][]byte, priceIds [][32]byte, minPublishTime uint64, maxPublishTime uint64) (*types.Transaction, error) {
	return _Contract.Contract.ParsePriceFeedUpdatesUnique(&_Contract.TransactOpts, updateData, priceIds, minPublishTime, maxPublishTime)
}

// UpdatePriceFeeds is a paid mutator transaction binding the contract method 0xef9e5e28.
//
// Solidity: function updatePriceFeeds(bytes[] updateData) payable returns()
func (_Contract *ContractTransactor) UpdatePriceFeeds(opts *bind.TransactOpts, updateData [][]byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "updatePriceFeeds", updateData)
}

// UpdatePriceFeeds is a paid mutator transaction binding the contract method 0xef9e5e28.
//
// Solidity: function updatePriceFeeds(bytes[] updateData) payable returns()
func (_Contract *ContractSession) UpdatePriceFeeds(updateData [][]byte) (*types.Transaction, error) {
	return _Contract.Contract.UpdatePriceFeeds(&_Contract.TransactOpts, updateData)
}

// UpdatePriceFeeds is a paid mutator transaction binding the contract method 0xef9e5e28.
//
// Solidity: function updatePriceFeeds(bytes[] updateData) payable returns()
func (_Contract *ContractTransactorSession) UpdatePriceFeeds(updateData [][]byte) (*types.Transaction, error) {
	return _Contract.Contract.UpdatePriceFeeds(&_Contract.TransactOpts, updateData)
}

// UpdatePriceFeedsIfNecessary is a paid mutator transaction binding the contract method 0xb9256d28.
//
// Solidity: function updatePriceFeedsIfNecessary(bytes[] updateData, bytes32[] priceIds, uint64[] publishTimes) payable returns()
func (_Contract *ContractTransactor) UpdatePriceFeedsIfNecessary(opts *bind.TransactOpts, updateData [][]byte, priceIds [][32]byte, publishTimes []uint64) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "updatePriceFeedsIfNecessary", updateData, priceIds, publishTimes)
}

// UpdatePriceFeedsIfNecessary is a paid mutator transaction binding the contract method 0xb9256d28.
//
// Solidity: function updatePriceFeedsIfNecessary(bytes[] updateData, bytes32[] priceIds, uint64[] publishTimes) payable returns()
func (_Contract *ContractSession) UpdatePriceFeedsIfNecessary(updateData [][]byte, priceIds [][32]byte, publishTimes []uint64) (*types.Transaction, error) {
	return _Contract.Contract.UpdatePriceFeedsIfNecessary(&_Contract.TransactOpts, updateData, priceIds, publishTimes)
}

// UpdatePriceFeedsIfNecessary is a paid mutator transaction binding the contract method 0xb9256d28.
//
// Solidity: function updatePriceFeedsIfNecessary(bytes[] updateData, bytes32[] priceIds, uint64[] publishTimes) payable returns()
func (_Contract *ContractTransactorSession) UpdatePriceFeedsIfNecessary(updateData [][]byte, priceIds [][32]byte, publishTimes []uint64) (*types.Transaction, error) {
	return _Contract.Contract.UpdatePriceFeedsIfNecessary(&_Contract.TransactOpts, updateData, priceIds, publishTimes)
}

// ContractBatchPriceFeedUpdateIterator is returned from FilterBatchPriceFeedUpdate and is used to iterate over the raw logs and unpacked data for BatchPriceFeedUpdate events raised by the Contract contract.
type ContractBatchPriceFeedUpdateIterator struct {
	Event *ContractBatchPriceFeedUpdate // Event containing the contract specifics and raw log

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
func (it *ContractBatchPriceFeedUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractBatchPriceFeedUpdate)
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
		it.Event = new(ContractBatchPriceFeedUpdate)
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
func (it *ContractBatchPriceFeedUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractBatchPriceFeedUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractBatchPriceFeedUpdate represents a BatchPriceFeedUpdate event raised by the Contract contract.
type ContractBatchPriceFeedUpdate struct {
	ChainId        uint16
	SequenceNumber uint64
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterBatchPriceFeedUpdate is a free log retrieval operation binding the contract event 0x943f0e8a16c19895fb87cbeb1a349ed86d7f31923089dd36c1a1ed5e300f267b.
//
// Solidity: event BatchPriceFeedUpdate(uint16 chainId, uint64 sequenceNumber)
func (_Contract *ContractFilterer) FilterBatchPriceFeedUpdate(opts *bind.FilterOpts) (*ContractBatchPriceFeedUpdateIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "BatchPriceFeedUpdate")
	if err != nil {
		return nil, err
	}
	return &ContractBatchPriceFeedUpdateIterator{contract: _Contract.contract, event: "BatchPriceFeedUpdate", logs: logs, sub: sub}, nil
}

// WatchBatchPriceFeedUpdate is a free log subscription operation binding the contract event 0x943f0e8a16c19895fb87cbeb1a349ed86d7f31923089dd36c1a1ed5e300f267b.
//
// Solidity: event BatchPriceFeedUpdate(uint16 chainId, uint64 sequenceNumber)
func (_Contract *ContractFilterer) WatchBatchPriceFeedUpdate(opts *bind.WatchOpts, sink chan<- *ContractBatchPriceFeedUpdate) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "BatchPriceFeedUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractBatchPriceFeedUpdate)
				if err := _Contract.contract.UnpackLog(event, "BatchPriceFeedUpdate", log); err != nil {
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

// ParseBatchPriceFeedUpdate is a log parse operation binding the contract event 0x943f0e8a16c19895fb87cbeb1a349ed86d7f31923089dd36c1a1ed5e300f267b.
//
// Solidity: event BatchPriceFeedUpdate(uint16 chainId, uint64 sequenceNumber)
func (_Contract *ContractFilterer) ParseBatchPriceFeedUpdate(log types.Log) (*ContractBatchPriceFeedUpdate, error) {
	event := new(ContractBatchPriceFeedUpdate)
	if err := _Contract.contract.UnpackLog(event, "BatchPriceFeedUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractPriceFeedUpdateIterator is returned from FilterPriceFeedUpdate and is used to iterate over the raw logs and unpacked data for PriceFeedUpdate events raised by the Contract contract.
type ContractPriceFeedUpdateIterator struct {
	Event *ContractPriceFeedUpdate // Event containing the contract specifics and raw log

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
func (it *ContractPriceFeedUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractPriceFeedUpdate)
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
		it.Event = new(ContractPriceFeedUpdate)
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
func (it *ContractPriceFeedUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractPriceFeedUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractPriceFeedUpdate represents a PriceFeedUpdate event raised by the Contract contract.
type ContractPriceFeedUpdate struct {
	Id          [32]byte
	PublishTime uint64
	Price       int64
	Conf        uint64
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterPriceFeedUpdate is a free log retrieval operation binding the contract event 0xd06a6b7f4918494b3719217d1802786c1f5112a6c1d88fe2cfec00b4584f6aec.
//
// Solidity: event PriceFeedUpdate(bytes32 indexed id, uint64 publishTime, int64 price, uint64 conf)
func (_Contract *ContractFilterer) FilterPriceFeedUpdate(opts *bind.FilterOpts, id [][32]byte) (*ContractPriceFeedUpdateIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "PriceFeedUpdate", idRule)
	if err != nil {
		return nil, err
	}
	return &ContractPriceFeedUpdateIterator{contract: _Contract.contract, event: "PriceFeedUpdate", logs: logs, sub: sub}, nil
}

// WatchPriceFeedUpdate is a free log subscription operation binding the contract event 0xd06a6b7f4918494b3719217d1802786c1f5112a6c1d88fe2cfec00b4584f6aec.
//
// Solidity: event PriceFeedUpdate(bytes32 indexed id, uint64 publishTime, int64 price, uint64 conf)
func (_Contract *ContractFilterer) WatchPriceFeedUpdate(opts *bind.WatchOpts, sink chan<- *ContractPriceFeedUpdate, id [][32]byte) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "PriceFeedUpdate", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractPriceFeedUpdate)
				if err := _Contract.contract.UnpackLog(event, "PriceFeedUpdate", log); err != nil {
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

// ParsePriceFeedUpdate is a log parse operation binding the contract event 0xd06a6b7f4918494b3719217d1802786c1f5112a6c1d88fe2cfec00b4584f6aec.
//
// Solidity: event PriceFeedUpdate(bytes32 indexed id, uint64 publishTime, int64 price, uint64 conf)
func (_Contract *ContractFilterer) ParsePriceFeedUpdate(log types.Log) (*ContractPriceFeedUpdate, error) {
	event := new(ContractPriceFeedUpdate)
	if err := _Contract.contract.UnpackLog(event, "PriceFeedUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
