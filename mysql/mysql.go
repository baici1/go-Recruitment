package mysql

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)


var db *sqlx.DB



//连接数据库
func InitDB() error {
	var err error
	dsn := "root:123456@tcp(127.0.0.1:3306)/qmx?charset=utf8mb4&parseTime=True"
	db,err=sqlx.Connect("mysql",dsn)//进行连接
	if err != nil {
		fmt.Printf("connect DB failed, err:%v\n", err)
		return err
	}
	db.SetMaxOpenConns(20)//设置与数据库建立连接的最大数目。
	db.SetMaxIdleConns(10)//设置连接池中的最大闲置连接数。
	return nil
}

type User struct {
	ID        int    `json:"id" valid:"required" form:"id"`
	Stu_id    string `json:"stu_id" valid:"required" form:"stu_id"`
	Password  string `json:"password" valid:"required" form:"password"`
}

//查询------查询单个数据
func Queryonedata(stu_id string ,password string) (User,error) {
	var u User
	sqlStr:="select id,stu_id,password from stu where stu_id=? and password=?"
	// stmt,err:=db.Prepare(sqlStr)
	// if err != nil {
	// 	fmt.Printf("prepare failed, err:%v\n", err)
	// 	return u
	// }
	//defer stmt.Close()
	
	err:=db.Get(&u,sqlStr,stu_id,password)
	if err != nil {
		if err != nil {
			fmt.Printf("get failed, err:%v\n", err)
			return u,err
		}
	}
	return u,nil
}

//增加----add单个数据
func Addonedata(stu_id string ,password string) error {
	sqlStr:="insert into stu (stu_id,password) values (?,?)"
	_,err:=db.Exec(sqlStr,stu_id,password)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return err
	}
	return nil
}
