// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// VerifierABI is the input ABI used to generate the binding from.
const VerifierABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"getStatus\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"destroy\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"hash\",\"type\":\"string\"},{\"name\":\"version\",\"type\":\"int256\"}],\"name\":\"verify\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"reqNode\",\"type\":\"address\"},{\"name\":\"hash\",\"type\":\"string\"},{\"name\":\"version\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// VerifierBin is the compiled bytecode used for deploying new contracts.
const VerifierBin = `0x6080604052600660055534801561001557600080fd5b5060405161060b38038061060b8339810180604052606081101561003857600080fd5b81516020830180519193928301929164010000000081111561005957600080fd5b8201602081018481111561006c57600080fd5b815164010000000081118282018710171561008657600080fd5b5050602091820151600080546001600160a01b0319166001600160a01b038816179055815191945092506100c091600191908501906100da565b5060025550506003805460ff191690556000600455610175565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061011b57805160ff1916838001178555610148565b82800160010185558215610148579182015b8281111561014857825182559160200191906001019061012d565b50610154929150610158565b5090565b61017291905b80821115610154576000815560010161015e565b90565b610487806101846000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c80634e69d5601461004657806383197ef0146100c35780639bd08f91146100cd575b600080fd5b61004e610175565b6040805160208082528351818301528351919283929083019185019080838360005b83811015610088578181015183820152602001610070565b50505050905090810190601f1680156100b55780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6100cb610204565b005b61004e600480360360408110156100e357600080fd5b8101906020810181356401000000008111156100fe57600080fd5b82018360208201111561011057600080fd5b8035906020019184600183028401116401000000008311171561013257600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295505091359250610227915050565b60035460609060ff166101dc57600554600454136101b257604051806040016040528060048152602001600160e21b6315d85a5d028152506101d7565b604051806040016040528060088152602001600160c21b6715995c9a599a5959028152505b6101ff565b604051806040016040528060068152602001600160d01b65557064617465028152505b905090565b6000546001600160a01b0316331415610225576000546001600160a01b0316ff5b565b60606002548214156102ff5760018054604080516020600284861615610100026000190190941693909304601f81018490048402820184019092528181526102ca938793919290918301828280156102c05780601f10610295576101008083540402835291602001916102c0565b820191906000526020600020905b8154815290600101906020018083116102a357829003601f168201915b5050505050610364565b156102d9576004805460010190555b506040805180820190915260068152600160d01b655468616e6b7302602082015261035e565b60025482131561033c57506003805460ff191660011790556040805180820190915260068152600160d01b655468616e6b7302602082015261035e565b506040805180820190915260068152600160d01b655570646174650260208201525b92915050565b600081518351146103775750600061035e565b816040516020018082805190602001908083835b602083106103aa5780518252601f19909201916020918201910161038b565b6001836020036101000a03801982511681845116808217855250505050505090500191505060405160208183030381529060405280519060200120836040516020018082805190602001908083835b602083106104185780518252601f1990920191602091820191016103f9565b6001836020036101000a0380198251168184511680821785525050505050509050019150506040516020818303038152906040528051906020012014905061035e56fea165627a7a72305820c549f4c6e7cfddcda0655f390eca7e1d8560173bb4f052d200e295ba0428d2c00029`

