package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"myBlog/models"
	"myBlog/utils"
	"net/http"
	"net/url"
)

//GetAvatar 通过用户名获取头像信息
func GetAvatar(c *gin.Context) {
	var res models.GetAvatarResponse

	username := c.Query("username")
	if imgUrl, err := utils.GetAvatarByStr(username); err != nil {
		res.Code, res.Massage = -1, "Create avatar failed & err: ["+err.Error()+"]"
	} else {
		res.Code, res.AvatarUrl = 0, imgUrl
	}
	c.JSON(http.StatusOK, res)
}

//GetImg 获取数据库中存储的文件
func GetImg(c *gin.Context) {
	if id := c.Query("id"); id != "" {

	} else if imgUrl := c.Query("imgName"); imgUrl != "" {
		if imgName, err := url.QueryUnescape(imgUrl); err != nil {
			c.JSON(http.StatusNotFound, gin.H{})
		} else {
			fmt.Println(imgName)
			c.File("./img/" + imgName)
		}
	} else {
		c.JSON(http.StatusNotFound, gin.H{})
	}
}
