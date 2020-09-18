package log_test

import (
	"strings"
	"testing"

	log "git.sr.ht/~spc/go-log"
)

func TestLog(t *testing.T) {
	tests := []struct {
		description string
		input       struct {
			msg    string
			logLvl log.Level
			atLvl  log.Level
		}
		want string
	}{
		{
			input: struct {
				msg    string
				logLvl log.Level
				atLvl  log.Level
			}{
				msg:    "visible",
				logLvl: log.LevelInfo,
				atLvl:  log.LevelInfo,
			},
			want: "visible\n",
		},
		{
			input: struct {
				msg    string
				logLvl log.Level
				atLvl  log.Level
			}{
				msg:    "visible",
				logLvl: log.LevelWarn,
				atLvl:  log.LevelWarn,
			},
			want: "visible\n",
		},
		{
			input: struct {
				msg    string
				logLvl log.Level
				atLvl  log.Level
			}{
				msg:    "visible",
				logLvl: log.LevelError,
				atLvl:  log.LevelError,
			},
			want: "visible\n",
		},
		{
			input: struct {
				msg    string
				logLvl log.Level
				atLvl  log.Level
			}{
				msg:    "visible",
				logLvl: log.LevelDebug,
				atLvl:  log.LevelDebug,
			},
			want: "visible\n",
		},
		{
			input: struct {
				msg    string
				logLvl log.Level
				atLvl  log.Level
			}{
				msg:    "visible",
				logLvl: log.LevelTrace,
				atLvl:  log.LevelTrace,
			},
			want: "visible\n",
		},
		{
			input: struct {
				msg    string
				logLvl log.Level
				atLvl  log.Level
			}{
				msg:    "visible",
				logLvl: log.LevelWarn,
				atLvl:  log.LevelInfo,
			},
			want: "visible\n",
		},
		{
			input: struct {
				msg    string
				logLvl log.Level
				atLvl  log.Level
			}{
				msg:    "visible",
				logLvl: log.LevelWarn,
				atLvl:  log.LevelError,
			},
			want: "",
		},
		{
			input: struct {
				msg    string
				logLvl log.Level
				atLvl  log.Level
			}{
				msg:    "visible",
				logLvl: log.LevelWarn,
				atLvl:  log.LevelDebug,
			},
			want: "",
		},
		{
			input: struct {
				msg    string
				logLvl log.Level
				atLvl  log.Level
			}{
				msg:    "visible",
				logLvl: log.LevelWarn,
				atLvl:  log.LevelTrace,
			},
			want: "",
		},
		{
			input: struct {
				msg    string
				logLvl log.Level
				atLvl  log.Level
			}{
				msg:    "visible",
				logLvl: log.LevelError,
				atLvl:  log.LevelInfo,
			},
			want: "visible\n",
		},
		{
			input: struct {
				msg    string
				logLvl log.Level
				atLvl  log.Level
			}{
				msg:    "visible",
				logLvl: log.LevelError,
				atLvl:  log.LevelWarn,
			},
			want: "visible\n",
		},
		{
			input: struct {
				msg    string
				logLvl log.Level
				atLvl  log.Level
			}{
				msg:    "visible",
				logLvl: log.LevelError,
				atLvl:  log.LevelDebug,
			},
			want: "",
		},
		{
			input: struct {
				msg    string
				logLvl log.Level
				atLvl  log.Level
			}{
				msg:    "visible",
				logLvl: log.LevelError,
				atLvl:  log.LevelTrace,
			},
			want: "",
		},
		{
			input: struct {
				msg    string
				logLvl log.Level
				atLvl  log.Level
			}{
				msg:    "visible",
				logLvl: log.LevelDebug,
				atLvl:  log.LevelInfo,
			},
			want: "visible\n",
		},
		{
			input: struct {
				msg    string
				logLvl log.Level
				atLvl  log.Level
			}{
				msg:    "visible",
				logLvl: log.LevelDebug,
				atLvl:  log.LevelWarn,
			},
			want: "visible\n",
		},
		{
			input: struct {
				msg    string
				logLvl log.Level
				atLvl  log.Level
			}{
				msg:    "visible",
				logLvl: log.LevelDebug,
				atLvl:  log.LevelError,
			},
			want: "visible\n",
		},
		{
			input: struct {
				msg    string
				logLvl log.Level
				atLvl  log.Level
			}{
				msg:    "visible",
				logLvl: log.LevelDebug,
				atLvl:  log.LevelTrace,
			},
			want: "",
		},
		{
			input: struct {
				msg    string
				logLvl log.Level
				atLvl  log.Level
			}{
				msg:    "visible",
				logLvl: log.LevelTrace,
				atLvl:  log.LevelInfo,
			},
			want: "visible\n",
		},
		{
			input: struct {
				msg    string
				logLvl log.Level
				atLvl  log.Level
			}{
				msg:    "visible",
				logLvl: log.LevelTrace,
				atLvl:  log.LevelWarn,
			},
			want: "visible\n",
		},
		{
			input: struct {
				msg    string
				logLvl log.Level
				atLvl  log.Level
			}{
				msg:    "visible",
				logLvl: log.LevelTrace,
				atLvl:  log.LevelError,
			},
			want: "visible\n",
		},
		{
			input: struct {
				msg    string
				logLvl log.Level
				atLvl  log.Level
			}{
				msg:    "visible",
				logLvl: log.LevelTrace,
				atLvl:  log.LevelDebug,
			},
			want: "visible\n",
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			var writer strings.Builder

			log.SetOutput(&writer)
			log.SetFlags(0)
			log.SetLevel(test.input.logLvl)

			switch test.input.atLvl {
			case log.LevelInfo:
				log.Info(test.input.msg)
			case log.LevelWarn:
				log.Warn(test.input.msg)
			case log.LevelError:
				log.Error(test.input.msg)
			case log.LevelDebug:
				log.Debug(test.input.msg)
			case log.LevelTrace:
				log.Trace(test.input.msg)
			}

			got := writer.String()
			if got != test.want {
				t.Errorf("%+v != %+v", got, test.want)
			}

			// Reset standard logger
			log.SetLevel(log.LevelInfo)
		})
	}
}

