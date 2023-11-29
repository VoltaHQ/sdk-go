package voltasdk

import "C"
import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/NuKeyHQ/sdk-go/contracts/EntryPoint"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
)

type VaultClient interface {
	NextNonce(sender common.Address) (*big.Int, error)
	BuildExecuteUserOp(ctx context.Context, params BuildExecuteUserOpParams) (*UserOperation, error)
	SignUserOp(userOp *UserOperation, keys ...*ecdsa.PrivateKey) error
}

func NewVaultClient(chain Blockchain) (VaultClient, error) {
	return newVaultClient(chain)
}

type vaultClient struct {
	chain      Blockchain
	ethClient  *ethclient.Client
	entryPoint *EntryPoint.EntryPoint
}

func newVaultClient(chain Blockchain) (*vaultClient, error) {
	if !chain.IsValid() {
		return nil, ErrInvalidBlockchain
	}

	ethClient, err := ethclient.DialContext(context.Background(), chain.NodeRPCURL())
	if err != nil {
		return nil, fmt.Errorf("failed to dial node: %w", err)
	}
	entryPoint, err := EntryPoint.NewEntryPoint(defaultEVMEntryPointAddress, ethClient)
	if err != nil {
		return nil, fmt.Errorf("failed to instantiate entrypoint: %w", err)
	}

	return &vaultClient{
		chain:      chain,
		ethClient:  ethClient,
		entryPoint: entryPoint,
	}, nil
}

func (c vaultClient) NextNonce(sender common.Address) (*big.Int, error) {
	nonce, err := c.entryPoint.GetNonce(
		&bind.CallOpts{},
		sender,
		big.NewInt(0),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get nonce: %w", err)
	}
	return nonce, nil
}

func (c vaultClient) BuildExecuteUserOp(ctx context.Context, params BuildExecuteUserOpParams) (*UserOperation, error) {
	nonce, err := c.NextNonce(params.Sender)
	if err != nil {
		return nil, err
	}

	maxFeePerGas, maxPriorityFeePerGas, err := c.getLatestGasValues(ctx)
	if err != nil {
		return nil, err
	}

	callData, err := accountABI.Pack("execute", params.To, params.Value, params.CallData)
	if err != nil {
		return nil, fmt.Errorf("failed to pack execute call: %w", err)
	}

	userOp := newUserOp()
	userOp.Sender = params.Sender
	userOp.Nonce = nonce
	userOp.CallData = callData
	userOp.Blockchain = c.chain
	userOp.EntryPointAddress = defaultEVMEntryPointAddress
	userOp.MaxPriorityFeePerGas = maxPriorityFeePerGas
	userOp.MaxFeePerGas = maxFeePerGas

	return userOp, nil
}

func (c vaultClient) getLatestGasValues(ctx context.Context) (maxFeePerGas, maxPriorityFeePerGas *big.Int, err error) {
	gasTipCap, err := c.ethClient.SuggestGasTipCap(ctx)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get gas tip cap: %w", err)
	}
	if gasTipCap.Cmp(big.NewInt(0)) == 0 {
		gasTipCap = big.NewInt(1)
	}
	lastBlock, err := c.ethClient.BlockByNumber(ctx, nil)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get last block: %w", err)
	}

	return new(big.Int).Add(gasTipCap, lastBlock.BaseFee()), gasTipCap, nil
}

func (c vaultClient) SignUserOp(op *UserOperation, keys ...*ecdsa.PrivateKey) error {
	return op.Sign(keys...)
}
