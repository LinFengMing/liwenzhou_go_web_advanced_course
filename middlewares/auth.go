package middlewares

import (
	"gin_demo/pkg/jwt"
	"strings"

	"gin_demo/controllers"

	"github.com/gin-gonic/gin"
)

func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		// 客戶端帶 Token 有三種方式：1.放在 Header 的 Authorization 2. 放在請求裡 3.放在 URI
		// 這邊假設 Token 放在 Header 的 Authorization 裡，並使用 Bearer token
		// 這邊的具體實現方式要依據你的實際業務情況決定
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			controllers.ResponseError(c, controllers.CodeNeedLogin)
			c.Abort()
			return
		}
		// 按空格分割
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			controllers.ResponseError(c, controllers.CodeInvalidToken)
			c.Abort()
			return
		}
		// parts[1] 是取得到的 tokenString，我們使用之前定義的解析 JWT 的函式來解析這個 tokenString
		mc, err := jwt.ParseToken(parts[1])
		if err != nil {
			controllers.ResponseError(c, controllers.CodeInvalidToken)
			c.Abort()
			return
		}
		// 將當前用戶的 userID 資訊保存到請求的上下文 c 上
		c.Set(controllers.CtxtUserIDKey, mc.UserID)
		c.Next() // 後續的處理函式可以用 c.Get(CtxtUserIDKey) 來取得當前請求的用戶資訊
	}
}
