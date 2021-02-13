package api

import (
	"sync"
	"go.uber.org/zap"
)

var lock = &sync.Mutex{};
var sugarInstance *zap.SugaredLogger;

func loggerInstance() *zap.SugaredLogger {
	if sugarInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		// SugaredLogger doesn't exist. Create one
		if sugarInstance == nil {
			logger, _ := zap.NewProduction()
			defer logger.Sync() // flushes buffer, if any
			sugarInstance = logger.Sugar()
		}
	}
	return sugarInstance
}
