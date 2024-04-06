package utils

import (
	"fmt"

	voltasdk "github.com/VoltaHQ/sdk-go"
)

func PrintUserOp(op *voltasdk.UserOperation) {
	fmt.Println(userOpToString(op))
}

func userOpToString(op *voltasdk.UserOperation) string {
	return fmt.Sprintf(`
		Sender:               %s,
		Nonce:                %s,
		InitCode:             0x%x,
		CallData:             0x%x,
		CallGasLimit:         %s,
		VerificationGasLimit: %s,
		PreVerificationGas:   %s,
		MaxFeePerGas:         %s,
		MaxPriorityFeePerGas: %s,
		PaymasterAndData:     0x%x,
		Signature:            0x%x,
`, op.Sender.Hex(), op.Nonce, op.InitCode, op.CallData, op.CallGasLimit, op.VerificationGasLimit,
		op.PreVerificationGas, op.MaxFeePerGas, op.MaxPriorityFeePerGas, op.PaymasterAndData, op.Signature)
}
