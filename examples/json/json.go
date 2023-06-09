package main

import logger "github.com/containerscrew/devstdout/pkg"

func main() {
	log := logger.NewLogger(
		logger.OptionsLogger{Level: "trace", AddSource: false, LoggerType: "json"},
	)

	log.Info(
		"image upload successful",
		logger.PrintMessage("image_id", "39ud88"),
	)

	log.Error("error message", logger.PrintMessage("image_id", "39ud88"))

	log.Info(
		"image upload successful",
		logger.PrintMessage("process_id", 9876544),
	)

	log.Success(
		"Success Message",
		logger.PrintMessage("internal_id", "34fces"),
	)
}
