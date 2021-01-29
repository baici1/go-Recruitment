package api

import (
	"go-Recruitment/mysql"
	"net/http"

	"github.com/gin-gonic/gin"
)


func Login(c *gin.Context) {
	var u mysql.User
	err:=c.ShouldBind(&u)
	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"msg":"获取账号信息出错",
		})
		return
	}
	_,err=mysql.Queryonedata(u.Stu_id,u.Password)
	if err==nil {
		c.JSON(http.StatusOK,gin.H{
			"msg":"登录成功",
		})
		return
	}else{
		c.JSON(400,gin.H{
			"msg":"未找到此用户",
		})
	}
}

func Register(c *gin.Context)  {
	var u mysql.User
	err:=c.ShouldBind(&u)
	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"msg":"获取账号信息出错",
		})
		return 
	}
	_,err=mysql.Queryonedata(u.Stu_id,u.Password)
	if err==nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"msg":"此用户已存在",
		})
		return
	}
	err=mysql.Addonedata(u.Stu_id,u.Password)
	if err != nil {
		c.JSON(500,gin.H{
			"msg":"注册失败",
		})
		return
	}else{
		c.JSON(200,gin.H{
			"msg":"注册成功",
		})
	}
}