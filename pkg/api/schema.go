package api

import "time"

type UserCredentials struct {
	Username string `json:"username" validate:"required,min=3,max=100"`
	Password string `json:"password" validate:"required,min=8,max=200"`
}

type Token struct {
	Token string `json:"token" validate:"required"`
}

// Transaction.Type enum
const (
	Deposit  = "deposit"
	Withdraw = "withdraw"
)

type Transaction struct {
	Username string `json:"username" validate:"required,min=3,max=100"`
	Amount int  `json:"amount" validate:"required,gt=0"`
	Type string `json:"type" validate:"required,type"`
	Timestamp time.Time `json:"timestamp"`
}

type ReportRequest struct {
	Username string `json:"username" validate:"required,min=3,max=100"`
	Start time.Time `json:"start"`
	End time.Time `json:"end"`
}

type Report struct {
	Username string `json:"username" validate:"required,min=3,max=100"`
	Start time.Time `json:"start"`
	End time.Time `json:"end"`
	Transactions []Transaction `json:"transactions"`
}
