package models

//LoginRequest 用户登录请求
type LoginRequest struct {
	Email    string `json:"email"`    // 用户邮箱，用户注册账号所绑定的邮箱
	Password string `json:"password"` // 用户密码，明文密码，没啥好说的。。。
}
