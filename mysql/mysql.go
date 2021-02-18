package mysql

import (
	"fmt"
	stu "go-Recruitment/dao"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)


var db *sqlx.DB

type User struct {
	Id        int    `json:"id" valid:"required" form:"id"`
	Stu_id    string `json:"stu_id" valid:"required" form:"stu_id"`
	Password  string `json:"password" valid:"required" form:"password"`
}

//连接数据库
func InitDB() error {
	var err error
	dsn := "qmx:123456@tcp(47.113.203.60:3306)/qmx?charset=utf8mb4&parseTime=True"
	db,err=sqlx.Connect("mysql",dsn)//进行连接
	if err != nil {
		fmt.Printf("connect DB failed, err:%v\n", err)
		return err
	}
	db.SetMaxOpenConns(20)//设置与数据库建立连接的最大数目。
	db.SetMaxIdleConns(10)//设置连接池中的最大闲置连接数。
	return nil
}



//查询------查询单个数据
func Queryonedata(stu_id string ,password string) (User,error) {
	var u User
	sqlStr:="select id,stu_id,password from stu where stu_id=? and password=?"
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

//更新数据
func UpdateoneForm(real_name string,group_id,sex int ,college,major,phone,qq string,result,code int,stu_id string) error {
	sqlStr:="update stu set real_name=?,group_id=?,sex=?,college=?,major=?,phone=?,qq=?,result=?,code=? where stu_id=?"
	_,err:=db.Exec(sqlStr,real_name,group_id,sex,college,major,phone,qq ,result,code,stu_id)
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		return err
	}
	return nil
}

//查询数据
func Querydata(stu_id string) (stu.User,error) {
	sqlStr:="select * from stu where stu_id=?"
	var all stu.User
	err:=db.Get(&all,sqlStr,stu_id)
	if err != nil {
		fmt.Printf("get failed, err:%v\n", err)
			return all,err
	}
	return all,nil
}

//获取mysql全部学生信息
func Queryalldata() ([]stu.User,error) {
	sqlStr:="select * from stu"
	var all []stu.User
	err:=db.Select(&all,sqlStr)
	// fmt.Println(all[1])
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return all,err
	}
	return all,nil
}

//删除单个成员信息
func Deletedata(stu_id string) error {
	sqlStr:="delete from stu where stu_id=?"
	_,err:=db.Exec(sqlStr,stu_id)
	if err != nil {
		fmt.Printf("delete failed, err:%v\n", err)
		return err
	}
	return nil
}