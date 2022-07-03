package dao

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/utils"
	"log"
	"myBlog/models"
	"myBlog/setting"
)

var db *gorm.DB

func init() {
	conf := setting.Config
	dsn := conf.MysqlUsername + ":" + conf.MysqlPassword + "@tcp(" + utils.ToString(conf.MysqlHost) + ":" + utils.ToString(conf.MysqlPort) + ")/" + conf.MysqlDatabase + "?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(&models.User{})
}
