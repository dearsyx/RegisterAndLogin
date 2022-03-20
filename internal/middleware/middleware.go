package middleware

import (
	"net/http"

	"goprojects.com/simple_regist_login/pkg/code"

	"goprojects.com/simple_regist_login/internal/token"

	"github.com/gin-gonic/gin"
)

//UserMiddleware 用户验证中间件
func UserMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 去cookie中找token
		tokenKey, err := c.Cookie(token.CookieTokenKey)
		if err == http.ErrNoCookie {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"code": code.RequestForbidden, "msg": "请先登录后再进行操作",
			})
			return
		}
		c.Set("user_token", tokenKey)
	}
}
