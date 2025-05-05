package logging

import "log"

var logger *log.Logger

func GetLogger() *log.Logger {
	if logger == nil {
		logger = log.Default()
		logger.Print("creating new logger")
	}
	return logger
}
