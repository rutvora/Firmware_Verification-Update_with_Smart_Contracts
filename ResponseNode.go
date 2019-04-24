package main

import (
	"./Contracts"
	"github.com/eclipse/paho.mqtt.golang"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"log"
	"math/big"
	"sync"
)

func SendResponse(sim *backends.SimulatedBackend, contract *contracts.Verifier, hash string, responseNodeVer *big.Int) {
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

func InitiateResponses(sim *backends.SimulatedBackend, contract *contracts.Verifier, waitgroup *sync.WaitGroup) {
	i := new(big.Int).SetInt64(1)
	SendResponse(sim, contract, "hash1", i)
	SendResponse(sim, contract, "hash1", i)
	SendResponse(sim, contract, "hash1", i)
	SendResponse(sim, contract, "hash1", i)
	SendResponse(sim, contract, "hash1", i)
	SendResponse(sim, contract, "hash1", i)
	//i.SetInt64(2)
	SendResponse(sim, contract, "hash1", i)
}

func newResponseNode(model string, hash string, version big.Int) {
	opts := mqtt.NewClientOptions().AddBroker("tcp://172.16.8.33:1883")
	client := mqtt.NewClient(opts)
	token := client.Connect()
	if token.Wait() && token.Error() != nil {
		log.Fatalln(token.Error())
	}

}
