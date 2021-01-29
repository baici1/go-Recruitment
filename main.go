package main

import (
	"fmt"
	"go-Recruitment/mysql"
)


func main() {
	//连接数据库
	err:=mysql.InitDB()
	if err != nil {
		fmt.Println("failed")
	}else{
		fmt.Println("success")
	}
	//router.InitRoute()
}