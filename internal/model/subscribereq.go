package model

type SubscribeReq struct {
	InstType   string `json:"instType"`
	Channel    string `json:"channel"`
	InstFamily string `json:"instFamily"`
	InstId     string `json:"instId"`
}
