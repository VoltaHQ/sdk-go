# Volta Wallet SDK Uniswap Example

## Steps
Make sure that the `SwapRouter` is approved on your token contract.  Feel free to use the `approve` example to do that for you.

The output of this example will be transaction details as well as the `digest` that needs to be signed.  We have provided a handly cli tool called [sign](https://github.com/VoltaHQ/sdk-go/blob/master/examples/sign/sign.go) which takes 2 command line params: <private key> <digest>

