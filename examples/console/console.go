package main

import (
	logger "github.com/containerscrew/devstdout/pkg"
)

func main() {
	log := logger.NewLogger(
		logger.OptionsLogger{Level: "info", AddSource: false, LoggerType: "console"},
	)

	log.Info("hello, world!")

	log.Debug(
		"testing message",
		logger.PrintMessage("test", "Debug test"),
	)

	log.Info(
		"testing message",
		logger.PrintMessage("test", "test"),
	)

	log.Success(
		"Success Message",
		logger.PrintMessage("2", "test"),
	)
}
