package helpers

import (
	"math/big"
)

func ConvertToBigInt(value string) *big.Int {
	n := new(big.Int)
	n, ok := n.SetString(value, 10)
	if !ok {
		return nil
	}
	return n
}
