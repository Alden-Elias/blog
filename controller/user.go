package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-module/carbon"
	"gorm.io/gorm"
	"myBlog/dao"
	"myBlog/models"
	"myBlog/setting"
	"myBlog/utils"
	"net/http"
	"strconv"
)

var (
	pageSize = int(setting.Config.PageSize)
)

//ListUser 返回所有用户
func ListUser(c *gin.Context) {
	var res models.ListUsersResponse
	p := c.Query("page")
	page, err := strconv.Atoi(p)
	if err != nil {
		res.Code = -1
		res.Massage = "Request parameter parsing error. & " + err.Error()
		c.JSON(http.StatusOK, res)
		return
	}
	if page <= 0 {
		res.Code = -1
		res.Massage = "'page' should not be less than or equal to 0"
		c.JSON(http.StatusOK, res)
		return
	}

	users, err, totalPage := dao.ListUsers(page)
	if err != nil {
		res.Code = -1
		res.Massage = err.Error()
		c.JSON(http.StatusOK, res)
		return
	}
	res.PageSize = pageSize
	res.Page = page
	res.PageTotal = totalPage
	res.Users = users

	c.JSON(http.StatusOK, res)
}

//GetUserByID 通特定条件获取用户信息
func GetUserByID(c *gin.Context) {
	id := c.Param("id")
	user, err := dao.GetUserByUid(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"massage": err.Error(),
		})
		return
	}
	resp := models.GetUserByIDResponse{
		UserDetail: *user,
	}
	c.JSON(http.StatusOK, resp)
}

//UserAdd 用户添加
func UserAdd(c *gin.Context) {
	var req models.UserAddRequest
	var res models.CodeAndMassage
	if err := c.BindJSON(&req); err != nil {
		res.Code, res.Massage = -1, "Request parameter parsing error. & "+err.Error()
		c.JSON(http.StatusOK, res)
		return
	} else if encrypt, err := utils.EncryptPwd(req.Password); err != nil {
		res.Code, res.Massage = -1, "Encrypt password failed  & error : ["+err.Error()+"]"
	} else {
		user := models.User{
			Email:    req.Email,
			NickName: req.NickName,
			Password: encrypt,
			Birthday: carbon.CreateFromDate(2002, 1, 1).Carbon2Time(),
		}
		if err := dao.InsertUser(&user); err != nil {
			res.Code, res.Massage = -1, "insert user into database failed & error : ["+err.Error()+"]"
		} else {
			res.Code = 0
		}
	}
	c.JSON(http.StatusOK, res)
}

//UserLogin 用户登录
//waring 不返回cookie无状态
func UserLogin(c *gin.Context) {
	var req models.LoginRequest
	var res models.CodeAndMassage
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
	} else {
		res.Code = 0
	}
	c.JSON(http.StatusOK, res)
}
