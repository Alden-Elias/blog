package models

import (
	"gorm.io/gorm"
	"time"
)

// User 用户信息表
type User struct {
	gorm.Model
	NickName  string    `gorm:"default:null"`                     //用户昵称
	Password  string    `gorm:"unique_index;not null"`            //明文密码
	Email     string    `gorm:"unique;unique_index;default:null"` //邮箱
	Telephone string    `gorm:"unique_index;default:null"`        //手机号码
	Birthday  time.Time //用户生日
	Sex       string    `gorm:"check:sex='男' OR sex='女' OR ISNULL(sex);default:null"` //用户性别
	IsAdmin   bool      `gorm:"default:false"`                                        //是否为管理员
	IsLock    bool      `gorm:"default:false"`                                        //是否被关小黑屋
	AvatarUrl string    `gorm:"default:null"`                                         // 头像链接
}
