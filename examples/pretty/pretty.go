package main

import (
	logger "github.com/containerscrew/devstdout/pkg"
)

func main() {
	log := logger.NewLogger(
		logger.OptionsLogger{Level: "debug", AddSource: false, LoggerType: "pretty"},
	)

	log.Debug(
		"testing message",
		logger.PrintMessage("hello", "world"),
	)

	log.Info(
		"testing message",
		logger.PrintMessage("bob", "marley"),
	)

	log.Warning("warning message!")

	log.Success(
		"Success Message",
		logger.PrintMessage("alice", "bob"),
	)

	log.Error("error in your app!", logger.PrintMessage("error", "your_error_is_here"))

	log.ErrorWithExit("fatal error, app must stop!", logger.PrintMessage("error", "your_error_is_here"))
}
