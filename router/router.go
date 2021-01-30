package router

import (
	JWT "go-Recruitment/HandlerFunc/jwt"
	"go-Recruitment/api"
	"go-Recruitment/tool/middleware"

	_ "go-Recruitment/docs"

	"github.com/gin-gonic/gin"
	gs "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func InitRoute() {
	r := gin.Default()
	r.Use(middleware.Cors())//跨域处理
	r.GET("/login",api.Login )//登录
	r.POST("/register",api.Register)//注册
	userGroup:=r.Group("/user")
	{
		userGroup.POST("/form",JWT.JWTAuthMiddleware,api.Postform)
		userGroup.GET("/alldata",JWT.JWTAuthMiddleware,api.Alldata)
	}
	r.GET("/download",api.Getexcel)
	r.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))
	r.Run(":8090")
}