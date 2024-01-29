package main

import (
	"crypto/ecdsa"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/fatih/color"
)

func main() {
	if len(os.Args) != 3 {
		color.White("usage: %s <private key> <digest to sign>", os.Args[0])
		return
	}
	signingKey := os.Args[1]
	digestToSign := os.Args[2]
	color.Green("Signing UserOp digest: %s", digestToSign)
	digest := common.FromHex(digestToSign)
	if len(digest) != 32 {
		color.Red("Invalid digest: %s", digestToSign)
		return
	}
	keys := strings.Split(signingKey, ",")

	sig, sigTimes, err := signOpRaw(digest, keys)
	if err != nil {
		color.Red("Signing Error: %s", err)
		return
	}
	total := time.Duration(0)
	for i, sigTime := range sigTimes {
		color.Green("Signature %d took %s", i+1, sigTime)
		total += sigTime
	}
	color.Green("Total time: %s", total)
	color.Green("Signature: 0x%s", common.Bytes2Hex(sig))

}

func signOpRaw(digest []byte, privateKeys []string) (sigs []byte, signingTimes []time.Duration, err error) {
	for _, pk := range privateKeys {
		eachOwnerStart := time.Now()
		var privateKey *ecdsa.PrivateKey
		privateKey, err = crypto.ToECDSA(common.FromHex(pk))
		if err != nil {
			return sigs, signingTimes, fmt.Errorf("invalid Private key (%s): %w", pk, err)
		}
		var sig []byte
		sig, err = crypto.Sign(digest[:], privateKey)
		if err != nil {
			return sigs, signingTimes, fmt.Errorf("failed to sign digest: %w", err)
		}
		// ethereum expects v to be either 1b or 1c.  so we have to add 1b to
		// the output of crypto.Sign
		sig[64] += 27
		signingTimes = append(signingTimes, time.Since(eachOwnerStart))
		sigs = append(sigs, sig...)
	}
	return sigs, signingTimes, nil
}
