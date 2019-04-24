package main

import (
	"./Contracts"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"log"
	"math/big"
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

func InitiateResponses(sim *backends.SimulatedBackend, contract *contracts.Verifier) {
	version := new(big.Int).SetInt64(1)

	go newResponseNode("model1", "hash1", version, sim)
	SendResponse(sim, contract, "hash1", version)
	SendResponse(sim, contract, "hash1", version)
	SendResponse(sim, contract, "hash1", version)
	SendResponse(sim, contract, "hash1", version)
	SendResponse(sim, contract, "hash1", version)
	SendResponse(sim, contract, "hash1", version)
	//version.SetInt64(2)
	SendResponse(sim, contract, "hash1", version)
}

func newResponseNode(model string, hash string, version *big.Int, sim *backends.SimulatedBackend) {
	opts := mqtt.NewClientOptions().AddBroker("tcp://localhost:1883")
	client := mqtt.NewClient(opts)
	token := client.Connect()
	if token.Wait() && token.Error() != nil {
		log.Fatalln(token.Error())
	}

	var messageHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
		contract, _ := contracts.NewVerifier(common.HexToAddress(string(msg.Payload())), sim)
		SendResponse(sim, contract, hash, version)
	}
	token = client.Subscribe(model, 0, messageHandler)
	//token.Wait()
	//if token.Wait() && token.Error() != nil {
	//	log.Fatalln(token.Error())
	//}
}
