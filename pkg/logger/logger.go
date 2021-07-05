package logger

import (
	"fmt"

	log "github.com/mgutz/logxi/v1"
)

func init() {
	log.KeyMap = &log.KeyMapping{
		Level:     "lvl",
		Message:   "msg",
		Name:      "name",
		PID:       "pid",
		Time:      "time",
		CallStack: "stack",
	}
}

// Logger is a wrapper around github.com/mgutz/logxi/v1
// that removes the returned error from Warn and Error
// and allows for sub-loggers
type Logger struct {
	path string
	log.Logger
}

// New creates a colorable default logger.
func New(name string) *Logger {
	return &Logger{Logger: log.New(name), path: name}
}

// New creates a sub-logger of the original one
func (l *Logger) New(path string) *Logger {
	if l.path != "" {
		path = fmt.Sprintf("%s.%s", l.path, path)
	}
	return &Logger{path: path, Logger: log.New(path)}
}

// Warn logs a warning statement. On terminals it logs file and line number.
func (l *Logger) Warn(msg string, args ...interface{}) {
	_ = l.Logger.Warn(msg, args...)
}

// Error logs an error statement with callstack.
func (l *Logger) Error(msg string, args ...interface{}) {
	_ = l.Logger.Error(msg, args...)
}
