package models

//GetAvatarResponse 获取头像回复
type GetAvatarResponse struct {
	CodeAndMassage
	AvatarUrl string `json:"avatarUrl,omitempty"` // 头像，头像
}
