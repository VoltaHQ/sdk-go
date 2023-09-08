package voltasdk

import (
	"errors"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
)

type BuildUserOperationParams struct {
	// Sender is the address of the smart contract wallet
	Sender string
	// CallData is the (hex-encoded) data that will be sent to/executed on the smart contract wallet
	// This should be a call to `execute(to, value, data)`
	CallData string
	// Blockchain is the blockchain that the smart contract wallet is on
	Blockchain Blockchain
	// EntryPointAddress is the address of the EntryPoint contract to use. If this is empty, the SDK will
	// use the default EntryPoint address for the given blockchain.
	EntryPointAddress string

	CallGasLimit         *big.Int
	VerificationGasLimit *big.Int
	PreVerificationGas   *big.Int
	MaxFeePerGas         *big.Int
	MaxPriorityFeePerGas *big.Int
}

var (
	ErrInvalidSenderAddress          = errors.New("invalid Sender address")
	ErrInvalidCallData               = errors.New("invalid call data")
	ErrInvalidBlockchain             = errors.New("invalid blockchain")
	ErrInvalidEntryPointAddress      = errors.New("invalid entrypoint address")
	ErrInvalidCallGasLimit           = errors.New("invalid call gas limit")
	ErrInvalidVerificationGasLimit   = errors.New("invalid verification gas limit")
	ErrorInvalidPreVerificationGas   = errors.New("invalid pre verification gas")
	ErrorInvalidMaxFeePerGas         = errors.New("invalid max fee per gas")
	ErrorInvalidMaxPriorityFeePerGas = errors.New("invalid max priority fee per gas")
)

func ValidateParams(params BuildUserOperationParams) error {
	if !common.IsHexAddress(params.Sender) {
		return ErrInvalidSenderAddress
	}
	_, err := hexutil.Decode(params.CallData)
	if err != nil {
		return ErrInvalidCallData
	}

	if !params.Blockchain.IsValid() {
		return ErrInvalidBlockchain
	}
	if params.EntryPointAddress != "" && !common.IsHexAddress(params.EntryPointAddress) {
		return ErrInvalidEntryPointAddress
	}

	// If any gas values are set to negative, return an error
	if params.CallGasLimit != nil && params.CallGasLimit.Cmp(big.NewInt(0)) == -1 {
		return ErrInvalidCallGasLimit
	}
	if params.VerificationGasLimit != nil && params.VerificationGasLimit.Cmp(big.NewInt(0)) == -1 {
		return ErrInvalidVerificationGasLimit
	}
	if params.PreVerificationGas != nil && params.PreVerificationGas.Cmp(big.NewInt(0)) == -1 {
		return ErrorInvalidPreVerificationGas
	}
	if params.MaxFeePerGas != nil && params.MaxFeePerGas.Cmp(big.NewInt(0)) == -1 {
		return ErrorInvalidMaxFeePerGas
	}
	if params.MaxPriorityFeePerGas != nil && params.MaxPriorityFeePerGas.Cmp(big.NewInt(0)) == -1 {
		return ErrorInvalidMaxPriorityFeePerGas
	}

	return nil
}
