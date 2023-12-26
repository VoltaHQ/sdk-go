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
	BuildExecuteUserOp(vaultInfo VaultInfo, call Call) (*UserOperation, error)
	BuildExecuteUserOpFromTx(vaultInfo VaultInfo, tx *types.Transaction) (*UserOperation, error)
	BuildExecuteBatchUserOp(vaultInfo VaultInfo, calls []Call) (*UserOperation, error)
	BuildEnableBatchUserOp(vaultInfo VaultInfo, enableCalls []UserCallData) (*UserOperation, error)
	BuildDisableBatchUserOp(vaultInfo VaultInfo, disableCalls []UserCallData) (*UserOperation, error)
	BuildCustomUserOp(vaultInfo VaultInfo, callData []byte) (*UserOperation, error)
	NextNonce(vaultInfo VaultInfo) (*big.Int, error)
}

func NewVaultClient(chain Blockchain, config ChainConfig) (*vaultClient, error) {
	if !chain.IsValid() {
		return nil, ErrInvalidBlockchain
	}
	return newVaultClient(chain.BundlerURL(), config, chain.ChainID())
}

func NewVaultClientFromBundlerUrl(url string, config ChainConfig) (*vaultClient, error) {
	return newVaultClient(url, config, nil)
}

type VaultInfo struct {
	Owners          []common.Address
	Threshold, Seed *big.Int
}

type vaultClient struct {
	chainId       big.Int
	chainConfig   ChainConfig
	factoryClient FactoryClient
	bundlerClient BundlerClient
	ethClient     *ethclient.Client
	entryPoint    *entrypoint.EntryPoint
}

func newVaultClient(bundlerUrl string, config ChainConfig, chainId *big.Int) (*vaultClient, error) {
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

	factory, err := NewFactoryClient(config.FactoryAddress, ethClient)
	if err != nil {
		return nil, fmt.Errorf("failed to instantiate factoryClient: %w", err)
	}

	bundler := NewBundlerClientWithUrl(bundlerUrl)

	return &vaultClient{
		chainId:       *chainId,
		chainConfig:   config,
		factoryClient: factory,
		bundlerClient: bundler,
		ethClient:     ethClient,
		entryPoint:    entryPoint,
	}, nil
}

func (c vaultClient) BuildExecuteUserOp(vaultInfo VaultInfo, call Call) (*UserOperation, error) {
	callData, err := accountABI.Pack("execute", call.Target, call.Value, call.Data)
	if err != nil {
		return nil, fmt.Errorf("failed to pack execute call: %w", err)
	}

	return c.buildUserOp(vaultInfo, callData)
}

func (c vaultClient) BuildExecuteUserOpFromTx(vaultInfo VaultInfo, tx *types.Transaction) (*UserOperation, error) {
	if tx.To() == nil {
		return nil, errors.New("tx must have a recipient")
	}

	return c.BuildExecuteUserOp(vaultInfo, Call{
		Target: *tx.To(),
		Value:  tx.Value(),
		Data:   tx.Data(),
	})
}

func (c vaultClient) BuildExecuteBatchUserOp(vaultInfo VaultInfo, calls []Call) (*UserOperation, error) {
	callData, err := accountABI.Pack("executeBatch", calls)
	if err != nil {
		return nil, fmt.Errorf("failed to pack executeBatch call: %w", err)
	}
	return c.buildUserOp(vaultInfo, callData)
}

func (c vaultClient) BuildEnableBatchUserOp(vaultInfo VaultInfo, enableCalls []UserCallData) (*UserOperation, error) {
	callData, err := accountABI.Pack("enableBatch", enableCalls)
	if err != nil {
		return nil, fmt.Errorf("failed to pack enableBatch call: %w", err)
	}

	return c.buildUserOp(vaultInfo, callData)
}

func (c vaultClient) BuildDisableBatchUserOp(vaultInfo VaultInfo, disableCalls []UserCallData) (*UserOperation, error) {
	callData, err := accountABI.Pack("disableBatch", disableCalls)
	if err != nil {
		return nil, fmt.Errorf("failed to pack disableBatch call: %w", err)
	}

	return c.buildUserOp(vaultInfo, callData)
}

func (c vaultClient) BuildCustomUserOp(vaultInfo VaultInfo, callData []byte) (*UserOperation, error) {
	return c.buildUserOp(vaultInfo, callData)
}

func (c vaultClient) buildUserOp(vaultInfo VaultInfo, callData []byte) (*UserOperation, error) {
	vaultAddress, err := c.GetVaultAddress(vaultInfo)
	if err != nil {
		return nil, fmt.Errorf("failed to get vault address: %w", err)
	}
	nonce, err := c.NextNonce(vaultInfo)
	if err != nil {
		return nil, fmt.Errorf("failed to get next nonce: %w", err)
	}

	var initCode []byte
	if nonce.Cmp(big.NewInt(0)) == 0 {
		initCode, err = c.factoryClient.BuildInitCode(
			c.chainConfig.AccountImplementation,
			c.chainConfig.SessionKeyValidatorImplementation,
			c.chainConfig.ExecutorImplementation,
			vaultInfo.Owners,
			vaultInfo.Threshold,
			vaultInfo.Seed,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to get init code: %w", err)
		}
	}

	ctx := context.TODO()
	gasTipCap, err := c.ethClient.SuggestGasTipCap(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get gas tip cap: %w", err)
	}
	if gasTipCap.Cmp(big.NewInt(0)) == 0 {
		gasTipCap = big.NewInt(1)
	}
	lastBlock, err := c.ethClient.BlockByNumber(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get last block: %w", err)
	}

	userOp := UserOperation{
		Sender:               vaultAddress,
		Nonce:                nonce,
		ChainID:              c.chainId,
		InitCode:             initCode,
		CallData:             callData,
		EntryPointAddress:    defaultEVMEntryPointAddress,
		CallGasLimit:         big.NewInt(1e6),
		VerificationGasLimit: big.NewInt(1054982),
		PreVerificationGas:   big.NewInt(700000),
		MaxFeePerGas:         big.NewInt(0).Add(gasTipCap, lastBlock.BaseFee()),
		MaxPriorityFeePerGas: gasTipCap,
	}
	pvg, vg, cgl, err := c.bundlerClient.EstimateUserOpGas(ctx, &userOp)
	if err == nil {
		userOp.PreVerificationGas = pvg
		userOp.VerificationGasLimit = vg
		userOp.CallGasLimit = cgl
	}

	return &userOp, nil
}

func (c vaultClient) NextNonce(vaultInfo VaultInfo) (*big.Int, error) {
	vaultAddress, err := c.GetVaultAddress(vaultInfo)
	if err != nil {
		return nil, fmt.Errorf("failed to get vault address: %w", err)
	}
	nonce, err := c.entryPoint.GetNonce(
		&bind.CallOpts{},
		vaultAddress,
		big.NewInt(0),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get nonce: %w", err)
	}
	return nonce, nil
}

func (c vaultClient) GetVaultAddress(vaultInfo VaultInfo) (common.Address, error) {
	return c.factoryClient.GetAccountAddress(
		c.chainConfig.AccountImplementation,
		c.chainConfig.SessionKeyValidatorImplementation,
		c.chainConfig.ExecutorImplementation,
		vaultInfo.Owners,
		vaultInfo.Threshold,
		vaultInfo.Seed,
	)
}
