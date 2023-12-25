package voltasdk

import (
	"errors"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

type Call struct {
	// Target is the address of the contract to call
	Target common.Address
	// Value is the amount of ETH to send with the call (in wei)
	Value *big.Int
	// Data is the data to send with the call
	Data []byte
}

type UserCallData struct {
	// Data is the data to send with the call
	Data []byte
}

var ErrInvalidBlockchain = errors.New("invalid blockchain")
