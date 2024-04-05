package main

import (
	"context"
	"math/big"
	"strings"

	voltasdk "github.com/VoltaHQ/sdk-go"
	"github.com/VoltaHQ/sdk-go/examples/utils"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/fatih/color"
)

const (
	// ABI for the approve function of ERC20
	approveFunctionABI = `[{"constant":false,"inputs":[{"name":"_spender","type":"address"},{"name":"_value","type":"uint256"}],"name":"approve","outputs":[{"name":"","type":"bool"}],"payable":false,"stateMutability":"nonpayable","type":"function"}]`
)

func main() {
	parsedABI, err := abi.JSON(strings.NewReader(approveFunctionABI))
	if err != nil {
		color.Red("Failed to parse contract ABI: %v", err)
		return
	}

	targetChain := "optimism-mainnet"
	voltaClient, err := voltasdk.NewVaultClient(voltasdk.Blockchain(targetChain))
	if err != nil {
		color.Red("Could not create the volta client: %v", err)
		return
	}

	bundlerClient, err := voltasdk.NewBundlerClient(voltasdk.Blockchain(targetChain))

	// tokenAddress := common.HexToAddress("0x4200000000000000000000000000000000000006") // weth
	// spenderAddress := common.HexToAddress("0xec8b0f7ffe3ae75d7ffab09429e3675bb63503e4") // uniswap universal router

	var (
		tokenAddress, spenderAddress, vaultAddress common.Address
		approveAmount                              *big.Float
	)
	vaultAddress, err = utils.GetAddressWithPrompt("Your Vault Address")
	if err != nil {
		color.Red("unable to get user address: %v", err)
		return
	}

	tokenAddress, err = utils.GetAddressWithPrompt("Token Address")
	if err != nil {
		color.Red("unable to get user address: %v", err)
		return
	}

	spenderAddress, err = utils.GetAddressWithPrompt("Spender Address")
	if err != nil {
		color.Red("unable to get user address: %v", err)
		return
	}

	approveAmount, err = utils.GetAmount()
	if err != nil {
		color.Red("unable to get correct approve amount: %v", err)
		return
	}
	weiAmount := utils.ToWei(approveAmount)

	var userOp *voltasdk.UserOperation
	// Pack the input arguments
	var callData []byte
	callData, err = parsedABI.Pack("approve", spenderAddress, weiAmount)
	if err != nil {
		color.Red("Failed to pack data for approve: %v", err)
	}

	userOp, err = voltaClient.BuildExecuteUserOp(vaultAddress, voltasdk.Call{
		Target: tokenAddress,
		Value:  big.NewInt(0),
		Data:   callData,
	})
	if err != nil {
		color.Red("Failed to build user op: %v", err)
	}

	maxFeePerGas, maxPriorityFeePerGas, err := utils.SuggestUserOpFeePerGas(context.Background(), voltaClient)
	if err == nil {
		userOp.MaxFeePerGas = maxFeePerGas
		userOp.MaxPriorityFeePerGas = maxPriorityFeePerGas
	}
	// for performance critical applications, the following should be cached/hardcoded.  they do not change unless
	// the underlying contract itself has changed
	userOp, err = utils.SuggestUserOpGas(context.Background(), userOp, bundlerClient)

	preVerificationGas, verificationGas, callGasLimit, err := bundlerClient.EstimateUserOpGas(context.Background(), userOp)
	if err != nil {
		color.Red("Could not estimate gas for user op: %v", err)
	}
	userOp.PreVerificationGas = utils.IncreaseByPercent(preVerificationGas, 25)
	userOp.VerificationGasLimit = verificationGas
	userOp.CallGasLimit = callGasLimit

	utils.PresentUserOpToUserAndSend(context.Background(), &bundlerClient, userOp)
}
