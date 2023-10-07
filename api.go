package voltasdk

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"math/big"
)

type jsonrpcMessage struct {
	Version string          `json:"jsonrpc"`
	Method  string          `json:"method"`
	Params  any             `json:"params"`
	Result  json.RawMessage `json:"result,omitempty"`
	Error   json.RawMessage `json:"error,omitempty"`
	ID      int             `json:"id,omitempty"`
}

type jsonrpcErrorMessage struct {
	Code    int    `json:"code,omitempty"`
	Data    string `json:"data,omitempty"`
	Message string `json:"message,omitempty"`
}

type BundlerAPIClient struct {
	urls map[Blockchain]string
	r    *resty.Client
}

func NewBundlerAPIClient(urls map[Blockchain]string) *BundlerAPIClient {
	return &BundlerAPIClient{
		urls: urls,
		r:    resty.New(),
	}
}

func (b BundlerAPIClient) SendUserOperation(ctx context.Context, blockchain Blockchain, userOp UserOperation,
	entryPoint string) (string, error) {

	params := []any{userOp, entryPoint}

	response, err := b.doJSONRPCRequest(ctx, blockchain, "eth_sendUserOperation", params)
	if err != nil {
		return "", err
	}

	var userOpHash string
	return userOpHash, json.Unmarshal(response.Result, &userOpHash)
}

type estimateUserOperationGasResult struct {
	PreVerificationGas string `json:"PreVerificationGas"`
	VerificationGas    string `json:"VerificationGas"`
	CallGasLimit       string `json:"CallGasLimit"`
}

func (b BundlerAPIClient) EstimateUserOperationGas(ctx context.Context, blockchain Blockchain,
	userOperation UserOperation, entryPoint string) (preVerificationGas, verificationGas, callGasLimit *big.Int, err error) {

	params := []any{userOperation, entryPoint}

	response, err := b.doJSONRPCRequest(ctx, blockchain, "eth_estimateUserOperationGas", params)
	if err != nil {
		return nil, nil, nil, err
	}

	var result estimateUserOperationGasResult
	if err = json.Unmarshal(response.Result, &result); err != nil {
		return nil, nil, nil, err
	}

	var ok bool
	preVerificationGas, ok = big.NewInt(0).SetString(result.PreVerificationGas, 16)
	if !ok {
		return nil, nil, nil, fmt.Errorf("failed to parse preVerificationGas")
	}
	verificationGas, ok = big.NewInt(0).SetString(result.VerificationGas, 16)
	if !ok {
		return nil, nil, nil, fmt.Errorf("failed to parse verificationGas")
	}
	callGasLimit, ok = big.NewInt(0).SetString(result.CallGasLimit, 16)
	if !ok {
		return nil, nil, nil, fmt.Errorf("failed to parse callGasLimit")
	}

	return preVerificationGas, verificationGas, callGasLimit, nil
}

func (b BundlerAPIClient) GetUserOperationReceipt(ctx context.Context, blockchain Blockchain,
	userOpHash string) (out GetUserOperationReceiptResponse, err error) {
	params := []any{userOpHash}

	response, err := b.doJSONRPCRequest(ctx, blockchain, "eth_getUserOperationReceipt", params)
	if err != nil {
		return out, err
	}

	return out, json.Unmarshal(response.Result, &out)
}

func (b BundlerAPIClient) doJSONRPCRequest(ctx context.Context, blockchain Blockchain, method string,
	params interface{}) (jsonrpcMessage, error) {

	url, ok := b.urls[blockchain]
	if !ok {
		return jsonrpcMessage{}, fmt.Errorf("no url for blockchain %s", blockchain)
	}

	var (
		request = jsonrpcMessage{
			Version: "2.0",
			Method:  method,
			Params:  params,
			ID:      1,
		}
		response jsonrpcMessage
	)
	// Do request, check error messages and status code
	resp, err := b.r.R().
		SetContext(ctx).
		SetBody(request).
		SetResult(&response).
		SetError(&response).
		Post(url)
	if err != nil {
		return jsonrpcMessage{}, err
	}

	if len(response.Error) > 0 {
		// Switch on first byte to determine if the error is an object or string
		// NOTE: this was necessary for Stackup's hosted bundler. We should check if this is still necessary
		// when we run our own bundler
		switch response.Error[0] {
		case '{':
			var errorMessage jsonrpcErrorMessage
			if err = json.Unmarshal(response.Error, &errorMessage); err != nil {
				return jsonrpcMessage{}, fmt.Errorf("failed to unmarshal error message: %w", err)
			}
			return jsonrpcMessage{}, fmt.Errorf("error response: %s statusCode: %d", errorMessage.Message, resp.StatusCode())
		case '"':
			var errString string
			if err = json.Unmarshal(response.Error, &errString); err != nil {
				return jsonrpcMessage{}, fmt.Errorf("failed to unmarshal error message: %w", err)
			}
			return jsonrpcMessage{}, fmt.Errorf("error response: %s statusCode: %d", errString, resp.StatusCode())
		default:
			return jsonrpcMessage{}, fmt.Errorf("unknown error response: %s statusCode: %d", response.Error, resp.StatusCode())
		}
	}

	return response, nil
}
