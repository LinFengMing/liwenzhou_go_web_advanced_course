package main

import (
	"log"
	"time"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		time.Sleep(5. * time.Second)
		c.String(200, "Wlecome Gin Server")
	})
	// ser := &http.Server{
	// 	Addr:    ":8080",
	// 	Handler: router,
	// }
	// go func() {
	// 	// 啟動一個 goroutine 啟動伺服器
	// 	if err := ser.ListenAndServe(); err != nil && err != http.ErrServerClosed {
	// 		log.Fatalf("listen: %s\n", err)
	// 	}
	// }()
	// // 等待中斷訊號來關閉伺服器，為關閉伺服器操作設定一個 5 秒的超時
	// quit := make(chan os.Signal, 5) // 建立一個接收訊號的通道
	// // kill 預設會發送 syscall.SIGTERM 訊號
	// // kill -2 會發送 syscall.SIGINT 訊號，我們常用的 Ctrl+C 也是發送這個訊號
	// // kill -9 會發送 syscall.SIGKILL 訊號，但是不能被捕獲，所以不需要增加它
	// // signal.Notify 把收到的 syscall.SIGINT 或 syscall.SIGTERM 訊號轉發給 quit 通道
	// signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM) // 此處不會阻塞
	// <-quit                                               // 阻塞在此，當接收到上述兩種訊號時才會往下執行
	// log.Println("Shutdown Server")
	// // 建立一個 5 秒超時的 context
	// ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// defer cancel()
	// // 5 秒內關閉服務，將未處理完的請求處理完再關閉服務，超過 5 秒就超時退出
	// if err := ser.Shutdown(ctx); err != nil {
	// 	log.Fatal("Server Shutdown:", err)
	// }
	// 預設 endless 伺服器會監聽以下訊號
	// syscall.SIGHUP, syscall.SIGUSR1, syscall.SIGUSR2, syscall.SIGINT, syscall.SIGTERM和syscall.SIGTSTP
	// 接收到 syscallHUP 訊號會觸發 `fork/restart` 實現優雅重啟 ( kill -1 pid 會發送 syscall.SIGHUP 訊號)
	// 接收到 syscall.SIGINT或syscall.SIGTERM 訊號會觸發優雅關機
	// 接收到 SIGUSR2 訊號會觸發 HammerTime
	// SIGUSR1 和 SIGTSTP 被用來觸發一些用戶自定義的 hook 函數
	if err := endless.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("Server err: %v", err)
	}
	log.Println("Server exiting")
}
