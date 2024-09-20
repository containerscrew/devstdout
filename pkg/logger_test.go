package devstdout

import (
	"testing"

	"golang.org/x/exp/slog"
)

func TestGetLevel(t *testing.T) {
	tests := []struct {
		input    string
		expected slog.Level
	}{
		{"SUCCESS", LevelSuccess},
		{"WARNING", slog.LevelWarn},
		{"TRACE", LevelTrace},
		{"DEBUG", slog.LevelDebug},
		{"INFO", slog.LevelInfo},
		{"unknown", slog.LevelInfo},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			result := getLevel(test.input)
			if result != test.expected {
				t.Errorf("getLevel(%s) = %v; want %v", test.input, result, test.expected)
			}
		})
	}
}
