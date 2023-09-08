package voltasdk

import "math/big"

type Blockchain string

const (
	BlockchainEthereumMainnet Blockchain = "ethereum-mainnet"
	BlockchainEthereumGoerli  Blockchain = "ethereum-goerli"
)

func (b Blockchain) IsValid() bool {
	return b.ChainID() != nil && b.DefaultEntryPointAddress() != ""
}

var chainIDs = map[Blockchain]int64{
	BlockchainEthereumMainnet: 1,
	BlockchainEthereumGoerli:  5,
}

func (b Blockchain) ChainID() *big.Int {
	if chainID, ok := chainIDs[b]; ok {
		return big.NewInt(chainID)
	}
	return nil
}

func (b Blockchain) IsEVM() bool {
	return b == BlockchainEthereumMainnet || b == BlockchainEthereumGoerli
}

func (b Blockchain) DefaultEntryPointAddress() string {
	if b.IsEVM() {
		return "0x5FF137D4b0FDCD49DcA30c7CF57E578a026d2789"
	}
	return ""
}
