package voltasdk

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"math/big"
)

type BundlerClient interface {
	SendUserOp(ctx context.Context, userOp *UserOperation) (userOpHash string, err error)
	EstimateUserOpGas(ctx context.Context, userOperation *UserOperation) (preVerificationGas, verificationGas, callGasLimit *big.Int, err error)
	GetUserOpReceipt(ctx context.Context, userOpHash string) (GetUserOpReceiptResponse, error)
	GetUserOperationGasPrice(ctx context.Context) (UserOperationGasPrice, error)
}

func NewBundlerClient(chain Blockchain) (BundlerClient, error) {
	return newBundlerClient(chain)
}

func NewBundlerClientWithUrl(bundlerUrl string) BundlerClient {
	return &bundlerClient{
		r: resty.New().SetBaseURL(bundlerUrl),
	}
}

type bundlerClient struct {
	r *resty.Client
}

func newBundlerClient(chain Blockchain) (*bundlerClient, error) {
	if !chain.IsValid() {
		return nil, ErrInvalidBlockchain
	}

	return &bundlerClient{
		r: resty.New().SetBaseURL(chain.BundlerURL()),
	}, nil
}

func (b bundlerClient) SendUserOp(ctx context.Context, userOp *UserOperation) (string, error) {
	params := []any{userOp, userOp.EntryPointAddress}

	response, err := b.doJSONRPCRequest(ctx, "eth_sendUserOperation", params)
	if err != nil {
		return "", err
	}

	var userOpHash string
	return userOpHash, json.Unmarshal(response.Result, &userOpHash)
}

type estimateUserOpGasResult struct {
	PreVerificationGas *big.Int `json:"PreVerificationGas"`
	VerificationGas    *big.Int `json:"VerificationGas"`
	CallGasLimit       *big.Int `json:"CallGasLimit"`
}

func (b bundlerClient) EstimateUserOpGas(ctx context.Context, userOperation *UserOperation) (preVerificationGas,
	verificationGas, callGasLimit *big.Int, err error) {

	params := []any{userOperation, userOperation.EntryPointAddress}

	response, err := b.doJSONRPCRequest(ctx, "eth_estimateUserOperationGas", params)
	if err != nil {
		return nil, nil, nil, err
	}

	var result estimateUserOpGasResult
	if err = json.Unmarshal(response.Result, &result); err != nil {
		return nil, nil, nil, err
	}

	return result.PreVerificationGas, result.VerificationGas, result.CallGasLimit, nil
}

type rawUserOperationGasPrice struct {
	Slow     rawGasPriceTier `json:"slow"`
	Standard rawGasPriceTier `json:"standard"`
	Fast     rawGasPriceTier `json:"fast"`
}

type rawGasPriceTier struct {
	MaxFeePerGas         string `json:"maxFeePerGas"`
	MaxPriorityFeePerGas string `json:"maxPriorityFeePerGas"`
}

// GasPriceTier contains the max fee and max priority fee per gas suggested by the bundler for a tier.
type GasPriceTier struct {
	MaxFeePerGas         *big.Int
	MaxPriorityFeePerGas *big.Int
}

// UserOperationGasPrice contains the slow, standard and fast gas price tiers returned by the bundler.
type UserOperationGasPrice struct {
	Slow     GasPriceTier
	Standard GasPriceTier
	Fast     GasPriceTier
}

func (b bundlerClient) GetUserOperationGasPrice(ctx context.Context) (UserOperationGasPrice, error) {
	response, err := b.doJSONRPCRequest(ctx, "pimlico_getUserOperationGasPrice", []any{})
	if err != nil {
		return UserOperationGasPrice{}, err
	}

	var rawResult rawUserOperationGasPrice
	if err := json.Unmarshal(response.Result, &rawResult); err != nil {
		return UserOperationGasPrice{}, err
	}

	convertTier := func(rawTier rawGasPriceTier, label string) (GasPriceTier, error) {
		if rawTier.MaxFeePerGas == "" || rawTier.MaxPriorityFeePerGas == "" {
			return GasPriceTier{}, fmt.Errorf("bundler gas price response missing %s tier", label)
		}

		maxFeePerGas, err := parseHexToBigInt(rawTier.MaxFeePerGas)
		if err != nil {
			return GasPriceTier{}, fmt.Errorf("failed to parse %s maxFeePerGas: %w", label, err)
		}

		maxPriorityFeePerGas, err := parseHexToBigInt(rawTier.MaxPriorityFeePerGas)
		if err != nil {
			return GasPriceTier{}, fmt.Errorf("failed to parse %s maxPriorityFeePerGas: %w", label, err)
		}

		return GasPriceTier{
			MaxFeePerGas:         maxFeePerGas,
			MaxPriorityFeePerGas: maxPriorityFeePerGas,
		}, nil
	}

	slow, err := convertTier(rawResult.Slow, "slow")
	if err != nil {
		return UserOperationGasPrice{}, err
	}

	standard, err := convertTier(rawResult.Standard, "standard")
	if err != nil {
		return UserOperationGasPrice{}, err
	}

	fast, err := convertTier(rawResult.Fast, "fast")
	if err != nil {
		return UserOperationGasPrice{}, err
	}

	return UserOperationGasPrice{
		Slow:     slow,
		Standard: standard,
		Fast:     fast,
	}, nil
}

func parseHexToBigInt(value string) (*big.Int, error) {
	parsed, ok := new(big.Int).SetString(value, 0)
	if !ok {
		return nil, fmt.Errorf("invalid hex value: %s", value)
	}
	return parsed, nil
}

func (b bundlerClient) GetUserOpReceipt(ctx context.Context, userOpHash string) (out GetUserOpReceiptResponse, err error) {
	params := []any{userOpHash}

	response, err := b.doJSONRPCRequest(ctx, "eth_getUserOperationReceipt", params)
	if err != nil {
		return out, err
	}

	return out, json.Unmarshal(response.Result, &out)
}

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
	Data    any    `json:"data,omitempty"`
	Message string `json:"message,omitempty"`
}

func (b bundlerClient) doJSONRPCRequest(ctx context.Context, method string, params interface{}) (jsonrpcMessage, error) {
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
		Post("/")
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
