package main

import (
	logger "github.com/containerscrew/devstdout/pkg"
)

func main() {
	log := logger.NewLogger(
		logger.OptionsLogger{Level: "warning", AddSource: false, LoggerType: "pretty"},
	)

	log.Debug(
		"testing message",
		logger.PrintMessage("test", "Debug test"),
	)

	log.Info(
		"testing message",
		logger.PrintMessage("test", "test"),
	)

	log.Warning("warning message!")

	log.Success(
		"Success Message",
		logger.PrintMessage("test", "test"),
	)
}
