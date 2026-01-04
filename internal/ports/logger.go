package ports

// Logger provides structured logging capabilities.
// This interface abstracts the logging library to enable testing and flexibility.
type Logger interface {
	// Info logs an informational message.
	Info(msg string, keysAndValues ...interface{})

	// Warn logs a warning message.
	Warn(msg string, keysAndValues ...interface{})

	// Error logs an error message.
	Error(msg string, keysAndValues ...interface{})

	// Debug logs a debug message (only shown in verbose mode).
	Debug(msg string, keysAndValues ...interface{})

	// With creates a child logger with additional context fields.
	With(keysAndValues ...interface{}) Logger
}

// LogLevel represents the minimum logging level.
type LogLevel int

const (
	LogLevelDebug LogLevel = iota
	LogLevelInfo
	LogLevelWarn
	LogLevelError
)
