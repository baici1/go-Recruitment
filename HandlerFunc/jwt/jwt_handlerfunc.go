package JWT

import (
	"fmt"
	"go-Recruitment/tool/jwt"

	"github.com/gin-gonic/gin"
)

//jwt中间件
func JWTAuthMiddleware(c *gin.Context) {
	authHandler := c.Request.Header.Get("Authorization")//获取请求头中的token
	if authHandler==""{
		c.JSON(404,gin.H{
			"code": "404",
			"msg":"请求头为空",
		})
		c.Abort()
		return
	}
	token, err := jwt.ParseToken(authHandler)//解析token
	fmt.Println(token)
	if err != nil {
		c.JSON(500, gin.H{
			"code": 500,
			"msg":  "无效的Token",
		})
		c.Abort()
		return
	}
	// 将当前请求的username信息保存到请求的上下文c上
	c.Next()
}