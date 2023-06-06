package main

import (
	"github.com/containerscrew/devstdout/logger"
	"os"
)

var appEnv = os.Getenv("APP_ENV")

func main() {
	log := logger.NewLogger(
		logger.OptionsLogger{Level: "success", AddSource: true},
		"prod",
	)

	log.Logger.Debug(
		"executing database query",
		logger.PrintMessage("query", "SELECT * FROM users"),
	)
	log.Logger.Info("image upload successful", logger.PrintMessage("image_id", "39ud88"))
	log.Logger.Info("image upload successful", logger.PrintMessage("image_id", 9876544))
	log.Success("Success Message")
}
