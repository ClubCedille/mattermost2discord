package api

import (
	"sync"
	"go.uber.org/zap"
)

var lock = &sync.Mutex{};
var sugarInstance *zap.SugaredLogger;

func loggerInstance() *zap.SugaredLogger {
	if sugarInstance == nil {
		lock.Lock() // Lock mutex to prevent overwriting Logger instance
		defer lock.Unlock() // Release mutex when done
		if sugarInstance == nil {
			logger, _ := zap.NewProduction()
			defer logger.Sync()
			sugarInstance = logger.Sugar() // Set singleton instance
		}
	}
	return sugarInstance
}
