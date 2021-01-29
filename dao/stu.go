package stu

type User struct {
	ID        int    `json:"id" valid:"required" form:"id"`
	Stu_id    string `json:"stu_id" valid:"required" form:"stu_id"`
	Password  string `json:"password" valid:"required" form:"password"`
	Real_name string `json:"real_name" valid:"required" form:"real_name"`
	Group_id  int    `json:"group_id" valid:"required" form:"group_id"`
	Sex       int    `json:"sex" valid:"required" form:"sex"`
	College   string `json:"college" valid:"required" form:"college"`
	Major     string `json:"major" valid:"required" form:"major"`
	Phone     string `json:"phone" valid:"required" form:"phone"`
	Qq        string `json:"qq" valid:"required" form:"qq"`
	Result    int    `json:"result" valid:"required" form:"result"`
	Code      int    `json:"code" valid:"required" form:"code"`
}