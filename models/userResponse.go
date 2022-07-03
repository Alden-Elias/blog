package models

// UserMain 用户主要信息
type UserMain struct {
	AvatarUrl string `json:"avatar"`   // 用户头像，用户头像地址
	Email     string `json:"email"`    // 用户邮箱，用户注册账号所绑定的邮箱
	NickName  string `json:"nickName"` // 用户昵称，用户可自定义的昵称
	Password  string `json:"password"` // 用户密码，明文密码，没啥好说的。。。
	ID        string `json:"uID"`      // 用户ID，用户唯一标识，具有唯一性且创建后不变
	IsAdmin   bool   `json:"isAdmin"`  // 管理员标识，是否为管理员
}

//UserDetail 用户详情
type UserDetail struct {
	Birthday  string            `json:"birthday,omitempty"`  // 生日，用户生日
	Sex       string            `json:"sex,omitempty"`       // 性别，用户性别（男 | 女 | 未知）
	Telephone string            `json:"telephone,omitempty"` // 手机号，用户设置的手机号
	UserMain  `json:"userMain"` // 用户详情
}

//GetUserByIDResponse 通过用户ID获取用户详细信息响应结构体
type GetUserByIDResponse struct {
	CodeAndMassage
	UserDetail UserDetail `json:"userDetail,omitempty"`
}

//ListUsersResponse 列出用户的响应结构体
type ListUsersResponse struct {
	CodeAndMassage
	Page      int         `json:"page,omitempty"`      // 获取页，无特殊要求下与请求参数相同
	PageSize  int         `json:"pageSize,omitempty"`  // 页面大小，页面下存放的最大数据条数
	PageTotal int         `json:"pageTotal,omitempty"` // 总页面数，总页面数量
	Users     []*UserMain `json:"users,omitempty"`     // 用户列表，指定分页下的用户列表
}
