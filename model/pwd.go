package model

//swagger:model
type PwdResponse struct {
	Response
	Data struct {
		//券码
		ExchangePwd string `json:"exchange_pwd"`
	} `json:"data"`
}
