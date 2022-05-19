// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package api

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

// MUUVMetaData contains all meta data concerning the MUUV contract.
var MUUVMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ab\",\"outputs\":[{\"internalType\":\"contractAntiBot\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"abEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_ab\",\"type\":\"address\"}],\"name\":\"setABAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_enabled\",\"type\":\"bool\"}],\"name\":\"setABEnabled\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_dev\",\"type\":\"address\"}],\"name\":\"setDevAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b506040518060400160405280600a81526020016926aaaaab102a37b5b2b760b11b8152506040518060400160405280600481526020016326aaaaab60e11b8152506200006c62000066620000d960201b60201c565b620000dd565b81516200008190600490602085019062000215565b5080516200009790600590602084019062000215565b505050620000c1620000ae620000d960201b60201c565b6b02e87669c308736a040000006200012d565b600780546001600160a01b031916331790556200031e565b3390565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6001600160a01b038216620001885760405162461bcd60e51b815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f206164647265737300604482015260640160405180910390fd5b80600360008282546200019c9190620002bb565b90915550506001600160a01b03821660009081526001602052604081208054839290620001cb908490620002bb565b90915550506040518181526001600160a01b038316906000907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9060200160405180910390a35050565b8280546200022390620002e2565b90600052602060002090601f01602090048101928262000247576000855562000292565b82601f106200026257805160ff191683800117855562000292565b8280016001018555821562000292579182015b828111156200029257825182559160200191906001019062000275565b50620002a0929150620002a4565b5090565b5b80821115620002a05760008155600101620002a5565b60008219821115620002dd57634e487b7160e01b600052601160045260246000fd5b500190565b600181811c90821680620002f757607f821691505b6020821081036200031857634e487b7160e01b600052602260045260246000fd5b50919050565b610da2806200032e6000396000f3fe608060405234801561001057600080fd5b50600436106101215760003560e01c806373e2423f116100ad578063a9059cbb11610071578063a9059cbb1461025f578063b473b03614610272578063d0d41fe114610285578063dd62ed3e14610298578063f2fde38b146102d157600080fd5b806373e2423f1461020c5780637568f8c7146102205780638da5cb5b1461023357806395d89b4114610244578063a457c2d71461024c57600080fd5b806323b872dd116100f457806323b872dd146101a4578063313ce567146101b757806339509351146101c657806370a08231146101d9578063715018a61461020257600080fd5b806306fdde0314610126578063095ea7b31461014457806318160ddd146101675780631fe5d6e314610179575b600080fd5b61012e6102e4565b60405161013b9190610b51565b60405180910390f35b610157610152366004610bc2565b610376565b604051901515815260200161013b565b6003545b60405190815260200161013b565b60065461018c906001600160a01b031681565b6040516001600160a01b03909116815260200161013b565b6101576101b2366004610bec565b61038e565b6040516012815260200161013b565b6101576101d4366004610bc2565b6103b2565b61016b6101e7366004610c28565b6001600160a01b031660009081526001602052604090205490565b61020a6103f1565b005b60065461015790600160a01b900460ff1681565b61020a61022e366004610c4a565b610430565b6000546001600160a01b031661018c565b61012e6104c5565b61015761025a366004610bc2565b6104d4565b61015761026d366004610bc2565b610566565b61020a610280366004610c28565b610574565b61020a610293366004610c28565b6105c0565b61016b6102a6366004610c6c565b6001600160a01b03918216600090815260026020908152604080832093909416825291909152205490565b61020a6102df366004610c28565b61060c565b6060600480546102f390610c9f565b80601f016020809104026020016040519081016040528092919081815260200182805461031f90610c9f565b801561036c5780601f106103415761010080835404028352916020019161036c565b820191906000526020600020905b81548152906001019060200180831161034f57829003601f168201915b5050505050905090565b6000336103848185856106a7565b5060019392505050565b60003361039c8582856107cb565b6103a785858561085d565b506001949350505050565b3360008181526002602090815260408083206001600160a01b038716845290915281205490919061038490829086906103ec908790610cef565b6106a7565b6000546001600160a01b031633146104245760405162461bcd60e51b815260040161041b90610d07565b60405180910390fd5b61042e6000610933565b565b6000546001600160a01b0316331461045a5760405162461bcd60e51b815260040161041b90610d07565b6006546001600160a01b03166104a75760405162461bcd60e51b8152602060048201526012602482015271185b9d1a48189bdd081b9bdd08199bdd5b9960721b604482015260640161041b565b60068054911515600160a01b0260ff60a01b19909216919091179055565b6060600580546102f390610c9f565b3360008181526002602090815260408083206001600160a01b0387168452909152812054909190838110156105595760405162461bcd60e51b815260206004820152602560248201527f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f77604482015264207a65726f60d81b606482015260840161041b565b6103a782868684036106a7565b60003361038481858561085d565b6000546001600160a01b0316331461059e5760405162461bcd60e51b815260040161041b90610d07565b600680546001600160a01b0319166001600160a01b0392909216919091179055565b6000546001600160a01b031633146105ea5760405162461bcd60e51b815260040161041b90610d07565b600780546001600160a01b0319166001600160a01b0392909216919091179055565b6000546001600160a01b031633146106365760405162461bcd60e51b815260040161041b90610d07565b6001600160a01b03811661069b5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b606482015260840161041b565b6106a481610933565b50565b6001600160a01b0383166107095760405162461bcd60e51b8152602060048201526024808201527f45524332303a20617070726f76652066726f6d20746865207a65726f206164646044820152637265737360e01b606482015260840161041b565b6001600160a01b03821661076a5760405162461bcd60e51b815260206004820152602260248201527f45524332303a20617070726f766520746f20746865207a65726f206164647265604482015261737360f01b606482015260840161041b565b6001600160a01b0383811660008181526002602090815260408083209487168084529482529182902085905590518481527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925910160405180910390a3505050565b6001600160a01b038381166000908152600260209081526040808320938616835292905220546000198114610857578181101561084a5760405162461bcd60e51b815260206004820152601d60248201527f45524332303a20696e73756666696369656e7420616c6c6f77616e6365000000604482015260640161041b565b61085784848484036106a7565b50505050565b600654600160a01b900460ff161561092357600654604051637e2f3afd60e01b81526001600160a01b0385811660048301528481166024830152604482018490526000921690637e2f3afd906064016020604051808303816000875af11580156108cb573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108ef9190610d3c565b9050801561090f5760075461090f9085906001600160a01b031683610983565b610857848461091e8486610d55565b610983565b61092e838383610983565b505050565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6001600160a01b0383166109e75760405162461bcd60e51b815260206004820152602560248201527f45524332303a207472616e736665722066726f6d20746865207a65726f206164604482015264647265737360d81b606482015260840161041b565b6001600160a01b038216610a495760405162461bcd60e51b815260206004820152602360248201527f45524332303a207472616e7366657220746f20746865207a65726f206164647260448201526265737360e81b606482015260840161041b565b6001600160a01b03831660009081526001602052604090205481811015610ac15760405162461bcd60e51b815260206004820152602660248201527f45524332303a207472616e7366657220616d6f756e7420657863656564732062604482015265616c616e636560d01b606482015260840161041b565b6001600160a01b03808516600090815260016020526040808220858503905591851681529081208054849290610af8908490610cef565b92505081905550826001600160a01b0316846001600160a01b03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef84604051610b4491815260200190565b60405180910390a3610857565b600060208083528351808285015260005b81811015610b7e57858101830151858201604001528201610b62565b81811115610b90576000604083870101525b50601f01601f1916929092016040019392505050565b80356001600160a01b0381168114610bbd57600080fd5b919050565b60008060408385031215610bd557600080fd5b610bde83610ba6565b946020939093013593505050565b600080600060608486031215610c0157600080fd5b610c0a84610ba6565b9250610c1860208501610ba6565b9150604084013590509250925092565b600060208284031215610c3a57600080fd5b610c4382610ba6565b9392505050565b600060208284031215610c5c57600080fd5b81358015158114610c4357600080fd5b60008060408385031215610c7f57600080fd5b610c8883610ba6565b9150610c9660208401610ba6565b90509250929050565b600181811c90821680610cb357607f821691505b602082108103610cd357634e487b7160e01b600052602260045260246000fd5b50919050565b634e487b7160e01b600052601160045260246000fd5b60008219821115610d0257610d02610cd9565b500190565b6020808252818101527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604082015260600190565b600060208284031215610d4e57600080fd5b5051919050565b600082821015610d6757610d67610cd9565b50039056fea26469706673582212207bd01544370b85eaf8e3c3f221011f203c494553271336ebb3e9338f2bb7011464736f6c634300080d0033",
}

