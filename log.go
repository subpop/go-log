// Package log implements a simple level logging package. It defines a type,
// Logger, with methods for formatting output at specific verbosity levels.
// It maintains API compatibility with the standard library "log" package by
// embedding a 'log.Logger' struct.
package log

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

// These flags define which text to prefix to each log entry generated by the Logger.
// Bits are or'ed together to control what's printed.
// There is no control over the order they appear (the order listed
// here) or the format they present (as described in the comments).
// The prefix is followed by a colon only when Llongfile or Lshortfile
// is specified.
// For example, flags Ldate | Ltime (or LstdFlags) produce,
//
//	2009/01/23 01:23:23 message
//
// while flags Ldate | Ltime | Lmicroseconds | Llongfile produce,
//
//	2009/01/23 01:23:23.123123 /a/b/c/d.go:23: message
//
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

// LevelError is the default level.
const (
	LevelError Level = iota
	LevelWarn
	LevelInfo
	LevelDebug
	LevelTrace

	LevelUnknown Level = -1
)

// ParseLevel returns the Level value represented by the string.
func ParseLevel(str string) (Level, error) {
	switch strings.ToUpper(str) {
	case "ERROR":
		return LevelError, nil
	case "WARN":
		return LevelWarn, nil
	case "INFO":
		return LevelInfo, nil
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
	case LevelError:
		return "ERROR"
	case LevelWarn:
		return "WARN"
	case LevelInfo:
		return "INFO"
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

var std = New(os.Stderr, "", log.LstdFlags, LevelError)

// Error prints to the logger if level is at least LevelError. Arguments are
// handled in the manner of fmt.Print.
func (l *Logger) Error(v ...interface{}) {
	if l.Level >= LevelError {
		_ = l.Output(3, fmt.Sprint(v...))
	}
}

// Errorf prints to the logger if level is at least LevelError. Arguments are
// handled in the manner of fmt.Printf.
func (l *Logger) Errorf(format string, v ...interface{}) {
	if l.Level >= LevelError {
		_ = l.Output(3, fmt.Sprintf(format, v...))
	}
}

// Errorln prints to the logger if level is at least LevelError. Arguments are
// handled in the manner of fmt.Println.
func (l *Logger) Errorln(v ...interface{}) {
	if l.Level >= LevelError {
		_ = l.Output(3, fmt.Sprintln(v...))
	}
}

// Warn prints to the logger if level is at least LevelWarn. Arguments are
// handled in the manner of fmt.Print.
func (l *Logger) Warn(v ...interface{}) {
	if l.Level >= LevelWarn {
		_ = l.Logger.Output(3, fmt.Sprint(v...))
	}
}

// Warnf prints to the logger if level is at least LevelWarn. Arguments are
// handled in the manner of fmt.Printf.
func (l *Logger) Warnf(format string, v ...interface{}) {
	if l.Level >= LevelWarn {
		_ = l.Logger.Output(3, fmt.Sprintf(format, v...))
	}
}

// Warnln prints to the logger if level is at least LevelWarn. Arguments are
// handled in the manner of fmt.Println.
func (l *Logger) Warnln(v ...interface{}) {
	if l.Level >= LevelWarn {
		_ = l.Logger.Output(3, fmt.Sprintln(v...))
	}
}

// Info prints to the logger if level is at least LevelInfo. Arguments are
// handled in the manner of fmt.Print.
func (l *Logger) Info(v ...interface{}) {
	if l.Level >= LevelInfo {
		_ = l.Logger.Output(3, fmt.Sprint(v...))
	}
}

// Infof prints to the logger if level is at least LevelInfo. Arguments are
// handled in the manner of fmt.Printf.
func (l *Logger) Infof(format string, v ...interface{}) {
	if l.Level >= LevelInfo {
		_ = l.Logger.Output(3, fmt.Sprintf(format, v...))
	}
}

// Infoln prints to the logger if level is at least LevelInfo. Arguments are
// handled in the manner of fmt.Println.
func (l *Logger) Infoln(v ...interface{}) {
	if l.Level >= LevelInfo {
		_ = l.Logger.Output(3, fmt.Sprintln(v...))
	}
}

// Debug prints to the logger if level is at least LevelDebug. Arguments are
// handled in the manner of fmt.Print.
func (l *Logger) Debug(v ...interface{}) {
	if l.Level >= LevelDebug {
		_ = l.Logger.Output(3, fmt.Sprint(v...))
	}
}

// Debugf prints to the logger if level is at least LevelDebug. Arguments are
// handled in the manner of fmt.Printf.
func (l *Logger) Debugf(format string, v ...interface{}) {
	if l.Level >= LevelDebug {
		_ = l.Logger.Output(3, fmt.Sprintf(format, v...))
	}
}

// Debugln prints to the logger if level is at least LevelDebug. Arguments are
// handled in the manner of fmt.Println.
func (l *Logger) Debugln(v ...interface{}) {
	if l.Level >= LevelDebug {
		_ = l.Logger.Output(3, fmt.Sprintln(v...))
	}
}

// Trace prints to the logger if level is at least LevelTrace. Arguments are
// handled in the manner of fmt.Print.
func (l *Logger) Trace(v ...interface{}) {
	if l.Level >= LevelTrace {
		_ = l.Logger.Output(3, fmt.Sprint(v...))
	}
}

// Tracef prints to the logger if level is at least LevelTrace. Arguments are
// handled in the manner of fmt.Printf.
func (l *Logger) Tracef(format string, v ...interface{}) {
	if l.Level >= LevelTrace {
		_ = l.Logger.Output(3, fmt.Sprintf(format, v...))
	}
}

// Traceln prints to the logger if level is at least LevelTrace. Arguments are
// handled in the manner of fmt.Println.
func (l *Logger) Traceln(v ...interface{}) {
	if l.Level >= LevelTrace {
		_ = l.Logger.Output(3, fmt.Sprintln(v...))
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

// Print calls Print to print to the standard logger.
// Arguments are handled in the manner of fmt.Print.
func Print(v ...interface{}) {
	std.Print(v...)
}

// Printf calls Printf to print to the standard logger.
// Arguments are handled in the manner of fmt.Printf.
func Printf(format string, v ...interface{}) {
	std.Printf(format, v...)
}

// Println calls Println to print to the standard logger.
// Arguments are handled in the manner of fmt.Println.
func Println(v ...interface{}) {
	std.Println(v...)
}

// Fatal is equivalent to Print() followed by a call to os.Exit(1).
func Fatal(v ...interface{}) {
	std.Fatal(v...)
}

// Fatalf is equivalent to Printf() followed by a call to os.Exit(1).
func Fatalf(format string, v ...interface{}) {
	std.Fatalf(format, v...)
}

// Fatalln is equivalent to Println() followed by a call to os.Exit(1).
func Fatalln(v ...interface{}) {
	std.Fatalln(v...)
}

// Panic is equivalent to Print() followed by a call to panic().
func Panic(v ...interface{}) {
	std.Panic(v...)
}

// Panicf is equivalent to Printf() followed by a call to panic().
func Panicf(format string, v ...interface{}) {
	std.Panicf(format, v...)
}

// Panicln is equivalent to Println() followed by a call to panic().
func Panicln(v ...interface{}) {
	std.Panicln(v...)
}

// Output wraps a call to the standard logger's Output function for API
// compatibility.
//
// Output writes the output for a logging event. The string s contains
// the text to print after the prefix specified by the flags of the
// Logger. A newline is appended if the last character of s is not
// already a newline. Calldepth is the count of the number of
// frames to skip when computing the file name and line number
// if Llongfile or Lshortfile is set; a value of 1 will print the details
// for the caller of Output.
func Output(calldepth int, s string) error {
	return std.Logger.Output(calldepth+2, s)
}
