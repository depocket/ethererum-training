package main

import (
	"api/contracts/bep20"
	"context"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	_ "github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	"log"
)

func readContract(client *ethclient.Client, abi abi.ABI, to *common.Address, callParam string, args ...interface{}) ([]interface{}, error) {
	msg, err := abi.Pack(callParam, args...)

	if err != nil {
		return nil, err
	}

	respBytes, err := client.CallContract(context.Background(), ethereum.CallMsg{
		To:   to,
		Data: msg,
	}, nil)

	if err != nil {
		return nil, err
	}

	return abi.Unpack(callParam, respBytes)
}

func main() {
	/// endpoints: https://docs.binance.org/smart-chain/developer/rpc.html
	client, err := ethclient.Dial("https://bsc-dataseed1.ninicoin.io/")
	bep20ABI, _ := bep20.GetAbi()

	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()
	r.GET("/token/:tokenAddress", func(c *gin.Context) {
		address := common.HexToAddress(c.Param("tokenAddress"))

		/// specification: https://github.com/binance-chain/BEPs/blob/master/BEP20.md#5--specification
		contractName, err := readContract(client, bep20ABI, &address, "name")

		if err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})

			return
		}

		contractDecimals, _ := readContract(client, bep20ABI, &address, "decimals")
		contractSymbol, _ := readContract(client, bep20ABI, &address, "symbol")
		contractTotalSupply, _ := readContract(client, bep20ABI, &address, "totalSupply")

		c.JSON(200, gin.H{
			"name":        contractName[0],
			"decimals":    contractDecimals[0],
			"symbol":      contractSymbol[0],
			"totalSupply": contractTotalSupply[0],
		})
	})
	r.GET("/token/:tokenAddress/balance/:accountAddress", func(c *gin.Context) {
		address := common.HexToAddress(c.Param("tokenAddress"))
		userAddress := common.HexToAddress(c.Param("accountAddress"))
		balanceOf, err := readContract(client, bep20ABI, &address, "balanceOf", userAddress)

		if err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})

			return
		}

		c.JSON(200, gin.H{
			"balance": balanceOf[0],
		})
	})

	err = r.Run(":3000")

	if err != nil {
		log.Fatal(err)
	}
}
