package model

type Wallet struct {
	WalletID      int    `json:"walletId"`
	OperationType string `json:"operationType"`
	Amount        int64  `json:"amount,omitempty"`
}
