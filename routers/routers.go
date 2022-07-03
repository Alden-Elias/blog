package routers

import (
	"github.com/gin-gonic/gin"
	"myBlog/controller"
	"net/http"
)

func Start() {
	engine := gin.Default()
	engine.Static("static", "/static")
	engine.StaticFS("/img", http.Dir("./img"))

	api := engine.Group("api")
	{
		//User 用户信息相关
		user := api.Group("/user")
		{
			user.GET("/list", controller.ListUser)
			user.GET("/:id", controller.GetUserByID)
			user.POST("/add", controller.UserAdd)
			user.POST("/login", controller.UserLogin)
		}

		//Register 用户注册相关
		register := api.Group("/register")
		{
			register.GET("/captcha", controller.GetCaptcha)
			register.POST("/verifyCaptcha", controller.VerifyCaptcha)
			register.POST("", controller.Register)
		}

		//Login 用户登录相关
		login := api.Group("/login")
		{
			login.POST("", controller.Login)
		}

		//Utils 工具类接口
		utils := api.Group("/utils")
		{
			utils.GET("/avatar", controller.GetAvatar)
			utils.GET("/img", controller.GetImg)
		}
	}
	engine.Run()
}
