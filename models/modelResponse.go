package models

type CodeAndMassage struct {
	Code    int    `json:"code"`              // 状态码，目前非0表示获取失败
	Massage string `json:"massage,omitempty"` // 错误信息，当返回错误时提供错误信息
}
