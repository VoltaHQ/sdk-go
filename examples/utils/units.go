package utils

import (
	"math/big"

	"github.com/ethereum/go-ethereum/params"
)

var weiMultiplier = new(big.Float).SetInt(big.NewInt(params.Ether))

func ToWei(ether *big.Float) *big.Int {
	// Convert Ether to Wei
	wei := new(big.Float).Mul(ether, weiMultiplier)
	weiInt, _ := wei.Int(nil) // Convert to big.Int
	return weiInt
}

func FromWei(wei *big.Int) *big.Float {
	// Convert Wei back to Ether
	etherFromWei := new(big.Float).Quo(new(big.Float).SetInt(wei), big.NewFloat(params.Ether))
	return etherFromWei
}

func IncreaseByPercent(value *big.Int, percent int64) *big.Int {
	multiplier := big.NewRat(100+percent, 100)
	newValue := new(big.Int).Mul(value, multiplier.Num())
	newValue.Div(newValue, multiplier.Denom())
	return newValue
}
