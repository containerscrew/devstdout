package main

import (
	"github.com/containerscrew/devstdout/logger"
	"os"
)

var appEnv = os.Getenv("APP_ENV")

func main() {
	opts := logger.WithOptions("debug", true)

	log := logger.NewLogger(opts)

	log.Logger.Debug(
		"executing database query",
		logger.PrintString("query", "SELECT * FROM users"),
	)
	log.Logger.Info("image upload successful", logger.PrintString("image_id", "39ud88"))
	log.Success("Success Message")
	// Me encantaría implementar este método jejejejejejejeejejejeje
	//log.Logger.Success()
}
