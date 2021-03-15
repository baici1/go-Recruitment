package stu

type User struct {
	ID        string `json:"id" valid:"required" form:"id"`
	Stu_id    string `json:"stu_id" valid:"required" form:"stu_id"`
	Password  string `json:"password" valid:"required" form:"password"`
	Real_name string `json:"real_name" valid:"required" form:"real_name"`
	Group_id  string `json:"group_id" valid:"required" form:"group_id"`
	Sex       string `json:"sex" valid:"required" form:"sex"`
	College   string `json:"college" valid:"required" form:"college"`
	Major     string `json:"major" valid:"required" form:"major"`
	Phone     string `json:"phone" valid:"required" form:"phone"`
	Qq        string `json:"qq" valid:"required" form:"qq"`
	Result    string `json:"result" valid:"required" form:"result"`
	Code      string `json:"code" valid:"required" form:"code"`
}

type Des struct {
	Real_name   string `json:"real_name" valid:"required" form:"real_name"`
	Group_id    string `json:"group_id" valid:"required" form:"group_id"`
	Sex         string `json:"sex" valid:"required" form:"sex"`
	Stu_id      string `json:"stu_id" valid:"required" form:"stu_id"`
	Comments    string `json:"comments" valid:"required" form:"comments"`
	Grades      int    `json:"grades" valid:"required" form:"grades"`
	Self_study  string `json:"self_study" valid:"required" form:"self_study"`
	Attach      string `json:"attach" valid:"required" form:"attach"`
	Development string `json:"development" valid:"required" form:"development"`
	Ready       string `json:"ready" valid:"required" form:"ready"`
	Degree      string `json:"degree" valid:"required" form:"degree"`
	One      string `json:"one" valid:"required" form:"one"`
	Two      string `json:"two" valid:"required" form:"two"`
}