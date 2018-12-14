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

type Trade struct {
	Timestamp       time.Time
	Symbol          string
	Side            string
	Size            float64
	Price           float64
	TickDirection   string
	TrdMatchID      string
	GrossValue      float64
	HomeNotional    float64
	ForeignNotional float64
}
