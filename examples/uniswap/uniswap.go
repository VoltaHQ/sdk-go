package main

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"time"

	voltasdk "github.com/VoltaHQ/sdk-go"
	"github.com/VoltaHQ/sdk-go/examples/utils"
	coreEntities "github.com/daoleno/uniswap-sdk-core/entities"
	"github.com/daoleno/uniswapv3-sdk/constants"
	"github.com/daoleno/uniswapv3-sdk/entities"
	"github.com/daoleno/uniswapv3-sdk/examples/contract"
	"github.com/daoleno/uniswapv3-sdk/periphery"
	sdkutils "github.com/daoleno/uniswapv3-sdk/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/fatih/color"
)

var (
	uniswapV3Factory  = common.HexToAddress("0x1F98431c8aD98523631AE4a59f267346ea31F984")
	swapRouterAddress = common.HexToAddress("0xE592427A0AEce92De3Edee1F18E0157C05861564")
	fromToken         = common.HexToAddress("0x4200000000000000000000000000000000000006")
	toToken           = common.HexToAddress("0x4200000000000000000000000000000000000042")
	// a pair is structured as BASE-QUOTE.  The above defaults to ETH-OP

	ethClient *ethclient.Client
)

func GetPoolAddress(token0, token1 common.Address, fee *big.Int) (common.Address, error) {
	f, err := contract.NewUniswapv3Factory(uniswapV3Factory, ethClient)
	if err != nil {
		return common.Address{}, fmt.Errorf("cannot get factory contract: %v", err)
	}
	poolAddr, err := f.GetPool(nil, token0, token1, fee)
	if err != nil {
		return common.Address{}, fmt.Errorf("cannot get pool: %v", err)
	}
	if poolAddr == (common.Address{}) {
		return common.Address{}, errors.New("pool does not exist")
	}

	return poolAddr, nil
}

func ConstructV3Pool(token0, token1 *coreEntities.Token, poolFee uint64) (*entities.Pool, common.Address, error) {
	poolAddress, err := GetPoolAddress(token0.Address, token1.Address, new(big.Int).SetUint64(poolFee))
	if err != nil {
		return nil, poolAddress, fmt.Errorf("cannot get pool address: %v", err)
	}

	color.Green("Pool Address: %s", poolAddress)

	contractPool, err := contract.NewUniswapv3Pool(poolAddress, ethClient)
	if err != nil {
		return nil, poolAddress, fmt.Errorf("cannot get pool contract: %v", err)
	}

	liquidity, err := contractPool.Liquidity(nil)
	if err != nil {
		return nil, poolAddress, fmt.Errorf("cannot get pool liquidity: %v", err)
	}

	slot0, err := contractPool.Slot0(nil)
	if err != nil {
		return nil, poolAddress, fmt.Errorf("cannot get pool slot0: %v", err)
	}

	poolTick, err := contractPool.Ticks(nil, big.NewInt(0))
	if err != nil {
		return nil, poolAddress, fmt.Errorf("cannot get pool tick: %v", err)
	}

	feeAmount := constants.FeeAmount(poolFee)
	ticks := []entities.Tick{
		{
			Index: entities.NearestUsableTick(sdkutils.MinTick,
				constants.TickSpacings[feeAmount]),
			LiquidityNet:   poolTick.LiquidityNet,
			LiquidityGross: poolTick.LiquidityGross,
		},
		{
			Index: entities.NearestUsableTick(sdkutils.MaxTick,
				constants.TickSpacings[feeAmount]),
			LiquidityNet:   poolTick.LiquidityGross,
			LiquidityGross: poolTick.LiquidityNet,
		},
	}

	ticks[1].LiquidityNet = big.NewInt(0).Sub(big.NewInt(0), poolTick.LiquidityNet)
	// create tick data provider
	p, err := entities.NewTickListDataProvider(ticks, constants.TickSpacings[feeAmount])
	// err = "tick net delta must be zero"
	if err != nil {
		return nil, poolAddress, fmt.Errorf("cannot create tick data provider: %v", err)
	}

	p2, err := entities.NewPool(token0, token1, constants.FeeAmount(poolFee),
		slot0.SqrtPriceX96, liquidity, int(slot0.Tick.Int64()), p)

	return p2, poolAddress, err

}

func main() {
	// 0.0001 * 1e18
	swapValue := big.NewInt(100000000000000)
	myAddress := common.HexToAddress("0x0a466bbDD9b54F47C74585D8aD4e35F8912aA087")
	targetChain := "optimism-mainnet"
	voltaClient, err := voltasdk.NewVaultClient(voltasdk.Blockchain(targetChain))
	if err != nil {
		color.Red("Could not create the volta client: %v", err)
		return
	}

	ethClient, err = ethclient.DialContext(context.Background(), "https://api-bundler.dev.nukey.fi/"+targetChain)
	bundlerClient, err := voltasdk.NewBundlerClient(voltasdk.Blockchain(targetChain))

	op := coreEntities.NewToken(10, toToken, 18, "OP", "OP")
	weth := coreEntities.NewToken(10, fromToken, 18, "WETH", "Wrapped ETH")

	pool, _, err := ConstructV3Pool(op, weth, uint64(constants.FeeMedium))
	if err != nil {
		color.Red("Could not create v3 pool: %v", err)
		return
	}

	//1%
	slippageTolerance := coreEntities.NewPercent(big.NewInt(1), big.NewInt(100))
	//after 15 minutes
	d := time.Now().Add(time.Minute * time.Duration(15)).Unix()
	deadline := big.NewInt(d)

	// single trade input
	// single-hop exact input
	r, err := entities.NewRoute([]*entities.Pool{pool}, op, weth)
	if err != nil {
		color.Red("Could not create new route: %v", err)
	}

	trade, err := entities.FromRoute(
		r,
		coreEntities.FromRawAmount(op, swapValue),
		coreEntities.ExactInput,
	)
	if err != nil {
		color.Red("Could not create trade from route: %v", err)
	}

	params, err := periphery.SwapCallParameters([]*entities.Trade{trade}, &periphery.SwapOptions{
		SlippageTolerance: slippageTolerance,
		Recipient:         myAddress,
		Deadline:          deadline,
	})
	if err != nil {
		color.Red("Could not create swap params: %v", err)
	}

	userOp, err := voltaClient.BuildExecuteUserOp(myAddress, voltasdk.Call{
		Target: swapRouterAddress, // swap router on optimism
		Value:  params.Value,
		Data:   params.Calldata,
	})
	if err != nil {
		color.Red("Could not build user op: %v", err)
	}

	maxFeePerGas, maxPriorityFeePerGas, err := utils.SuggestUserOpFeePerGas(context.Background(), voltaClient)
	if err == nil {
		userOp.MaxFeePerGas = maxFeePerGas
		userOp.MaxPriorityFeePerGas = maxPriorityFeePerGas
	}
	userOp, err = utils.SuggestUserOpGas(context.Background(), userOp, bundlerClient)

	preVerificationGas, verificationGas, callGasLimit, err := bundlerClient.EstimateUserOpGas(context.Background(), userOp)
	if err != nil {
		color.Red("Could not estimate gas for user op: %v", err)
	}
	userOp.PreVerificationGas = utils.IncreaseByPercent(preVerificationGas, 10)
	userOp.VerificationGasLimit = verificationGas
	userOp.CallGasLimit = callGasLimit

	utils.PresentUserOpToUserAndSend(context.Background(), &bundlerClient, userOp)
}