func TestLogf(t *testing.T) {
	tests := []struct {
		description string
		input       struct {
			msg    string
			logLvl log.Level
			atLvl  log.Level
		}
		want string
	}{
		{
			input: struct {
				msg    string
				logLvl log.Level
				atLvl  log.Level
			}{
				msg:    "visible",
				logLvl: log.LevelInfo,
				atLvl:  log.LevelInfo,
			},
			want: "visible\n",
		},
		{
			input: struct {
				msg    string
				logLvl log.Level
				atLvl  log.Level
			}{
				msg:    "visible",
				logLvl: log.LevelWarn,
				atLvl:  log.LevelWarn,
			},
			want: "visible\n",
		},
		{
			input: struct {
				msg    string
				logLvl log.Level
				atLvl  log.Level
			}{
				msg:    "visible",
				logLvl: log.LevelError,
				atLvl:  log.LevelError,
			},
			want: "visible\n",
		},
		{
			input: struct {
				msg    string
				logLvl log.Level
				atLvl  log.Level
			}{
				msg:    "visible",
				logLvl: log.LevelDebug,
				atLvl:  log.LevelDebug,
			},
			want: "visible\n",
		},
		{
			input: struct {
				msg    string
				logLvl log.Level
				atLvl  log.Level
			}{
				msg:    "visible",
				logLvl: log.LevelTrace,
				atLvl:  log.LevelTrace,
			},
			want: "visible\n",
		},
		{
			input: struct {
				msg    string
				logLvl log.Level
				atLvl  log.Level
			}{
				msg:    "visible",
				logLvl: log.LevelWarn,
				atLvl:  log.LevelInfo,
			},
			want: "visible\n",
		},
		{
			input: struct {
				msg    string
				logLvl log.Level
				atLvl  log.Level
			}{
				msg:    "visible",
				logLvl: log.LevelWarn,
				atLvl:  log.LevelError,
			},
			want: "",
		},
		{
			input: struct {
				msg    string
				logLvl log.Level
				atLvl  log.Level
			}{
				msg:    "visible",
				logLvl: log.LevelWarn,
				atLvl:  log.LevelDebug,
			},
			want: "",
		},
		{
			input: struct {
				msg    string
				logLvl log.Level
				atLvl  log.Level
			}{
				msg:    "visible",
				logLvl: log.LevelWarn,
				atLvl:  log.LevelTrace,
			},
			want: "",
		},
		{
			input: struct {
				msg    string
				logLvl log.Level
				atLvl  log.Level
			}{
				msg:    "visible",
				logLvl: log.LevelError,
				atLvl:  log.LevelInfo,
			},
			want: "visible\n",
		},
		{
			input: struct {
				msg    string
				logLvl log.Level
				atLvl  log.Level
			}{
				msg:    "visible",
				logLvl: log.LevelError,
				atLvl:  log.LevelWarn,
			},
			want: "visible\n",
		},
		{
			input: struct {
				msg    string
				logLvl log.Level
				atLvl  log.Level
			}{
				msg:    "visible",
				logLvl: log.LevelError,
				atLvl:  log.LevelDebug,
			},
			want: "",
		},
		{
			input: struct {
				msg    string
				logLvl log.Level
				atLvl  log.Level
			}{
				msg:    "visible",
				logLvl: log.LevelError,
				atLvl:  log.LevelTrace,
			},
			want: "",
		},
		{
			input: struct {
				msg    string
				logLvl log.Level
				atLvl  log.Level
			}{
				msg:    "visible",
				logLvl: log.LevelDebug,
				atLvl:  log.LevelInfo,
			},
			want: "visible\n",
		},
		{
			input: struct {
				msg    string
				logLvl log.Level
				atLvl  log.Level
			}{
				msg:    "visible",
				logLvl: log.LevelDebug,
				atLvl:  log.LevelWarn,
			},
			want: "visible\n",
		},
		{
			input: struct {
				msg    string
				logLvl log.Level
				atLvl  log.Level
			}{
				msg:    "visible",
				logLvl: log.LevelDebug,
				atLvl:  log.LevelError,
			},
			want: "visible\n",
		},
		{
			input: struct {
				msg    string
				logLvl log.Level
				atLvl  log.Level
			}{
				msg:    "visible",
				logLvl: log.LevelDebug,
				atLvl:  log.LevelTrace,
			},
			want: "",
		},
		{
			input: struct {
				msg    string
				logLvl log.Level
				atLvl  log.Level
			}{
				msg:    "visible",
				logLvl: log.LevelTrace,
				atLvl:  log.LevelInfo,
			},
			want: "visible\n",
		},
		{
			input: struct {
				msg    string
				logLvl log.Level
				atLvl  log.Level
			}{
				msg:    "visible",
				logLvl: log.LevelTrace,
				atLvl:  log.LevelWarn,
			},
			want: "visible\n",
		},
		{
			input: struct {
				msg    string
				logLvl log.Level
				atLvl  log.Level
			}{
				msg:    "visible",
				logLvl: log.LevelTrace,
				atLvl:  log.LevelError,
			},
			want: "visible\n",
		},
		{
			input: struct {
				msg    string
				logLvl log.Level
				atLvl  log.Level
			}{
				msg:    "visible",
				logLvl: log.LevelTrace,
				atLvl:  log.LevelDebug,
			},
			want: "visible\n",
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			var writer strings.Builder

			log.SetOutput(&writer)
			log.SetFlags(0)
			log.SetLevel(test.input.logLvl)

			switch test.input.atLvl {
			case log.LevelInfo:
				log.Infof("%s\n", test.input.msg)
			case log.LevelWarn:
				log.Warnf("%s\n", test.input.msg)
			case log.LevelError:
				log.Errorf("%s\n", test.input.msg)
			case log.LevelDebug:
				log.Debugf("%s\n", test.input.msg)
			case log.LevelTrace:
				log.Tracef("%s\n", test.input.msg)
			}

			got := writer.String()
			if got != test.want {
				t.Errorf("%+v != %+v", got, test.want)
			}

			// Reset standard logger
			log.SetLevel(log.LevelInfo)
		})
	}
}

