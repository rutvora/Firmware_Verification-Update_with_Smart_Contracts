package main

import (
	"./Contracts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
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
