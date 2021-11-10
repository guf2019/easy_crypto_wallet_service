package utils

import (
	"fmt"
	"math"
	"math/big"
)

var WeiInEther int = 18

// NanoUnixToUnix converts NanoUnix date format to Unix date format
func NanoUnixToUnix(t int64) int64 {
	stringT := fmt.Sprint(t)
	if len(stringT) > 10 {
		return t / int64(math.Pow(10, float64(len(stringT)-10)))
	}
	return t
}

// UnixToNanoUnix converts Unix date format to NanoUnix date format
func UnixToNanoUnix(t int64) int64 {
	stringT := fmt.Sprint(t)
	if len(stringT) > 10 {
		return t
	}
	return t * 1000
}

func EtherToWei(val *big.Float) *big.Float {
	return new(big.Float).Mul(val, big.NewFloat(math.Pow10(WeiInEther)))
}

func WeiToEther(val *big.Float) *big.Float {
	return new(big.Float).Quo(val, big.NewFloat(math.Pow10(WeiInEther)))
}
