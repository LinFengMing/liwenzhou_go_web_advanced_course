package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func func1(c *gin.Context) {
	fmt.Println("func1")
}

func func2(c *gin.Context) {
	fmt.Println("func2 before")
	c.Next()
	fmt.Println("func2 after")
}

func func3(c *gin.Context) {
	fmt.Println("func3")
	// c.Abort()
}

func func4(c *gin.Context) {
	fmt.Println("func4")
	c.Set("name", "jiro")
}

func func5(c *gin.Context) {
	fmt.Println("func5")
	value, ok := c.Get("name")
	if ok {
		valueStr := value.(string) // 型別轉換
		fmt.Println(valueStr)
	}
}

func main() {
	r := gin.Default()
	r.GET("/hello", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})
	shopGroup := r.Group("/shop", func1, func2)
	shopGroup.Use(func3)
	{
		shopGroup.GET("/index", func4, func5)
	}
	r.Run()
}
