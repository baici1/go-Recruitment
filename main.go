package main

// @title go-Recruitment
// @version 1.0
// @description 招新网站api
// @termsOfService http://swagger.io/terms/

// @contact.name yay
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 127.0.0.1:8081

import (
	"fmt"
	"go-Recruitment/mysql"
	"go-Recruitment/router"
)


func main() {
	//连接数据库
	err:=mysql.InitDB()
	if err != nil {
		fmt.Println("failed")
	}else{
		fmt.Println("success")
	}
	router.InitRoute()
}