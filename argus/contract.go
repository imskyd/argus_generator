// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package argus

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

// CallData is an auto generated low-level Go binding around an user-defined struct.
type CallData struct {
	Flag  *big.Int
	To    common.Address
	Value *big.Int
	Data  []byte
	Hint  []byte
	Extra []byte
}

// TransactionData is an auto generated low-level Go binding around an user-defined struct.
type TransactionData struct {
	From     common.Address
	Delegate common.Address
	Flag     *big.Int
	To       common.Address
	Value    *big.Int
	Data     []byte
	Hint     []byte
	Extra    []byte
}

// TransactionResult is an auto generated low-level Go binding around an user-defined struct.
type TransactionResult struct {
	Success bool
	Data    []byte
	Hint    []byte
}

// ArgusMetaData contains all meta data concerning the Argus contract.
var ArgusMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"authorizer\",\"type\":\"address\"}],\"name\":\"AuthorizerSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegate\",\"type\":\"address\"}],\"name\":\"DelegateAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegate\",\"type\":\"address\"}],\"name\":\"DelegateRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"NewOwnerSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"PendingOwnerSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"roleManager\",\"type\":\"address\"}],\"name\":\"RoleManagerSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"delegate\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"flag\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"hint\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"extra\",\"type\":\"bytes\"}],\"indexed\":false,\"internalType\":\"structTransactionData\",\"name\":\"transaction\",\"type\":\"tuple\"}],\"name\":\"TransactionExecuted\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"NAME\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VERSION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_delegate\",\"type\":\"address\"}],\"name\":\"addDelegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_delegates\",\"type\":\"address[]\"}],\"name\":\"addDelegates\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"authorizer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"flag\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"hint\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"extra\",\"type\":\"bytes\"}],\"internalType\":\"structCallData\",\"name\":\"callData\",\"type\":\"tuple\"}],\"name\":\"execTransaction\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"hint\",\"type\":\"bytes\"}],\"internalType\":\"structTransactionResult\",\"name\":\"result\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"flag\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"hint\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"extra\",\"type\":\"bytes\"}],\"internalType\":\"structCallData[]\",\"name\":\"callDataList\",\"type\":\"tuple[]\"}],\"name\":\"execTransactions\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"hint\",\"type\":\"bytes\"}],\"internalType\":\"structTransactionResult[]\",\"name\":\"resultList\",\"type\":\"tuple[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAccountAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllDelegates\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_delegate\",\"type\":\"address\"}],\"name\":\"hasDelegate\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_roleManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_authorizer\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_delegate\",\"type\":\"address\"}],\"name\":\"removeDelegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_delegates\",\"type\":\"address[]\"}],\"name\":\"removeDelegates\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"roleManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"safe\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_authorizer\",\"type\":\"address\"}],\"name\":\"setAuthorizer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"setPendingOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_roleManager\",\"type\":\"address\"}],\"name\":\"setRoleManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// ArgusABI is the input ABI used to generate the binding from.
// Deprecated: Use ArgusMetaData.ABI instead.
var ArgusABI = ArgusMetaData.ABI

// Argus is an auto generated Go binding around an Ethereum contract.
type Argus struct {
	ArgusCaller     // Read-only binding to the contract
	ArgusTransactor // Write-only binding to the contract
	ArgusFilterer   // Log filterer for contract events
}