// MUUVABI is the input ABI used to generate the binding from.
// Deprecated: Use MUUVMetaData.ABI instead.
var MUUVABI = MUUVMetaData.ABI

// MUUVBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MUUVMetaData.Bin instead.
var MUUVBin = MUUVMetaData.Bin

// DeployMUUV deploys a new Ethereum contract, binding an instance of MUUV to it.
func DeployMUUV(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MUUV, error) {
	parsed, err := MUUVMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MUUVBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MUUV{MUUVCaller: MUUVCaller{contract: contract}, MUUVTransactor: MUUVTransactor{contract: contract}, MUUVFilterer: MUUVFilterer{contract: contract}}, nil
}

// MUUV is an auto generated Go binding around an Ethereum contract.
type MUUV struct {
	MUUVCaller     // Read-only binding to the contract
	MUUVTransactor // Write-only binding to the contract
	MUUVFilterer   // Log filterer for contract events
}

// MUUVCaller is an auto generated read-only Go binding around an Ethereum contract.
type MUUVCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MUUVTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MUUVTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MUUVFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MUUVFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MUUVSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MUUVSession struct {
	Contract     *MUUV             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MUUVCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MUUVCallerSession struct {
	Contract *MUUVCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// MUUVTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MUUVTransactorSession struct {
	Contract     *MUUVTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MUUVRaw is an auto generated low-level Go binding around an Ethereum contract.
type MUUVRaw struct {
	Contract *MUUV // Generic contract binding to access the raw methods on
}

// MUUVCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MUUVCallerRaw struct {
	Contract *MUUVCaller // Generic read-only contract binding to access the raw methods on
}

// MUUVTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MUUVTransactorRaw struct {
	Contract *MUUVTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMUUV creates a new instance of MUUV, bound to a specific deployed contract.
func NewMUUV(address common.Address, backend bind.ContractBackend) (*MUUV, error) {
	contract, err := bindMUUV(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MUUV{MUUVCaller: MUUVCaller{contract: contract}, MUUVTransactor: MUUVTransactor{contract: contract}, MUUVFilterer: MUUVFilterer{contract: contract}}, nil
}

// NewMUUVCaller creates a new read-only instance of MUUV, bound to a specific deployed contract.
func NewMUUVCaller(address common.Address, caller bind.ContractCaller) (*MUUVCaller, error) {
	contract, err := bindMUUV(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MUUVCaller{contract: contract}, nil
}

// NewMUUVTransactor creates a new write-only instance of MUUV, bound to a specific deployed contract.
func NewMUUVTransactor(address common.Address, transactor bind.ContractTransactor) (*MUUVTransactor, error) {
	contract, err := bindMUUV(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MUUVTransactor{contract: contract}, nil
}

// NewMUUVFilterer creates a new log filterer instance of MUUV, bound to a specific deployed contract.
func NewMUUVFilterer(address common.Address, filterer bind.ContractFilterer) (*MUUVFilterer, error) {
	contract, err := bindMUUV(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MUUVFilterer{contract: contract}, nil
}

// bindMUUV binds a generic wrapper to an already deployed contract.
func bindMUUV(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MUUVABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MUUV *MUUVRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MUUV.Contract.MUUVCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MUUV *MUUVRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MUUV.Contract.MUUVTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MUUV *MUUVRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MUUV.Contract.MUUVTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MUUV *MUUVCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MUUV.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MUUV *MUUVTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MUUV.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MUUV *MUUVTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MUUV.Contract.contract.Transact(opts, method, params...)
}

// Ab is a free data retrieval call binding the contract method 0x1fe5d6e3.
//
// Solidity: function ab() view returns(address)
func (_MUUV *MUUVCaller) Ab(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MUUV.contract.Call(opts, &out, "ab")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Ab is a free data retrieval call binding the contract method 0x1fe5d6e3.
//
// Solidity: function ab() view returns(address)
func (_MUUV *MUUVSession) Ab() (common.Address, error) {
	return _MUUV.Contract.Ab(&_MUUV.CallOpts)
}

// Ab is a free data retrieval call binding the contract method 0x1fe5d6e3.
//
// Solidity: function ab() view returns(address)
func (_MUUV *MUUVCallerSession) Ab() (common.Address, error) {
	return _MUUV.Contract.Ab(&_MUUV.CallOpts)
}

// AbEnabled is a free data retrieval call binding the contract method 0x73e2423f.
//
// Solidity: function abEnabled() view returns(bool)
func (_MUUV *MUUVCaller) AbEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _MUUV.contract.Call(opts, &out, "abEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AbEnabled is a free data retrieval call binding the contract method 0x73e2423f.
//
// Solidity: function abEnabled() view returns(bool)
func (_MUUV *MUUVSession) AbEnabled() (bool, error) {
	return _MUUV.Contract.AbEnabled(&_MUUV.CallOpts)
}

// AbEnabled is a free data retrieval call binding the contract method 0x73e2423f.
//
// Solidity: function abEnabled() view returns(bool)
func (_MUUV *MUUVCallerSession) AbEnabled() (bool, error) {
	return _MUUV.Contract.AbEnabled(&_MUUV.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_MUUV *MUUVCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MUUV.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_MUUV *MUUVSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _MUUV.Contract.Allowance(&_MUUV.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_MUUV *MUUVCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _MUUV.Contract.Allowance(&_MUUV.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_MUUV *MUUVCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MUUV.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_MUUV *MUUVSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _MUUV.Contract.BalanceOf(&_MUUV.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_MUUV *MUUVCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _MUUV.Contract.BalanceOf(&_MUUV.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_MUUV *MUUVCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _MUUV.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_MUUV *MUUVSession) Decimals() (uint8, error) {
	return _MUUV.Contract.Decimals(&_MUUV.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_MUUV *MUUVCallerSession) Decimals() (uint8, error) {
	return _MUUV.Contract.Decimals(&_MUUV.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_MUUV *MUUVCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _MUUV.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_MUUV *MUUVSession) Name() (string, error) {
	return _MUUV.Contract.Name(&_MUUV.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_MUUV *MUUVCallerSession) Name() (string, error) {
	return _MUUV.Contract.Name(&_MUUV.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MUUV *MUUVCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MUUV.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MUUV *MUUVSession) Owner() (common.Address, error) {
	return _MUUV.Contract.Owner(&_MUUV.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MUUV *MUUVCallerSession) Owner() (common.Address, error) {
	return _MUUV.Contract.Owner(&_MUUV.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_MUUV *MUUVCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _MUUV.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_MUUV *MUUVSession) Symbol() (string, error) {
	return _MUUV.Contract.Symbol(&_MUUV.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_MUUV *MUUVCallerSession) Symbol() (string, error) {
	return _MUUV.Contract.Symbol(&_MUUV.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_MUUV *MUUVCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MUUV.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_MUUV *MUUVSession) TotalSupply() (*big.Int, error) {
	return _MUUV.Contract.TotalSupply(&_MUUV.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_MUUV *MUUVCallerSession) TotalSupply() (*big.Int, error) {
	return _MUUV.Contract.TotalSupply(&_MUUV.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_MUUV *MUUVTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MUUV.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_MUUV *MUUVSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MUUV.Contract.Approve(&_MUUV.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_MUUV *MUUVTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MUUV.Contract.Approve(&_MUUV.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_MUUV *MUUVTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _MUUV.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_MUUV *MUUVSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _MUUV.Contract.DecreaseAllowance(&_MUUV.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_MUUV *MUUVTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _MUUV.Contract.DecreaseAllowance(&_MUUV.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_MUUV *MUUVTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _MUUV.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_MUUV *MUUVSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _MUUV.Contract.IncreaseAllowance(&_MUUV.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_MUUV *MUUVTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _MUUV.Contract.IncreaseAllowance(&_MUUV.TransactOpts, spender, addedValue)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MUUV *MUUVTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MUUV.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MUUV *MUUVSession) RenounceOwnership() (*types.Transaction, error) {
	return _MUUV.Contract.RenounceOwnership(&_MUUV.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MUUV *MUUVTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _MUUV.Contract.RenounceOwnership(&_MUUV.TransactOpts)
}

// SetABAddress is a paid mutator transaction binding the contract method 0xb473b036.
//
// Solidity: function setABAddress(address _ab) returns()
func (_MUUV *MUUVTransactor) SetABAddress(opts *bind.TransactOpts, _ab common.Address) (*types.Transaction, error) {
	return _MUUV.contract.Transact(opts, "setABAddress", _ab)
}

// SetABAddress is a paid mutator transaction binding the contract method 0xb473b036.
//
// Solidity: function setABAddress(address _ab) returns()
func (_MUUV *MUUVSession) SetABAddress(_ab common.Address) (*types.Transaction, error) {
	return _MUUV.Contract.SetABAddress(&_MUUV.TransactOpts, _ab)
}

// SetABAddress is a paid mutator transaction binding the contract method 0xb473b036.
//
// Solidity: function setABAddress(address _ab) returns()
func (_MUUV *MUUVTransactorSession) SetABAddress(_ab common.Address) (*types.Transaction, error) {
	return _MUUV.Contract.SetABAddress(&_MUUV.TransactOpts, _ab)
}

// SetABEnabled is a paid mutator transaction binding the contract method 0x7568f8c7.
//
// Solidity: function setABEnabled(bool _enabled) returns()
func (_MUUV *MUUVTransactor) SetABEnabled(opts *bind.TransactOpts, _enabled bool) (*types.Transaction, error) {
	return _MUUV.contract.Transact(opts, "setABEnabled", _enabled)
}

// SetABEnabled is a paid mutator transaction binding the contract method 0x7568f8c7.
//
// Solidity: function setABEnabled(bool _enabled) returns()
func (_MUUV *MUUVSession) SetABEnabled(_enabled bool) (*types.Transaction, error) {
	return _MUUV.Contract.SetABEnabled(&_MUUV.TransactOpts, _enabled)
}

// SetABEnabled is a paid mutator transaction binding the contract method 0x7568f8c7.
//
// Solidity: function setABEnabled(bool _enabled) returns()
func (_MUUV *MUUVTransactorSession) SetABEnabled(_enabled bool) (*types.Transaction, error) {
	return _MUUV.Contract.SetABEnabled(&_MUUV.TransactOpts, _enabled)
}

// SetDevAddress is a paid mutator transaction binding the contract method 0xd0d41fe1.
//
// Solidity: function setDevAddress(address _dev) returns()
func (_MUUV *MUUVTransactor) SetDevAddress(opts *bind.TransactOpts, _dev common.Address) (*types.Transaction, error) {
	return _MUUV.contract.Transact(opts, "setDevAddress", _dev)
}

// SetDevAddress is a paid mutator transaction binding the contract method 0xd0d41fe1.
//
// Solidity: function setDevAddress(address _dev) returns()
func (_MUUV *MUUVSession) SetDevAddress(_dev common.Address) (*types.Transaction, error) {
	return _MUUV.Contract.SetDevAddress(&_MUUV.TransactOpts, _dev)
}

// SetDevAddress is a paid mutator transaction binding the contract method 0xd0d41fe1.
//
// Solidity: function setDevAddress(address _dev) returns()
func (_MUUV *MUUVTransactorSession) SetDevAddress(_dev common.Address) (*types.Transaction, error) {
	return _MUUV.Contract.SetDevAddress(&_MUUV.TransactOpts, _dev)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_MUUV *MUUVTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MUUV.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_MUUV *MUUVSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MUUV.Contract.Transfer(&_MUUV.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_MUUV *MUUVTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MUUV.Contract.Transfer(&_MUUV.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_MUUV *MUUVTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MUUV.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_MUUV *MUUVSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MUUV.Contract.TransferFrom(&_MUUV.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_MUUV *MUUVTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MUUV.Contract.TransferFrom(&_MUUV.TransactOpts, from, to, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MUUV *MUUVTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _MUUV.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MUUV *MUUVSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MUUV.Contract.TransferOwnership(&_MUUV.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MUUV *MUUVTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MUUV.Contract.TransferOwnership(&_MUUV.TransactOpts, newOwner)
}

// MUUVApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the MUUV contract.
type MUUVApprovalIterator struct {
	Event *MUUVApproval // Event containing the contract specifics and raw log

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
func (it *MUUVApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MUUVApproval)
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
		it.Event = new(MUUVApproval)
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
func (it *MUUVApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MUUVApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MUUVApproval represents a Approval event raised by the MUUV contract.
type MUUVApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_MUUV *MUUVFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*MUUVApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _MUUV.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &MUUVApprovalIterator{contract: _MUUV.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_MUUV *MUUVFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *MUUVApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _MUUV.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MUUVApproval)
				if err := _MUUV.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_MUUV *MUUVFilterer) ParseApproval(log types.Log) (*MUUVApproval, error) {
	event := new(MUUVApproval)
	if err := _MUUV.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MUUVOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the MUUV contract.
type MUUVOwnershipTransferredIterator struct {
	Event *MUUVOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *MUUVOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MUUVOwnershipTransferred)
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
		it.Event = new(MUUVOwnershipTransferred)
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
func (it *MUUVOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MUUVOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MUUVOwnershipTransferred represents a OwnershipTransferred event raised by the MUUV contract.
type MUUVOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MUUV *MUUVFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*MUUVOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MUUV.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MUUVOwnershipTransferredIterator{contract: _MUUV.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MUUV *MUUVFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MUUVOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MUUV.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MUUVOwnershipTransferred)
				if err := _MUUV.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MUUV *MUUVFilterer) ParseOwnershipTransferred(log types.Log) (*MUUVOwnershipTransferred, error) {
	event := new(MUUVOwnershipTransferred)
	if err := _MUUV.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MUUVTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the MUUV contract.
type MUUVTransferIterator struct {
	Event *MUUVTransfer // Event containing the contract specifics and raw log

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
func (it *MUUVTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MUUVTransfer)
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
		it.Event = new(MUUVTransfer)
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
func (it *MUUVTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MUUVTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MUUVTransfer represents a Transfer event raised by the MUUV contract.
type MUUVTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_MUUV *MUUVFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*MUUVTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MUUV.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &MUUVTransferIterator{contract: _MUUV.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_MUUV *MUUVFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *MUUVTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MUUV.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MUUVTransfer)
				if err := _MUUV.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_MUUV *MUUVFilterer) ParseTransfer(log types.Log) (*MUUVTransfer, error) {
	event := new(MUUVTransfer)
	if err := _MUUV.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
