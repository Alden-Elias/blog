package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-module/carbon"
	"myBlog/dao"
	"myBlog/models"
	"myBlog/utils"
	"net/http"
)

func GetCaptcha(c *gin.Context) {
	var res models.GetCaptchaResponse

	id, b64s, err := utils.CaptGet()
	if err != nil {
		res.Code = -1
		res.Massage = err.Error()
		c.JSON(http.StatusOK, res)
	}
	res.CaptchaImg = b64s
	res.CaptchaID = id
	c.JSON(http.StatusOK, res)
}

func VerifyCaptcha(c *gin.Context) {
	var res models.CodeAndMassage
	var req models.VerifyCaptchaRequest

	if err := c.BindJSON(&req); err != nil {
		res.Code, res.Massage = -1, "Request parameter parsing error."
	} else if !utils.CaptVerify(req.CaptchaID, req.Captcha) {
		res.Code, res.Massage = -1, "verify failed & id: "+req.CaptchaID
	} else {
		res.Code = 0
	}
	c.JSON(http.StatusOK, res)
}

func Register(c *gin.Context) {
	var req models.RegisterRequest
	var res models.CodeAndMassage

	if c.BindJSON(&req) != nil {
		res.Code, res.Massage = -1, "Request parameter parsing error."
	} else if !utils.CaptVerify(req.CaptchaID, req.Captcha) {
		res.Code, res.Massage = -1, "verify failed & id: "+req.CaptchaID
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
