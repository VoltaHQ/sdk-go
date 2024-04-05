package utils

import (
	"context"
	"fmt"
	"math/big"

	voltasdk "github.com/VoltaHQ/sdk-go"
)

func SuggestUserOpFeePerGas(ctx context.Context, vaultClient voltasdk.VaultClient) (maxFeePerGas, maxPriorityFeePerGas *big.Int, err error) {
	gasTipCap, err := vaultClient.SuggestGasTipCap(ctx)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get gas tip cap: %w", err)
	}
	if gasTipCap.Cmp(big.NewInt(0)) == 0 {
		gasTipCap = big.NewInt(1)
	}
	latestHeader, err := vaultClient.HeaderByNumber(ctx, nil)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get last header: %w", err)
	}
	return big.NewInt(0).Add(gasTipCap, latestHeader.BaseFee), gasTipCap, nil
}

func SuggestUserOpGas(ctx context.Context, userOp *voltasdk.UserOperation, bundlerClient voltasdk.BundlerClient) (*voltasdk.UserOperation, error) {
	pvg, vg, cgl, err := bundlerClient.EstimateUserOpGas(ctx, userOp)
	if err == nil {
		userOp.PreVerificationGas = pvg
		userOp.VerificationGasLimit = vg
		userOp.CallGasLimit = cgl
	} else {
		userOp.PreVerificationGas = big.NewInt(700000)
		userOp.VerificationGasLimit = big.NewInt(1054982)
		userOp.CallGasLimit = big.NewInt(1e6)
	}
	return userOp, nil
}
