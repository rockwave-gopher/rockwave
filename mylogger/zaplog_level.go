package main

import (
	"fmt"
	"go.uber.org/zap"
	"net/http"
	"time"
)

/**
get a level
curl http://localhost:9090/handle/level  get a logger level
set a new level
curl -XPUT --data '{"level":"debug"}' http://localhost:9090/handle/level

log level debug info warning error fatal
 */
func main() {
	alevel := zap.NewAtomicLevel()

	http.HandleFunc("/handle/level", alevel.ServeHTTP)
	//start a web server get a new level
	go func() {
		if err := http.ListenAndServe(":9090", nil); err != nil {
			panic(err)
		}
	}()

	// 默认是Info级别
	logcfg := zap.NewProductionConfig()

	logcfg.Level = alevel

	logger, err := logcfg.Build()

	if err != nil {
		fmt.Println("err", err)
	}

	defer logger.Sync()

	for i := 0; i < 50; i++ {
		time.Sleep(5 * time.Second)
		logger.Debug("debug log", zap.String("level", alevel.String()))
		logger.Info("Info log", zap.String("level", alevel.String()))
	}
}
