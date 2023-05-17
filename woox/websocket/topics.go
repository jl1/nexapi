package websocket

import (
	"fmt"

	"github.com/go-playground/validator"
)

func (w *WooXWebsocketClient) GetOrderbookTopic(symbol string) (string, error) {
	if symbol == "" {
		return "", fmt.Errorf("the symbol field must be provided")
	}
	return fmt.Sprintf("%s@orderbook100", symbol), nil
}

func (w *WooXWebsocketClient) GetTradeTopic(symbol string) (string, error) {
	if symbol == "" {
		return "", fmt.Errorf("the symbol field must be provided")
	}
	return fmt.Sprintf("%s@trade", symbol), nil
}

func (w *WooXWebsocketClient) GetTickerTopic(symbol string) (string, error) {
	if symbol == "" {
		return "", fmt.Errorf("the symbol field must be provided")
	}
	return fmt.Sprintf("%s@ticker", symbol), nil
}

func (w *WooXWebsocketClient) GetAllTickersTopic() (string, error) {
	return "tickers", nil
}

func (w *WooXWebsocketClient) GetBboTopic(symbol string) (string, error) {
	if symbol == "" {
		return "", fmt.Errorf("the symbol field must be provided")
	}
	return fmt.Sprintf("%s@bbo", symbol), nil
}

func (w *WooXWebsocketClient) GetAllBbosTopic() (string, error) {
	return "bbos", nil
}

type KlineTopicParam struct {
	Symbol string `validate:"required"`
	Time   string `validate:"required,oneof=1m 5m 15m 30m 1h 1d 1w 1M"`
}

func (w *WooXWebsocketClient) GetKlineTopic(params *KlineTopicParam) (string, error) {
	err := validator.New().Struct(params)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s@kline_%s", params.Symbol, params.Time), nil
}

func (w *WooXWebsocketClient) GetIndexPriceTopic(symbol string) (string, error) {
	if symbol == "" {
		return "", fmt.Errorf("the symbol field must be provided")
	}
	return fmt.Sprintf("%s@indexprice", symbol), nil
}

func (w *WooXWebsocketClient) GetMarkPriceTopic(symbol string) (string, error) {
	if symbol == "" {
		return "", fmt.Errorf("the symbol field must be provided")
	}
	return fmt.Sprintf("%s@markprice", symbol), nil
}

func (w *WooXWebsocketClient) GetMarkPricesTopic() (string, error) {
	return "markprices", nil
}

func (w *WooXWebsocketClient) GetOpenInterestTopic(symbol string) (string, error) {
	if symbol == "" {
		return "", fmt.Errorf("the symbol field must be provided")
	}
	return fmt.Sprintf("%s@openinterest", symbol), nil
}

func (w *WooXWebsocketClient) GetEstFundingRateTopic(symbol string) (string, error) {
	if symbol == "" {
		return "", fmt.Errorf("the symbol field must be provided")
	}
	return fmt.Sprintf("%s@estfundingrate", symbol), nil
}
