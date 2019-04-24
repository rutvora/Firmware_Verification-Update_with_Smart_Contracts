package main

import (
	"./Contracts"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"log"
	"math/big"
	"sync"
	"time"
)

//Sends out the request to other devices for verification using MQTT
func broadcastRequest(model string, msg string) {
	opts := mqtt.NewClientOptions().AddBroker("tcp://localhost:1883")
	client := mqtt.NewClient(opts)
	token := client.Connect()
	if token.Wait() && token.Error() != nil {
		log.Fatalln(token.Error())
	}
	token = client.Publish(model, 0, false, msg)
	if token.Wait() && token.Error() != nil {
		log.Fatalln(token.Error())
	}
}

//Creates a new smart contract
func createContract(sim *backends.SimulatedBackend, model string, hash string, reqNodeVer *big.Int) (*bind.TransactOpts, *contracts.Verifier, string) {
	var auth *bind.TransactOpts
	for key := range initAddresses {
		auth = key
		delete(initAddresses, auth)
		break
	}
	addr, _, contract, err := contracts.DeployVerifier(auth, sim, auth.From, hash, reqNodeVer)
	broadcastRequest(model, addr.String())

	if err != nil {
		log.Fatalln(err)
	}
	sim.Commit()
	return auth, contract, auth.From.String()
}

//	Checks the verification status of the smart contract
//	(smart contracts only send responses and never initiate a message)
func checkContract(auth *bind.TransactOpts, reqNode string, contract *contracts.Verifier, waitGroup *sync.WaitGroup) {
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
	_, _ = contract.Destroy(auth)
	waitGroup.Done()
}

//Initiate Request Nodes for simulation
func InitiateRequestNodes(sim *backends.SimulatedBackend, waitGroup *sync.WaitGroup) {
	version := new(big.Int).SetInt64(1)
	newRequestNode("model1", "hash1", version, sim, waitGroup)
	//newRequestNode("model2", "hash2", version, sim, waitGroup)
}

//Adds a request node in the simulation
func newRequestNode(model string, hash string, version *big.Int, sim *backends.SimulatedBackend, waitGroup *sync.WaitGroup) {
	auth, contract, addr := createContract(sim, model, hash, version)
	go checkContract(auth, addr, contract, waitGroup)
	waitGroup.Add(1)
}
