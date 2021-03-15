package api_stu

import (
	stu "go-Recruitment/dao"
	"go-Recruitment/mysql"

	"github.com/gin-gonic/gin"
)

func Updatedes(c *gin.Context) {
	var all stu.Des
	err:=c.ShouldBind(&all)
	if err != nil {
		c.JSON(500,gin.H{
			"msg":"获取信息失败",
			"code":"500",
		})
		return 
	}
	err=mysql.Updatedes(all.Comments,all.Self_study,all.Attach,all.Development,all.Ready,all.Degree,all.Stu_id,all.Grades)
	if err != nil {
		c.JSON(403,gin.H{
			"msg":"提交表单数据出错",
			"code":403,
		})
		return 
	}
	c.JSON(200,gin.H{
		"code":20000,
		"msg":"修改成功",
		"data":all,
	})
}

func Getdes(c *gin.Context)  {
	var all stu.Des
	err:=c.ShouldBind(&all)
	if err != nil {
		c.JSON(500,gin.H{
			"msg":"获取信息失败",
			"code":"500",
			
		})
		return 
	}
	all,err=mysql.Querydes(all.Stu_id)
	if err != nil {
		c.JSON(403,gin.H{
			"msg":"提交表单数据出错",
			"code":403,
		})
		return
	}
	c.JSON(200,gin.H{
		"msg":"获取成功",
		"code":20000,
		"data":all,
	})
}