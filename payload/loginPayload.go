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
