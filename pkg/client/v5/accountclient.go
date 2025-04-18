package v5

import (
	"github.com/339-Labs/okx-api-sdk-go/config"
	"github.com/339-Labs/okx-api-sdk-go/internal"
	"github.com/339-Labs/okx-api-sdk-go/internal/baseclient"
)

type AccountClient struct {
	restOkxClient *baseclient.OkxRestBaseClient
}

func (a *AccountClient) Init(config *config.OkxConfig) *AccountClient {
	a.restOkxClient = new(baseclient.OkxRestBaseClient).Init(config)
	return a
}

func (a *AccountClient) Instruments(params map[string]string) (string, error) {
	rsp, err := a.restOkxClient.DoGet("/api/v5/account/instruments", params)
	return rsp, err
}

func (a *AccountClient) Balance(params map[string]string) (string, error) {
	rsp, err := a.restOkxClient.DoGet("/api/v5/account/balance", params)
	return rsp, err
}

func (a *AccountClient) SetPositionMode(params map[string]string) (string, error) {
	postBody, jsonErr := internal.ToJson(params)
	if jsonErr != nil {
		return "", jsonErr
	}
	rsp, err := a.restOkxClient.DoPost("/api/v5/account/set-position-mode", postBody)
	return rsp, err
}

func (a *AccountClient) SetLeverage(params map[string]string) (string, error) {
	postBody, jsonErr := internal.ToJson(params)
	if jsonErr != nil {
		return "", jsonErr
	}
	rsp, err := a.restOkxClient.DoPost("/api/v5/account/set-leverage", postBody)
	return rsp, err
}

func (a *AccountClient) MaxSize(params map[string]string) (string, error) {
	rsp, err := a.restOkxClient.DoGet("/api/v5/account/max-size", params)
	return rsp, err
}

func (a *AccountClient) MaxAvailSize(params map[string]string) (string, error) {
	rsp, err := a.restOkxClient.DoGet("/api/v5/account/max-avail-size", params)
	return rsp, err
}

func (a *AccountClient) MarginBalance(params map[string]string) (string, error) {
	postBody, jsonErr := internal.ToJson(params)
	if jsonErr != nil {
		return "", jsonErr
	}
	rsp, err := a.restOkxClient.DoPost("/api/v5/account/margin-balance", postBody)
	return rsp, err
}
