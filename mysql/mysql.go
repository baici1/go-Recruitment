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
func Queryonedata(stu_id string ) (User,error) {
	var u User
	sqlStr:="select id,stu_id from stu where stu_id=? "
	err:=db.Get(&u,sqlStr,stu_id)
	if err != nil {
		if err != nil {
			fmt.Printf("Queryonedata get failed, err:%v\n", err)
			return u,err
		}
	}
	return u,nil
}

//增加----add单个数据
func Addonedata(stu_id string) error {
	sqlStr:="insert into stu (stu_id) values (?)"
	_,err:=db.Exec(sqlStr,stu_id)
	if err != nil {
		fmt.Printf("Addonedata insert failed, err:%v\n", err)
		return err
	}
	return nil
}

//增加----add一条数据
func Addalldata(stu_id ,real_name ,group_id,sex,college,major,phone,qq string ) error {
	sqlStr:="insert into stu (stu_id ,real_name ,group_id,sex,college,major,phone,qq) values (?,?,?,?,?,?,?,?)"
	sqlstr:="insert into description (stu_id) values (?)"//,comments,code,self_study,attach,development,ready,degree
	_,err:=db.Exec(sqlStr,stu_id ,real_name ,group_id,sex,college,major,phone,qq)
	_,err=db.Exec(sqlstr,stu_id)
	if err != nil {
		fmt.Printf("Addalldata insert failed, err:%v\n", err)
		return err
	}
	return nil
}

//更新数据
func UpdateoneForm(real_name string,group_id,sex string ,college,major,phone,qq string,result,code string,stu_id string) error {
	sqlStr:="update stu set real_name=?,group_id=?,sex=?,college=?,major=?,phone=?,qq=?,result=?,code=? where stu_id=?"
	_,err:=db.Exec(sqlStr,real_name,group_id,sex,college,major,phone,qq ,result,code,stu_id)
	if err != nil {
		fmt.Printf("UpdateoneForm update failed, err:%v\n", err)
		return err
	}
	return nil
}

//查询数据
func Querydata(real_name string) ([]stu.User,error) {
	sqlStr:="select * from stu where real_name like ?"
	var all []stu.User
	err:=db.Select(&all,sqlStr,"%"+real_name+"%")
	if err != nil {
		fmt.Printf("Querydata get failed, err:%v\n", err)
			return all,err
	}
	return all,nil
}

//获取mysql全部学生信息
func Queryalldata(limit string,offset int) ([]stu.User,error) {
	sqlStr:="select * from stu limit ? offset ?"
	var all []stu.User
	err:=db.Select(&all,sqlStr,limit,offset)
	// fmt.Println(all[1])
	if err != nil {
		fmt.Printf("Queryalldata query failed, err:%v\n", err)
		return all,err
	}
	return all,nil
}

//删除单个成员信息
func Deletedata(stu_id string) error {
	sqlStr:="delete from stu where stu_id=?"
	sql:="delete from description where stu_id=?"
	_,err:=db.Exec(sqlStr,stu_id)
	_,err=db.Exec(sql,stu_id)
	if err != nil {
		fmt.Printf("Deletedata delete failed, err:%v\n", err)
		return err
	}
	return nil
}

//查询------查询单个字段
func Queryfield(stu_id ,phone,qq string ) (stu.User,error) {
	var u stu.User
	sqlStr:="select result from stu where stu_id=? and phone=? and qq=?"
	err:=db.Get(&u,sqlStr,stu_id,phone,qq)
	if err != nil {
		
			fmt.Printf("Queryfield get failed, err:%v\n", err)
			return u,err
		
	}
	return u,nil
}


//增加stu的描述信息
func Updatedes(comments,self_study,attach,development,ready,degree,stu_id string,grades int) error {
	sqlStr:="update description set comments=?,grades=?,self_study=?,attach=?,development=?,ready=?,degree=? where stu_id=?"
	_,err:=db.Exec(sqlStr,comments,grades,self_study,attach,development,ready,degree,stu_id)
	if err != nil {
		fmt.Printf("Updatedes update failed, err:%v\n", err)
		return err
	}
	return nil
}
//查询des所有字段和stu某个字段
func Querydes(stu_id string) (stu.Des,error) {
	var des stu.Des
	sqlStr:="SELECT stu.sex,stu.group_id, stu.real_name, des.stu_id,des.comments,des.grades,des.self_study,des.attach,des.development,des.ready,des.degree,des.one,des.two FROM stu stu LEFT JOIN description des ON stu.stu_id=des.stu_id where des.stu_id=?;"
	err:=db.Get(&des,sqlStr,stu_id)
	if err != nil {
		fmt.Printf("Querydes get failed, err:%v\n", err)
			return des,err
	}
	return des,nil

}

//修改值
func Updateone(one string,stu_id string) error {

	sqlStr:="update description set one=? where stu_id=?"
	_,err:=db.Exec(sqlStr,one,stu_id)
	if err != nil {
		fmt.Printf("updateone update failed, err:%v\n", err)
		return err
	}
	return nil;
}

//修改值
func Updatetwo(two string,stu_id string) error {
	
	sqlStr:="update description set two=? where stu_id=?"
	_,err:=db.Exec(sqlStr,two,stu_id)
	if err != nil {
		fmt.Printf("updatetwo update failed, err:%v\n", err)
		return err
	}
	return nil;
}

//查询过一面的人
func Queryone() ([]stu.Des,error) {
	var des []stu.Des
	sqlStr:="SELECT stu.sex,stu.group_id, stu.real_name, des.stu_id,des.comments,des.grades,des.self_study,des.attach,des.development,des.ready,des.degree,des.one,des.two FROM stu stu LEFT JOIN description des ON stu.stu_id=des.stu_id where des.one=?;"
	err:=db.Select(&des,sqlStr,"1")
	if err != nil {
		fmt.Printf("updatetwo update failed, err:%v\n", err)
		return des,err;
	}
	return des,nil;
}

//查询过二面的人
func Querytwo() ([]stu.Des,error) {
	var des []stu.Des
	sqlStr:="SELECT stu.sex,stu.group_id, stu.real_name, des.stu_id,des.comments,des.grades,des.self_study,des.attach,des.development,des.ready,des.degree,des.one,des.two FROM stu stu LEFT JOIN description des ON stu.stu_id=des.stu_id where des.one=? and des.two=?;"
	err:=db.Select(&des,sqlStr,"1","1")
	if err != nil {
		fmt.Printf("updatetwo update failed, err:%v\n", err)
		return des,err;
	}
	return des,nil;
}
