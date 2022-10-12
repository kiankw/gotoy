package main

import (
	"net/http"

	"github.com/kiankw/gotoy/logzap/logzap"

	"go.uber.org/zap"
)

// var logger *zap.Logger
// var sugarLogger *zap.SugaredLogger

func main() {
	logger := logzap.NewLogger()
	defer logger.Sync()
	for i := 0; i < 100; i++ {
		logger.Info("Success", zap.Int("i", i))
		logger.Error("Success", zap.Int("i", i))

		// simpleHttpGet("www.baidu.com", logger)
		// simpleHttpGet("http://www.baidu.com", logger)
	}
}

func simpleHttpGet(url string, logger *zap.Logger) {
	resp, err := http.Get(url)
	if err != nil {
		logger.Error(
			"Error fetching url..",
			zap.String("url", url),
			zap.Error(err))
	} else {
		logger.Info("Success..",
			zap.String("statusCode", resp.Status),
			zap.String("url", url))
		resp.Body.Close()
	}
}
