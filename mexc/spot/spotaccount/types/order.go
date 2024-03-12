package types

import "github.com/jl1/rbt/mexc/utils"

type CreateOrderParam struct {
	Symbol        string   `url:"symbol"`
	Side          string   `url:"side"`                    // ENUM: Order Side
	Type          string   `url:"type"`                    // ENUM: Order Type
	Quantity      *float64 `url:"quantity,omitempty"`      // DECIMAL
	QuoteOrderQty *float64 `url:"quoteOrderQty,omitempty"` // DECIMAL
	Price         *float64 `url:"price,omitempty"`         // DECIMAL
}

type CreateOrderParams struct {
	CreateOrderParam
	utils.DefaultParam
}

// {"symbol":"USDCUSDT","orderId":"C01__379608025012453377","orderListId":-1,"price":"1.0505","origQty":"32.36","type":"MARKET","side":"BUY","transactTime":1706287841805}
type CreateOrderResp struct {
	Symbol       string `json:"symbol"`
	OrderID      string `json:"orderId"`
	OrderListId  int64  `json:"orderListId"`
	Price        string `json:"price"`
	OrigQty      string `json:"origQty"`
	Type         string `json:"type"`
	Side         string `json:"side"`
	TransactTime int64  `json:"transactTime"`
}

type QueryOrderParam struct {
	Symbol  string `url:"symbol"`
	OrderID string `url:"orderId"`
}

type QueryOrderParams struct {
	QueryOrderParam
	utils.DefaultParam
}

type Order struct {
	Symbol              string `json:"symbol"`
	OrigClientOrderID   string `json:"origClientOrderId"`
	OrderID             string `json:"orderId"`
	ClientOrderID       string `json:"clientOrderId"`
	Price               string `json:"price"`
	OrigQty             string `json:"origQty"`
	ExecutedQty         string `json:"executedQty"`
	CummulativeQuoteQty string `json:"cummulativeQuoteQty"`
	Status              string `json:"status"`
	TimeInForce         string `json:"timeInForce"`
	Type                string `json:"type"`
	Side                string `json:"side"`
	StopPrice           string `json:"stopPrice"`
	Time                int64  `json:"time"`
	UpdateTime          int64  `json:"updateTime"`
	IsWorking           bool   `json:"isWorking"`
	OrigQuoteOrderQty   string `json:"origQuoteOrderQty"`
}
