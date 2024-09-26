package main

import devstdout "github.com/containerscrew/devstdout/pkg"

func main() {
	log := devstdout.NewLogger(
		devstdout.OptionsLogger{Level: "trace", AddSource: false, LoggerType: "json"},
	)

	log.Info(
		"image upload successful",
		devstdout.Argument("image_id", "39ud88"),
	)

	log.Error("error message", devstdout.Argument("image_id", "39ud88"))

	log.Info("testing int message", devstdout.Argument("number", 33))

	log.Info(
		"image upload successful",
		devstdout.Argument("process_id", 9876544),
	)

	log.Success(
		"Success Message",
		devstdout.Argument("internal_id", "34fces"),
	)

	log.ErrorWithExit("fatal error, app must stop!", devstdout.Argument("error", "your_error_is_here"))
}
