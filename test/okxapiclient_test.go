package test

import (
	"github.com/339-Labs/okx-api-sdk-go/config"
	"github.com/339-Labs/okx-api-sdk-go/internal"
	"github.com/339-Labs/okx-api-sdk-go/pkg/client"
	"testing"
)

func Test_post(t *testing.T) {

	params := internal.NewParams()
	params["instType"] = "SPOT"

	config := config.NewOkxConfig(config.OkxApiKey, config.OkxApiSecretKey, config.OkxPassphrase, 1000, "", "")
	okx := new(client.OkxApiClient).Init(config)
	rsp, err := okx.Get("/api/v5/account/instruments", params)
	if err != nil {
		t.Error(err)
	}
	t.Log(rsp)
}
