package http

type AuthenticateData struct {
  UserAppID string `json:"userAppID"`
  CountryCode string `json:"countryCode"`
}

type AuthenticateResponse struct {
  Success  bool    `json:"success" valid:"required"`
  Code    string  `json:"code"`
  Message string   `json:"message"`
  Data    AuthenticateData 
}