// DeployVerifier deploys a new Ethereum contract, binding an instance of Verifier to it.
func DeployVerifier(auth *bind.TransactOpts, backend bind.ContractBackend, reqNode common.Address, hash string, version *big.Int) (common.Address, *types.Transaction, *Verifier, error) {
	parsed, err := abi.JSON(strings.NewReader(VerifierABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(VerifierBin), backend, reqNode, hash, version)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Verifier{VerifierCaller: VerifierCaller{contract: contract}, VerifierTransactor: VerifierTransactor{contract: contract}, VerifierFilterer: VerifierFilterer{contract: contract}}, nil
}

// Verifier is an auto generated Go binding around an Ethereum contract.
type Verifier struct {
	VerifierCaller     // Read-only binding to the contract
	VerifierTransactor // Write-only binding to the contract
	VerifierFilterer   // Log filterer for contract events
}

// VerifierCaller is an auto generated read-only Go binding around an Ethereum contract.
type VerifierCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VerifierTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VerifierTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VerifierFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VerifierFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VerifierSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VerifierSession struct {
	Contract     *Verifier         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VerifierCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VerifierCallerSession struct {
	Contract *VerifierCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// VerifierTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VerifierTransactorSession struct {
	Contract     *VerifierTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// VerifierRaw is an auto generated low-level Go binding around an Ethereum contract.
type VerifierRaw struct {
	Contract *Verifier // Generic contract binding to access the raw methods on
}

// VerifierCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VerifierCallerRaw struct {
	Contract *VerifierCaller // Generic read-only contract binding to access the raw methods on
}

// VerifierTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VerifierTransactorRaw struct {
	Contract *VerifierTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVerifier creates a new instance of Verifier, bound to a specific deployed contract.
func NewVerifier(address common.Address, backend bind.ContractBackend) (*Verifier, error) {
	contract, err := bindVerifier(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Verifier{VerifierCaller: VerifierCaller{contract: contract}, VerifierTransactor: VerifierTransactor{contract: contract}, VerifierFilterer: VerifierFilterer{contract: contract}}, nil
}

// NewVerifierCaller creates a new read-only instance of Verifier, bound to a specific deployed contract.
func NewVerifierCaller(address common.Address, caller bind.ContractCaller) (*VerifierCaller, error) {
	contract, err := bindVerifier(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VerifierCaller{contract: contract}, nil
}

// NewVerifierTransactor creates a new write-only instance of Verifier, bound to a specific deployed contract.
func NewVerifierTransactor(address common.Address, transactor bind.ContractTransactor) (*VerifierTransactor, error) {
	contract, err := bindVerifier(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VerifierTransactor{contract: contract}, nil
}

// NewVerifierFilterer creates a new log filterer instance of Verifier, bound to a specific deployed contract.
func NewVerifierFilterer(address common.Address, filterer bind.ContractFilterer) (*VerifierFilterer, error) {
	contract, err := bindVerifier(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VerifierFilterer{contract: contract}, nil
}

// bindVerifier binds a generic wrapper to an already deployed contract.
func bindVerifier(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VerifierABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Verifier *VerifierRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Verifier.Contract.VerifierCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Verifier *VerifierRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Verifier.Contract.VerifierTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Verifier *VerifierRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Verifier.Contract.VerifierTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Verifier *VerifierCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Verifier.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Verifier *VerifierTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Verifier.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Verifier *VerifierTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Verifier.Contract.contract.Transact(opts, method, params...)
}

// GetStatus is a free data retrieval call binding the contract method 0x4e69d560.
//
// Solidity: function getStatus() constant returns(string)
func (_Verifier *VerifierCaller) GetStatus(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Verifier.contract.Call(opts, out, "getStatus")
	return *ret0, err
}

// GetStatus is a free data retrieval call binding the contract method 0x4e69d560.
//
// Solidity: function getStatus() constant returns(string)
func (_Verifier *VerifierSession) GetStatus() (string, error) {
	return _Verifier.Contract.GetStatus(&_Verifier.CallOpts)
}

// GetStatus is a free data retrieval call binding the contract method 0x4e69d560.
//
// Solidity: function getStatus() constant returns(string)
func (_Verifier *VerifierCallerSession) GetStatus() (string, error) {
	return _Verifier.Contract.GetStatus(&_Verifier.CallOpts)
}

// Destroy is a paid mutator transaction binding the contract method 0x83197ef0.
//
// Solidity: function destroy() returns()
func (_Verifier *VerifierTransactor) Destroy(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Verifier.contract.Transact(opts, "destroy")
}

// Destroy is a paid mutator transaction binding the contract method 0x83197ef0.
//
// Solidity: function destroy() returns()
func (_Verifier *VerifierSession) Destroy() (*types.Transaction, error) {
	return _Verifier.Contract.Destroy(&_Verifier.TransactOpts)
}

// Destroy is a paid mutator transaction binding the contract method 0x83197ef0.
//
// Solidity: function destroy() returns()
func (_Verifier *VerifierTransactorSession) Destroy() (*types.Transaction, error) {
	return _Verifier.Contract.Destroy(&_Verifier.TransactOpts)
}

// Verify is a paid mutator transaction binding the contract method 0x9bd08f91.
//
// Solidity: function verify(string hash, int256 version) returns(string)
func (_Verifier *VerifierTransactor) Verify(opts *bind.TransactOpts, hash string, version *big.Int) (*types.Transaction, error) {
	return _Verifier.contract.Transact(opts, "verify", hash, version)
}

// Verify is a paid mutator transaction binding the contract method 0x9bd08f91.
//
// Solidity: function verify(string hash, int256 version) returns(string)
func (_Verifier *VerifierSession) Verify(hash string, version *big.Int) (*types.Transaction, error) {
	return _Verifier.Contract.Verify(&_Verifier.TransactOpts, hash, version)
}

// Verify is a paid mutator transaction binding the contract method 0x9bd08f91.
//
// Solidity: function verify(string hash, int256 version) returns(string)
func (_Verifier *VerifierTransactorSession) Verify(hash string, version *big.Int) (*types.Transaction, error) {
	return _Verifier.Contract.Verify(&_Verifier.TransactOpts, hash, version)
}
