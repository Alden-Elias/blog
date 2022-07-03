package models

//VerifyCaptchaRequest 验证码确认请求
type VerifyCaptchaRequest struct {
	Captcha   string `json:"captcha"`   // 验证码，用户填写的验证码
	CaptchaID string `json:"captchaID"` // 图像验证码ID，获取图像验证码时返回的ID号
}

//RegisterRequest 注册请求参数
type RegisterRequest struct {
	Captcha   string `json:"captcha"`   // 验证码，用户输入的验证码
	CaptchaID string `json:"captchaID"` // 图像验证码ID，获取图像验证码时返回的ID号
	Email     string `json:"email"`     // 用户邮箱，用户注册账号所绑定的邮箱
	NickName  string `json:"nickName"`  // 用户昵称，用户可自定义的昵称
	Password  string `json:"password"`  // 用户密码，明文密码，没啥好说的。。。
}
