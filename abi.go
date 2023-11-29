package voltasdk

import (
	volta "github.com/NuKeyHQ/sdk-go/contracts/voltaaccount"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"strings"
)

var (
	accountABI abi.ABI
)

func init() {
	var err error
	accountABI, err = abi.JSON(strings.NewReader(volta.VoltaMetaData.ABI))
	if err != nil {
		panic("failed to parse volta account abi: " + err.Error())
	}
}
