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

// VerificationContractGeneratorABI is the input ABI used to generate the binding from.
const VerificationContractGeneratorABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"hash\",\"type\":\"string\"},{\"name\":\"version\",\"type\":\"int256\"}],\"name\":\"generateVerificationContract\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// VerificationContractGeneratorBin is the compiled bytecode used for deploying new contracts.
const VerificationContractGeneratorBin = `0x608060405234801561001057600080fd5b5061083e806100206000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c8063a34b811c14610030575b600080fd5b6100d86004803603604081101561004657600080fd5b81019060208101813564010000000081111561006157600080fd5b82018360208201111561007357600080fd5b8035906020019184600183028401116401000000008311171561009557600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955050913592506100f4915050565b604080516001600160a01b039092168252519081900360200190f35b600080338484604051610106906101ad565b6001600160a01b038416815260408101829052606060208083018281528551928401929092528451608084019186019080838360005b8381101561015457818101518382015260200161013c565b50505050905090810190601f1680156101815780820380516001836020036101000a031916815260200191505b50945050505050604051809103906000f0801580156101a4573d6000803e3d6000fd5b50949350505050565b610658806101bb8339019056fe6080604052600660055534801561001557600080fd5b506040516106583803806106588339810180604052606081101561003857600080fd5b81516020830180519193928301929164010000000081111561005957600080fd5b8201602081018481111561006c57600080fd5b815164010000000081118282018710171561008657600080fd5b5050602091820151600080546001600160a01b0319166001600160a01b038816179055815191945092506100c091600191908501906100da565b5060025550506003805460ff191690556000600455610175565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061011b57805160ff1916838001178555610148565b82800160010185558215610148579182015b8281111561014857825182559160200191906001019061012d565b50610154929150610158565b5090565b61017291905b80821115610154576000815560010161015e565b90565b6104d4806101846000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c80634e69d5601461004657806383197ef0146100c35780639bd08f91146100cd575b600080fd5b61004e610175565b6040805160208082528351818301528351919283929083019185019080838360005b83811015610088578181015183820152602001610070565b50505050905090810190601f1680156100b55780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6100cb610251565b005b61004e600480360360408110156100e357600080fd5b8101906020810181356401000000008111156100fe57600080fd5b82018360208201111561011057600080fd5b8035906020019184600183028401116401000000008311171561013257600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295505091359250610274915050565b6000546060906001600160a01b031633146101c4575060408051808201909152600c81527f556e617574686f72697a65640000000000000000000000000000000000000000602082015261024e565b60035460ff1661022857600554600454136101fe57604051806040016040528060048152602001600160e21b6315d85a5d02815250610223565b604051806040016040528060088152602001600160c21b6715995c9a599a5959028152505b61024b565b604051806040016040528060068152602001600160d01b65557064617465028152505b90505b90565b6000546001600160a01b0316331415610272576000546001600160a01b0316ff5b565b606060025482141561034c5760018054604080516020600284861615610100026000190190941693909304601f81018490048402820184019092528181526103179387939192909183018282801561030d5780601f106102e25761010080835404028352916020019161030d565b820191906000526020600020905b8154815290600101906020018083116102f057829003601f168201915b50505050506103b1565b15610326576004805460010190555b506040805180820190915260068152600160d01b655468616e6b730260208201526103ab565b60025482131561038957506003805460ff191660011790556040805180820190915260068152600160d01b655468616e6b730260208201526103ab565b506040805180820190915260068152600160d01b655570646174650260208201525b92915050565b600081518351146103c4575060006103ab565b816040516020018082805190602001908083835b602083106103f75780518252601f1990920191602091820191016103d8565b6001836020036101000a03801982511681845116808217855250505050505090500191505060405160208183030381529060405280519060200120836040516020018082805190602001908083835b602083106104655780518252601f199092019160209182019101610446565b6001836020036101000a038019825116818451168082178552505050505050905001915050604051602081830303815290604052805190602001201490506103ab56fea165627a7a723058200d35c837861d3985e83e74652583505ab35ad2e511e95b82614a2921ee6696c80029a165627a7a72305820fae4f9365e52f54b71208059d507bb6b41b293633b065ed6408507f65cc4ddce0029`

