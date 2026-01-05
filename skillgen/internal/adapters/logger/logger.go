package logger

import (
	"fmt"
	"log"
	"strings"

	"github.com/adaptive-enforcement-lab/claude-skills/skillgen/internal/ports"
)

// Logger implements ports.Logger using standard library log package.
type Logger struct {
	level  ports.LogLevel
	prefix string
}

// NewLogger creates a new logger with the specified log level.
func NewLogger(level ports.LogLevel) *Logger {
	return &Logger{
		level: level,
	}
}

// Info logs an informational message.
func (l *Logger) Info(msg string, keysAndValues ...interface{}) {
	if l.level > ports.LogLevelInfo {
		return
	}
	log.Printf("[INFO] %s%s", l.prefix, l.formatMessage(msg, keysAndValues...))
}

// Warn logs a warning message.
func (l *Logger) Warn(msg string, keysAndValues ...interface{}) {
	if l.level > ports.LogLevelWarn {
		return
	}
	log.Printf("[WARN] %s%s", l.prefix, l.formatMessage(msg, keysAndValues...))
}

// Error logs an error message.
func (l *Logger) Error(msg string, keysAndValues ...interface{}) {
	if l.level > ports.LogLevelError {
		return
	}
	log.Printf("[ERROR] %s%s", l.prefix, l.formatMessage(msg, keysAndValues...))
}

// Debug logs a debug message.
func (l *Logger) Debug(msg string, keysAndValues ...interface{}) {
	if l.level > ports.LogLevelDebug {
		return
	}
	log.Printf("[DEBUG] %s%s", l.prefix, l.formatMessage(msg, keysAndValues...))
}

// With creates a child logger with additional context fields.
func (l *Logger) With(keysAndValues ...interface{}) ports.Logger {
	child := &Logger{
		level:  l.level,
		prefix: l.prefix + l.formatContext(keysAndValues...),
	}
	return child
}

// formatMessage formats a log message with key-value pairs.
func (l *Logger) formatMessage(msg string, keysAndValues ...interface{}) string {
	if len(keysAndValues) == 0 {
		return msg
	}

	var pairs []string
	for i := 0; i < len(keysAndValues); i += 2 {
		if i+1 < len(keysAndValues) {
			key := fmt.Sprintf("%v", keysAndValues[i])
			value := fmt.Sprintf("%v", keysAndValues[i+1])
			pairs = append(pairs, fmt.Sprintf("%s=%s", key, value))
		}
	}

	if len(pairs) > 0 {
		return fmt.Sprintf("%s %s", msg, strings.Join(pairs, " "))
	}

	return msg
}

// formatContext formats key-value pairs for logger context.
func (l *Logger) formatContext(keysAndValues ...interface{}) string {
	if len(keysAndValues) == 0 {
		return ""
	}

	var pairs []string
	for i := 0; i < len(keysAndValues); i += 2 {
		if i+1 < len(keysAndValues) {
			key := fmt.Sprintf("%v", keysAndValues[i])
			value := fmt.Sprintf("%v", keysAndValues[i+1])
			pairs = append(pairs, fmt.Sprintf("[%s=%s]", key, value))
		}
	}

	return strings.Join(pairs, " ") + " "
}
