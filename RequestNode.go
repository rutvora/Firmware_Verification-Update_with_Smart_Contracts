package main

import (
	"./Contracts"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"log"
	"math/big"
	"time"
)

func CreateContract(sim *backends.SimulatedBackend, hash string, reqNodeVer *big.Int) (*bind.TransactOpts, *contracts.Verifier, string) {
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

func CheckContract(auth *bind.TransactOpts, reqNode string, contract *contracts.Verifier) {
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
