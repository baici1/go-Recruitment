package router

import (
	"go-Recruitment/api/api"
	"go-Recruitment/api/api_stu"
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
		userGroup.POST("/form",api.Postform)//提交表单信息
		userGroup.GET("/alldata",api.Alldata)
	}
	r.POST("/form",api.Postformdata)//提交网站招新表单
	r.GET("/download",api.Getexcel)//下载excel
	r.GET("/getalldata",api.GetAlldatamysql)//获取所有的信息
	r.GET("/delete",api.DeleteData)//删除
	r.POST("/result",api.Getresult)//获取结果
	r.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))//swagger
	des:=r.Group("/des")
	{
		des.POST("/update",api_stu.Updatedes)
		des.GET("/get",api_stu.Getdes)
	}
	r.Run(":8081")
}