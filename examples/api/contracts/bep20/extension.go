package bep20

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"strings"
)

func GetAbi() (abi.ABI, error) {
	return abi.JSON(strings.NewReader(Bep20MetaData.ABI))
}
