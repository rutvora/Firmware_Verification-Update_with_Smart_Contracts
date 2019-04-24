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
	waitGroup.Done()
}

func InitiateRequests(sim *backends.SimulatedBackend, waitGroup *sync.WaitGroup) *contracts.Verifier {
	version := new(big.Int).SetInt64(1)
	auth, contract, addr := createContract(sim, "model1", "hash1", version)
	go checkContract(auth, addr, contract, waitGroup)
	waitGroup.Add(1)
	return contract
}
