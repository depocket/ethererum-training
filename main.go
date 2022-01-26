package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	store "goethers-sample/contracts"
	"log"
)

func main() {
	client, err := ethclient.Dial("http://127.0.0.1:7545/")

	if err != nil {
		log.Fatal(err)
	}

	address := common.HexToAddress("0xa6472718168c23AcE2bcd473646936a3131e0BD1")
	instance, err := store.NewStore(address, client)

	if err != nil {
		log.Fatal(err)
	}

	greet, err := instance.Greet(nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(greet) // "1.0"
}
