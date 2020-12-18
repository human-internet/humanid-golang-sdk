package payload

type LoginDataResp struct {
	WebLoginUrl string `json:"webLoginUrl"`
}

type LoginResp struct {
	Success bool          `json:"success"`
	Code    string        `json:"code"`
	Message string        `json:"message"`
	Data    LoginDataResp `json:"data"`
}

type VerifyTokenDataResp struct {
	UserAppId string `json:"userAppId"`
	CountryCode string `json:"countryCide`
}

type VerifyTokenResp struct {
	Success bool          `json:"success"`
	Code    string        `json:"code"`
	Message string        `json:"message"`
	Data    VerifyTokenDataResp `json:"data"`
}

type VerifyTokenReq struct {
	ExchangeToken    string        `json:"exchangeToken"`
}

