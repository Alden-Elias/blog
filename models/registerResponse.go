package models

//GetCaptchaResponse 获取图像验证码返回信息
type GetCaptchaResponse struct {
	CaptchaID  string `json:"captchaID,omitempty"`  // 验证码ID，图像验证码ID号
	CaptchaImg string `json:"captchaImg,omitempty"` // 图像验证码，图像验证码
	CodeAndMassage
}
