package utils

import (
	"context"

	voltasdk "github.com/NuKeyHQ/sdk-go"
	"github.com/ethereum/go-ethereum/common"
	"github.com/fatih/color"
)

func PresentUserOpToUserAndSend(ctx context.Context, bundlerClient *voltasdk.BundlerClient, userOp *voltasdk.UserOperation) (signature []byte, err error) {
	digest, err := userOp.GetDigest()
	if err != nil {
		color.Red("unable to get signing digest: %s", err)
		return
	}
	color.Green("Please sign this digest: 0x%x", digest)
	PrintUserOp(userOp)

	signatures, err := GetHexWithPrompt("Signatures")
	if err != nil {
		color.Red("invalid signature: %s", err)
		return
	}
	signature = common.FromHex(signatures)

	if bundlerClient != nil {
		userOp.Signature = signature
		err = SendUserOp(ctx, bundlerClient, userOp)
	}

	return
}

func SendUserOp(ctx context.Context, bundlerClient *voltasdk.BundlerClient, userOp *voltasdk.UserOperation) (err error) {
	userOpHash, err := (*bundlerClient).SendUserOp(ctx, userOp)
	if err != nil {
		color.Red("unable to send user op: %s", err)
		return
	}
	color.Green("UserOp sent.")
	color.Green("Check the status of the UserOp here: https://jiffyscan.xyz/userOpHash/%s", userOpHash)
	color.Green("OpHash: %s", userOpHash)
	return err
}
