package config

import (
	"io"
	"log"
	"os"
)

type Logger struct {
  debug *log.Logger
  info *log.Logger
  warning *log.Logger 
  error *log.Logger
  writer io.Writer
}

// NewLogger creates a new Logger instance with the specified prefix for log messages.
func NewLogger(prefix string) *Logger {
  writer := io.Writer(os.Stdout)
  logger := log.New(writer, prefix, log.Ldate|log.Ltime)

  // Create separate loggers for different log levels (debug, info, warning, error) with appropriate prefixes.
  return &Logger{
    debug: log.New(writer, "> DEBUG: ", logger.Flags()),
    info: log.New(writer, "> INFO: ", logger.Flags()),
    warning: log.New(writer, "> WARNING: ", logger.Flags()),
    error: log.New(writer, "> ERROR: ", logger.Flags()),
    writer: writer,
  }
}

// Non-formatted Logs
func (l *Logger) Debug(v ...interface{}) {
  l.debug.Println(v...)
}
func (l *Logger) Info(v ...interface{}) {
    l.info.Println(v...)
}
func (l *Logger) Warn(v ...interface{}) {
  l.warning.Println(v...)
}
func (l *Logger) Error(v ...interface{}) {
  l.error.Println(v...)
}

// Formated Logs
func (l *Logger) Debugf(format string, v ... interface{}) {
  l.debug.Printf(format, v...)
}
func (l *Logger) Infof(format string, v ... interface{}) {
  l.info.Printf(format, v...)
}
func (l *Logger) Warnf(format string, v ... interface{}) {
  l.warning.Printf(format, v...)
}
func (l *Logger) Errorf(format string, v ... interface{}) {
  l.error.Printf(format, v...)
}
