package v5

import (
	//"github.com/339-Labs/okx-api-sdk-go/config"
	"github.com/339-Labs/okx-api-sdk-go/config"
	"github.com/339-Labs/okx-api-sdk-go/internal/baseclient"
)

type MarketClient struct {
	restOkxClient *baseclient.OkxRestBaseClient
}

func (t *MarketClient) Init(config *config.OkxConfig) *MarketClient {
	t.restOkxClient = new(baseclient.OkxRestBaseClient).Init(config)
	return t
}

func (a *MarketClient) Tickers(params map[string]string) (string, error) {
	rsp, err := a.restOkxClient.DoGetNoAuth("/api/v5/market/tickers", params)
	return rsp, err
}

func (a *MarketClient) Ticker(params map[string]string) (string, error) {
	rsp, err := a.restOkxClient.DoGetNoAuth("/api/v5/market/ticker", params)
	return rsp, err
}

func (a *MarketClient) Books(params map[string]string) (string, error) {
	rsp, err := a.restOkxClient.DoGetNoAuth("/api/v5/market/books", params)
	return rsp, err
}

func (a *MarketClient) BooksFull(params map[string]string) (string, error) {
	rsp, err := a.restOkxClient.DoGetNoAuth("/api/v5/market/books-full", params)
	return rsp, err
}

func (a *MarketClient) Candles(params map[string]string) (string, error) {
	rsp, err := a.restOkxClient.DoGetNoAuth("/api/v5/market/candles", params)
	return rsp, err
}
