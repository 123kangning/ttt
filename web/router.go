package web

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"qin/pkg/jwt"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	baseGroup := r.Group("/qin")

	baseGroup.POST("/user/signIn/", SignIn)
	baseGroup.POST("/user/login/", Login)

	userG := baseGroup.Group("/user", AuthMiddleware())
	userG.POST("/signOut/", SignOut)

	return r
}
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从HTTP请求头中读取token
		tokenString := c.GetHeader("Authorization")
		// 确保token非空
		if tokenString == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}
		// 解析token
		claim, err := jwt.ParseToken(tokenString)

		// 处理token解析错误
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}
		// 将解析后的 username 存储到gin的上下文中，以便后续使用
		c.Set("username", claim.UserName)
		// 继续处理请求
		c.Next()
	}
}
