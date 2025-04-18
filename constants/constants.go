package constants

const (
	/*
	 * http headers
	 */
	ContentType         = "Content-Type"
	OkxAccessKey        = "OK-ACCESS-KEY"
	OkxAccessSign       = "OK-ACCESS-SIGN"
	OkxAccessTimestamp  = "OK-ACCESS-TIMESTAMP"
	OkxAccessPassphrase = "OK-ACCESS-PASSPHRASE"
	ApplicationJson     = "application/json"

	EN_US  = "en_US"
	ZH_CN  = "zh_CN"
	LOCALE = "locale="

	/*
	 * http methods
	 */
	GET  = "GET"
	POST = "POST"

	/*
	 * websocket
	 */
	WsAuthMethod        = "GET"
	WsAuthPath          = "/users/self/verify"
	WsOpLogin           = "login"
	WsOpUnsubscribe     = "unsubscribe"
	WsOpSubscribe       = "subscribe"
	TimerIntervalSecond = 5
	ReconnectWaitSecond = 60

	/*
	 * SignType
	 */
	RSA    = "RSA"
	SHA256 = "SHA256"
)
