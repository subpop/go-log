// Package log implements a simple level logging package. It defines a type,
// Logger, with methods for formatting output at specific verbosity levels.
// It maintains API compatibility with the standard library "log" package by
// embedding a 'log.Logger' struct.
package log

import (
	"io"
	"log"
	"os"
	"strings"
)

// These variables are exported in order to maintain package API compatibility.
const (
	Ldate         = log.Ldate
	Ltime         = log.Ltime
	Lmicroseconds = log.Lmicroseconds
	Llongfile     = log.Llongfile
	Lshortfile    = log.Lshortfile
	LUTC          = log.LUTC
	LstdFlags     = log.LstdFlags
)

// A ParseError occurs when a string value cannot be parsed to a valid level.
type ParseError string

func (e ParseError) Error() string {
	return "invalid level value: " + string(e)
}

// Level defines various levels of logging output
type Level int

// LevelInfo is the default level.
const (
	LevelInfo Level = iota
	LevelWarn
	LevelError
	LevelDebug
	LevelTrace

	LevelUnknown Level = -1
)

// ParseLevel returns the Level value represented by the string.
func ParseLevel(str string) (Level, error) {
	switch strings.ToUpper(str) {
	case "INFO":
		return LevelInfo, nil
	case "WARN":
		return LevelWarn, nil
	case "ERROR":
		return LevelError, nil
	case "DEBUG":
		return LevelDebug, nil
	case "TRACE":
		return LevelTrace, nil
	}
	return LevelUnknown, ParseError(str)
}

// FormatLevel converts the Level to a string.
func FormatLevel(lvl Level) string {
	return lvl.String()
}

func (l Level) String() string {
	switch l {
	case LevelInfo:
		return "INFO"
	case LevelWarn:
		return "WARN"
	case LevelError:
		return "ERROR"
	case LevelDebug:
		return "DEBUG"
	case LevelTrace:
		return "TRACE"
	default:
		return ""
	}
}

// A Logger represents a standard library Logger, configured for output at a
// given verbosity level. See standard library log.Logger for detailed usage.
type Logger struct {
	*log.Logger
	Level Level
}

// New creates a new Logger. The out, prefix and flag variables are passed
// directly to the embedded standard library log.New function. The level
// argument defines a minimum verbosity level.
func New(out io.Writer, prefix string, flag int, level Level) *Logger {
	return &Logger{
		Logger: log.New(out, prefix, flag),
		Level:  level,
	}
}

var std = New(os.Stderr, "", log.LstdFlags, LevelInfo)

// Info prints to the logger if level is at least LevelInfo. Arguments are
// handled in the manner of fmt.Print.
func (l *Logger) Info(v ...interface{}) {
	if l.Level >= LevelInfo {
		l.Logger.Print(v...)
	}
}

// Infof prints to the logger if level is at least LevelInfo. Arguments are
// handled in the manner of fmt.Printf.
func (l *Logger) Infof(format string, v ...interface{}) {
	if l.Level >= LevelInfo {
		l.Logger.Printf(format, v...)
	}
}

// Infoln prints to the logger if level is at least LevelInfo. Arguments are
// handled in the manner of fmt.Println.
func (l *Logger) Infoln(v ...interface{}) {
	if l.Level >= LevelInfo {
		l.Logger.Println(v...)
	}
}

// Warn prints to the logger if level is at least LevelWarn. Arguments are
// handled in the manner of fmt.Print.
func (l *Logger) Warn(v ...interface{}) {
	if l.Level >= LevelWarn {
		l.Logger.Print(v...)
	}
}

// Warnf prints to the logger if level is at least LevelWarn. Arguments are
// handled in the manner of fmt.Printf.
func (l *Logger) Warnf(format string, v ...interface{}) {
	if l.Level >= LevelWarn {
		l.Logger.Printf(format, v...)
	}
}

// Warnln prints to the logger if level is at least LevelWarn. Arguments are
// handled in the manner of fmt.Println.
func (l *Logger) Warnln(v ...interface{}) {
	if l.Level >= LevelWarn {
		l.Logger.Println(v...)
	}
}

// Error prints to the logger if level is at least LevelError. Arguments are
// handled in the manner of fmt.Print.
func (l *Logger) Error(v ...interface{}) {
	if l.Level >= LevelError {
		l.Logger.Print(v...)
	}
}

// Errorf prints to the logger if level is at least LevelError. Arguments are
// handled in the manner of fmt.Printf.
func (l *Logger) Errorf(format string, v ...interface{}) {
	if l.Level >= LevelError {
		l.Logger.Printf(format, v...)
	}
}

// Errorln prints to the logger if level is at least LevelError. Arguments are
// handled in the manner of fmt.Println.
func (l *Logger) Errorln(v ...interface{}) {
	if l.Level >= LevelError {
		l.Logger.Println(v...)
	}
}

// Debug prints to the logger if level is at least LevelDebug. Arguments are
// handled in the manner of fmt.Print.
func (l *Logger) Debug(v ...interface{}) {
	if l.Level >= LevelDebug {
		l.Logger.Print(v...)
	}
}

// Debugf prints to the logger if level is at least LevelDebug. Arguments are
// handled in the manner of fmt.Printf.
func (l *Logger) Debugf(format string, v ...interface{}) {
	if l.Level >= LevelDebug {
		l.Logger.Printf(format, v...)
	}
}

