package baseclient

import (
	"fmt"
	"github.com/339-Labs/okx-api-sdk-go/internal/model"
	"github.com/339-Labs/okx-api-sdk-go/logging/logger"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"github.com/robfig/cron"

	"github.com/339-Labs/okx-api-sdk-go/config"
	"github.com/339-Labs/okx-api-sdk-go/constants"
	"github.com/339-Labs/okx-api-sdk-go/internal"
)

type OkxWsBaseClient struct {
	NeedLogin        bool
	Connection       bool
	LoginStatus      bool
	Listener         OnReceive
	ErrorListener    OnReceive
	Ticker           *time.Ticker
	SendMutex        *sync.Mutex
	WebSocketClient  *websocket.Conn
	LastReceivedTime time.Time
	AllSuribe        *model.Set
	Signer           *Signer
	ScribeMap        map[model.SubscribeReq]OnReceive
	Config           *config.OkxConfig
}

func (p *OkxWsBaseClient) Init(config *config.OkxConfig) *OkxWsBaseClient {
	p.Connection = false
	p.AllSuribe = model.NewSet()
	p.Signer = new(Signer).Init(config.SecretKey)
	p.ScribeMap = make(map[model.SubscribeReq]OnReceive)
	p.SendMutex = &sync.Mutex{}
	p.Ticker = time.NewTicker(constants.TimerIntervalSecond * time.Second)
	p.LastReceivedTime = time.Now()
	p.Config = config
	return p
}

func (p *OkxWsBaseClient) SetListener(msgListener OnReceive, errorListener OnReceive) {
	p.Listener = msgListener
	p.ErrorListener = errorListener
}

func (p *OkxWsBaseClient) Connect() {
	p.tickerLoop()
	p.ExecuterPing()
}

func (p *OkxWsBaseClient) ConnectWebSocket() {
	var err error
	logger.Info("WebSocket connecting...")
	p.WebSocketClient, _, err = websocket.DefaultDialer.Dial(p.Config.WsUrl, nil)
	if err != nil {
		fmt.Printf("WebSocket connected error: %s\n", err)
		return
	}
	logger.Info("WebSocket connected")
	p.Connection = true
}

func (p *OkxWsBaseClient) Login() {
	timesStamp := internal.TimesStampSec()
	sign := p.Signer.Sign(constants.WsAuthMethod, constants.WsAuthPath, "", timesStamp)
	if constants.RSA == p.Config.SignType {
		sign = p.Signer.SignByRSA(constants.WsAuthMethod, constants.WsAuthPath, "", timesStamp)
	}

	loginReq := model.WsLoginReq{
		ApiKey:     p.Config.ApiKey,
		Passphrase: p.Config.PASSPHRASE,
		Timestamp:  timesStamp,
		Sign:       sign,
	}
	var args []interface{}
	args = append(args, loginReq)

	baseReq := model.WsBaseReq{
		Op:   constants.WsOpLogin,
		Args: args,
	}
	p.SendByType(baseReq)
}

func (p *OkxWsBaseClient) StartReadLoop() {
	go p.ReadLoop()
}

func (p *OkxWsBaseClient) ExecuterPing() {
	c := cron.New()
	_ = c.AddFunc("*/15 * * * * *", p.ping)
	c.Start()
}
func (p *OkxWsBaseClient) ping() {
	p.Send("ping")
}

func (p *OkxWsBaseClient) SendByType(req model.WsBaseReq) {
	json, _ := internal.ToJson(req)
	p.Send(json)
}

func (p *OkxWsBaseClient) Send(data string) {
	if p.WebSocketClient == nil {
		logger.Error("WebSocket sent error: no connection available")
		return
	}
	logger.Info("sendMessage:%s", data)
	p.SendMutex.Lock()
	err := p.WebSocketClient.WriteMessage(websocket.TextMessage, []byte(data))
	p.SendMutex.Unlock()
	if err != nil {
		logger.Error("WebSocket sent error: data=%s, error=%s", data, err)
	}
}

func (p *OkxWsBaseClient) tickerLoop() {
	logger.Info("tickerLoop started")
	for {
		select {
		case <-p.Ticker.C:
			elapsedSecond := time.Now().Sub(p.LastReceivedTime).Seconds()

			if elapsedSecond > constants.ReconnectWaitSecond {
				logger.Info("WebSocket reconnect...")
				p.disconnectWebSocket()
				p.ConnectWebSocket()
			}
		}
	}
}

func (p *OkxWsBaseClient) disconnectWebSocket() {
	if p.WebSocketClient == nil {
		return
	}

	fmt.Println("WebSocket disconnecting...")
	err := p.WebSocketClient.Close()
	if err != nil {
		logger.Error("WebSocket disconnect error: %s\n", err)
		return
	}

	logger.Info("WebSocket disconnected")
}

func (p *OkxWsBaseClient) ReadLoop() {
	for {

		if p.WebSocketClient == nil {
			logger.Info("Read error: no connection available")
			//time.Sleep(TimerIntervalSecond * time.Second)
			continue
		}

		_, buf, err := p.WebSocketClient.ReadMessage()
		if err != nil {
			logger.Info("Read error: %s", err)
			continue
		}
		p.LastReceivedTime = time.Now()
		message := string(buf)

		logger.Info("rev:" + message)

		if message == "pong" {
			logger.Info("Keep connected:" + message)
			continue
		}
		jsonMap := internal.JSONToMap(message)

		v, e := jsonMap["code"]

		if e && v != "0" {
			p.ErrorListener(message)
			continue
		}

		v, e = jsonMap["event"]
		if e && v == "login" {
			logger.Info("login msg:" + message)
			p.LoginStatus = true
			continue
		}

		v, e = jsonMap["data"]
		if e {
			listener := p.GetListener(jsonMap["arg"])
			listener(message)
			continue
		}
		p.handleMessage(message)
	}

}

func (p *OkxWsBaseClient) GetListener(argJson interface{}) OnReceive {

	mapData := argJson.(map[string]interface{})

	subscribeReq := model.SubscribeReq{
		InstType: fmt.Sprintf("%v", mapData["instType"]),
		Channel:  fmt.Sprintf("%v", mapData["channel"]),
		InstId:   fmt.Sprintf("%v", mapData["instId"]),
	}

	v, e := p.ScribeMap[subscribeReq]

	if !e {
		return p.Listener
	}
	return v
}

type OnReceive func(message string)

func (p *OkxWsBaseClient) handleMessage(msg string) {
	fmt.Println("default:" + msg)
}
