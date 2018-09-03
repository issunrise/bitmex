package simple

import (
	"time"
)

type Funding struct {
	Timestamp        time.Time `json:"Timestamp,omitempty"`
	Symbol           string    `json:"Symbol,omitempty"`
	FundingInterval  time.Time `json:"FundingInterval,omitempty"`
	FundingRate      float64   `json:"FundingRate,omitempty"`
	FundingRateDaily float64   `json:"FundingRateDaily,omitempty"`
}

type Insurance struct {
	Currency      string    "currency"
	Timestamp     time.Time "timestamp"
	WalletBalance float64   "walletBalance"
}
