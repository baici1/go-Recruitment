# swagger

## 介绍

> Swagger本质上是一种用于描述使用JSON表示的RESTful API的接口描述语言。Swagger与一组开源软件工具一起使用，以设计、构建、记录和使用RESTful Web服务。Swagger包括自动文档，代码生成和测试用例生成。
>
> 在前后端分离的项目开发过程中，如果后端同学能够提供一份清晰明了的接口文档，那么就能极大地提高大家的沟通效率和开发效率。可是编写接口文档历来都是令人头痛的，而且后续接口文档的维护也十分耗费精力。
>
> 最好是有一种方案能够既满足我们输出文档的需要又能随代码的变更自动更新，而Swagger正是那种能帮我们解决接口文档问题的工具。

## 下载

`go get -u github.com/swaggo/swag/cmd/swag`

`go get -u github.com/swaggo/gin-swagger`

`go get -u github.com/swaggo/gin-swagger/swaggerFiles`

## 步骤

想要使用`gin-swagger`为你的代码自动生成接口文档，一般需要下面三个步骤：

1. 按照swagger要求给接口代码添加声明式注释.
2. 使用swag工具扫描代码自动生成API接口文档数据
3. 使用gin-swagger渲染在线接口文档页面

### 1.注释

在程序入口main函数上以注释的方式写下项目相关介绍信息。

```go
package main

// @title 这里写标题
// @version 1.0
// @description 这里写描述信息
// @termsOfService http://swagger.io/terms/

// @contact.name 这里写联系人信息
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 这里写接口服务的host
func main() {
	r := gin.New()

	// liwenzhou.com ...

	r.Run()
}
```

在你代码中处理请求的接口函数（通常位于controller层）按如下方式写上注释：

```go
//@Description 用于招新网站的登录
//@Summary 获取账号进行登录
//@Accept multipart/form-data
//@Produce application/json
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

```

### 2.生成接口文档数据

编写完注释后，使用以下命令安装swag工具：

```bash
go get -u github.com/swaggo/swag/cmd/swag
```

在项目根目录执行以下命令，使用swag工具生成接口文档数据。

```bash
swag init
```

执行完上述命令后，如果你写的注释格式没问题，此时你的项目根目录下会多出一个`docs`文件夹。



### 3.引入gin-swagger渲染文档数据

然后在项目代码中注册路由的地方按如下方式引入`gin-swagger`相关内容：

```go
import (

	_ "model名称/docs"  // 千万不要忘了导入把你上一步生成的docs

	gs "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	"github.com/gin-gonic/gin"
)
```

注册swagger api相关路由

```go
r.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))
```

把你的项目程序运行起来，打开浏览器访问http://localhost:8080/swagger/index.html就能看到Swagger 2.0 Api文档了。

![image-20210131204920565](https://cdn.jsdelivr.net/gh/baici1/image-host/img/20210131204920.png)



## swagger注解说明

![image-20210131205539126](https://cdn.jsdelivr.net/gh/baici1/image-host/img/20210131205539.png)

![image-20210201132355738](https://cdn.jsdelivr.net/gh/baici1/image-host/img/20210201132355.png)

注意点:

* `Param`
  * 参数，表示需要传递到服务器端的参数，有五列参数，使用空格或者 tab 分割，五个分别表示的含义如下
    1. 参数名
    2. 参数类型，可以有的值是 formData、query、path、body、header
       * formData 表示是 post 请求的数据，
       * query 表示带在 url 之后的参数
       * path 表示请求路径上得参数，
       * body 表示是一个 raw 数据请求
       * header 表示带在 header 信息中得参数。
       * ==获取参数的方式与gin框架获取参数方式一样== 
    3. 参数类型
    4. 是否必须
    5. 注释
* `@Success`
  * 成功返回给客户端的信息，三个参数，
    * 第一个是 status code。
    * 第二个参数是返回的类型，必须使用 {} 包含，
    * 第三个是返回的对象或者字符串信息
* `@router`
  * 路由信息，包含两个参数，使用空格分隔.
    * 第一个是请求的路由地址，支持正则和自定义路由
    * 第二个参数是支持的请求方法,放在 [] 之中，如果有多个方法，那么使用 , 分隔。
* `@Failure`
  * 失败返回的信息，包含两个参数，使用空格分隔，
    * 第一个表示 status code
    * 第二个表示错误信息
* `@Accept `  ----------------[Content-type](https://www.jianshu.com/p/ba40da728806)
  * 代表的是http实体首部字段,接收的类型
    * application/x-www-form-urlencoded
      * 1）浏览器的原生form表单
      * 2） 提交的数据按照 key1=val1&key2=val2 的方式进行编码，key和val都进行了URL转码
    * multipart/form-data
      * 常见的 POST 数据提交的方式。
    * application/json
      * 消息主体是序列化后的 JSON 字符串,这个类型越来越多地被大家所使用
    * text/xml
      * 是一种使用 HTTP 作为传输协议，XML 作为编码方式的远程调用规范
* `@Produce`  application/json
  * 指定返回的内容类型

