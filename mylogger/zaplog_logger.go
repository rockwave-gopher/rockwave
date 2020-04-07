package main

import (
	"go.uber.org/zap"
	"time"
)

/**
high performance than sugar logger
 */
func main() {
	logger, _ := zap.NewProduction()
	url:="http://www.google.com"
	defer logger.Sync()
	logger.Info("failed to fetch URL",
		// Structured context as strongly typed Field values.
		zap.String("url", url),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second),
	)

}
