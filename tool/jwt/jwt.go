package jwt

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

// 指定加密密钥
var jwtSecret=[]byte("yay")

type MyClaims struct {
	Id       int  `json:"id" valid:"required"`
	Stu_id     string `json:"stu_id" valid:"required"`
	Password string `json:"password" valid:"required"`
	jwt.StandardClaims
}
//设置时间
const TokenExpireDuration = time.Hour * 24

//生成jwt
func GetToken(id int,stu_id string,password string) (string,error)  {
	c:=MyClaims{
		Id:id,
		Stu_id:stu_id,
		Password:password,
		StandardClaims:jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(),//过期时间
			Issuer:"yay",//签发人
		},
	}
	token:=jwt.NewWithClaims(jwt.SigningMethodHS256,c)//生成token
	return token.SignedString(jwtSecret)//转化为string
}

//解析JWT
func ParseToken(tokenString string) (*MyClaims ,error)  {
	token,err:=	jwt.ParseWithClaims(tokenString,&MyClaims{},func(token *jwt.Token)(interface{}, error){
		return jwtSecret,nil
	})//解析token//用于解析鉴权的声明，方法内部主要是具体的解码和校验的过程，最终返回*Token
	if err != nil {
		return nil, err
	}
	// 从tokenClaims中获取到Claims对象，并使用断言，将该对象转换为我们自己定义的Claims
		// 要传入指针，项目中结构体都是用指针传递，节省空间。
	if claims,ok:= token.Claims.(*MyClaims);ok&&token.Valid{
		return claims,err
	}
	return nil,err
}