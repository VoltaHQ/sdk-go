package voltasdk

import "C"
import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
)

type VoltaSessionKeyClient interface {
	BuildUserOperation(ctx context.Context, params BuildUserOperationParams) (*UserOperation, error)
	SignUserOperation(userOp *UserOperation, key *ecdsa.PrivateKey) error
	SendUserOperation(ctx context.Context, userOp *UserOperation) (userOpHash string, err error)
}

type client struct {
	bundlerClient *bundlerAPIClient
	clients       map[Blockchain]*ethclient.Client
}

var (
	_ VoltaSessionKeyClient = (*client)(nil)

	defaultBundlerURLs = map[Blockchain]string{
		BlockchainEthereumMainnet: "TODO",
	}
)

func NewClient() VoltaSessionKeyClient {
	return &client{
		bundlerClient: newBundlerAPIClient(defaultBundlerURLs),
		clients:       make(map[Blockchain]*ethclient.Client),
	}
}

func (c client) BuildUserOperation(ctx context.Context, params BuildUserOperationParams) (*UserOperation, error) {
	// Validate params / set defaults
	err := ValidateParams(params)
	if err != nil {
		return nil, err
	}
	entryPointAddress := params.EntryPointAddress
	if entryPointAddress == "" {
		entryPointAddress = params.Blockchain.DefaultEntryPointAddress()
	}

	// TODO: Get Nonce from entrypoint contract (contract calls via node API)

	// Estimate gas
	userOp := UserOperation{
		Sender: params.Sender,
		//Nonce:            nonce,
		CallData:          params.CallData,
		Blockchain:        params.Blockchain,
		EntryPointAddress: entryPointAddress,
	}
	preVerificationGas, verificationGas, callGasLimit, err := c.bundlerClient.EstimateUserOperationGas(ctx,
		params.Blockchain, userOp, entryPointAddress)
	if err != nil {
		return nil, err
	}
	userOp.PreVerificationGas = preVerificationGas
	userOp.VerificationGasLimit = verificationGas
	userOp.CallGasLimit = callGasLimit
	// TODO: get maxFeePerGas and maxPriorityFeePerGas from node

	return &userOp, nil
}

func (c client) SignUserOperation(op *UserOperation, key *ecdsa.PrivateKey) error {
	return op.sign(key)
}

func (c client) SendUserOperation(ctx context.Context, op *UserOperation) (userOpHash string, err error) {
	if op.Signature == "" {
		return "", fmt.Errorf("user operation must be signed before sending")
	}

	return c.bundlerClient.SendUserOperation(ctx, op.Blockchain, *op, op.EntryPointAddress)
}
