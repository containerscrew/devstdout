package main

import (
	"github.com/containerscrew/devstdout/logger"
)

func main() {
	log := logger.NewLogger(
		logger.OptionsLogger{Level: "info", AddSource: false, LoggerType: "pretty"},
	)

	log.Debug(
		"executing database query",
		logger.PrintMessage("query", "Debug message"),
	)

	log.Info(
		"image upload successful",
		logger.PrintMessage("image_id", "39ud88"),
	)

	log.Info(
		"image upload successful",
		logger.PrintMessage("process_id", 9876544),
	)

	log.Success(
		"Success Message",
		logger.PrintMessage("internal_id", "34fces"),
	)
}
