package routes

import (
	"gin_demo/controllers"
	"gin_demo/logger"
	"gin_demo/pkg/jwt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func Setup(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	r.POST("/signup", controllers.SignUpHandler)
	r.POST("/login", controllers.LoginHandler)
	r.GET("/pong", JWTAuthMiddleware(), func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "404",
		})
	})
	return r
}

func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		// 客戶端帶 Token 有三種方式：1.放在 Header 的 Authorization 2. 放在請求裡 3.放在 URI
		// 這邊假設 Token 放在 Header 的 Authorization 裡，並使用 Bearer token
		// 這邊的具體實現方式要依據你的實際業務情況決定
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 2003,
				"msg":  "Header Authorization is empty",
			})
			c.Abort()
			return
		}
		// 按空格分割
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 2004,
				"msg":  "Authorization format must be Bearer {token}",
			})
			c.Abort()
			return
		}
		// parts[1] 是取得到的 tokenString，我們使用之前定義的解析 JWT 的函式來解析這個 tokenString
		mc, err := jwt.ParseToken(parts[1])
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 2005,
				"msg":  "無效的 Token",
			})
			c.Abort()
			return
		}
		// 將當前用戶的 userID 資訊保存到請求的上下文 c 上
		c.Set("userID", mc.UserID)
		c.Next() // 後續的處理函式可以用 c.Get("userID") 來取得當前請求的用戶資訊
	}
}
