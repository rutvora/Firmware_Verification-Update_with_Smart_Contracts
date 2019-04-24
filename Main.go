package main

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
	"math/big"
)

var initAddresses map[*bind.TransactOpts]core.GenesisAccount
var maxAccounts int

func initiateSimulatedBackend() *backends.SimulatedBackend {
	alloc := make(core.GenesisAlloc)
	for i := 0; i < maxAccounts; i++ {
		key, _ := crypto.GenerateKey()
		auth := bind.NewKeyedTransactor(key)

		initAddresses[auth] = core.GenesisAccount{Balance: big.NewInt(999999999)}
		alloc[auth.From] = initAddresses[auth]
	}

	sim := backends.NewSimulatedBackend(alloc, 999999999)
	return sim
}

func main() {
	//Simulated Backend to avoid mining during demo
	maxAccounts = 10
	initAddresses = make(map[*bind.TransactOpts]core.GenesisAccount, maxAccounts)
	sim := initiateSimulatedBackend()

	//Initiate request Node
	i := new(big.Int)
	i.SetInt64(1)
	auth, contract, addr := CreateContract(sim, "hash1", i)
	go CheckContract(auth, addr, contract)

	//Initiate Response Nodes
	SendResponse(sim, contract, "hash1", i)
	SendResponse(sim, contract, "hash1", i)
	SendResponse(sim, contract, "hash1", i)
	SendResponse(sim, contract, "hash1", i)
	SendResponse(sim, contract, "hash1", i)
	SendResponse(sim, contract, "hash1", i)
	//i.SetInt64(2)
	SendResponse(sim, contract, "hash1", i)

	//fmt.Println("Finished")
	select {}
}
