package main

import (
	"./Contracts"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
	"math/big"
	"time"
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

func reqNodeCreateContract(sim *backends.SimulatedBackend, hash string, reqNodeVer *big.Int) (*bind.TransactOpts, *contracts.Verifier, string) {
	var auth *bind.TransactOpts
	for key := range initAddresses {
		auth = key
		delete(initAddresses, auth)
		break
	}
	_, _, contract, err := contracts.DeployVerifier(auth, sim, auth.From, hash, reqNodeVer)

	if err != nil {
		log.Fatalln(err)
	}
	sim.Commit()
	return auth, contract, auth.From.String()
}

func reqNodeCheckContract(auth *bind.TransactOpts, reqNode string, contract *contracts.Verifier) {
LOOP:
	for {
		time.Sleep(time.Millisecond * 1)
		status, _ := contract.GetStatus(&bind.CallOpts{
			Pending: true,
		})
		//i, _ := contract.HashVerificationCount(nil)
		fmt.Println(reqNode + " " + status)
		if status == "Verified" || status == "Update" {
			_, _ = contract.Destroy(auth)
			break LOOP
		}
	}
}

func responseNodeFunc(sim *backends.SimulatedBackend, contract *contracts.Verifier, hash string, responseNodeVer *big.Int) {
	var auth *bind.TransactOpts
	for key := range initAddresses {
		auth = key
		delete(initAddresses, auth)
		break
	}

	_, err := contract.Verify(auth, hash, responseNodeVer)
	if err != nil {
		log.Fatalln(err)
	}
	sim.Commit()
}
func main() {
	//Simulated Backend to avoid mining during demo
	maxAccounts = 10
	initAddresses = make(map[*bind.TransactOpts]core.GenesisAccount, maxAccounts)
	sim := initiateSimulatedBackend()

	//Initiate request Node
	i := new(big.Int)
	i.SetInt64(1)
	auth, contract, addr := reqNodeCreateContract(sim, "hash1", i)
	go reqNodeCheckContract(auth, addr, contract)

	//Initiate Response Nodes
	responseNodeFunc(sim, contract, "hash1", i)
	responseNodeFunc(sim, contract, "hash1", i)
	responseNodeFunc(sim, contract, "hash1", i)
	responseNodeFunc(sim, contract, "hash1", i)
	responseNodeFunc(sim, contract, "hash1", i)
	responseNodeFunc(sim, contract, "hash1", i)
	//i.SetInt64(2)
	responseNodeFunc(sim, contract, "hash1", i)

	//fmt.Println("Finished")
	select {}
}
