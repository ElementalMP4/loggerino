package main

import (
	"loggerino/log"
)

func main() {
	// Using the default logger
	log.Ok("demo-logger", "This is a success message")
	log.Info("demo-logger", "This is an info message")
	log.Warn("demo-logger", "This is a warning message")
	log.Error("demo-logger", "This is an error message")
	log.Debug("demo-logger", "This is a debug message")

	log.Okf("demo-logger", "OK with format: %s", "success")
	log.Infof("demo-logger", "Info with format: %s %d", "hello", 42)
	log.Warnf("demo-logger", "Warning with format: %v", []int{1, 2, 3})
	log.Errorf("demo-logger", "Error with format: %s", "something went wrong")
	log.Debugf("demo-logger", "Debug with format: %t", true)

	// Using a custom logger instance
	logger := log.New()
	logger.SetFile("./output.log")

	logger.Ok("custom-logger", "This is a success message")
	logger.Info("custom-logger", "Custom logger info")
	logger.Warn("custom-logger", "Custom logger warning")
	logger.Error("custom-logger", "Custom logger error")
	logger.Debug("custom-logger", "Custom logger debug")

	logger.Okf("custom-logger", "OK with format: %s", "success")
	logger.Infof("custom-logger", "Custom info with format: %d", 123)
	logger.Warnf("custom-logger", "Custom warning with format: %f", 3.14)
	logger.Errorf("custom-logger", "Custom error with format: %q", "error")
	logger.Debugf("custom-logger", "Custom debug with format: %v", map[string]int{"key": 1})
}
