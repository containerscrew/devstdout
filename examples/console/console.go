package main

import (
	devstdout "github.com/containerscrew/devstdout/pkg"
)

func main() {
	log := devstdout.NewLogger(
		devstdout.OptionsLogger{Level: "info", AddSource: false, LoggerType: "console"},
	)

	log.Info("hello, world!")

	log.Debug(
		"testing message",
		devstdout.Argument("test", "Debug test"),
	)

	log.Info(
		"testing message",
		devstdout.Argument("test", "test"),
	)

	log.Success(
		"Success Message",
		devstdout.Argument("2", "test"),
	)
}
