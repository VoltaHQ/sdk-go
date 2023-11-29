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
		BlockchainPolygonMumbai: "https://api-bundler.dev.nukey.fi/",
	}

	// Note: once we deploy bundlers with RPC forwarding, we can get rid of this
	// and simply use the bundler URL
	nodeRPCURLs = map[Blockchain]string{
		BlockchainPolygonMumbai: "put_node_rpc_url_here",
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

func (b Blockchain) NodeRPCURL() string {
	if url, ok := nodeRPCURLs[b]; ok {
		return url
	}
	return ""
}

func (b Blockchain) IsValid() bool {
	return b.ChainID() != nil && b.BundlerURL() != "" && b.NodeRPCURL() != ""
}
