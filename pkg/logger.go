package devstdout

import (
	"context"
	"os"
	"strings"

	"golang.org/x/exp/slog"
)

const (
	LevelTrace   = slog.Level(-8)
	LevelSuccess = slog.Level(6)
)

type OptionsLogger struct {
	Level      string
	AddSource  bool
	LoggerType string
}

type CustomLogger struct {
	logger  *slog.Logger
	opts    *slog.HandlerOptions
	ctx     context.Context
	options OptionsLogger
}

var LevelNames = map[slog.Leveler]string{
	LevelTrace:   "TRACE",
	LevelSuccess: "SUCCESS",
}

type LogMessageType interface {
	int | int64 | float64 | string | uint32 | uint64
}

func Argument[T LogMessageType](key string, value T) any {
	switch any(value).(type) {
	case int:
		return slog.Int(key, any(value).(int))
	case string:
		return slog.String(key, any(value).(string))
	case uint64:
		return slog.Uint64(key, any(value).(uint64))
	default:
		return nil
	}
}

func getLevel(l string) slog.Level {
	switch strings.ToUpper(l) {
	case "SUCCESS":
		return LevelSuccess
	case "WARNING":
		return slog.LevelWarn
	case "TRACE":
		return LevelTrace
	case "DEBUG":
		return slog.LevelDebug
	default:
		return slog.LevelInfo
	}
}

func (c *CustomLogger) withOptions() {
	c.opts = &slog.HandlerOptions{
		Level:     getLevel(c.options.Level),
		AddSource: c.options.AddSource,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {

			// Replace msg key with message string
			// if a.Key == slog.MessageKey {
			// 	a.Key = "message"
			// 	return a
			// }

			if a.Key == slog.LevelKey {
				level := a.Value.Any().(slog.Level)
				levelLabel, exists := LevelNames[level]
				if !exists {
					levelLabel = level.String()
				}

				a.Value = slog.StringValue(strings.ToLower(levelLabel))
			}
			return a
		},
	}

	switch c.options.LoggerType {
	case "console":
		c.logger = slog.New(slog.NewTextHandler(os.Stdout, c.opts))
	case "pretty":
		c.logger = slog.New(newPrettyHandler(os.Stdout, c.opts))
	case "json":
		c.logger = slog.New(slog.NewJSONHandler(os.Stdout, c.opts))
	default:
		// Default logger always json
		c.logger = slog.New(slog.NewJSONHandler(os.Stdout, c.opts))
	}
}

func NewLogger(options OptionsLogger) *CustomLogger {
	c := &CustomLogger{ctx: context.Background(), options: options}
	c.withOptions()

	return c
}

func (c *CustomLogger) Success(msg string, args ...any) {
	c.logger.Log(c.ctx, LevelSuccess, msg, args...)
}

func (c *CustomLogger) Debug(msg string, args ...any) {
	c.logger.Log(c.ctx, slog.LevelDebug, msg, args...)
}

func (c *CustomLogger) Info(msg string, args ...any) {
	c.logger.Log(c.ctx, slog.LevelInfo, msg, args...)
}

func (c *CustomLogger) Warning(msg string, args ...any) {
	c.logger.Log(c.ctx, slog.LevelWarn, msg, args...)
}

func (c *CustomLogger) Error(msg string, args ...any) {
	c.logger.Log(c.ctx, slog.LevelError, msg, args...)
}

func (c *CustomLogger) ErrorWithExit(msg string, args ...any) {
	c.logger.Log(c.ctx, slog.LevelError, msg, args...)
	os.Exit(1)
}
