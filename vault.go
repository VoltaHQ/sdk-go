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
	BuildEnableBatchUserOp(vault common.Address, enableCalls []UserCallData) (*UserOperation, error)
	BuildDisableBatchUserOp(vault common.Address, disableCalls []UserCallData) (*UserOperation, error)
	BuildCustomUserOp(vault common.Address, callData []byte) (*UserOperation, error)
	NextNonce(sender common.Address) (*big.Int, error)
	SuggestUserOpGasPrice(ctx context.Context, userOp *UserOperation) error
}

func NewVaultClient(chain Blockchain) (VaultClient, error) {
	if !chain.IsValid() {
		return nil, ErrInvalidBlockchain
	}
	return newVaultClient(chain.BundlerURL(), chain.ChainID())
}

func NewVaultClientFromBundlerUrl(url string) (*vaultClient, error) {
	return newVaultClient(url, nil)
}

type vaultClient struct {
	chainId    big.Int
	ethClient  *ethclient.Client
	entryPoint *entrypoint.EntryPoint
}

func newVaultClient(bundlerUrl string, chainId *big.Int) (*vaultClient, error) {
	ethClient, err := ethclient.DialContext(context.Background(), bundlerUrl)
	if err != nil {
		return nil, fmt.Errorf("failed to dial node: %w", err)
	}
	if chainId == nil {
		chainId, err = ethClient.ChainID(context.Background())
		if err != nil {
			return nil, ErrInvalidBlockchain
		}
	}

	entryPoint, err := entrypoint.NewEntryPoint(defaultEVMEntryPointAddress, ethClient)
	if err != nil {
		return nil, fmt.Errorf("failed to instantiate entrypoint: %w", err)
	}

	return &vaultClient{
		chainId:    *chainId,
		ethClient:  ethClient,
		entryPoint: entryPoint,
	}, nil
}

func (c vaultClient) BuildExecuteUserOp(vault common.Address, call Call) (*UserOperation, error) {
	callData, err := accountABI.Pack("execute", call.Target, call.Value, call.Data)
	if err != nil {
		return nil, fmt.Errorf("failed to pack execute call: %w", err)
	}

	nonce, err := c.NextNonce(vault)
	if err != nil {
		return nil, fmt.Errorf("failed to get next nonce: %w", err)
	}

	userOp := newUserOp(vault, nonce)
	userOp.CallData = callData
	userOp.ChainID = c.chainId
	userOp.EntryPointAddress = defaultEVMEntryPointAddress

	err = c.SuggestUserOpGasPrice(context.TODO(), userOp)
	if err != nil {
		return nil, fmt.Errorf("failed to get suggest gas price: %w", err)
	}

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

	nonce, err := c.NextNonce(sender)
	if err != nil {
		return nil, fmt.Errorf("failed to get next nonce: %w", err)
	}
	userOp := newUserOp(sender, nonce)
	userOp.CallData = callData
	userOp.ChainID = c.chainId
	userOp.EntryPointAddress = defaultEVMEntryPointAddress

	err = c.SuggestUserOpGasPrice(context.TODO(), userOp)
	if err != nil {
		return nil, fmt.Errorf("failed to get suggest gas price: %w", err)
	}

	return userOp, nil
}

func (c vaultClient) BuildEnableBatchUserOp(sender common.Address, enableCalls []UserCallData) (*UserOperation, error) {
	callData, err := accountABI.Pack("enableBatch", enableCalls)
	if err != nil {
		return nil, fmt.Errorf("failed to pack enableBatch call: %w", err)
	}

	nonce, err := c.NextNonce(sender)
	if err != nil {
		return nil, fmt.Errorf("failed to get next nonce: %w", err)
	}
	userOp := newUserOp(sender, nonce)
	userOp.CallData = callData
	userOp.Blockchain = c.chain
	userOp.EntryPointAddress = defaultEVMEntryPointAddress

	err = c.SuggestUserOpGasPrice(context.TODO(), userOp)
	if err != nil {
		return nil, fmt.Errorf("failed to get suggest gas price: %w", err)
	}

	return userOp, nil
}

func (c vaultClient) BuildDisableBatchUserOp(sender common.Address, disableCalls []UserCallData) (*UserOperation, error) {
	callData, err := accountABI.Pack("disableBatch", disableCalls)
	if err != nil {
		return nil, fmt.Errorf("failed to pack disableBatch call: %w", err)
	}

	nonce, err := c.NextNonce(sender)
	if err != nil {
		return nil, fmt.Errorf("failed to get next nonce: %w", err)
	}
	userOp := newUserOp(sender, nonce)
	userOp.CallData = callData
	userOp.Blockchain = c.chain
	userOp.EntryPointAddress = defaultEVMEntryPointAddress

	err = c.SuggestUserOpGasPrice(context.TODO(), userOp)
	if err != nil {
		return nil, fmt.Errorf("failed to get suggest gas price: %w", err)
	}

	return userOp, nil
}

func (c vaultClient) BuildCustomUserOp(vault common.Address, callData []byte) (*UserOperation, error) {
	nonce, err := c.NextNonce(vault)
	if err != nil {
		return nil, fmt.Errorf("failed to get next nonce: %w", err)
	}
	userOp := newUserOp(vault, nonce)
	userOp.CallData = callData
	userOp.ChainID = c.chainId
	userOp.EntryPointAddress = defaultEVMEntryPointAddress

	err = c.SuggestUserOpGasPrice(context.TODO(), userOp)
	if err != nil {
		return nil, fmt.Errorf("failed to get suggest gas price: %w", err)
	}

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