// Debugln prints to the logger if level is at least LevelDebug. Arguments are
// handled in the manner of fmt.Println.
func (l *Logger) Debugln(v ...interface{}) {
	if l.Level >= LevelDebug {
		l.Logger.Println(v...)
	}
}

// Trace prints to the logger if level is at least LevelTrace. Arguments are
// handled in the manner of fmt.Print.
func (l *Logger) Trace(v ...interface{}) {
	if l.Level >= LevelTrace {
		l.Logger.Print(v...)
	}
}

// Tracef prints to the logger if level is at least LevelTrace. Arguments are
// handled in the manner of fmt.Printf.
func (l *Logger) Tracef(format string, v ...interface{}) {
	if l.Level >= LevelTrace {
		l.Logger.Printf(format, v...)
	}
}

// Traceln prints to the logger if level is at least LevelTrace. Arguments are
// handled in the manner of fmt.Println.
func (l *Logger) Traceln(v ...interface{}) {
	if l.Level >= LevelTrace {
		l.Logger.Println(v...)
	}
}

// SetOutput sets the output destination for the standard logger.
func SetOutput(w io.Writer) {
	std.Logger.SetOutput(w)
}

// Flags returns the output flags for the standard logger.
func Flags() int {
	return std.Flags()
}

// SetFlags sets the output flags for the standard logger.
func SetFlags(flag int) {
	std.Logger.SetFlags(flag)
}

// Prefix returns the output prefix for the standard logger.
func Prefix() string {
	return std.Logger.Prefix()
}

// SetPrefix sets the output prefix for the standard logger.
func SetPrefix(prefix string) {
	std.Logger.SetPrefix(prefix)
}

// Writer returns the output destination for the standard logger.
func Writer() io.Writer {
	return std.Logger.Writer()
}

// CurrentLevel returns the output level for the standard logger.
func CurrentLevel() Level {
	return std.Level
}

// SetLevel sets the output level for the standard logger.
func SetLevel(l Level) {
	std.Level = l
}

// These functions write to the standard logger.

// Info prints to the standard logger if level is at least LevelInfo.
// Arguments are handled in the manner of fmt.Print.
func Info(v ...interface{}) {
	std.Info(v...)
}

// Infof prints to the standard logger if level is at least LevelInfo.
// Arguments are handled in the manner of fmt.Printf.
func Infof(format string, v ...interface{}) {
	std.Infof(format, v...)
}

// Infoln prints to the standard logger if level is at least LevelInfo.
// Arguments are handled in the manner of fmt.Println.
func Infoln(v ...interface{}) {
	std.Infoln(v...)
}

// Warn prints to the standard logger if level is at least LevelWarn.
// Arguments are handled in the manner of fmt.Print.
func Warn(v ...interface{}) {
	std.Warn(v...)
}

// Warnf prints to the standard logger if level is at least LevelWarn.
// Arguments are handled in the manner of fmt.Printf.
func Warnf(format string, v ...interface{}) {
	std.Warnf(format, v...)
}

// Warnln prints to the standard logger if level is at least LevelWarn.
// Arguments are handled in the manner of fmt.Println.
func Warnln(v ...interface{}) {
	std.Warnln(v...)
}

// Error prints to the standard logger if level is at least LevelError.
// Arguments are handled in the manner of fmt.Print.
func Error(v ...interface{}) {
	std.Error(v...)
}

// Errorf prints to the standard logger if level is at least LevelError.
// Arguments are handled in the manner of fmt.Printf.
func Errorf(format string, v ...interface{}) {
	std.Errorf(format, v...)
}

// Errorln prints to the standard logger if level is at least LevelError.
// Arguments are handled in the manner of fmt.Println.
func Errorln(v ...interface{}) {
	std.Errorln(v...)
}

// Debug prints to the standard logger if level is at least LevelDebug.
// Arguments are handled in the manner of fmt.Print.
func Debug(v ...interface{}) {
	std.Debug(v...)
}

// Debugf prints to the standard logger if level is at least LevelDebug.
// Arguments are handled in the manner of fmt.Printf.
func Debugf(format string, v ...interface{}) {
	std.Debugf(format, v...)
}

// Debugln prints to the standard logger if level is at least LevelDebug.
// Arguments are handled in the manner of fmt.Println.
func Debugln(v ...interface{}) {
	std.Debugln(v...)
}

// Trace prints to the standard logger if level is at least LevelTrace.
// Arguments are handled in the manner of fmt.Print.
func Trace(v ...interface{}) {
	std.Trace(v...)
}

// Tracef prints to the standard logger if level is at least LevelTrace.
// Arguments are handled in the manner of fmt.Printf.
func Tracef(format string, v ...interface{}) {
	std.Tracef(format, v...)
}

// Traceln prints to the standard logger if level is at least LevelTrace.
// Arguments are handled in the manner of fmt.Println.
func Traceln(v ...interface{}) {
	std.Traceln(v...)
}

// Output simply wraps a call to the standard logger's Output function to
// maintain API compatibility.
func Output(calldepth int, s string) error {
	return std.Logger.Output(calldepth+1, s)
}
