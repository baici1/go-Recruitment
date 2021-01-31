package api

import (
	"fmt"
	stu "go-Recruitment/dao"
	"go-Recruitment/mysql"
	"go-Recruitment/tool/jwt"
	"net/http"
	"strconv"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/gin-gonic/gin"
)

//@Title 登录
//@Description 用于招新网站的登录
//@Param stu_id formData string true "学号"
//@Param password formData string true "密码"
//@Success 200 {json} json ""msg":"登录成功","token":token,"code":200,"data":{}"
//@Failure 500 "获取账号信息出错"
//@Failure 404 "未找到此用户"
//@Router /login [POST]
func Login(c *gin.Context) {
	
	var u mysql.User
	err:=c.ShouldBind(&u)
	if err != nil {
		c.JSON(500,gin.H{
			"msg":"获取账号信息出错",
			"code":500,
		})
		return
	}
	u,err=mysql.Queryonedata(u.Stu_id,u.Password)
	if err==nil {
		token,_:=jwt.GetToken(u.Id,u.Stu_id,u.Password)
		c.JSON(http.StatusOK,gin.H{
			"msg":"登录成功",
			"token":token,
			"code":200,
			"data":gin.H{
				"id":u.Id,
				"stu_id":u.Stu_id,
				"password":u.Password,
			},
		})
		return
	}else{
		c.JSON(404,gin.H{
			"msg":"未找到此用户",
			"code":404,
		})
	}
}
//注册
func Register(c *gin.Context)  {
	 var f stu.User
	var u mysql.User
	err:=c.ShouldBind(&u)
	 
	if err != nil {
		c.JSON(500,gin.H{
			"msg":"获取账号信息出错",
			"code":500,
		})
		return 
	}
	_,err=mysql.Queryonedata(u.Stu_id,u.Password)
	
	if err==nil {
		c.JSON(500,gin.H{
			"msg":"此用户已存在",
			"code":500,
		})
		return
	}
	err=mysql.Addonedata(u.Stu_id,u.Password)
	if err != nil {
		c.JSON(500,gin.H{
			"msg":"注册失败",
			"code":500,
		})
		return
	}else{

		token,_:=jwt.GetToken(u.Id,u.Stu_id,u.Password)
		c.JSON(200,gin.H{
			"msg":"注册成功",
			"code":200,
			"token":token,
			"data":gin.H{
				"id":u.Id,
				"stu_id":u.Stu_id,
				"password":u.Password,
			},
		})
		mysql.UpdateoneForm(f.Real_name,f.Group_id,f.Sex,f.College,f.Major,f.Phone,f.Qq,f.Result,f.Code,u.Stu_id)
	}
}

//报名信息
func Postform(c *gin.Context)  {
	var f stu.User
	err:=c.ShouldBind(&f)
	if err != nil {
		c.JSON(500,gin.H{
			"msg":"获取表单数据出错",
			"code":500,
		})
		return 
	}
	err=mysql.UpdateoneForm(f.Real_name,f.Group_id,f.Sex,f.College,f.Major,f.Phone,f.Qq,f.Result,f.Code,f.Stu_id)
	if err != nil {
		c.JSON(500,gin.H{
			"msg":"获取表单数据出错",
			"code":500,
		})
		return 
	}
	c.JSON(200,gin.H{
		"msg":"成功",
		"code":200,
		"data":gin.H{
			"real_name":f.Real_name,
			"group_id":f.Group_id,
			"sex":f.Sex,
			"college":f.College,
			"major":f.Major,
			"phone":f.Phone,
			"qq":f.Qq,
			"result":f.Result,
			"code":f.Code,
		},
		
	})
}

//获取所有信息
func Alldata(c *gin.Context){
	var all stu.User
	err:=c.ShouldBind(&all)
	if err != nil {
		c.JSON(500,gin.H{
			"msg":"获取信息失败",
			"code":500,
		})
		return 
	}
	all,err=mysql.Querydata(all.Stu_id)
	if err != nil {
		c.JSON(404,gin.H{
			"msg":"未找到此用户",
			"code":404,
		})
		return 
	}
	c.JSON(200,gin.H{
		"msg":"获取成功",
		"code":200,
		"data":all,
	})
}


//将数据库中数据生成excel

func Getexcel(c *gin.Context)  {
	// var rows=[...]string{"A","B","C","D","E","F","G","H","I","J","K","L"}
	all,err:=mysql.Queryalldata()
	if err != nil {
		fmt.Println("failed")
	}
	xlsx := excelize.NewFile()
	xlsx.SetCellValue("Sheet1","A1","ID")
	xlsx.SetCellValue("Sheet1","B1","Stu_id")
	xlsx.SetCellValue("Sheet1","C1","Password")
	xlsx.SetCellValue("Sheet1","D1","Real_name")
	xlsx.SetCellValue("Sheet1","E1","Group_id")
	xlsx.SetCellValue("Sheet1","F1","Sex")
	xlsx.SetCellValue("Sheet1","G1","College")
	xlsx.SetCellValue("Sheet1","H1","Major")
	xlsx.SetCellValue("Sheet1","I1","Phone")
	xlsx.SetCellValue("Sheet1","J1","Qq")
	xlsx.SetCellValue("Sheet1","K1","Result")
	xlsx.SetCellValue("Sheet1","L1","Code")
	for key,value:=range all{
		fmt.Println("A"+strconv.Itoa(key+2))
		var num string
		num=strconv.Itoa(key+2)
		xlsx.SetCellValue("Sheet1","A"+num,value.ID)
		xlsx.SetCellValue("Sheet1","B"+num,value.Stu_id)
		xlsx.SetCellValue("Sheet1","C"+num,value.Password)
		xlsx.SetCellValue("Sheet1","D"+num,value.Real_name)
		xlsx.SetCellValue("Sheet1","E"+num,value.Group_id)
		xlsx.SetCellValue("Sheet1","F"+num,value.Sex)
		xlsx.SetCellValue("Sheet1","G"+num,value.College)
		xlsx.SetCellValue("Sheet1","H"+num,value.Major)
		xlsx.SetCellValue("Sheet1","I"+num,value.Phone)
		xlsx.SetCellValue("Sheet1","J"+num,value.Qq)
		xlsx.SetCellValue("Sheet1","K"+num,value.Result)
		xlsx.SetCellValue("Sheet1","L"+num,value.Code)
	}
	// xlsx.SetCellValue("Sheet1", "A2", "我要下载一个excel文件")
	// xlsx.SetCellValue("Sheet1", "A1", "有没有看到我帅气的脸庞")

	//保存文件方式
	//_ = xlsx.SaveAs("./aaa.xlsx")

	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Disposition", "attachment; filename="+"Workbook.xlsx")
	c.Header("Content-Transfer-Encoding", "binary")

	//回写到web 流媒体 形成下载
	_ = xlsx.Write(c.Writer)
}

