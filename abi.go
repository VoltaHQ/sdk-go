package voltasdk

import (
	"github.com/NuKeyHQ/sdk-go/contracts/VoltaAccount"
	factory "github.com/NuKeyHQ/sdk-go/contracts/VoltaFactory"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"strings"
)

var (
	accountABI, factoryABI abi.ABI
)

func init() {
	var err error
	accountABI, err = abi.JSON(strings.NewReader(account.AccountMetaData.ABI))
	if err != nil {
		panic("failed to parse volta account abi: " + err.Error())
	}

	factoryABI, err = abi.JSON(strings.NewReader(factory.FactoryMetaData.ABI))
	if err != nil {
		panic("failed to parse volta factory abi: " + err.Error())
	}
}
