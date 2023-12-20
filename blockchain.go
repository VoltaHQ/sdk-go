package voltasdk

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

type Blockchain string

const (
	BlockchainAvalancheMainnet Blockchain = "avalanche-mainnet"
	BlockchainPolygonMumbai    Blockchain = "polygon-mumbai"
)

var (
	chainIDs = map[Blockchain]int64{
		BlockchainAvalancheMainnet: 43114,
		BlockchainPolygonMumbai:    80001,
	}

	bundlerURLs = map[Blockchain]string{
		BlockchainAvalancheMainnet: "http://localhost:4337/avalanche_mainnet",
		BlockchainPolygonMumbai:    "https://api-bundler.dev.nukey.fi/",
	}

	defaultEVMEntryPointAddress = common.HexToAddress("0x5FF137D4b0FDCD49DcA30c7CF57E578a026d2789")
)

func (b Blockchain) ChainID() *big.Int {
	if chainID, ok := chainIDs[b]; ok {
		return big.NewInt(chainID)
	}
	return nil
}

func (b Blockchain) BundlerURL() string {
	if url, ok := bundlerURLs[b]; ok {
		return url
	}
	return ""
}

func (b Blockchain) IsValid() bool {
	return b.ChainID() != nil && b.BundlerURL() != ""
}
