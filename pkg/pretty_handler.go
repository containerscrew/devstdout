package logger

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/fatih/color"
	"golang.org/x/exp/slog"
)

type PrettyHandlerOptions struct {
	SlogOpts slog.HandlerOptions
}

type PrettyHandler struct {
	slog.Handler
	l *log.Logger
}

// Function to convert the map into a "key=value" formatted string
func formatFields(fields map[string]interface{}) string {
	var result string
	for key, value := range fields {
		result += fmt.Sprintf("%s=%v ", key, value) // %v allows printing any type of value
	}
	// Remove the extra trailing space
	if len(result) > 0 {
		result = result[:len(result)-1]
	}
	return result
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

	timeStr := r.Time.Format("[15:05:05]")
	msg := color.CyanString(r.Message)

	// Use the function to format the map as "key=value"
	formattedFields := formatFields(fields)

	h.l.Println(timeStr, level, msg, color.WhiteString(formattedFields))

	return nil
}

func newPrettyHandler(out io.Writer, opts *slog.HandlerOptions) *PrettyHandler {
	return &PrettyHandler{
		Handler: slog.NewTextHandler(out, opts),
		l:       log.New(out, "", 0),
	}
}
