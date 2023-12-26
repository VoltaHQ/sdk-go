package voltasdk

import (
	"fmt"
	factory "github.com/NuKeyHQ/sdk-go/contracts/VoltaFactory"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
)

type FactoryClient interface {
	GetAccountAddress(accountImplementation common.Address, sessionKeyValidatorImplementation common.Address, executorImplementation common.Address,
		owners []common.Address, minNumSigner, seed *big.Int) (common.Address, error)
	BuildInitCode(accountImplementation common.Address, sessionKeyValidatorImplementation common.Address, executorImplementation common.Address,
		owners []common.Address, threshold *big.Int, seed *big.Int) ([]byte, error)
}

type factoryClient struct {
	factory   *factory.Factory
	ethClient *ethclient.Client
	address   common.Address
}

func NewFactoryClient(factoryAddress common.Address, ethClient *ethclient.Client) (FactoryClient, error) {
	factory, err := factory.NewFactory(factoryAddress, ethClient)
	if err != nil {
		return nil, fmt.Errorf("failed to instantiate factory: %w", err)
	}

	return &factoryClient{
		factory:   factory,
		ethClient: ethClient,
		address:   factoryAddress,
	}, nil
}

func (c factoryClient) GetAccountAddress(
	accountImplementation common.Address,
	sessionKeyValidatorImplementation common.Address,
	executorImplementation common.Address,
	owners []common.Address,
	minNumSigner,
	seed *big.Int,
) (common.Address, error) {
	address, err := c.factory.GetAddress(&bind.CallOpts{}, accountImplementation, sessionKeyValidatorImplementation,
		executorImplementation, owners, minNumSigner, seed)
	if err != nil {
		return common.Address{}, fmt.Errorf("failed to get address: %w", err)
	}
	return address, nil
}

func (c factoryClient) BuildInitCode(
	accountImplementation common.Address,
	sessionKeyValidatorImplementation common.Address,
	executorImplementation common.Address,
	owners []common.Address,
	threshold *big.Int,
	seed *big.Int,
) ([]byte, error) {
	// Get init code for the factory
	data, err := factoryABI.Pack("createAccount", accountImplementation, sessionKeyValidatorImplementation,
		executorImplementation, owners, threshold, seed,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to pack init code: %w", err)
	}

	return append(c.address.Bytes(), data...), nil
}