func TestLogln(t *testing.T) {
	tests := []struct {
		description string
		input       struct {
			msg    string
			logLvl log.Level
			atLvl  log.Level
		}
		want string
	}{
		{
			input: struct {
				msg    string
				logLvl log.Level
				atLvl  log.Level
			}{
				msg:    "visible",
				logLvl: log.LevelInfo,
				atLvl:  log.LevelInfo,
			},
			want: "visible\n",
		},
		{
			input: struct {
				msg    string
				logLvl log.Level
				atLvl  log.Level
			}{
				msg:    "visible",
				logLvl: log.LevelWarn,
				atLvl:  log.LevelWarn,
			},
			want: "visible\n",
		},
		{
			input: struct {
				msg    string
				logLvl log.Level
				atLvl  log.Level
			}{
				msg:    "visible",
				logLvl: log.LevelError,
				atLvl:  log.LevelError,
			},
			want: "visible\n",
		},
		{
			input: struct {
				msg    string
				logLvl log.Level
				atLvl  log.Level
			}{
				msg:    "visible",
				logLvl: log.LevelDebug,
				atLvl:  log.LevelDebug,
			},
			want: "visible\n",
		},
		{
			input: struct {
				msg    string
				logLvl log.Level
				atLvl  log.Level
			}{
				msg:    "visible",
				logLvl: log.LevelTrace,
				atLvl:  log.LevelTrace,
			},
			want: "visible\n",
		},
		{
			input: struct {
				msg    string
				logLvl log.Level
				atLvl  log.Level
			}{
				msg:    "visible",
				logLvl: log.LevelWarn,
				atLvl:  log.LevelInfo,
			},
			want: "visible\n",
		},
		{
			input: struct {
				msg    string
				logLvl log.Level
				atLvl  log.Level
			}{
				msg:    "visible",
				logLvl: log.LevelWarn,
				atLvl:  log.LevelError,
			},
			want: "",
		},
		{
			input: struct {
				msg    string
				logLvl log.Level
				atLvl  log.Level
			}{
				msg:    "visible",
				logLvl: log.LevelWarn,
				atLvl:  log.LevelDebug,
			},
			want: "",
		},
		{
			input: struct {
				msg    string
				logLvl log.Level
				atLvl  log.Level
			}{
				msg:    "visible",
				logLvl: log.LevelWarn,
				atLvl:  log.LevelTrace,
			},
			want: "",
		},
		{
			input: struct {
				msg    string
				logLvl log.Level
				atLvl  log.Level
			}{
				msg:    "visible",
				logLvl: log.LevelError,
				atLvl:  log.LevelInfo,
			},
			want: "visible\n",
		},
		{
			input: struct {
				msg    string
				logLvl log.Level
				atLvl  log.Level
			}{
				msg:    "visible",
				logLvl: log.LevelError,
				atLvl:  log.LevelWarn,
			},
			want: "visible\n",
		},
		{
			input: struct {
				msg    string
				logLvl log.Level
				atLvl  log.Level
			}{
				msg:    "visible",
				logLvl: log.LevelError,
				atLvl:  log.LevelDebug,
			},
			want: "",
		},
		{
			input: struct {
				msg    string
				logLvl log.Level
				atLvl  log.Level
			}{
				msg:    "visible",
				logLvl: log.LevelError,
				atLvl:  log.LevelTrace,
			},
			want: "",
		},
		{
			input: struct {
				msg    string
				logLvl log.Level
				atLvl  log.Level
			}{
				msg:    "visible",
				logLvl: log.LevelDebug,
				atLvl:  log.LevelInfo,
			},
			want: "visible\n",
		},
		{
			input: struct {
				msg    string
				logLvl log.Level
				atLvl  log.Level
			}{
				msg:    "visible",
				logLvl: log.LevelDebug,
				atLvl:  log.LevelWarn,
			},
			want: "visible\n",
		},
		{
			input: struct {
				msg    string
				logLvl log.Level
				atLvl  log.Level
			}{
				msg:    "visible",
				logLvl: log.LevelDebug,
				atLvl:  log.LevelError,
			},
			want: "visible\n",
		},
		{
			input: struct {
				msg    string
				logLvl log.Level
				atLvl  log.Level
			}{
				msg:    "visible",
				logLvl: log.LevelDebug,
				atLvl:  log.LevelTrace,
			},
			want: "",
		},
		{
			input: struct {
				msg    string
				logLvl log.Level
				atLvl  log.Level
			}{
				msg:    "visible",
				logLvl: log.LevelTrace,
				atLvl:  log.LevelInfo,
			},
			want: "visible\n",
		},
		{
			input: struct {
				msg    string
				logLvl log.Level
				atLvl  log.Level
			}{
				msg:    "visible",
				logLvl: log.LevelTrace,
				atLvl:  log.LevelWarn,
			},
			want: "visible\n",
		},
		{
			input: struct {
				msg    string
				logLvl log.Level
				atLvl  log.Level
			}{
				msg:    "visible",
				logLvl: log.LevelTrace,
				atLvl:  log.LevelError,
			},
			want: "visible\n",
		},
		{
			input: struct {
				msg    string
				logLvl log.Level
				atLvl  log.Level
			}{
				msg:    "visible",
				logLvl: log.LevelTrace,
				atLvl:  log.LevelDebug,
			},
			want: "visible\n",
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			var writer strings.Builder

			log.SetOutput(&writer)
			log.SetFlags(0)
			log.SetLevel(test.input.logLvl)

			switch test.input.atLvl {
			case log.LevelInfo:
				log.Infoln(test.input.msg)
			case log.LevelWarn:
				log.Warnln(test.input.msg)
			case log.LevelError:
				log.Errorln(test.input.msg)
			case log.LevelDebug:
				log.Debugln(test.input.msg)
			case log.LevelTrace:
				log.Traceln(test.input.msg)
			}

			got := writer.String()
			if got != test.want {
				t.Errorf("%+v != %+v", got, test.want)
			}

			// Reset standard logger
			log.SetLevel(log.LevelInfo)
		})
	}
}
