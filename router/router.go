package router

import (
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
	r.POST("/login",api.Login )//登录
	r.POST("/register",api.Register)//注册
	userGroup:=r.Group("/user")
	{
		userGroup.POST("/form",api.Postform)
		userGroup.GET("/alldata",api.Alldata)
	}
	r.POST("/form",api.Postformdata)
	r.GET("/download",api.Getexcel)
	r.GET("/getalldata",api.GetAlldatamysql)
	r.GET("/delete",api.DeleteData)
	r.POST("/result",api.Getresult)
	r.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))
	r.Run(":8081")
}