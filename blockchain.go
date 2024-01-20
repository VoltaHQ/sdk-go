package voltasdk

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type Blockchain string

const (
	BlockchainAvalancheMainnet Blockchain = "avalanche-mainnet"
	BlockchainPolygonMumbai    Blockchain = "polygon-mumbai"
	BlockchainEthereumGoerli   Blockchain = "ethereum-goerli"
	BlockchainOptimismSepolia  Blockchain = "optimism-sepolia"
	BlockchainOptimismMainnet  Blockchain = "optimism-mainnet"
)

var (
	chainIDs = map[Blockchain]int64{
		BlockchainAvalancheMainnet: 43114,
		BlockchainPolygonMumbai:    80001,
		BlockchainEthereumGoerli:   5,
		BlockchainOptimismSepolia:  11155420,
		BlockchainOptimismMainnet:  10,
	}

	bundlerURLs = map[Blockchain]string{
		BlockchainAvalancheMainnet: "https://api-bundler.dev.nukey.fi/avalanche-mainnet",
		BlockchainPolygonMumbai:    "https://api-bundler.dev.nukey.fi/polygon-mumbai",
		BlockchainEthereumGoerli:   "https://api-bundler.dev.nukey.fi/ethereum-goerli",
		BlockchainOptimismSepolia:  "https://api-bundler.dev.nukey.fi/optimism-sepolia",
		BlockchainOptimismMainnet:  "https://api-bundler.dev.nukey.fi/optimism-mainnet",
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
