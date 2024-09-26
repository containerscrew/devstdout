package main

import (
	devstdout "github.com/containerscrew/devstdout/pkg"
)

func main() {
	log := devstdout.NewLogger(
		devstdout.OptionsLogger{Level: "debug", AddSource: false, LoggerType: "pretty"},
	)

	log.Debug(
		"testing message",
		devstdout.Argument("hello", "world"),
	)

	log.Info(
		"testing message",
		devstdout.Argument("bob", "marley"),
	)

	log.Warning("warning message!")

	log.Success(
		"Success Message",
		devstdout.Argument("alice", "bob"),
	)

	log.Error("error in your app!", devstdout.Argument("error", "your_error_is_here"))

	log.ErrorWithExit("fatal error, app must stop!", devstdout.Argument("error", "your_error_is_here"))
}
