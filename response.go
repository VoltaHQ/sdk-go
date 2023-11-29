package voltasdk

type Log struct {
	Address          string   `json:"address"`
	Topics           []string `json:"topics"`
	Data             string   `json:"data"`
	BlockNumber      string   `json:"blockNumber"`
	TransactionHash  string   `json:"transactionHash"`
	TransactionIndex string   `json:"transactionIndex"`
	BlockHash        string   `json:"blockHash"`
	LogIndex         string   `json:"logIndex"`
	Removed          bool     `json:"removed"`
}
type Receipt struct {
	BlockHash         string `json:"blockHash"`
	BlockNumber       string `json:"blockNumber"`
	From              string `json:"from"`
	CumulativeGasUsed string `json:"cumulativeGasUsed"`
	GasUsed           string `json:"gasUsed"`
	LogsBloom         string `json:"logsBloom"`
	TransactionHash   string `json:"transactionHash"`
	TransactionIndex  string `json:"transactionIndex"`
	EffectiveGasPrice string `json:"effectiveGasPrice"`
	Logs              []Log  `json:"logs"`
}

type GetUserOpReceiptResponse struct {
	UserOpHash    string  `json:"userOpHash"`
	Sender        string  `json:"sender"`
	Paymaster     string  `json:"paymaster"`
	Nonce         string  `json:"nonce"`
	Success       bool    `json:"success"`
	ActualGasCost string  `json:"actualGasCost"`
	ActualGasUsed string  `json:"actualGasUsed"`
	From          string  `json:"from"`
	Receipt       Receipt `json:"receipt"`
	Logs          []Log   `json:"logs"`
}
