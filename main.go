package main

import (
	"context"
	"fmt"
	"gin_demo/dao/mysql"
	"gin_demo/dao/redis"
	"gin_demo/logger"
	"gin_demo/routes"
	"gin_demo/settings"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func main() {
	// 1. 載入設定
	if err := settings.Init(); err != nil {
		fmt.Printf("Init settings failed, err:%v\n", err)
		return
	}
	// 2. 初始化日誌
	if err := logger.Init(settings.Conf.LogConfig); err != nil {
		fmt.Printf("Init logger failed, err:%v\n", err)
		return
	}
	defer zap.L().Sync()
	zap.L().Debug("logger init success")
	// 3. 初始化 MySQL 連接
	if err := mysql.Init(settings.Conf.MySQLConfig); err != nil {
		fmt.Printf("Init mysql failed, err:%v\n", err)
		return
	}
	defer mysql.Close()
	// 4. 初始化 Redis 連接
	if err := redis.Init(settings.Conf.RedisConfig); err != nil {
		fmt.Printf("Init redis failed, err:%v\n", err)
		return
	}
	defer redis.Close()
	// 5. 註冊路由
	router := routes.Setup()
	// 6. 啟動服務(優雅關機)
	ser := &http.Server{
		Addr:    fmt.Sprintf(":%d", viper.GetInt("app.port")),
		Handler: router,
	}
	go func() {
		// 啟動一個 goroutine 啟動伺服器
		if err := ser.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	// 等待中斷訊號來關閉伺服器，為關閉伺服器操作設定一個 5 秒的超時
	quit := make(chan os.Signal, 5) // 建立一個接收訊號的通道
	// kill 預設會發送 syscall.SIGTERM 訊號
	// kill -2 會發送 syscall.SIGINT 訊號，我們常用的 Ctrl+C 也是發送這個訊號
	// kill -9 會發送 syscall.SIGKILL 訊號，但是不能被捕獲，所以不需要增加它
	// signal.Notify 把收到的 syscall.SIGINT 或 syscall.SIGTERM 訊號轉發給 quit 通道
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM) // 此處不會阻塞
	<-quit                                               // 阻塞在此，當接收到上述兩種訊號時才會往下執行
	zap.L().Info("Shutdown Server")
	// 建立一個 5 秒超時的 context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// 5 秒內關閉服務，將未處理完的請求處理完再關閉服務，超過 5 秒就超時退出
	if err := ser.Shutdown(ctx); err != nil {
		zap.L().Fatal("Server Shutdown", zap.Error(err))
	}
	zap.L().Info("Server exiting")
}