// DeployVerificationContractGenerator deploys a new Ethereum contract, binding an instance of VerificationContractGenerator to it.
func DeployVerificationContractGenerator(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *VerificationContractGenerator, error) {
	parsed, err := abi.JSON(strings.NewReader(VerificationContractGeneratorABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(VerificationContractGeneratorBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &VerificationContractGenerator{VerificationContractGeneratorCaller: VerificationContractGeneratorCaller{contract: contract}, VerificationContractGeneratorTransactor: VerificationContractGeneratorTransactor{contract: contract}, VerificationContractGeneratorFilterer: VerificationContractGeneratorFilterer{contract: contract}}, nil
}

// VerificationContractGenerator is an auto generated Go binding around an Ethereum contract.
type VerificationContractGenerator struct {
	VerificationContractGeneratorCaller     // Read-only binding to the contract
	VerificationContractGeneratorTransactor // Write-only binding to the contract
	VerificationContractGeneratorFilterer   // Log filterer for contract events
}

// VerificationContractGeneratorCaller is an auto generated read-only Go binding around an Ethereum contract.
type VerificationContractGeneratorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VerificationContractGeneratorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VerificationContractGeneratorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VerificationContractGeneratorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VerificationContractGeneratorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VerificationContractGeneratorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VerificationContractGeneratorSession struct {
	Contract     *VerificationContractGenerator // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                  // Call options to use throughout this session
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// VerificationContractGeneratorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VerificationContractGeneratorCallerSession struct {
	Contract *VerificationContractGeneratorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                        // Call options to use throughout this session
}

// VerificationContractGeneratorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VerificationContractGeneratorTransactorSession struct {
	Contract     *VerificationContractGeneratorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                        // Transaction auth options to use throughout this session
}

// VerificationContractGeneratorRaw is an auto generated low-level Go binding around an Ethereum contract.
type VerificationContractGeneratorRaw struct {
	Contract *VerificationContractGenerator // Generic contract binding to access the raw methods on
}

// VerificationContractGeneratorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VerificationContractGeneratorCallerRaw struct {
	Contract *VerificationContractGeneratorCaller // Generic read-only contract binding to access the raw methods on
}

// VerificationContractGeneratorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VerificationContractGeneratorTransactorRaw struct {
	Contract *VerificationContractGeneratorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVerificationContractGenerator creates a new instance of VerificationContractGenerator, bound to a specific deployed contract.
func NewVerificationContractGenerator(address common.Address, backend bind.ContractBackend) (*VerificationContractGenerator, error) {
	contract, err := bindVerificationContractGenerator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VerificationContractGenerator{VerificationContractGeneratorCaller: VerificationContractGeneratorCaller{contract: contract}, VerificationContractGeneratorTransactor: VerificationContractGeneratorTransactor{contract: contract}, VerificationContractGeneratorFilterer: VerificationContractGeneratorFilterer{contract: contract}}, nil
}

// NewVerificationContractGeneratorCaller creates a new read-only instance of VerificationContractGenerator, bound to a specific deployed contract.
func NewVerificationContractGeneratorCaller(address common.Address, caller bind.ContractCaller) (*VerificationContractGeneratorCaller, error) {
	contract, err := bindVerificationContractGenerator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VerificationContractGeneratorCaller{contract: contract}, nil
}

// NewVerificationContractGeneratorTransactor creates a new write-only instance of VerificationContractGenerator, bound to a specific deployed contract.
func NewVerificationContractGeneratorTransactor(address common.Address, transactor bind.ContractTransactor) (*VerificationContractGeneratorTransactor, error) {
	contract, err := bindVerificationContractGenerator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VerificationContractGeneratorTransactor{contract: contract}, nil
}

// NewVerificationContractGeneratorFilterer creates a new log filterer instance of VerificationContractGenerator, bound to a specific deployed contract.
func NewVerificationContractGeneratorFilterer(address common.Address, filterer bind.ContractFilterer) (*VerificationContractGeneratorFilterer, error) {
	contract, err := bindVerificationContractGenerator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VerificationContractGeneratorFilterer{contract: contract}, nil
}

// bindVerificationContractGenerator binds a generic wrapper to an already deployed contract.
func bindVerificationContractGenerator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VerificationContractGeneratorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VerificationContractGenerator *VerificationContractGeneratorRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _VerificationContractGenerator.Contract.VerificationContractGeneratorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VerificationContractGenerator *VerificationContractGeneratorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VerificationContractGenerator.Contract.VerificationContractGeneratorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VerificationContractGenerator *VerificationContractGeneratorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VerificationContractGenerator.Contract.VerificationContractGeneratorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VerificationContractGenerator *VerificationContractGeneratorCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _VerificationContractGenerator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VerificationContractGenerator *VerificationContractGeneratorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VerificationContractGenerator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VerificationContractGenerator *VerificationContractGeneratorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VerificationContractGenerator.Contract.contract.Transact(opts, method, params...)
}

// GenerateVerificationContract is a paid mutator transaction binding the contract method 0xa34b811c.
//
// Solidity: function generateVerificationContract(string hash, int256 version) returns(address)
func (_VerificationContractGenerator *VerificationContractGeneratorTransactor) GenerateVerificationContract(opts *bind.TransactOpts, hash string, version *big.Int) (*types.Transaction, error) {
	return _VerificationContractGenerator.contract.Transact(opts, "generateVerificationContract", hash, version)
}

// GenerateVerificationContract is a paid mutator transaction binding the contract method 0xa34b811c.
//
// Solidity: function generateVerificationContract(string hash, int256 version) returns(address)
func (_VerificationContractGenerator *VerificationContractGeneratorSession) GenerateVerificationContract(hash string, version *big.Int) (*types.Transaction, error) {
	return _VerificationContractGenerator.Contract.GenerateVerificationContract(&_VerificationContractGenerator.TransactOpts, hash, version)
}

