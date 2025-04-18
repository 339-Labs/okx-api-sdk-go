package v5

import (
	"github.com/339-Labs/okx-api-sdk-go/config"
	"github.com/339-Labs/okx-api-sdk-go/internal"
	"github.com/339-Labs/okx-api-sdk-go/internal/baseclient"
)

type TradeClient struct {
	restOkxClient *baseclient.OkxRestBaseClient
}

func (t *TradeClient) Init(config *config.OkxConfig) *TradeClient {
	t.restOkxClient = new(baseclient.OkxRestBaseClient).Init(config)
	return t
}

func (t *TradeClient) Order(params map[string]string) (string, error) {
	postBody, jsonErr := internal.ToJson(params)
	if jsonErr != nil {
		return "", jsonErr
	}
	rsp, err := t.restOkxClient.DoPost("/api/v5/trade/order", postBody)
	return rsp, err
}

func (t *TradeClient) BatchOrders(params map[string]string) (string, error) {
	postBody, jsonErr := internal.ToJson(params)
	if jsonErr != nil {
		return "", jsonErr
	}
	rsp, err := t.restOkxClient.DoPost("/api/v5/trade/batch-orders", postBody)
	return rsp, err
}

func (t *TradeClient) CancelOrders(params map[string]string) (string, error) {
	postBody, jsonErr := internal.ToJson(params)
	if jsonErr != nil {
		return "", jsonErr
	}
	rsp, err := t.restOkxClient.DoPost("/api/v5/trade/cancel-order", postBody)
	return rsp, err
}

func (t *TradeClient) CancelBatchOrders(params map[string]string) (string, error) {
	postBody, jsonErr := internal.ToJson(params)
	if jsonErr != nil {
		return "", jsonErr
	}
	rsp, err := t.restOkxClient.DoPost("/api/v5/trade/cancel-batch-orders", postBody)
	return rsp, err
}

func (t *TradeClient) AmendBatchOrders(params map[string]string) (string, error) {
	postBody, jsonErr := internal.ToJson(params)
	if jsonErr != nil {
		return "", jsonErr
	}
	rsp, err := t.restOkxClient.DoPost("/api/v5/trade/amend-batch-orders", postBody)
	return rsp, err
}

func (t *TradeClient) ClosePosition(params map[string]string) (string, error) {
	postBody, jsonErr := internal.ToJson(params)
	if jsonErr != nil {
		return "", jsonErr
	}
	rsp, err := t.restOkxClient.DoPost("/api/v5/trade/close-position", postBody)
	return rsp, err
}

func (t *TradeClient) GetOrder(params map[string]string) (string, error) {
	rsp, err := t.restOkxClient.DoGet("/api/v5/trade/order", params)
	return rsp, err
}

func (t *TradeClient) GetPendingOrder(params map[string]string) (string, error) {
	rsp, err := t.restOkxClient.DoGet("/api/v5/trade/orders-pending", params)
	return rsp, err
}
