package voltasdk

import (
	"context"
	"errors"
	"fmt"
	"github.com/NuKeyHQ/sdk-go/contracts/EntryPoint"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
)

type VaultClient interface {
	BuildExecuteUserOp(vault common.Address, call Call) (*UserOperation, error)
	BuildExecuteUserOpFromTx(vault common.Address, tx *types.Transaction) (*UserOperation, error)
	BuildExecuteBatchUserOp(vault common.Address, calls []Call) (*UserOperation, error)
	NextNonce(sender common.Address) (*big.Int, error)
	SuggestUserOpGasPrice(ctx context.Context, userOp *UserOperation) error
}

func NewVaultClient(chain Blockchain) (VaultClient, error) {
	return newVaultClient(chain)
}

type vaultClient struct {
	chain      Blockchain
	ethClient  *ethclient.Client
	entryPoint *entrypoint.EntryPoint
}

func newVaultClient(chain Blockchain) (*vaultClient, error) {
	if !chain.IsValid() {
		return nil, ErrInvalidBlockchain
	}

	ethClient, err := ethclient.DialContext(context.Background(), chain.BundlerURL())
	if err != nil {
		return nil, fmt.Errorf("failed to dial node: %w", err)
	}
	entryPoint, err := entrypoint.NewEntryPoint(defaultEVMEntryPointAddress, ethClient)
	if err != nil {
		return nil, fmt.Errorf("failed to instantiate entrypoint: %w", err)
	}

	return &vaultClient{
		chain:      chain,
		ethClient:  ethClient,
		entryPoint: entryPoint,
	}, nil
}

func (c vaultClient) BuildExecuteUserOp(vault common.Address, call Call) (*UserOperation, error) {
	callData, err := accountABI.Pack("execute", call.Target, call.Value, call.Data)
	if err != nil {
		return nil, fmt.Errorf("failed to pack execute call: %w", err)
	}

	userOp := newUserOp(vault, big.NewInt(0))
	userOp.CallData = callData
	userOp.Blockchain = c.chain
	userOp.EntryPointAddress = defaultEVMEntryPointAddress

	return userOp, nil
}

func (c vaultClient) BuildExecuteUserOpFromTx(vault common.Address, tx *types.Transaction) (*UserOperation, error) {
	if tx.To() == nil {
		return nil, errors.New("tx must have a recipient")
	}

	return c.BuildExecuteUserOp(vault, Call{
		Target: *tx.To(),
		Value:  tx.Value(),
		Data:   tx.Data(),
	})
}

func (c vaultClient) BuildExecuteBatchUserOp(sender common.Address, calls []Call) (*UserOperation, error) {
	callData, err := accountABI.Pack("executeBatch", calls)
	if err != nil {
		return nil, fmt.Errorf("failed to pack executeBatch call: %w", err)
	}

	userOp := newUserOp(sender, big.NewInt(0))
	userOp.CallData = callData
	userOp.Blockchain = c.chain
	userOp.EntryPointAddress = defaultEVMEntryPointAddress

	return userOp, nil
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

func (c vaultClient) SuggestUserOpGasPrice(ctx context.Context, userOp *UserOperation) error {
	gasTipCap, err := c.ethClient.SuggestGasTipCap(ctx)
	if err != nil {
		return fmt.Errorf("failed to get gas tip cap: %w", err)
	}
	if gasTipCap.Cmp(big.NewInt(0)) == 0 {
		gasTipCap = big.NewInt(1)
	}
	lastBlock, err := c.ethClient.BlockByNumber(ctx, nil)
	if err != nil {
		return fmt.Errorf("failed to get last block: %w", err)
	}

	userOp.MaxPriorityFeePerGas = gasTipCap
	userOp.MaxFeePerGas = big.NewInt(0).Add(gasTipCap, lastBlock.BaseFee())

	return nil
}
