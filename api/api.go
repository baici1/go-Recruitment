package api

import (
	"fmt"
	stu "go-Recruitment/dao"
	"go-Recruitment/mysql"
	"go-Recruitment/tool/jwt"
	"net/http"
	"os"
	"strconv"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/gin-gonic/gin"
)

//@Title 登录
//@Description 用于招新网站的登录
//@Summary 获取账号进行登录
//@Accept multipart/form-data
//@Produce application/json
//@Param stu_id formData string true "学号"
//@Success 200 {json} json ""msg":"登录成功","token":token,"code":20000,"data":{}"
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
	u,err=mysql.Queryonedata(u.Stu_id)
	if err==nil {
		token,_:=jwt.GetToken(u.Id,u.Stu_id,u.Password)
		c.JSON(http.StatusOK,gin.H{
			"msg":"登录成功",
			"token":token,
			"code":20000,
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






//@Title 注册
//@Description 用于招新网站的注册
//@Summary 用于注册
//@Accept multipart/form-data
//@Produce application/json
//@Param stu_id formData string true "学号"
//@Success 200 {json} json ""msg":"注册成功","token":token,"code":20000,"data":{}"
//@Failure 500 "获取账号信息出错"
//@Failure 404 "未找到此用户"
//@Failure 403 "用户已存在"
//@Router /register [POST]
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
	_,err=mysql.Queryonedata(u.Stu_id)
	
	if err==nil {
		c.JSON(500,gin.H{
			"msg":"此用户已存在",
			"code":403,
		})
		return
	}
	err=mysql.Addonedata(u.Stu_id)
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
			"code":20000,
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

//@Title 报名信息
//@Description 用于招新网站后台的表单信息
//@Summary 更改表单信息(用于后台)
//@Accept multipart/form-data
//@Produce application/json
//@Param stu_id formData string true "学号"
//@Param real_name formData string false "真实姓名"
//@Param group_id formData string false "1开发组0智能组"
//@Param sex formData string false "0男1女"
//@Param college formData string false "学院"
//@Param major formData string false "专业"
//@Param phone formData string false "手机"
//@Param qq formData string false "qq"
//@Param result formData string false "结果:1录取   0未录取"
//@Param code formData string false "成绩"
//@Success 200 {json} json ""msg":"获取成功""
//@Failure 500 "获取表单数据出错"
//@Failure 403 "提交表单数据出错"
//@Router /user/form [POST]
func Postform(c *gin.Context)  {
	var f stu.User
	err:=c.ShouldBind(&f)
	// fmt.Println(f)
	// fmt.Println(err)
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
			"msg":"提交表单数据出错",
			"code":403,
		})
		return 
	}
	c.JSON(200,gin.H{
		"msg":"成功",
		"code":20000,
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


//@Title 报名信息
//@Description 用于招新网站的表单信息
//@Summary 获取表单信息
//@Accept multipart/form-data
//@Produce application/json
//@Param stu_id formData string true "学号"
//@Param real_name formData string false "真实姓名"
//@Param group_id formData string false "1开发组0智能组"
//@Param sex formData string false "0男1女"
//@Param college formData string false "学院"
//@Param major formData string false "专业"
//@Param phone formData string false "手机"
//@Param qq formData string false "qq"
//@Param result formData string false "结果:1录取   0未录取"
//@Param code formData string false "成绩"
//@Success 200 {json} json ""msg":"获取成功""
//@Failure 500 "获取表单数据出错"
//@Failure 403 "提交表单数据出错"
//@Router /form [POST]
func Postformdata(c *gin.Context)  {
	var f stu.User
	err:=c.ShouldBind(&f)
	// fmt.Println(f)
	// fmt.Println(err)
	if err != nil {
		c.JSON(500,gin.H{
			"msg":"获取表单数据出错",
			"code":500,
		})
		return 
	}
	_,err=mysql.Queryonedata(f.Stu_id)
	if err == nil {
		c.JSON(400,gin.H{
			"msg":"用户重复提交",
		})
		return 
	}
	err=mysql.Addalldata(f.Stu_id,f.Real_name,f.Group_id,f.Sex,f.College,f.Major,f.Phone,f.Qq)
	if err != nil {
		c.JSON(403,gin.H{
			"msg":"提交表单数据出错",
			"code":403,
		})
		return 
	}
	c.JSON(200,gin.H{
		"msg":"成功",
		"code":20000,
		"data":gin.H{
			"stu_id":f.Stu_id,
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
	mysql.UpdateoneForm(f.Real_name,f.Group_id,f.Sex,f.College,f.Major,f.Phone,f.Qq,f.Result,f.Code,f.Stu_id)
}



//@Title 获取个人全部信息
//@Description 用于获取个人全部信息
//@Summary 用于获取个人全部信息
//@Accept multipart/form-data
//@Produce application/json
//@Param real_name query string true "名字"
//@Success 200 {json} json ""msg":"获取成功""
//@Failure 404 "未找到此用户"
//@Router /user/alldata [GET]
func Alldata(c *gin.Context){
	var all []stu.User
	real_name:=c.Query("real_name")
	// if err != nil {
	// 	c.JSON(500,gin.H{
	// 		"msg":"获取信息失败",
	// 		"code":500,
	// 	})
	// 	return 
	// }
	// fmt.Println(stuID)
	all,err:=mysql.Querydata(real_name)
	if err != nil {
		c.JSON(404,gin.H{
			"msg":"未找到此用户",
			"code":404,
		})
		return 
	}
	c.JSON(200,gin.H{
		"msg":"获取成功",
		"code":20000,
		"data":all,
	})
}


//@Title 获取数据库数据下载成excel
//@Description 用于数据库数据下载成excel
//@Summary 用于数据库数据下载成excel(直接调用api)
//@Accept multipart/form-data
//@Produce application/json
//@Success 200 {json} json ""msg":"下载成功""
//@Failure 500 "获取全部信息失败"
//@Failure 408 "请求时间超时,下载失败"
//@Router /download [GET]
func Getexcel(c *gin.Context)  {
	// var rows=[...]string{"A","B","C","D","E","F","G","H","I","J","K","L"}
	all,err:=mysql.Queryalldata("100",0)
	if err != nil {
		c.JSON(500,gin.H{
			"msg":"获取全部信息失败",
			"code":500,
		})
		return
	}
	xlsx := excelize.NewFile()
	xlsx.SetCellValue("Sheet1","A1","ID")// 设置单元格的值
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
		// fmt.Println("A"+strconv.Itoa(key+2))
		if value.Sex=="1" {
			value.Sex="女"
		}else {
			value.Sex="男"
		}
		if value.Group_id=="1" {
			value.Group_id="开发组"
		}else {
			value.Group_id="智能组"
		}
		if value.Result=="1" {
			value.Result="录取"
		}else {
			value.Result="未录取"
		}
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

	//保存文件方式
	if err := xlsx.SaveAs("Workbook.xlsx"); err != nil {
		fmt.Println(err)
	}
	// var filename = "./Workbook.xlsx"
	// file, err := os.Create(filename)
	  _, err = os.Open("Workbook.xlsx")//打开文件
	// content, err := ioutil.ReadFile("./Workbook.xlsx")
	// buf := bufio.NewWriter(file) //创建新的 Writer 对象
	// 	buf.WriteString("test")
	// 	buf.Flush()
	//  defer file.Close()
	// content, err := ioutil.ReadFile("./main.go")
	// content, err := ioutil.ReadFile("./Workbook.xlsx")
	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Disposition", "attachment; filename="+"Workbook.xlsx")
	c.Header("Content-Transfer-Encoding", "binary")
	//  c.Writer.Write(content)
	 c.File("Workbook.xlsx")
	// //回写到web 流媒体 形成下载
	//  err = xlsx.Write(c.Writer)
	if err != nil {
		c.JSON(408,gin.H{
			"msg":"请求时间超时,下载失败",
			"code":408,
		})
		return
	}
	// c.JSON(200,gin.H{
	// 	"code":20000,
	// 	"msg":"下载成功",
	// })
}

//@Title 获取mysql全部信息
//@Summary 获取mysql全部信息
//@Description 用于后台页面获取全部信息
//@Success 200 {json} json ""msg":"获取全部学生信息成功","
//@Failure 500 "获取全部信息失败"
//@Router /getalldata [GET]
func GetAlldatamysql(c *gin.Context){
	pagesize := c.DefaultQuery("pagesize", "20")
	page := c.DefaultQuery("page", "1")
	limit:=pagesize
	p,_ := strconv.Atoi(page)
	pa,_ := strconv.Atoi(pagesize)
	offset:=(p-1)*pa
	fmt.Println(limit,offset)
	alldata,_:=mysql.Queryalldata("1000",0)
	all,err:=mysql.Queryalldata(limit,offset)
	fmt.Println(len(all))
	if err != nil {
		c.JSON(500,gin.H{
			"msg":"获取全部信息失败",
			"code":500,
		})
		return
	}
	c.JSON(200,gin.H{
		"msg":"获取全部学生信息成功",
		"code":20000,
		"data":all,
		"total":len(alldata),
	})
}

//@Summary 删除单个信息
//@Description 用于删除单个信息
//@Accept multipart/form-data
//@Produce application/json
//@Param stu_id query string true "学号"
//@Success 200 {json} json ""msg":"删除信息成功""
//@Failure 500 "获取全部信息失败"
//@Failure 408 "请求时间超时,删除信息失败"
//@Router /delete [GET]
func DeleteData(c *gin.Context)  {
	var all stu.User
	err:=c.ShouldBind(&all)
	if err != nil {
		c.JSON(500,gin.H{
			"msg":"获取信息失败",
			"code":500,
		})
		return 
	}
	err=mysql.Deletedata(all.Stu_id)
	if err != nil {
		c.JSON(408,gin.H{
			"msg":"删除信息失败",
			"code":408,
		})
		return 
	}
	c.JSON(200,gin.H{
		"msg":"删除信息成功",
		"code":20000,
	})
}


//@Summary 查询结果字段
//Description 用于结果查询
//@Accept multipart/form-data
//@Produce application/json
//@Param stu_id query string true "学号"
//@Param phone query string true "手机"
//@Param qq query string true "qq"
//@Success 200 {json} json ""msg":"查询结果成功 0:未进行测试 1:已录取 2:笔试录取 3:面试录取""
//@Failure 500 "获取结果失败"
//@Failure 408 "请求时间超时,获取结果失败"
//@Router /result [POST]
func Getresult(c *gin.Context)  {
	var all stu.User
	err:=c.ShouldBind(&all)
	if err != nil {
		c.JSON(500,gin.H{
			"msg":"获取信息失败",
			"code":500,
		})
		return 
	}
	all,err=mysql.Queryfield(all.Stu_id,all.Phone,all.Qq)
	if err != nil {
		c.JSON(408,gin.H{
			"msg":"请求时间超时,获取结果失败",
			"code":408,
		})
		return 
	}
	c.JSON(200,gin.H{
		"msg":"查询结果成功",
		"code":20000,
		"data":all.Result,
	})
}