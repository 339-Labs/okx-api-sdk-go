package ws

import (
	"strings"

	"github.com/339-Labs/okx-api-sdk-go/config"
	"github.com/339-Labs/okx-api-sdk-go/constants"
	"github.com/339-Labs/okx-api-sdk-go/internal/baseclient"
	"github.com/339-Labs/okx-api-sdk-go/internal/model"
	"github.com/339-Labs/okx-api-sdk-go/logging/logger"
)

type OkxWSClient struct {
	okxWsBaseClient *baseclient.OkxWsBaseClient
	NeedLogin       bool
}

func (p *OkxWSClient) Init(config *config.OkxConfig, needLogin bool, listener baseclient.OnReceive, errorListener baseclient.OnReceive) *OkxWSClient {
	p.okxWsBaseClient = new(baseclient.OkxWsBaseClient).Init(config)
	p.okxWsBaseClient.SetListener(listener, errorListener)
	p.okxWsBaseClient.ConnectWebSocket()
	p.okxWsBaseClient.StartReadLoop()
	p.okxWsBaseClient.ExecuterPing()

	if needLogin {
		logger.Info("login in ...")
		p.okxWsBaseClient.Login()
		for {
			if !p.okxWsBaseClient.LoginStatus {
				continue
			}
			break
		}
		logger.Info("login in ... success")
	}
	return p

}

func (p *OkxWSClient) Connect() *OkxWSClient {
	p.okxWsBaseClient.Connect()
	return p
}

func (p *OkxWSClient) UnSubscribe(list []model.SubscribeReq) {

	var args []interface{}
	for i := 0; i < len(list); i++ {
		delete(p.okxWsBaseClient.ScribeMap, list[i])
		p.okxWsBaseClient.AllSuribe.Add(list[i])
		p.okxWsBaseClient.AllSuribe.Remove(list[i])
		args = append(args, list[i])
	}

	wsBaseReq := model.WsBaseReq{
		Op:   constants.WsOpUnsubscribe,
		Args: args,
	}

	p.SendMessageByType(wsBaseReq)
}

func (p *OkxWSClient) SubscribeDef(list []model.SubscribeReq) {

	var args []interface{}
	for i := 0; i < len(list); i++ {
		req := toUpperReq(list[i])
		args = append(args, req)
	}
	wsBaseReq := model.WsBaseReq{
		Op:   constants.WsOpSubscribe,
		Args: args,
	}

	p.SendMessageByType(wsBaseReq)
}

func toUpperReq(req model.SubscribeReq) model.SubscribeReq {
	req.Channel = strings.ToLower(req.Channel)
	req.InstType = strings.ToUpper(req.InstType)
	req.InstId = strings.ToUpper(req.InstId)
	req.InstFamily = strings.ToUpper(req.InstFamily)
	return req

}

func (p *OkxWSClient) Subscribe(list []model.SubscribeReq, listener baseclient.OnReceive) {

	var args []interface{}
	for i := 0; i < len(list); i++ {
		req := toUpperReq(list[i])
		args = append(args, req)

		p.okxWsBaseClient.ScribeMap[req] = listener
		p.okxWsBaseClient.AllSuribe.Add(req)
		args = append(args, req)
	}

	wsBaseReq := model.WsBaseReq{
		Op:   constants.WsOpSubscribe,
		Args: args,
	}

	p.okxWsBaseClient.SendByType(wsBaseReq)
}

func (p *OkxWSClient) SendMessage(msg string) {
	p.okxWsBaseClient.Send(msg)
}

func (p *OkxWSClient) SendMessageByType(req model.WsBaseReq) {
	p.okxWsBaseClient.SendByType(req)
}
