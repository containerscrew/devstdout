package logger

import (
	"context"
	"encoding/json"
	"github.com/fatih/color"
	"golang.org/x/exp/slog"
	"io"
	"log"
	"os"
)

const (
	LevelTrace   = slog.Level(-8)
	LevelSuccess = slog.Level(6)
)

type CustomLogger struct {
	Logger *slog.Logger
	opts   *slog.HandlerOptions
	ctx    context.Context
}

var LevelNames = map[slog.Leveler]string{
	LevelTrace:   "TRACE",
	LevelSuccess: "SUCCESS",
}

type PrettyHandlerOptions struct {
	SlogOpts slog.HandlerOptions
}

type PrettyHandler struct {
	slog.Handler
	l *log.Logger
}

// Mirar generics tete
func PrintString(key, value string) any {
	return slog.String(key, value)
}

func PrintInt(key string, value int) any {
	return slog.Int(key, value)
}

func (c *CustomLogger) withOptions(level string, addSource bool) {
	c.opts = &slog.HandlerOptions{
		Level:     LevelTrace,
		AddSource: addSource,
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

func (h *PrettyHandler) Handle(_ context.Context, r slog.Record) error {
	level := r.Level.String()

	switch r.Level {
	case slog.LevelDebug:
		level = color.MagentaString(level)
	case slog.LevelInfo:
		level = color.BlueString(level)
	case slog.LevelWarn:
		level = color.YellowString(level)
	case slog.LevelError:
		level = color.RedString(level)
	case LevelSuccess:
		level = color.GreenString(LevelNames[LevelSuccess])
	}

	fields := make(map[string]interface{}, r.NumAttrs())
	r.Attrs(func(a slog.Attr) bool {
		fields[a.Key] = a.Value.Any()
		return true
	})

	//b, err := json.MarshalIndent(fields, "", "  ")
	data, err := json.Marshal(fields)
	if err != nil {
		return err
	}

	timeStr := r.Time.Format("[15:05:05]")
	msg := color.CyanString(r.Message)

	h.l.Println(timeStr, level, msg, color.WhiteString(string(data)))

	return nil
}

func NewPrettyHandler(out io.Writer, opts *slog.HandlerOptions) *PrettyHandler {
	h := &PrettyHandler{
		Handler: slog.NewJSONHandler(out, opts),
		l:       log.New(out, "", 0),
	}

	return h
}

func NewLogger(level, env string, addSource bool) *CustomLogger {
	c := &CustomLogger{ctx: context.Background()}
	c.withOptions(level, addSource)
	c.Logger = slog.New(NewPrettyHandler(os.Stdout, c.opts))

	if env == "prod" {
		c.Logger = slog.New(slog.NewJSONHandler(os.Stdout, c.opts))
	}

	return c
}

func (c *CustomLogger) Success(msg string) {
	c.Logger.Log(c.ctx, LevelSuccess, msg)
}
