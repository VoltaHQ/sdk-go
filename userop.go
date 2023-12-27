package voltasdk

import (
	"crypto/ecdsa"
	"encoding/json"
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
		{Name: "callData", InternalType: "Data", Type: "bytes"},
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
	Sender               common.Address
	Nonce                *big.Int // Comes from the EntryPoint contract
	InitCode             []byte   // This is for creating a new wallet. SDK will not support this yet.
	CallData             []byte
	CallGasLimit         *big.Int // Gas allocated for execution step in handleOps
	VerificationGasLimit *big.Int // Gas allocated for verification step in handleOps
	PreVerificationGas   *big.Int // Gas allocated for other overhead in handleOps (and fee paid to bundler?)
	MaxFeePerGas         *big.Int
	MaxPriorityFeePerGas *big.Int
	PaymasterAndData     []byte
	Signature            []byte
	// These fields are not serialized in the UserOp, but are used to generate the UserOpHash.
	EntryPointAddress common.Address
	ChainID           big.Int
}

func newUserOp(sender common.Address, nonce *big.Int) *UserOperation {
	return &UserOperation{
		Sender:               sender,
		Nonce:                nonce,
		ChainID:              *big.NewInt(0),
		CallGasLimit:         big.NewInt(0),
		VerificationGasLimit: big.NewInt(0),
		PreVerificationGas:   big.NewInt(0),
		MaxFeePerGas:         big.NewInt(0),
		MaxPriorityFeePerGas: big.NewInt(0),
	}
}

func (op *UserOperation) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Sender               string `json:"sender"`
		Nonce                uint64 `json:"nonce"`
		InitCode             string `json:"initCode"`
		CallData             string `json:"callData"`
		CallGasLimit         uint64 `json:"callGasLimit"`
		VerificationGasLimit uint64 `json:"verificationGasLimit"`
		PreVerificationGas   uint64 `json:"preVerificationGas"`
		MaxFeePerGas         uint64 `json:"maxFeePerGas"`
		MaxPriorityFeePerGas uint64 `json:"maxPriorityFeePerGas"`
		PaymasterAndData     string `json:"paymasterAndData"`
		Signature            string `json:"signature"`
	}{
		Sender:               op.Sender.Hex(),
		Nonce:                op.Nonce.Uint64(),
		InitCode:             hexutil.Encode(op.InitCode),
		CallData:             hexutil.Encode(op.CallData),
		CallGasLimit:         op.CallGasLimit.Uint64(),
		VerificationGasLimit: op.VerificationGasLimit.Uint64(),
		PreVerificationGas:   op.PreVerificationGas.Uint64(),
		MaxFeePerGas:         op.MaxFeePerGas.Uint64(),
		MaxPriorityFeePerGas: op.MaxPriorityFeePerGas.Uint64(),
		PaymasterAndData:     hexutil.Encode(op.PaymasterAndData),
		Signature:            hexutil.Encode(op.Signature),
	})
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
		op.Sender,
		op.Nonce,
		op.InitCode,
		op.CallData,
		op.CallGasLimit,
		op.VerificationGasLimit,
		op.PreVerificationGas,
		op.MaxFeePerGas,
		op.MaxPriorityFeePerGas,
		op.PaymasterAndData,
		op.Signature,
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
	// TODO: check values, args.Pack panics if any are nil
	return args.Pack(
		op.Sender,
		op.Nonce,
		crypto.Keccak256Hash(op.InitCode),
		crypto.Keccak256Hash(op.CallData),
		op.CallGasLimit,
		op.VerificationGasLimit,
		op.PreVerificationGas,
		op.MaxFeePerGas,
		op.MaxPriorityFeePerGas,
		crypto.Keccak256Hash(op.PaymasterAndData),
	)
}

func (op *UserOperation) GetOpHash() (common.Hash, error) {
	packed, err := op.packForSignature()
	if err != nil {
		return common.Hash{}, err
	}

	return crypto.Keccak256Hash(
		crypto.Keccak256(packed),
		common.LeftPadBytes(op.EntryPointAddress.Bytes(), 32),
		common.LeftPadBytes(op.ChainID.Bytes(), 32),
	), nil
}

func (op *UserOperation) GetDigest() (common.Hash, error) {
	hash, err := op.GetOpHash()
	if err != nil {
		return common.Hash{}, err
	}
	return crypto.Keccak256Hash([]byte("\x19Ethereum Signed Message:\n32"), hash.Bytes()), nil
}

func (op *UserOperation) Sign(keys ...*ecdsa.PrivateKey) error {
	digest, err := op.GetDigest()
	if err != nil {
		return fmt.Errorf("error getting digest: %w", err)
	}

	var sig []byte
	for _, key := range keys {
		sig, err = crypto.Sign(digest.Bytes(), key)
		if err != nil {
			return fmt.Errorf("error signing user operation: %w", err)
		}
		op.Signature = append(op.Signature, sig...)
	}

	return nil
}
