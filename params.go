package voltasdk

import (
	"errors"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

type BaseUserOpParams struct {
	// Sender is the address of the smart contract wallet
	Sender common.Address
}

type BuildExecuteUserOpParams struct {
	BaseUserOpParams
	// To is the address of the contract to call
	To common.Address
	// Value is the amount of ETH to send with the call (in wei)
	Value *big.Int
	// CallData is the data to send with the call
	CallData []byte
}

var ErrInvalidBlockchain = errors.New("invalid blockchain")
