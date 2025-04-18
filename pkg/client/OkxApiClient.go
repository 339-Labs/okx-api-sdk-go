package client

import (
	"github.com/339-Labs/okx-api-sdk-go/config"
	"github.com/339-Labs/okx-api-sdk-go/internal"
	"github.com/339-Labs/okx-api-sdk-go/internal/baseclient"
)

type OkxApiClient struct {
	OkxRestClient *baseclient.OkxRestBaseClient
}

func (p *OkxApiClient) Init(config *config.OkxConfig) *OkxApiClient {
	p.OkxRestClient = new(baseclient.OkxRestBaseClient).Init(config)
	return p
}

func (p *OkxApiClient) Post(url string, params map[string]string) (string, error) {
	postBody, jsonErr := internal.ToJson(params)
	if jsonErr != nil {
		return "", jsonErr
	}
	resp, err := p.OkxRestClient.DoPost(url, postBody)
	return resp, err
}

func (p *OkxApiClient) Get(url string, params map[string]string) (string, error) {
	resp, err := p.OkxRestClient.DoGet(url, params)
	return resp, err
}
