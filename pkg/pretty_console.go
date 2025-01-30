package devstdout

import (
	"context"
	"fmt"
	"io"
	"os"
	"sync"

	"github.com/fatih/color"
	"golang.org/x/exp/slog"
)

type PrettyHandlerOptions struct {
	SlogOpts slog.HandlerOptions
}

type PrettyHandler struct {
	slog.Handler
	mu sync.Mutex // Mutex for thread-safe logging
}

// Convert map fields into a formatted string "key=value"
func formatFields(fields map[string]interface{}) string {
	var result string
	for key, value := range fields {
		result += fmt.Sprintf("%s=%v ", key, value)
	}
	// Trim trailing space
	if len(result) > 0 {
		result = result[:len(result)-1]
	}
	return result
}

func (h *PrettyHandler) Handle(_ context.Context, r slog.Record) error {
	h.mu.Lock()         // Lock to ensure thread-safe logging
	defer h.mu.Unlock() // Unlock when done

	level := r.Level.String()

	// Apply color formatting based on log level
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

	// Collect attributes into a map
	fields := make(map[string]interface{}, r.NumAttrs())
	r.Attrs(func(a slog.Attr) bool {
		fields[a.Key] = a.Value.Any()
		return true
	})

	// Format timestamp
	timeStr := r.Time.Format("[2006-01-02 15:05:05]")
	msg := color.CyanString(r.Message)

	// Format fields as key=value
	formattedFields := formatFields(fields)

	// Print log entry using fmt.Printf (directly writing to stdout)
	fmt.Printf("%s %s %s %s\n", timeStr, level, msg, color.WhiteString(formattedFields))

	// Force flush stdout to prevent logs from being out of order
	os.Stdout.Sync()

	return nil
}

// Create a new PrettyHandler
func newPrettyHandler(out io.Writer, opts *slog.HandlerOptions) *PrettyHandler {
	return &PrettyHandler{
		Handler: slog.NewTextHandler(out, opts),
	}
}