// GenerateVerificationContract is a paid mutator transaction binding the contract method 0xa34b811c.
//
// Solidity: function generateVerificationContract(string hash, int256 version) returns(address)
func (_VerificationContractGenerator *VerificationContractGeneratorTransactorSession) GenerateVerificationContract(hash string, version *big.Int) (*types.Transaction, error) {
	return _VerificationContractGenerator.Contract.GenerateVerificationContract(&_VerificationContractGenerator.TransactOpts, hash, version)
}

// VerifierABI is the input ABI used to generate the binding from.
const VerifierABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"getStatus\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"destroy\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"hash\",\"type\":\"string\"},{\"name\":\"version\",\"type\":\"int256\"}],\"name\":\"verify\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"reqNode\",\"type\":\"address\"},{\"name\":\"hash\",\"type\":\"string\"},{\"name\":\"version\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// VerifierBin is the compiled bytecode used for deploying new contracts.
const VerifierBin = `0x6080604052600660055534801561001557600080fd5b506040516106583803806106588339810180604052606081101561003857600080fd5b81516020830180519193928301929164010000000081111561005957600080fd5b8201602081018481111561006c57600080fd5b815164010000000081118282018710171561008657600080fd5b5050602091820151600080546001600160a01b0319166001600160a01b038816179055815191945092506100c091600191908501906100da565b5060025550506003805460ff191690556000600455610175565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061011b57805160ff1916838001178555610148565b82800160010185558215610148579182015b8281111561014857825182559160200191906001019061012d565b50610154929150610158565b5090565b61017291905b80821115610154576000815560010161015e565b90565b6104d4806101846000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c80634e69d5601461004657806383197ef0146100c35780639bd08f91146100cd575b600080fd5b61004e610175565b6040805160208082528351818301528351919283929083019185019080838360005b83811015610088578181015183820152602001610070565b50505050905090810190601f1680156100b55780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6100cb610251565b005b61004e600480360360408110156100e357600080fd5b8101906020810181356401000000008111156100fe57600080fd5b82018360208201111561011057600080fd5b8035906020019184600183028401116401000000008311171561013257600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295505091359250610274915050565b6000546060906001600160a01b031633146101c4575060408051808201909152600c81527f556e617574686f72697a65640000000000000000000000000000000000000000602082015261024e565b60035460ff1661022857600554600454136101fe57604051806040016040528060048152602001600160e21b6315d85a5d02815250610223565b604051806040016040528060088152602001600160c21b6715995c9a599a5959028152505b61024b565b604051806040016040528060068152602001600160d01b65557064617465028152505b90505b90565b6000546001600160a01b0316331415610272576000546001600160a01b0316ff5b565b606060025482141561034c5760018054604080516020600284861615610100026000190190941693909304601f81018490048402820184019092528181526103179387939192909183018282801561030d5780601f106102e25761010080835404028352916020019161030d565b820191906000526020600020905b8154815290600101906020018083116102f057829003601f168201915b50505050506103b1565b15610326576004805460010190555b506040805180820190915260068152600160d01b655468616e6b730260208201526103ab565b60025482131561038957506003805460ff191660011790556040805180820190915260068152600160d01b655468616e6b730260208201526103ab565b506040805180820190915260068152600160d01b655570646174650260208201525b92915050565b600081518351146103c4575060006103ab565b816040516020018082805190602001908083835b602083106103f75780518252601f1990920191602091820191016103d8565b6001836020036101000a03801982511681845116808217855250505050505090500191505060405160208183030381529060405280519060200120836040516020018082805190602001908083835b602083106104655780518252601f199092019160209182019101610446565b6001836020036101000a038019825116818451168082178552505050505050905001915050604051602081830303815290604052805190602001201490506103ab56fea165627a7a723058200d35c837861d3985e83e74652583505ab35ad2e511e95b82614a2921ee6696c80029`

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

// GetStatus is a paid mutator transaction binding the contract method 0x4e69d560.
//
// Solidity: function getStatus() returns(string)
func (_Verifier *VerifierTransactor) GetStatus(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Verifier.contract.Transact(opts, "getStatus")
}

// GetStatus is a paid mutator transaction binding the contract method 0x4e69d560.
//
// Solidity: function getStatus() returns(string)
func (_Verifier *VerifierSession) GetStatus() (*types.Transaction, error) {
	return _Verifier.Contract.GetStatus(&_Verifier.TransactOpts)
}

// GetStatus is a paid mutator transaction binding the contract method 0x4e69d560.
//
// Solidity: function getStatus() returns(string)
func (_Verifier *VerifierTransactorSession) GetStatus() (*types.Transaction, error) {
	return _Verifier.Contract.GetStatus(&_Verifier.TransactOpts)
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
