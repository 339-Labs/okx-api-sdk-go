package baseclient

import (
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/339-Labs/okx-api-sdk-go/config"
	"github.com/339-Labs/okx-api-sdk-go/constants"
	"github.com/339-Labs/okx-api-sdk-go/internal"
)

type OkxRestBaseClient struct {
	HttpClient http.Client
	Signer     *Signer
	Config     *config.OkxConfig
	Auth       bool
}

func (p *OkxRestBaseClient) Init(config *config.OkxConfig) *OkxRestBaseClient {
	p.Config = config

	p.Signer = new(Signer).Init(config.SecretKey)
	p.HttpClient = http.Client{
		Timeout: time.Duration(config.TimeoutSecond) * time.Second,
	}
	return p
}

func (p *OkxRestBaseClient) DoPost(uri string, params string) (string, error) {
	timesStamp := internal.TimesStampISO()
	//body, _ := internal.BuildJsonParams(params)

	sign := p.Signer.Sign(constants.POST, uri, params, timesStamp)
	if constants.RSA == p.Config.SignType {
		sign = p.Signer.SignByRSA(constants.POST, uri, params, timesStamp)
	}
	requestUrl := p.Config.BaseUrl + uri

	buffer := strings.NewReader(params)
	request, err := http.NewRequest(constants.POST, requestUrl, buffer)

	internal.Headers(request, p.Config.ApiKey, timesStamp, sign, p.Config.PASSPHRASE)
	if err != nil {
		return "", err
	}
	response, err := p.HttpClient.Do(request)

	if err != nil {
		return "", err
	}

	defer response.Body.Close()

	bodyStr, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	responseBodyString := string(bodyStr)
	return responseBodyString, err
}

func (p *OkxRestBaseClient) DoGet(uri string, params map[string]string) (string, error) {
	timesStamp := internal.TimesStampISO()
	body := internal.BuildGetParams(params)
	//fmt.Println(body)

	sign := p.Signer.Sign(constants.GET, uri, body, timesStamp)

	requestUrl := p.Config.BaseUrl + uri + body
	request, err := http.NewRequest(constants.GET, requestUrl, nil)
	if err != nil {
		return "", err
	}
	internal.Headers(request, p.Config.ApiKey, timesStamp, sign, p.Config.PASSPHRASE)

	response, err := p.HttpClient.Do(request)

	if err != nil {
		return "", err
	}

	defer response.Body.Close()

	bodyStr, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	responseBodyString := string(bodyStr)
	return responseBodyString, err
}

func (p *OkxRestBaseClient) DoGetNoAuth(uri string, params map[string]string) (string, error) {
	body := internal.BuildGetParams(params)
	requestUrl := p.Config.BaseUrl + uri + body
	request, err := http.NewRequest(constants.GET, requestUrl, nil)
	if err != nil {
		return "", err
	}
	response, err := p.HttpClient.Do(request)

	if err != nil {
		return "", err
	}

	defer response.Body.Close()

	bodyStr, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	responseBodyString := string(bodyStr)
	return responseBodyString, err
}
