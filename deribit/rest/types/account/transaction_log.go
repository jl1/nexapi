package account

type TransLog struct {
	Amount           float64 `json:"amount"`
	Balance          float64 `json:"balance"`
	Cashflow         float64 `json:"cashflow"`
	Change           float64 `json:"change"`
	Commission       float64 `json:"commission"`
	Currency         string  `json:"currency"`
	Equity           float64 `json:"equity"`
	Id               int64   `json:"id"`
	InstrumentName   string  `json:"instrument_name"`
	InterestPL       float64 `json:"interest_pl"`
	MarkPrice        float64 `json:"mark_price"`
	OrderId          string  `json:"order_id"`
	Position         float64 `json:"position"`
	Price            float64 `json:"price"`
	PriceCurrency    string  `json:"price_currency"`
	ProfitAsCashflow bool    `json:"profit_as_cashflow"`
	SessionRPL       float64 `json:"session_rpl"`
	SessionUPL       float64 `json:"session_upl"`
	Side             string  `json:"side"`
	Timestamp        int64   `json:"timestamp"`
	TotalInterestPL  float64 `json:"total_interest_pl"`
	TradeId          string  `json:"trade_id"`
	Type             string  `json:"type"`
	UserId           int64   `json:"user_id"`
	UserRole         string  `json:"user_role"`
	UserSeq          int64   `json:"user_seq"`
	Username         string  `json:"username"`
}

type GetTransactionLogParams struct {
	Currency       string `json:"currency"`
	StartTimestamp int64  `json:"start_timestamp,omitempty"`
	EndTimestamp   int64  `json:"end_timestamp,omitempty"`
	Query          string `json:"query,omitempty"`
	Count          int64  `json:"count,omitempty"`
	Continuation   int64  `json:"continuation,omitempty"`
}

type GetTransactionLogResponse struct {
	Continuation int64       `json:"continuation"`
	Logs         []*TransLog `json:"logs"`
}
