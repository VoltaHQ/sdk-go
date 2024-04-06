# Volta Wallet SDK Examples

There are 3 examples provided:
1. [uniswap](https://github.com/VoltaHQ/sdk-go/blob/master/examples/uniswap/uniswap.go): shows how to trade 2 ERC20 tokens on uniswap
2. [approve](https://github.com/VoltaHQ/sdk-go/blob/master/examples/approve/approve.go): shows how to call the approve function on an erc20
3. [sign](https://github.com/VoltaHQ/sdk-go/blob/master/examples/sign/sign.go): this is a helper tool to sign a given UserOp digest with a provided private key

## Typical Process

1. Make sure you have created your vault
2. Add the appropriate policies you want to test
3. After executing one of the examples it will provide you with a `digest`.  This digest needs to be signed by a private key that's authorized by the vault policy
4. This signature is provided to the running example
5. For a multi sign example, each of the private keys have to sign the same digest, and then all the signatures are to be provided - in a comma separated format - to the running example

