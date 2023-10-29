package main

import logging "github.com/jodydadescott/jody-go-logger"

func main() {

	// type Config struct {
	// 	LogLevel           LogLevel
	// 	SamplingInitial    int
	// 	SamplingThereafter int
	// 	Development        bool
	// 	Encoding           string
	// }

	config := &logging.Config{
		LogLevel:    logging.TraceLevel,
		Development: true,
	}

	logging.SetConfig(config)

	logging.Debug("Debug Message")
	logging.Trace("Trace Message")
	logging.Info("Info Message")

}
