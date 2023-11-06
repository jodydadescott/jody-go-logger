package main

import (
	logger "github.com/jodydadescott/jody-go-logger"
	"go.uber.org/zap"
)

func main() {

	logger.SetConfig(&logger.Config{
		LogLevel: logger.DebugLevel,
		Encoding: "json",
	})

	zap.L().Debug("Test Message")

}
