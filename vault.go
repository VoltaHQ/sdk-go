package voltasdk

import (
	"context"
	"errors"
	"fmt"
	"github.com/VoltaHQ/sdk-go/contracts/EntryPoint"
	account "github.com/VoltaHQ/sdk-go/contracts/VoltaAccount"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
)

type VaultClient interface {
	BuildExecuteUserOp(sender common.Address, call Call) (*UserOperation, error)
	BuildExecuteUserOpFromTx(sender common.Address, tx *types.Transaction) (*UserOperation, error)
	BuildExecuteBatchUserOp(sender common.Address, batchCall account.ExecuteBatchCall) (*UserOperation, error)
	BuildCustomUserOp(sender common.Address, callData []byte) (*UserOperation, error)
	NextNonce(sender common.Address) (*big.Int, error)
	SuggestGasTipCap(ctx context.Context) (*big.Int, error)
	BlockByNumber(ctx context.Context, number *big.Int) (*types.Block, error)
	HeaderByNumber(ctx context.Context, number *big.Int) (*types.Header, error)
}

func NewVaultClient(chain Blockchain) (*vaultClient, error) {
	if !chain.IsValid() {
		return nil, ErrInvalidBlockchain
	}
	return newVaultClient(chain.BundlerURL(), chain.ChainID(), defaultEVMEntryPointAddress)
}

func NewVaultClientFromBundlerUrl(url string) (*vaultClient, error) {
	return NewVaultClientWithEntryPoint(url, defaultEVMEntryPointAddress)
}

func NewVaultClientWithEntryPoint(bundlerUrl string, entryPoint common.Address) (*vaultClient, error) {
	return newVaultClient(bundlerUrl, nil, entryPoint)
}

type vaultClient struct {
	chainId    big.Int
	ethClient  *ethclient.Client
	entryPoint *entrypoint.EntryPoint
}

func newVaultClient(bundlerUrl string, chainId *big.Int, entryPoint common.Address) (*vaultClient, error) {
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

	boundEntryPoint, err := entrypoint.NewEntryPoint(entryPoint, ethClient)
	if err != nil {
		return nil, fmt.Errorf("failed to instantiate entrypoint: %w", err)
	}

	return &vaultClient{
		chainId:    *chainId,
		ethClient:  ethClient,
		entryPoint: boundEntryPoint,
	}, nil
}

func (c vaultClient) BuildExecuteUserOp(sender common.Address, call Call) (*UserOperation, error) {
	callData, err := accountABI.Pack("execute", call.Target, call.Value, call.Data)
	if err != nil {
		return nil, fmt.Errorf("failed to pack execute call: %w", err)
	}

	return c.getUserOp(sender, callData)
}

func (c vaultClient) BuildExecuteUserOpFromTx(sender common.Address, tx *types.Transaction) (*UserOperation, error) {
	if tx.To() == nil {
		return nil, errors.New("tx must have a recipient")
	}

	return c.BuildExecuteUserOp(sender, Call{
		Target: *tx.To(),
		Value:  tx.Value(),
		Data:   tx.Data(),
	})
}

func (c vaultClient) BuildExecuteBatchUserOp(sender common.Address, batchCall account.ExecuteBatchCall) (*UserOperation, error) {
	callData, err := accountABI.Pack("executeBatch", batchCall)
	if err != nil {
		return nil, fmt.Errorf("failed to pack executeBatch call: %w", err)
	}
	return c.getUserOp(sender, callData)
}

func (c vaultClient) BuildCustomUserOp(sender common.Address, callData []byte) (*UserOperation, error) {
	return c.getUserOp(sender, callData)
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

func (c vaultClient) SuggestGasTipCap(ctx context.Context) (*big.Int, error) {
	return c.ethClient.SuggestGasTipCap(ctx)
}

func (c vaultClient) BlockByNumber(ctx context.Context, number *big.Int) (*types.Block, error) {
	return c.ethClient.BlockByNumber(ctx, number)
}

func (c vaultClient) HeaderByNumber(ctx context.Context, number *big.Int) (*types.Header, error) {
	return c.ethClient.HeaderByNumber(ctx, number)
}

func (c vaultClient) getUserOp(sender common.Address, callData []byte) (*UserOperation, error) {
	nonce, err := c.NextNonce(sender)
	if err != nil {
		return nil, err
	}

	userOp := newUserOp(sender, nonce)
	userOp.CallData = callData
	userOp.EntryPointAddress = defaultEVMEntryPointAddress
	userOp.ChainID = c.chainId

	return userOp, nil
}
