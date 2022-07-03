package controller

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"myBlog/dao"
	"myBlog/models"
	"myBlog/utils"
	"net/http"
)

func Login(c *gin.Context) {
	var req models.LoginRequest
	var res models.LoginResponse

	if err := c.BindJSON(&req); err != nil {
		res.Code, res.Massage = -1, "Request parameter parsing error. & "+err.Error()
	} else if user, err := dao.GetUserByEmail(req.Email); err != nil {
		if err == gorm.ErrRecordNotFound {
			res.Code, res.Massage = -2, "The account does not exist, do you want to register?"
		} else {
			res.Code, res.Massage = -1, "Get user failed"
		}
	} else if !utils.CheckEncryptPwdMatch(req.Password, user.Password) {
		res.Code, res.Massage = -3, "Wrong password"
	} else if token, expiredTime, err := utils.GetToken(int(user.ID), user.Email); err != nil {
		res.Code, res.Massage = -3, "Get token filed & err: ["+err.Error()+"]"
	} else {
		res.Code, res.Token, res.ExpireTime = 0, token, expiredTime
	}
	c.JSON(http.StatusOK, res)
}
