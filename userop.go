package voltasdk

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"math/big"
)

var (
	address, _ = abi.NewType("address", "", nil)
	uint256, _ = abi.NewType("uint256", "", nil)
	bytes32, _ = abi.NewType("bytes32", "", nil)

	// UserOpPrimitives is the primitive ABI types for each UserOperation field.
	UserOpPrimitives = []abi.ArgumentMarshaling{
		{Name: "sender", InternalType: "Sender", Type: "address"},
		{Name: "nonce", InternalType: "Nonce", Type: "uint256"},
		{Name: "initCode", InternalType: "InitCode", Type: "bytes"},
		{Name: "callData", InternalType: "CallData", Type: "bytes"},
		{Name: "callGasLimit", InternalType: "CallGasLimit", Type: "uint256"},
		{Name: "verificationGasLimit", InternalType: "VerificationGasLimit", Type: "uint256"},
		{Name: "preVerificationGas", InternalType: "PreVerificationGas", Type: "uint256"},
		{Name: "maxFeePerGas", InternalType: "MaxFeePerGas", Type: "uint256"},
		{Name: "maxPriorityFeePerGas", InternalType: "MaxPriorityFeePerGas", Type: "uint256"},
		{Name: "paymasterAndData", InternalType: "PaymasterAndData", Type: "bytes"},
		{Name: "signature", InternalType: "Signature", Type: "bytes"},
	}

	// UserOpType is the ABI type of a UserOperation.
	UserOpType, _ = abi.NewType("tuple", "op", UserOpPrimitives)
)

// UserOperation is the struct that represents an ERC-4337 UserOp. All string fields are hex-encoded.
type UserOperation struct {
	Sender               string   `json:"sender,omitempty"`
	Nonce                *big.Int `json:"nonce,omitempty"`    // Comes from the EntryPoint contract
	InitCode             string   `json:"initCode,omitempty"` // This is for creating a new wallet. SDK will not support this yet.
	CallData             string   `json:"callData,omitempty"`
	CallGasLimit         *big.Int `json:"callGasLimit,omitempty"`         // Gas allocated for execution step in handleOps
	VerificationGasLimit *big.Int `json:"verificationGasLimit,omitempty"` // Gas allocated for verification step in handleOps
	PreVerificationGas   *big.Int `json:"preVerificationGas,omitempty"`   // Gas allocated for other overhead in handleOps (and fee paid to bundler?)
	MaxFeePerGas         *big.Int `json:"maxFeePerGas,omitempty"`
	MaxPriorityFeePerGas *big.Int `json:"maxPriorityFeePerGas,omitempty"`
	PaymasterAndData     string   `json:"paymasterAndDataa,omitempty"`
	Signature            string   `json:"signature,omitempty"`
	// These fields are not serialized in the UserOp, but are used to generate the UserOpHash.
	EntryPointAddress string     `json:"-"`
	Blockchain        Blockchain `json:"-"`
}

func (op *UserOperation) pack() ([]byte, error) {
	args := abi.Arguments{
		{Name: "UserOp", Type: UserOpType},
	}
	packed, _ := args.Pack(&struct {
		Sender               common.Address
		Nonce                *big.Int
		InitCode             []byte
		CallData             []byte
		CallGasLimit         *big.Int
		VerificationGasLimit *big.Int
		PreVerificationGas   *big.Int
		MaxFeePerGas         *big.Int
		MaxPriorityFeePerGas *big.Int
		PaymasterAndData     []byte
		Signature            []byte
	}{
		common.HexToAddress(op.Sender),
		op.Nonce,
		hexutil.MustDecode(op.InitCode),
		hexutil.MustDecode(op.CallData),
		op.CallGasLimit,
		op.VerificationGasLimit,
		op.PreVerificationGas,
		op.MaxFeePerGas,
		op.MaxPriorityFeePerGas,
		hexutil.MustDecode(op.PaymasterAndData),
		hexutil.MustDecode(op.Signature),
	})

	enc := hexutil.Encode(packed)
	enc = "0x" + enc[66:]
	return hexutil.MustDecode(enc), nil
}

func (op *UserOperation) packForSignature() ([]byte, error) {
	args := abi.Arguments{
		{Name: "sender", Type: address},
		{Name: "nonce", Type: uint256},
		{Name: "hashInitCode", Type: bytes32},
		{Name: "hashCallData", Type: bytes32},
		{Name: "callGasLimit", Type: uint256},
		{Name: "verificationGasLimit", Type: uint256},
		{Name: "preVerificationGas", Type: uint256},
		{Name: "maxFeePerGas", Type: uint256},
		{Name: "maxPriorityFeePerGas", Type: uint256},
		{Name: "hashPaymasterAndData", Type: bytes32},
	}
	return args.Pack(
		op.Sender,
		op.Nonce,
		crypto.Keccak256Hash(hexutil.MustDecode(op.InitCode)),
		crypto.Keccak256Hash(hexutil.MustDecode(op.CallData)),
		op.CallGasLimit,
		op.VerificationGasLimit,
		op.PreVerificationGas,
		op.MaxFeePerGas,
		op.MaxPriorityFeePerGas,
		crypto.Keccak256Hash(hexutil.MustDecode(op.PaymasterAndData)),
	)
}

func (op *UserOperation) getHash(entryPoint common.Address, chainID *big.Int) (common.Hash, error) {
	packed, err := op.packForSignature()
	if err != nil {
		return common.Hash{}, err
	}

	return crypto.Keccak256Hash(
		crypto.Keccak256(packed),
		common.LeftPadBytes(entryPoint.Bytes(), 32),
		common.LeftPadBytes(chainID.Bytes(), 32),
	), nil
}

func (op *UserOperation) Sign(key *ecdsa.PrivateKey) error {
	hash, err := op.getHash(common.HexToAddress(op.EntryPointAddress), op.Blockchain.ChainID())
	if err != nil {
		return fmt.Errorf("error getting user operation hash: %w", err)
	}

	digest := crypto.Keccak256Hash([]byte("\x19Ethereum Signed Message:\n32"), hash.Bytes())
	sig, err := crypto.Sign(digest.Bytes(), key)
	if err != nil {
		return fmt.Errorf("error signing user operation: %w", err)
	}

	op.Signature = hexutil.Encode(sig)
	return nil
}
