package router

import (
	"go-Recruitment/api"
	"go-Recruitment/tool/middleware"

	"github.com/gin-gonic/gin"
)

func InitRoute() {
	r := gin.Default()
	r.Use(middleware.Cors())//跨域处理
	r.GET("/login",api.Login )//登录
	r.POST("/register",api.Register)//注册
	r.Run(":8090")
}