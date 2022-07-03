package dao

import (
	"errors"
	"fmt"
	"myBlog/models"
	"myBlog/setting"
)

func init() {
	db.Model(&models.User{}).Set("gorm:table_options", "AUTO_INCREMENT=10000")
}

//InsertUser 添加用户
func InsertUser(user *models.User) error {
	return db.Create(user).Error
}

//UpdateUser 更新用户信息
func UpdateUser(user *models.User) error {
	return db.Save(user).Error
}

//GetUserByUid 通过用户ID获取用户信息
func GetUserByUid(uid string) (*models.UserDetail, error) {
	var user models.UserDetail
	err := db.Model(&models.User{}).First(&user, uid).Error
	return &user, err
}

//GetUserByUsername 通过用户名获取用户信息
func GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	err := db.First(&user, "username=?", username).Error
	return &user, err
}

//GetUserByEmail 通过用户注册邮箱获取用户
func GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := db.First(&user, "email=?", email).Error
	return &user, err
}

//GetUsersByNickName 通过用户昵称查找用户
func GetUsersByNickName(nickName string) (*[]models.User, error) {
	var users []models.User
	err := db.Find(&users, "nick_name=?", nickName).Error
	return &users, err
}

//LockUser 将用户关小黑屋
func LockUser(user *models.User) error {
	return db.Model(user).UpdateColumn("is_lock", true).Error
}

//UnlockUser 从小黑屋放出用户
func UnlockUser(user *models.User) error {
	return db.Model(user).UpdateColumn("is_lock", false).Error
}

//ListUsers 列出全部用户（非管理员）
func ListUsers(page int) (users []*models.UserMain, err error, totalPage int) {
	if page <= 0 {
		err = errors.New("'page' should not be less than or equal to 0")
		return
	}

	var total int64
	if err = db.Model(&models.User{}).Where("is_admin=?", false).Count(&total).Error; err != nil {
		return
	}

	pageSize := int(setting.Config.PageSize)

	totalPage = (int(total) + pageSize - 1) / pageSize
	if page > totalPage {
		err = fmt.Errorf("'page' should be less than or equal to %d", totalPage)
		return
	}

	err = db.Model(&models.User{}).Limit(pageSize).Offset(pageSize*(page-1)).Find(&users, "is_admin=?", false).Error
	return
}
