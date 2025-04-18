package test

import (
	"fmt"
	"github.com/339-Labs/okx-api-sdk-go/config"
	"github.com/339-Labs/okx-api-sdk-go/internal"
	v5 "github.com/339-Labs/okx-api-sdk-go/pkg/client/v5"
	"testing"
)

func Test_Instruments(t *testing.T) {
	config := config.NewOkxConfig(config.OkxApiKey, config.OkxApiSecretKey, config.OkxPassphrase, 1000, "", "")
	params := internal.NewParams()
	params["instType"] = "SPOT"
	cl := new(v5.AccountClient).Init(config)
	rsp, err := cl.Instruments(params)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(rsp)
}

func Test_Tickers(t *testing.T) {
	config := config.NewOkxConfig(config.OkxApiKey, config.OkxApiSecretKey, config.OkxPassphrase, 1000, "", "")
	params := internal.NewParams()
	params["instType"] = "SPOT"
	cl := new(v5.MarketClient).Init(config)
	rsp, err := cl.Tickers(params)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(rsp)
}
