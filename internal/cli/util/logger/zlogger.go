package zlogger

import (
	"os"

	"github.com/0chain/gosdk/core/logger"
)

var defaultLogLevel = logger.DEBUG
var Logger logger.Logger

func init() {
	Logger.Init(defaultLogLevel, "Zs3server Testing")
}

func SetLogFile(logFile string, verbose bool) {
	f, err := os.OpenFile(logFile, os.O_TRUNC|os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return
	}
	Logger.SetLogFile(f, verbose)
}
