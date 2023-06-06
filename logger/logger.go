package logger

import (
	"context"
	"golang.org/x/exp/slog"
	"os"
	"strings"
)

const (
	LevelTrace   = slog.Level(-8)
	LevelSuccess = slog.Level(6)
)

type OptionsLogger struct {
	Level     string
	AddSource bool
}

type CustomLogger struct {
	Logger  *slog.Logger
	opts    *slog.HandlerOptions
	ctx     context.Context
	options OptionsLogger
}

var LevelNames = map[slog.Leveler]string{
	LevelTrace:   "TRACE",
	LevelSuccess: "SUCCESS",
}

type LogMessageType interface {
	int | int64 | float64 | string
}

func PrintMessage[T LogMessageType](key string, value T) any {
	switch any(value).(type) {
	case int:
		return slog.Int(key, any(value).(int))
	case string:
		return slog.String(key, any(value).(string))
	default:
		return nil
	}
}

func getLevel(l string) slog.Level {
	switch strings.ToUpper(l) {
	case "SUCCESS":
		return LevelSuccess
	case "TRACE":
		return LevelTrace
	default:
		return slog.LevelInfo
	}
}

func (c *CustomLogger) withOptions() {
	c.opts = &slog.HandlerOptions{
		Level:     getLevel(c.options.Level),
		AddSource: c.options.AddSource,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.LevelKey {
				level := a.Value.Any().(slog.Level)
				levelLabel, exists := LevelNames[level]
				if !exists {
					levelLabel = level.String()
				}

				a.Value = slog.StringValue(levelLabel)
			}
			return a
		},
	}
}

func NewLogger(options OptionsLogger, env string) *CustomLogger {
	c := &CustomLogger{ctx: context.Background(), options: options}
	c.withOptions()
	c.Logger = slog.New(newPrettyHandler(os.Stdout, c.opts))

	if env == "prod" {
		c.Logger = slog.New(slog.NewJSONHandler(os.Stdout, c.opts))
	}

	return c
}

func (c *CustomLogger) Success(msg string, args ...any) {
	c.Logger.Log(c.ctx, LevelSuccess, msg, args...)
}
