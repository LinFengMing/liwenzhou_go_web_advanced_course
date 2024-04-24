package main

import (
	"net/http"

	"go.uber.org/zap"
)

var logger *zap.Logger

func InitLogger() {
	logger, _ = zap.NewProduction()
}

func simpleHttpGet(url string) {
	resp, err := http.Get(url)
	if err != nil {
		logger.Error("Error fetching URL",
			zap.String("url", url),
			zap.Error(err),
		)
	} else {
		logger.Info("Successfully fetched URL",
			zap.String("url", url),
			zap.Int("status", resp.StatusCode),
		)
		resp.Body.Close()
	}
}

func main() {
	InitLogger()
	defer logger.Sync()
	simpleHttpGet("www.google.com")
	simpleHttpGet("http://www.google.com")
}
