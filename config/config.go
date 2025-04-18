package config

import "github.com/339-Labs/okx-api-sdk-go/constants"

type OkxConfig struct {
	BaseUrl       string
	WsUrl         string
	ApiKey        string
	SecretKey     string
	PASSPHRASE    string
	TimeoutSecond int
	SignType      string // 可选: "HMAC_SHA256" or "RSA"
}

func NewOkxConfig(ApiKey string, SecretKey string, PASSPHRASE string, TimeoutSecond int, SignType string, WsUrl string) *OkxConfig {
	if SignType == "" {
		SignType = constants.SHA256
	}
	if WsUrl == "" {
		WsUrl = "wss://ws.okx.com:8443/ws/v5/public"
	}
	return &OkxConfig{
		BaseUrl:       "https://www.okx.com",
		WsUrl:         WsUrl,
		ApiKey:        ApiKey,
		SecretKey:     SecretKey,
		PASSPHRASE:    PASSPHRASE,
		TimeoutSecond: TimeoutSecond,
		SignType:      SignType,
	}
}
