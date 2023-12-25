package voltasdk

import (
	"github.com/NuKeyHQ/sdk-go/contracts/VoltaAccount"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"strings"
)

var (
	accountABI abi.ABI
)

func init() {
	var err error
	accountABI, err = abi.JSON(strings.NewReader(account.AccountMetaData.ABI))
	if err != nil {
		panic("failed to parse volta account abi: " + err.Error())
	}
}
