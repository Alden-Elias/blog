package models

//UserAddRequest 用户添加请求模型
type UserAddRequest struct {
	Email    string `json:"email"`    // 用户邮箱，用户注册账号所绑定的邮箱
	NickName string `json:"nickName"` // 用户昵称，用户可自定义的昵称
	Password string `json:"password"` // 用户密码，明文密码，没啥好说的。。。
}