// ArgusCaller is an auto generated read-only Go binding around an Ethereum contract.
type ArgusCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArgusTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ArgusTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArgusFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ArgusFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArgusSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ArgusSession struct {
	Contract     *Argus            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ArgusCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ArgusCallerSession struct {
	Contract *ArgusCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ArgusTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ArgusTransactorSession struct {
	Contract     *ArgusTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ArgusRaw is an auto generated low-level Go binding around an Ethereum contract.
type ArgusRaw struct {
	Contract *Argus // Generic contract binding to access the raw methods on
}

// ArgusCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ArgusCallerRaw struct {
	Contract *ArgusCaller // Generic read-only contract binding to access the raw methods on
}

// ArgusTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ArgusTransactorRaw struct {
	Contract *ArgusTransactor // Generic write-only contract binding to access the raw methods on
}

// NewArgus creates a new instance of Argus, bound to a specific deployed contract.
func NewArgus(address common.Address, backend bind.ContractBackend) (*Argus, error) {
	contract, err := bindArgus(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Argus{ArgusCaller: ArgusCaller{contract: contract}, ArgusTransactor: ArgusTransactor{contract: contract}, ArgusFilterer: ArgusFilterer{contract: contract}}, nil
}

// NewArgusCaller creates a new read-only instance of Argus, bound to a specific deployed contract.
func NewArgusCaller(address common.Address, caller bind.ContractCaller) (*ArgusCaller, error) {
	contract, err := bindArgus(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ArgusCaller{contract: contract}, nil
}

// NewArgusTransactor creates a new write-only instance of Argus, bound to a specific deployed contract.
func NewArgusTransactor(address common.Address, transactor bind.ContractTransactor) (*ArgusTransactor, error) {
	contract, err := bindArgus(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ArgusTransactor{contract: contract}, nil
}

// NewArgusFilterer creates a new log filterer instance of Argus, bound to a specific deployed contract.
func NewArgusFilterer(address common.Address, filterer bind.ContractFilterer) (*ArgusFilterer, error) {
	contract, err := bindArgus(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ArgusFilterer{contract: contract}, nil
}

// bindArgus binds a generic wrapper to an already deployed contract.
func bindArgus(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ArgusABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Argus *ArgusRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Argus.Contract.ArgusCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Argus *ArgusRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Argus.Contract.ArgusTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Argus *ArgusRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Argus.Contract.ArgusTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Argus *ArgusCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Argus.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Argus *ArgusTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Argus.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Argus *ArgusTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Argus.Contract.contract.Transact(opts, method, params...)
}

// NAME is a free data retrieval call binding the contract method 0xa3f4df7e.
//
// Solidity: function NAME() view returns(bytes32)
func (_Argus *ArgusCaller) NAME(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Argus.contract.Call(opts, &out, "NAME")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// NAME is a free data retrieval call binding the contract method 0xa3f4df7e.
//
// Solidity: function NAME() view returns(bytes32)
func (_Argus *ArgusSession) NAME() ([32]byte, error) {
	return _Argus.Contract.NAME(&_Argus.CallOpts)
}

// NAME is a free data retrieval call binding the contract method 0xa3f4df7e.
//
// Solidity: function NAME() view returns(bytes32)
func (_Argus *ArgusCallerSession) NAME() ([32]byte, error) {
	return _Argus.Contract.NAME(&_Argus.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(uint256)
func (_Argus *ArgusCaller) VERSION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Argus.contract.Call(opts, &out, "VERSION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(uint256)
func (_Argus *ArgusSession) VERSION() (*big.Int, error) {
	return _Argus.Contract.VERSION(&_Argus.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(uint256)
func (_Argus *ArgusCallerSession) VERSION() (*big.Int, error) {
	return _Argus.Contract.VERSION(&_Argus.CallOpts)
}

// Authorizer is a free data retrieval call binding the contract method 0xd09edf31.
//
// Solidity: function authorizer() view returns(address)
func (_Argus *ArgusCaller) Authorizer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Argus.contract.Call(opts, &out, "authorizer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Authorizer is a free data retrieval call binding the contract method 0xd09edf31.
//
// Solidity: function authorizer() view returns(address)
func (_Argus *ArgusSession) Authorizer() (common.Address, error) {
	return _Argus.Contract.Authorizer(&_Argus.CallOpts)
}

// Authorizer is a free data retrieval call binding the contract method 0xd09edf31.
//
// Solidity: function authorizer() view returns(address)
func (_Argus *ArgusCallerSession) Authorizer() (common.Address, error) {
	return _Argus.Contract.Authorizer(&_Argus.CallOpts)
}

// GetAccountAddress is a free data retrieval call binding the contract method 0x0e2562d9.
//
// Solidity: function getAccountAddress() view returns(address account)
func (_Argus *ArgusCaller) GetAccountAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Argus.contract.Call(opts, &out, "getAccountAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAccountAddress is a free data retrieval call binding the contract method 0x0e2562d9.
//
// Solidity: function getAccountAddress() view returns(address account)
func (_Argus *ArgusSession) GetAccountAddress() (common.Address, error) {
	return _Argus.Contract.GetAccountAddress(&_Argus.CallOpts)
}

// GetAccountAddress is a free data retrieval call binding the contract method 0x0e2562d9.
//
// Solidity: function getAccountAddress() view returns(address account)
func (_Argus *ArgusCallerSession) GetAccountAddress() (common.Address, error) {
	return _Argus.Contract.GetAccountAddress(&_Argus.CallOpts)
}

// GetAllDelegates is a free data retrieval call binding the contract method 0x7a4f3764.
//
// Solidity: function getAllDelegates() view returns(address[])
func (_Argus *ArgusCaller) GetAllDelegates(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Argus.contract.Call(opts, &out, "getAllDelegates")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetAllDelegates is a free data retrieval call binding the contract method 0x7a4f3764.
//
// Solidity: function getAllDelegates() view returns(address[])
func (_Argus *ArgusSession) GetAllDelegates() ([]common.Address, error) {
	return _Argus.Contract.GetAllDelegates(&_Argus.CallOpts)
}

// GetAllDelegates is a free data retrieval call binding the contract method 0x7a4f3764.
//
// Solidity: function getAllDelegates() view returns(address[])
func (_Argus *ArgusCallerSession) GetAllDelegates() ([]common.Address, error) {
	return _Argus.Contract.GetAllDelegates(&_Argus.CallOpts)
}

// HasDelegate is a free data retrieval call binding the contract method 0x480005cd.
//
// Solidity: function hasDelegate(address _delegate) view returns(bool)
func (_Argus *ArgusCaller) HasDelegate(opts *bind.CallOpts, _delegate common.Address) (bool, error) {
	var out []interface{}
	err := _Argus.contract.Call(opts, &out, "hasDelegate", _delegate)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasDelegate is a free data retrieval call binding the contract method 0x480005cd.
//
// Solidity: function hasDelegate(address _delegate) view returns(bool)
func (_Argus *ArgusSession) HasDelegate(_delegate common.Address) (bool, error) {
	return _Argus.Contract.HasDelegate(&_Argus.CallOpts, _delegate)
}

// HasDelegate is a free data retrieval call binding the contract method 0x480005cd.
//
// Solidity: function hasDelegate(address _delegate) view returns(bool)
func (_Argus *ArgusCallerSession) HasDelegate(_delegate common.Address) (bool, error) {
	return _Argus.Contract.HasDelegate(&_Argus.CallOpts, _delegate)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Argus *ArgusCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Argus.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Argus *ArgusSession) Owner() (common.Address, error) {
	return _Argus.Contract.Owner(&_Argus.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Argus *ArgusCallerSession) Owner() (common.Address, error) {
	return _Argus.Contract.Owner(&_Argus.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_Argus *ArgusCaller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Argus.contract.Call(opts, &out, "pendingOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_Argus *ArgusSession) PendingOwner() (common.Address, error) {
	return _Argus.Contract.PendingOwner(&_Argus.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_Argus *ArgusCallerSession) PendingOwner() (common.Address, error) {
	return _Argus.Contract.PendingOwner(&_Argus.CallOpts)
}

// RoleManager is a free data retrieval call binding the contract method 0x00435da5.
//
// Solidity: function roleManager() view returns(address)
func (_Argus *ArgusCaller) RoleManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Argus.contract.Call(opts, &out, "roleManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RoleManager is a free data retrieval call binding the contract method 0x00435da5.
//
// Solidity: function roleManager() view returns(address)
func (_Argus *ArgusSession) RoleManager() (common.Address, error) {
	return _Argus.Contract.RoleManager(&_Argus.CallOpts)
}

// RoleManager is a free data retrieval call binding the contract method 0x00435da5.
//
// Solidity: function roleManager() view returns(address)
func (_Argus *ArgusCallerSession) RoleManager() (common.Address, error) {
	return _Argus.Contract.RoleManager(&_Argus.CallOpts)
}

// Safe is a free data retrieval call binding the contract method 0x186f0354.
//
// Solidity: function safe() view returns(address)
func (_Argus *ArgusCaller) Safe(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Argus.contract.Call(opts, &out, "safe")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Safe is a free data retrieval call binding the contract method 0x186f0354.
//
// Solidity: function safe() view returns(address)
func (_Argus *ArgusSession) Safe() (common.Address, error) {
	return _Argus.Contract.Safe(&_Argus.CallOpts)
}

// Safe is a free data retrieval call binding the contract method 0x186f0354.
//
// Solidity: function safe() view returns(address)
func (_Argus *ArgusCallerSession) Safe() (common.Address, error) {
	return _Argus.Contract.Safe(&_Argus.CallOpts)
}

// AcceptOwner is a paid mutator transaction binding the contract method 0xebbc4965.
//
// Solidity: function acceptOwner() returns()
func (_Argus *ArgusTransactor) AcceptOwner(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Argus.contract.Transact(opts, "acceptOwner")
}

// AcceptOwner is a paid mutator transaction binding the contract method 0xebbc4965.
//
// Solidity: function acceptOwner() returns()
func (_Argus *ArgusSession) AcceptOwner() (*types.Transaction, error) {
	return _Argus.Contract.AcceptOwner(&_Argus.TransactOpts)
}

// AcceptOwner is a paid mutator transaction binding the contract method 0xebbc4965.
//
// Solidity: function acceptOwner() returns()
func (_Argus *ArgusTransactorSession) AcceptOwner() (*types.Transaction, error) {
	return _Argus.Contract.AcceptOwner(&_Argus.TransactOpts)
}

// AddDelegate is a paid mutator transaction binding the contract method 0xe71bdf41.
//
// Solidity: function addDelegate(address _delegate) returns()
func (_Argus *ArgusTransactor) AddDelegate(opts *bind.TransactOpts, _delegate common.Address) (*types.Transaction, error) {
	return _Argus.contract.Transact(opts, "addDelegate", _delegate)
}

// AddDelegate is a paid mutator transaction binding the contract method 0xe71bdf41.
//
// Solidity: function addDelegate(address _delegate) returns()
func (_Argus *ArgusSession) AddDelegate(_delegate common.Address) (*types.Transaction, error) {
	return _Argus.Contract.AddDelegate(&_Argus.TransactOpts, _delegate)
}

// AddDelegate is a paid mutator transaction binding the contract method 0xe71bdf41.
//
// Solidity: function addDelegate(address _delegate) returns()
func (_Argus *ArgusTransactorSession) AddDelegate(_delegate common.Address) (*types.Transaction, error) {
	return _Argus.Contract.AddDelegate(&_Argus.TransactOpts, _delegate)
}

// AddDelegates is a paid mutator transaction binding the contract method 0xeb700373.
//
// Solidity: function addDelegates(address[] _delegates) returns()
func (_Argus *ArgusTransactor) AddDelegates(opts *bind.TransactOpts, _delegates []common.Address) (*types.Transaction, error) {
	return _Argus.contract.Transact(opts, "addDelegates", _delegates)
}

// AddDelegates is a paid mutator transaction binding the contract method 0xeb700373.
//
// Solidity: function addDelegates(address[] _delegates) returns()
func (_Argus *ArgusSession) AddDelegates(_delegates []common.Address) (*types.Transaction, error) {
	return _Argus.Contract.AddDelegates(&_Argus.TransactOpts, _delegates)
}

// AddDelegates is a paid mutator transaction binding the contract method 0xeb700373.
//
// Solidity: function addDelegates(address[] _delegates) returns()
func (_Argus *ArgusTransactorSession) AddDelegates(_delegates []common.Address) (*types.Transaction, error) {
	return _Argus.Contract.AddDelegates(&_Argus.TransactOpts, _delegates)
}

// ExecTransaction is a paid mutator transaction binding the contract method 0x47ab9149.
//
// Solidity: function execTransaction((uint256,address,uint256,bytes,bytes,bytes) callData) returns((bool,bytes,bytes) result)
func (_Argus *ArgusTransactor) ExecTransaction(opts *bind.TransactOpts, callData CallData) (*types.Transaction, error) {
	return _Argus.contract.Transact(opts, "execTransaction", callData)
}

// ExecTransaction is a paid mutator transaction binding the contract method 0x47ab9149.
//
// Solidity: function execTransaction((uint256,address,uint256,bytes,bytes,bytes) callData) returns((bool,bytes,bytes) result)
func (_Argus *ArgusSession) ExecTransaction(callData CallData) (*types.Transaction, error) {
	return _Argus.Contract.ExecTransaction(&_Argus.TransactOpts, callData)
}

// ExecTransaction is a paid mutator transaction binding the contract method 0x47ab9149.
//
// Solidity: function execTransaction((uint256,address,uint256,bytes,bytes,bytes) callData) returns((bool,bytes,bytes) result)
func (_Argus *ArgusTransactorSession) ExecTransaction(callData CallData) (*types.Transaction, error) {
	return _Argus.Contract.ExecTransaction(&_Argus.TransactOpts, callData)
}

// ExecTransactions is a paid mutator transaction binding the contract method 0xf5ff8ef2.
//
// Solidity: function execTransactions((uint256,address,uint256,bytes,bytes,bytes)[] callDataList) returns((bool,bytes,bytes)[] resultList)
func (_Argus *ArgusTransactor) ExecTransactions(opts *bind.TransactOpts, callDataList []CallData) (*types.Transaction, error) {
	return _Argus.contract.Transact(opts, "execTransactions", callDataList)
}

// ExecTransactions is a paid mutator transaction binding the contract method 0xf5ff8ef2.
//
// Solidity: function execTransactions((uint256,address,uint256,bytes,bytes,bytes)[] callDataList) returns((bool,bytes,bytes)[] resultList)
func (_Argus *ArgusSession) ExecTransactions(callDataList []CallData) (*types.Transaction, error) {
	return _Argus.Contract.ExecTransactions(&_Argus.TransactOpts, callDataList)
}

// ExecTransactions is a paid mutator transaction binding the contract method 0xf5ff8ef2.
//
// Solidity: function execTransactions((uint256,address,uint256,bytes,bytes,bytes)[] callDataList) returns((bool,bytes,bytes)[] resultList)
func (_Argus *ArgusTransactorSession) ExecTransactions(callDataList []CallData) (*types.Transaction, error) {
	return _Argus.Contract.ExecTransactions(&_Argus.TransactOpts, callDataList)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _owner, address _roleManager, address _authorizer) returns()
func (_Argus *ArgusTransactor) Initialize(opts *bind.TransactOpts, _owner common.Address, _roleManager common.Address, _authorizer common.Address) (*types.Transaction, error) {
	return _Argus.contract.Transact(opts, "initialize", _owner, _roleManager, _authorizer)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _owner, address _roleManager, address _authorizer) returns()
func (_Argus *ArgusSession) Initialize(_owner common.Address, _roleManager common.Address, _authorizer common.Address) (*types.Transaction, error) {
	return _Argus.Contract.Initialize(&_Argus.TransactOpts, _owner, _roleManager, _authorizer)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _owner, address _roleManager, address _authorizer) returns()
func (_Argus *ArgusTransactorSession) Initialize(_owner common.Address, _roleManager common.Address, _authorizer common.Address) (*types.Transaction, error) {
	return _Argus.Contract.Initialize(&_Argus.TransactOpts, _owner, _roleManager, _authorizer)
}

// Initialize0 is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _owner) returns()
func (_Argus *ArgusTransactor) Initialize0(opts *bind.TransactOpts, _owner common.Address) (*types.Transaction, error) {
	return _Argus.contract.Transact(opts, "initialize0", _owner)
}

// Initialize0 is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _owner) returns()
func (_Argus *ArgusSession) Initialize0(_owner common.Address) (*types.Transaction, error) {
	return _Argus.Contract.Initialize0(&_Argus.TransactOpts, _owner)
}

// Initialize0 is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _owner) returns()
func (_Argus *ArgusTransactorSession) Initialize0(_owner common.Address) (*types.Transaction, error) {
	return _Argus.Contract.Initialize0(&_Argus.TransactOpts, _owner)
}

// RemoveDelegate is a paid mutator transaction binding the contract method 0x67e7646f.
//
// Solidity: function removeDelegate(address _delegate) returns()
func (_Argus *ArgusTransactor) RemoveDelegate(opts *bind.TransactOpts, _delegate common.Address) (*types.Transaction, error) {
	return _Argus.contract.Transact(opts, "removeDelegate", _delegate)
}

// RemoveDelegate is a paid mutator transaction binding the contract method 0x67e7646f.
//
// Solidity: function removeDelegate(address _delegate) returns()
func (_Argus *ArgusSession) RemoveDelegate(_delegate common.Address) (*types.Transaction, error) {
	return _Argus.Contract.RemoveDelegate(&_Argus.TransactOpts, _delegate)
}

// RemoveDelegate is a paid mutator transaction binding the contract method 0x67e7646f.
//
// Solidity: function removeDelegate(address _delegate) returns()
func (_Argus *ArgusTransactorSession) RemoveDelegate(_delegate common.Address) (*types.Transaction, error) {
	return _Argus.Contract.RemoveDelegate(&_Argus.TransactOpts, _delegate)
}

// RemoveDelegates is a paid mutator transaction binding the contract method 0x74723555.
//
// Solidity: function removeDelegates(address[] _delegates) returns()
func (_Argus *ArgusTransactor) RemoveDelegates(opts *bind.TransactOpts, _delegates []common.Address) (*types.Transaction, error) {
	return _Argus.contract.Transact(opts, "removeDelegates", _delegates)
}

// RemoveDelegates is a paid mutator transaction binding the contract method 0x74723555.
//
// Solidity: function removeDelegates(address[] _delegates) returns()
func (_Argus *ArgusSession) RemoveDelegates(_delegates []common.Address) (*types.Transaction, error) {
	return _Argus.Contract.RemoveDelegates(&_Argus.TransactOpts, _delegates)
}

// RemoveDelegates is a paid mutator transaction binding the contract method 0x74723555.
//
// Solidity: function removeDelegates(address[] _delegates) returns()
func (_Argus *ArgusTransactorSession) RemoveDelegates(_delegates []common.Address) (*types.Transaction, error) {
	return _Argus.Contract.RemoveDelegates(&_Argus.TransactOpts, _delegates)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Argus *ArgusTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Argus.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Argus *ArgusSession) RenounceOwnership() (*types.Transaction, error) {
	return _Argus.Contract.RenounceOwnership(&_Argus.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Argus *ArgusTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Argus.Contract.RenounceOwnership(&_Argus.TransactOpts)
}

// SetAuthorizer is a paid mutator transaction binding the contract method 0x058a628f.
//
// Solidity: function setAuthorizer(address _authorizer) returns()
func (_Argus *ArgusTransactor) SetAuthorizer(opts *bind.TransactOpts, _authorizer common.Address) (*types.Transaction, error) {
	return _Argus.contract.Transact(opts, "setAuthorizer", _authorizer)
}

// SetAuthorizer is a paid mutator transaction binding the contract method 0x058a628f.
//
// Solidity: function setAuthorizer(address _authorizer) returns()
func (_Argus *ArgusSession) SetAuthorizer(_authorizer common.Address) (*types.Transaction, error) {
	return _Argus.Contract.SetAuthorizer(&_Argus.TransactOpts, _authorizer)
}

// SetAuthorizer is a paid mutator transaction binding the contract method 0x058a628f.
//
// Solidity: function setAuthorizer(address _authorizer) returns()
func (_Argus *ArgusTransactorSession) SetAuthorizer(_authorizer common.Address) (*types.Transaction, error) {
	return _Argus.Contract.SetAuthorizer(&_Argus.TransactOpts, _authorizer)
}

// SetPendingOwner is a paid mutator transaction binding the contract method 0xc42069ec.
//
// Solidity: function setPendingOwner(address to) returns()
func (_Argus *ArgusTransactor) SetPendingOwner(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _Argus.contract.Transact(opts, "setPendingOwner", to)
}

// SetPendingOwner is a paid mutator transaction binding the contract method 0xc42069ec.
//
// Solidity: function setPendingOwner(address to) returns()
func (_Argus *ArgusSession) SetPendingOwner(to common.Address) (*types.Transaction, error) {
	return _Argus.Contract.SetPendingOwner(&_Argus.TransactOpts, to)
}

// SetPendingOwner is a paid mutator transaction binding the contract method 0xc42069ec.
//
// Solidity: function setPendingOwner(address to) returns()
func (_Argus *ArgusTransactorSession) SetPendingOwner(to common.Address) (*types.Transaction, error) {
	return _Argus.Contract.SetPendingOwner(&_Argus.TransactOpts, to)
}

// SetRoleManager is a paid mutator transaction binding the contract method 0xf1d588c5.
//
// Solidity: function setRoleManager(address _roleManager) returns()
func (_Argus *ArgusTransactor) SetRoleManager(opts *bind.TransactOpts, _roleManager common.Address) (*types.Transaction, error) {
	return _Argus.contract.Transact(opts, "setRoleManager", _roleManager)
}

// SetRoleManager is a paid mutator transaction binding the contract method 0xf1d588c5.
//
// Solidity: function setRoleManager(address _roleManager) returns()
func (_Argus *ArgusSession) SetRoleManager(_roleManager common.Address) (*types.Transaction, error) {
	return _Argus.Contract.SetRoleManager(&_Argus.TransactOpts, _roleManager)
}

// SetRoleManager is a paid mutator transaction binding the contract method 0xf1d588c5.
//
// Solidity: function setRoleManager(address _roleManager) returns()
func (_Argus *ArgusTransactorSession) SetRoleManager(_roleManager common.Address) (*types.Transaction, error) {
	return _Argus.Contract.SetRoleManager(&_Argus.TransactOpts, _roleManager)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Argus *ArgusTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Argus.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Argus *ArgusSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Argus.Contract.TransferOwnership(&_Argus.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Argus *ArgusTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Argus.Contract.TransferOwnership(&_Argus.TransactOpts, newOwner)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Argus *ArgusTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Argus.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Argus *ArgusSession) Receive() (*types.Transaction, error) {
	return _Argus.Contract.Receive(&_Argus.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Argus *ArgusTransactorSession) Receive() (*types.Transaction, error) {
	return _Argus.Contract.Receive(&_Argus.TransactOpts)
}

// ArgusAuthorizerSetIterator is returned from FilterAuthorizerSet and is used to iterate over the raw logs and unpacked data for AuthorizerSet events raised by the Argus contract.
type ArgusAuthorizerSetIterator struct {
	Event *ArgusAuthorizerSet // Event containing the contract specifics and raw log

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
func (it *ArgusAuthorizerSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArgusAuthorizerSet)
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
		it.Event = new(ArgusAuthorizerSet)
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
func (it *ArgusAuthorizerSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArgusAuthorizerSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArgusAuthorizerSet represents a AuthorizerSet event raised by the Argus contract.
type ArgusAuthorizerSet struct {
	Authorizer common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAuthorizerSet is a free log retrieval operation binding the contract event 0xb251079a3e59729d2256949e48e44b7959908cdf34789078b6a1462ec3276720.
//
// Solidity: event AuthorizerSet(address indexed authorizer)
func (_Argus *ArgusFilterer) FilterAuthorizerSet(opts *bind.FilterOpts, authorizer []common.Address) (*ArgusAuthorizerSetIterator, error) {

	var authorizerRule []interface{}
	for _, authorizerItem := range authorizer {
		authorizerRule = append(authorizerRule, authorizerItem)
	}

	logs, sub, err := _Argus.contract.FilterLogs(opts, "AuthorizerSet", authorizerRule)
	if err != nil {
		return nil, err
	}
	return &ArgusAuthorizerSetIterator{contract: _Argus.contract, event: "AuthorizerSet", logs: logs, sub: sub}, nil
}

// WatchAuthorizerSet is a free log subscription operation binding the contract event 0xb251079a3e59729d2256949e48e44b7959908cdf34789078b6a1462ec3276720.
//
// Solidity: event AuthorizerSet(address indexed authorizer)
func (_Argus *ArgusFilterer) WatchAuthorizerSet(opts *bind.WatchOpts, sink chan<- *ArgusAuthorizerSet, authorizer []common.Address) (event.Subscription, error) {

	var authorizerRule []interface{}
	for _, authorizerItem := range authorizer {
		authorizerRule = append(authorizerRule, authorizerItem)
	}

	logs, sub, err := _Argus.contract.WatchLogs(opts, "AuthorizerSet", authorizerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArgusAuthorizerSet)
				if err := _Argus.contract.UnpackLog(event, "AuthorizerSet", log); err != nil {
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

// ParseAuthorizerSet is a log parse operation binding the contract event 0xb251079a3e59729d2256949e48e44b7959908cdf34789078b6a1462ec3276720.
//
// Solidity: event AuthorizerSet(address indexed authorizer)
func (_Argus *ArgusFilterer) ParseAuthorizerSet(log types.Log) (*ArgusAuthorizerSet, error) {
	event := new(ArgusAuthorizerSet)
	if err := _Argus.contract.UnpackLog(event, "AuthorizerSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ArgusDelegateAddedIterator is returned from FilterDelegateAdded and is used to iterate over the raw logs and unpacked data for DelegateAdded events raised by the Argus contract.
type ArgusDelegateAddedIterator struct {
	Event *ArgusDelegateAdded // Event containing the contract specifics and raw log

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
func (it *ArgusDelegateAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArgusDelegateAdded)
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
		it.Event = new(ArgusDelegateAdded)
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
func (it *ArgusDelegateAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArgusDelegateAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArgusDelegateAdded represents a DelegateAdded event raised by the Argus contract.
type ArgusDelegateAdded struct {
	Delegate common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDelegateAdded is a free log retrieval operation binding the contract event 0x3be018e2afae1fe2bb23bb3cc019524cfe3872f2863d1c80b7a228c57bc1065a.
//
// Solidity: event DelegateAdded(address indexed delegate)
func (_Argus *ArgusFilterer) FilterDelegateAdded(opts *bind.FilterOpts, delegate []common.Address) (*ArgusDelegateAddedIterator, error) {

	var delegateRule []interface{}
	for _, delegateItem := range delegate {
		delegateRule = append(delegateRule, delegateItem)
	}

	logs, sub, err := _Argus.contract.FilterLogs(opts, "DelegateAdded", delegateRule)
	if err != nil {
		return nil, err
	}
	return &ArgusDelegateAddedIterator{contract: _Argus.contract, event: "DelegateAdded", logs: logs, sub: sub}, nil
}

// WatchDelegateAdded is a free log subscription operation binding the contract event 0x3be018e2afae1fe2bb23bb3cc019524cfe3872f2863d1c80b7a228c57bc1065a.
//
// Solidity: event DelegateAdded(address indexed delegate)
func (_Argus *ArgusFilterer) WatchDelegateAdded(opts *bind.WatchOpts, sink chan<- *ArgusDelegateAdded, delegate []common.Address) (event.Subscription, error) {

	var delegateRule []interface{}
	for _, delegateItem := range delegate {
		delegateRule = append(delegateRule, delegateItem)
	}

	logs, sub, err := _Argus.contract.WatchLogs(opts, "DelegateAdded", delegateRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArgusDelegateAdded)
				if err := _Argus.contract.UnpackLog(event, "DelegateAdded", log); err != nil {
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

// ParseDelegateAdded is a log parse operation binding the contract event 0x3be018e2afae1fe2bb23bb3cc019524cfe3872f2863d1c80b7a228c57bc1065a.
//
// Solidity: event DelegateAdded(address indexed delegate)
func (_Argus *ArgusFilterer) ParseDelegateAdded(log types.Log) (*ArgusDelegateAdded, error) {
	event := new(ArgusDelegateAdded)
	if err := _Argus.contract.UnpackLog(event, "DelegateAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ArgusDelegateRemovedIterator is returned from FilterDelegateRemoved and is used to iterate over the raw logs and unpacked data for DelegateRemoved events raised by the Argus contract.
type ArgusDelegateRemovedIterator struct {
	Event *ArgusDelegateRemoved // Event containing the contract specifics and raw log

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
func (it *ArgusDelegateRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArgusDelegateRemoved)
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
		it.Event = new(ArgusDelegateRemoved)
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
func (it *ArgusDelegateRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArgusDelegateRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArgusDelegateRemoved represents a DelegateRemoved event raised by the Argus contract.
type ArgusDelegateRemoved struct {
	Delegate common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDelegateRemoved is a free log retrieval operation binding the contract event 0x5a362b199d4fb2bd4e51a3436d86df24cb19b1b0890da774e2b54253d53aa4eb.
//
// Solidity: event DelegateRemoved(address indexed delegate)
func (_Argus *ArgusFilterer) FilterDelegateRemoved(opts *bind.FilterOpts, delegate []common.Address) (*ArgusDelegateRemovedIterator, error) {

	var delegateRule []interface{}
	for _, delegateItem := range delegate {
		delegateRule = append(delegateRule, delegateItem)
	}

	logs, sub, err := _Argus.contract.FilterLogs(opts, "DelegateRemoved", delegateRule)
	if err != nil {
		return nil, err
	}
	return &ArgusDelegateRemovedIterator{contract: _Argus.contract, event: "DelegateRemoved", logs: logs, sub: sub}, nil
}

// WatchDelegateRemoved is a free log subscription operation binding the contract event 0x5a362b199d4fb2bd4e51a3436d86df24cb19b1b0890da774e2b54253d53aa4eb.
//
// Solidity: event DelegateRemoved(address indexed delegate)
func (_Argus *ArgusFilterer) WatchDelegateRemoved(opts *bind.WatchOpts, sink chan<- *ArgusDelegateRemoved, delegate []common.Address) (event.Subscription, error) {

	var delegateRule []interface{}
	for _, delegateItem := range delegate {
		delegateRule = append(delegateRule, delegateItem)
	}

	logs, sub, err := _Argus.contract.WatchLogs(opts, "DelegateRemoved", delegateRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArgusDelegateRemoved)
				if err := _Argus.contract.UnpackLog(event, "DelegateRemoved", log); err != nil {
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

// ParseDelegateRemoved is a log parse operation binding the contract event 0x5a362b199d4fb2bd4e51a3436d86df24cb19b1b0890da774e2b54253d53aa4eb.
//
// Solidity: event DelegateRemoved(address indexed delegate)
func (_Argus *ArgusFilterer) ParseDelegateRemoved(log types.Log) (*ArgusDelegateRemoved, error) {
	event := new(ArgusDelegateRemoved)
	if err := _Argus.contract.UnpackLog(event, "DelegateRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ArgusNewOwnerSetIterator is returned from FilterNewOwnerSet and is used to iterate over the raw logs and unpacked data for NewOwnerSet events raised by the Argus contract.
type ArgusNewOwnerSetIterator struct {
	Event *ArgusNewOwnerSet // Event containing the contract specifics and raw log

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
func (it *ArgusNewOwnerSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArgusNewOwnerSet)
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
		it.Event = new(ArgusNewOwnerSet)
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
func (it *ArgusNewOwnerSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArgusNewOwnerSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArgusNewOwnerSet represents a NewOwnerSet event raised by the Argus contract.
type ArgusNewOwnerSet struct {
	Owner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterNewOwnerSet is a free log retrieval operation binding the contract event 0x038720101b9ced74445432ced46c7e5e4c80202669153dd67d226c66a0aa477b.
//
// Solidity: event NewOwnerSet(address indexed owner)
func (_Argus *ArgusFilterer) FilterNewOwnerSet(opts *bind.FilterOpts, owner []common.Address) (*ArgusNewOwnerSetIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Argus.contract.FilterLogs(opts, "NewOwnerSet", ownerRule)
	if err != nil {
		return nil, err
	}
	return &ArgusNewOwnerSetIterator{contract: _Argus.contract, event: "NewOwnerSet", logs: logs, sub: sub}, nil
}

// WatchNewOwnerSet is a free log subscription operation binding the contract event 0x038720101b9ced74445432ced46c7e5e4c80202669153dd67d226c66a0aa477b.
//
// Solidity: event NewOwnerSet(address indexed owner)
func (_Argus *ArgusFilterer) WatchNewOwnerSet(opts *bind.WatchOpts, sink chan<- *ArgusNewOwnerSet, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Argus.contract.WatchLogs(opts, "NewOwnerSet", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArgusNewOwnerSet)
				if err := _Argus.contract.UnpackLog(event, "NewOwnerSet", log); err != nil {
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

// ParseNewOwnerSet is a log parse operation binding the contract event 0x038720101b9ced74445432ced46c7e5e4c80202669153dd67d226c66a0aa477b.
//
// Solidity: event NewOwnerSet(address indexed owner)
func (_Argus *ArgusFilterer) ParseNewOwnerSet(log types.Log) (*ArgusNewOwnerSet, error) {
	event := new(ArgusNewOwnerSet)
	if err := _Argus.contract.UnpackLog(event, "NewOwnerSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ArgusPendingOwnerSetIterator is returned from FilterPendingOwnerSet and is used to iterate over the raw logs and unpacked data for PendingOwnerSet events raised by the Argus contract.
type ArgusPendingOwnerSetIterator struct {
	Event *ArgusPendingOwnerSet // Event containing the contract specifics and raw log

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
func (it *ArgusPendingOwnerSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArgusPendingOwnerSet)
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
		it.Event = new(ArgusPendingOwnerSet)
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
func (it *ArgusPendingOwnerSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArgusPendingOwnerSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArgusPendingOwnerSet represents a PendingOwnerSet event raised by the Argus contract.
type ArgusPendingOwnerSet struct {
	To  common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterPendingOwnerSet is a free log retrieval operation binding the contract event 0x68f49b346b94582a8b5f9d10e3fe3365318fe8f191ff8dce7c59c6cad06b02f5.
//
// Solidity: event PendingOwnerSet(address indexed to)
func (_Argus *ArgusFilterer) FilterPendingOwnerSet(opts *bind.FilterOpts, to []common.Address) (*ArgusPendingOwnerSetIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Argus.contract.FilterLogs(opts, "PendingOwnerSet", toRule)
	if err != nil {
		return nil, err
	}
	return &ArgusPendingOwnerSetIterator{contract: _Argus.contract, event: "PendingOwnerSet", logs: logs, sub: sub}, nil
}

// WatchPendingOwnerSet is a free log subscription operation binding the contract event 0x68f49b346b94582a8b5f9d10e3fe3365318fe8f191ff8dce7c59c6cad06b02f5.
//
// Solidity: event PendingOwnerSet(address indexed to)
func (_Argus *ArgusFilterer) WatchPendingOwnerSet(opts *bind.WatchOpts, sink chan<- *ArgusPendingOwnerSet, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Argus.contract.WatchLogs(opts, "PendingOwnerSet", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArgusPendingOwnerSet)
				if err := _Argus.contract.UnpackLog(event, "PendingOwnerSet", log); err != nil {
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

// ParsePendingOwnerSet is a log parse operation binding the contract event 0x68f49b346b94582a8b5f9d10e3fe3365318fe8f191ff8dce7c59c6cad06b02f5.
//
// Solidity: event PendingOwnerSet(address indexed to)
func (_Argus *ArgusFilterer) ParsePendingOwnerSet(log types.Log) (*ArgusPendingOwnerSet, error) {
	event := new(ArgusPendingOwnerSet)
	if err := _Argus.contract.UnpackLog(event, "PendingOwnerSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ArgusRoleManagerSetIterator is returned from FilterRoleManagerSet and is used to iterate over the raw logs and unpacked data for RoleManagerSet events raised by the Argus contract.
type ArgusRoleManagerSetIterator struct {
	Event *ArgusRoleManagerSet // Event containing the contract specifics and raw log

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
func (it *ArgusRoleManagerSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArgusRoleManagerSet)
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
		it.Event = new(ArgusRoleManagerSet)
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
func (it *ArgusRoleManagerSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArgusRoleManagerSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArgusRoleManagerSet represents a RoleManagerSet event raised by the Argus contract.
type ArgusRoleManagerSet struct {
	RoleManager common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRoleManagerSet is a free log retrieval operation binding the contract event 0x765235f6b1f9df25a0fa901c365a8db93771de0abb8f48ffed12959c5c4d59b9.
//
// Solidity: event RoleManagerSet(address indexed roleManager)
func (_Argus *ArgusFilterer) FilterRoleManagerSet(opts *bind.FilterOpts, roleManager []common.Address) (*ArgusRoleManagerSetIterator, error) {

	var roleManagerRule []interface{}
	for _, roleManagerItem := range roleManager {
		roleManagerRule = append(roleManagerRule, roleManagerItem)
	}

	logs, sub, err := _Argus.contract.FilterLogs(opts, "RoleManagerSet", roleManagerRule)
	if err != nil {
		return nil, err
	}
	return &ArgusRoleManagerSetIterator{contract: _Argus.contract, event: "RoleManagerSet", logs: logs, sub: sub}, nil
}

// WatchRoleManagerSet is a free log subscription operation binding the contract event 0x765235f6b1f9df25a0fa901c365a8db93771de0abb8f48ffed12959c5c4d59b9.
//
// Solidity: event RoleManagerSet(address indexed roleManager)
func (_Argus *ArgusFilterer) WatchRoleManagerSet(opts *bind.WatchOpts, sink chan<- *ArgusRoleManagerSet, roleManager []common.Address) (event.Subscription, error) {

	var roleManagerRule []interface{}
	for _, roleManagerItem := range roleManager {
		roleManagerRule = append(roleManagerRule, roleManagerItem)
	}

	logs, sub, err := _Argus.contract.WatchLogs(opts, "RoleManagerSet", roleManagerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArgusRoleManagerSet)
				if err := _Argus.contract.UnpackLog(event, "RoleManagerSet", log); err != nil {
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

// ParseRoleManagerSet is a log parse operation binding the contract event 0x765235f6b1f9df25a0fa901c365a8db93771de0abb8f48ffed12959c5c4d59b9.
//
// Solidity: event RoleManagerSet(address indexed roleManager)
func (_Argus *ArgusFilterer) ParseRoleManagerSet(log types.Log) (*ArgusRoleManagerSet, error) {
	event := new(ArgusRoleManagerSet)
	if err := _Argus.contract.UnpackLog(event, "RoleManagerSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ArgusTransactionExecutedIterator is returned from FilterTransactionExecuted and is used to iterate over the raw logs and unpacked data for TransactionExecuted events raised by the Argus contract.
type ArgusTransactionExecutedIterator struct {
	Event *ArgusTransactionExecuted // Event containing the contract specifics and raw log

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
func (it *ArgusTransactionExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArgusTransactionExecuted)
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
		it.Event = new(ArgusTransactionExecuted)
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
func (it *ArgusTransactionExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArgusTransactionExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArgusTransactionExecuted represents a TransactionExecuted event raised by the Argus contract.
type ArgusTransactionExecuted struct {
	To          common.Address
	Selector    [4]byte
	Value       *big.Int
	Transaction TransactionData
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTransactionExecuted is a free log retrieval operation binding the contract event 0x000f8708785e90dac3172a799e69be02455af248d971acdd26dc38e07cdb596b.
//
// Solidity: event TransactionExecuted(address indexed to, bytes4 indexed selector, uint256 indexed value, (address,address,uint256,address,uint256,bytes,bytes,bytes) transaction)
func (_Argus *ArgusFilterer) FilterTransactionExecuted(opts *bind.FilterOpts, to []common.Address, selector [][4]byte, value []*big.Int) (*ArgusTransactionExecutedIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var selectorRule []interface{}
	for _, selectorItem := range selector {
		selectorRule = append(selectorRule, selectorItem)
	}
	var valueRule []interface{}
	for _, valueItem := range value {
		valueRule = append(valueRule, valueItem)
	}

	logs, sub, err := _Argus.contract.FilterLogs(opts, "TransactionExecuted", toRule, selectorRule, valueRule)
	if err != nil {
		return nil, err
	}
	return &ArgusTransactionExecutedIterator{contract: _Argus.contract, event: "TransactionExecuted", logs: logs, sub: sub}, nil
}

// WatchTransactionExecuted is a free log subscription operation binding the contract event 0x000f8708785e90dac3172a799e69be02455af248d971acdd26dc38e07cdb596b.
//
// Solidity: event TransactionExecuted(address indexed to, bytes4 indexed selector, uint256 indexed value, (address,address,uint256,address,uint256,bytes,bytes,bytes) transaction)
func (_Argus *ArgusFilterer) WatchTransactionExecuted(opts *bind.WatchOpts, sink chan<- *ArgusTransactionExecuted, to []common.Address, selector [][4]byte, value []*big.Int) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var selectorRule []interface{}
	for _, selectorItem := range selector {
		selectorRule = append(selectorRule, selectorItem)
	}
	var valueRule []interface{}
	for _, valueItem := range value {
		valueRule = append(valueRule, valueItem)
	}

	logs, sub, err := _Argus.contract.WatchLogs(opts, "TransactionExecuted", toRule, selectorRule, valueRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArgusTransactionExecuted)
				if err := _Argus.contract.UnpackLog(event, "TransactionExecuted", log); err != nil {
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

// ParseTransactionExecuted is a log parse operation binding the contract event 0x000f8708785e90dac3172a799e69be02455af248d971acdd26dc38e07cdb596b.
//
// Solidity: event TransactionExecuted(address indexed to, bytes4 indexed selector, uint256 indexed value, (address,address,uint256,address,uint256,bytes,bytes,bytes) transaction)
func (_Argus *ArgusFilterer) ParseTransactionExecuted(log types.Log) (*ArgusTransactionExecuted, error) {
	event := new(ArgusTransactionExecuted)
	if err := _Argus.contract.UnpackLog(event, "TransactionExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
