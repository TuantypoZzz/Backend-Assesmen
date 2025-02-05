package model

type GeneralResponse struct {
	Code   int         `json:"code"`
	Remark string      `json:"remark"`
	Data   interface{} `json:"data,omitempty"`
}
