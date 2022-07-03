package models

//LoginResponse 用户登录响应
type LoginResponse struct {
	CodeAndMassage
	ExpireTime string `json:"expireTime ,omitempty"` // 过期时间，用户口令过期时间
	Token      string `json:"token,omitempty"`       // 口令，用户令牌
}
