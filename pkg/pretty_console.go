package devstdout

import (
	"context"
	"fmt"
	"io"
	"log"
	"sync"

	"github.com/fatih/color"
	"golang.org/x/exp/slog"
)

type PrettyHandlerOptions struct {
	SlogOpts slog.HandlerOptions
}

type PrettyHandler struct {
	slog.Handler
	mu *sync.Mutex  // Mutex to ensure thread-safety
	l  *log.Logger  // Custom logger
}

// Function to convert the map into a "key=value" formatted string
func formatFields(fields map[string]interface{}) string {
	var result string
	for key, value := range fields {
		result += fmt.Sprintf("%s=%v ", key, value)
	}
	// Remove the extra trailing space
	if len(result) > 0 {
		result = result[:len(result)-1]
	}
	return result
}

func (h *PrettyHandler) Handle(_ context.Context, r slog.Record) error {
	h.mu.Lock()         // Lock to ensure thread-safe writing
	defer h.mu.Unlock() // Unlock when done

	level := r.Level.String()

	// Apply color based on log level
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

	// Collect attributes from the log record into a map
	fields := make(map[string]interface{}, r.NumAttrs())
	r.Attrs(func(a slog.Attr) bool {
		fields[a.Key] = a.Value.Any()
		return true
	})

	// Format time for logging
	timeStr := r.Time.Format("[2006-01-02 15:05:05]")
	msg := color.CyanString(r.Message)

	// Use the function to format the map as "key=value"
	formattedFields := formatFields(fields)

	// Use fmt.Printf for formatted output, it provides better control over the format
	fmt.Printf("%s %s %s %s\n", timeStr, level, msg, color.WhiteString(formattedFields))

	return nil
}

// Create a new PrettyHandler
func newPrettyHandler(out io.Writer, opts *slog.HandlerOptions) *PrettyHandler {
	return &PrettyHandler{
		Handler: slog.NewTextHandler(out, opts),
		mu:      &sync.Mutex{},
		l:       log.New(out, "", 0),
	}
}
