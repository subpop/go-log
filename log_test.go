package log_test

import (
	"strings"
	"testing"

	log "github.com/subpop/go-log"
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
				logLvl: log.LevelWarn,
				atLvl:  log.LevelInfo,
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
				logLvl: log.LevelInfo,
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
				logLvl: log.LevelInfo,
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
				logLvl: log.LevelInfo,
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
				logLvl: log.LevelInfo,
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
			case log.LevelError:
				log.Error(test.input.msg)
			case log.LevelWarn:
				log.Warn(test.input.msg)
			case log.LevelInfo:
				log.Info(test.input.msg)
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
			log.SetLevel(log.LevelError)
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
				logLvl: log.LevelWarn,
				atLvl:  log.LevelInfo,
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
				logLvl: log.LevelInfo,
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
				logLvl: log.LevelInfo,
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
				logLvl: log.LevelInfo,
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
				logLvl: log.LevelInfo,
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
			case log.LevelError:
				log.Errorf("%s\n", test.input.msg)
			case log.LevelWarn:
				log.Warnf("%s\n", test.input.msg)
			case log.LevelInfo:
				log.Infof("%s\n", test.input.msg)
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
			log.SetLevel(log.LevelError)
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
				logLvl: log.LevelWarn,
				atLvl:  log.LevelInfo,
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
				logLvl: log.LevelInfo,
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
				logLvl: log.LevelInfo,
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
				logLvl: log.LevelInfo,
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
				logLvl: log.LevelInfo,
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
			case log.LevelError:
				log.Errorln(test.input.msg)
			case log.LevelWarn:
				log.Warnln(test.input.msg)
			case log.LevelInfo:
				log.Infoln(test.input.msg)
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
			log.SetLevel(log.LevelError)
		})
	}
}

func TestParseLevel(t *testing.T) {
	tests := []struct {
		description string
		input       string
		want        log.Level
		wantError   error
	}{
		{"info", "ERROR", log.LevelError, nil},
		{"warn", "WARN", log.LevelWarn, nil},
		{"error", "INFO", log.LevelInfo, nil},
		{"debug", "DEBUG", log.LevelDebug, nil},
		{"trace", "TRACE", log.LevelTrace, nil},
		{"unknown", "VERBOSE", log.LevelUnknown, log.ParseError("VERBOSE")},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			got, err := log.ParseLevel(test.input)

			if test.wantError != nil {
				if err != test.wantError {
					t.Errorf("%v != %v", err, test.wantError)
				}
			} else {
				if err != nil {
					t.Fatal(err)
				}
				if got != test.want {
					t.Errorf("%v != %v", got, test.want)
				}
			}
		})
	}
}

func TestFormatLevel(t *testing.T) {
	tests := []struct {
		description string
		input       log.Level
		want        string
	}{
		{"info", log.LevelError, "ERROR"},
		{"warn", log.LevelWarn, "WARN"},
		{"error", log.LevelInfo, "INFO"},
		{"debug", log.LevelDebug, "DEBUG"},
		{"trace", log.LevelTrace, "TRACE"},
		{"unknown", log.LevelUnknown, ""},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			got := log.FormatLevel(test.input)

			if got != test.want {
				t.Errorf("%+v != %+v", got, test.want)
			}
		})
	}
}

func TestPrint(t *testing.T) {
	tests := []struct {
		description string
		input       string
		want        string
	}{
		{
			input: "text",
			want:  "text\n",
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			var writer strings.Builder
			log.SetOutput(&writer)
			log.SetFlags(0)

			log.Print(test.input)

			got := writer.String()
			if got != test.want {
				t.Errorf("%#v != %#v", got, test.want)
			}
		})
	}
}

func TestPrintf(t *testing.T) {
	tests := []struct {
		description string
		input       string
		want        string
	}{
		{
			input: "text",
			want:  "text\n",
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			var writer strings.Builder
			log.SetOutput(&writer)
			log.SetFlags(0)

			log.Printf("%v", test.input)

			got := writer.String()
			if got != test.want {
				t.Errorf("%#v != %#v", got, test.want)
			}
		})
	}
}

func TestPrintln(t *testing.T) {
	tests := []struct {
		description string
		input       string
		want        string
	}{
		{
			input: "text",
			want:  "text\n",
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			var writer strings.Builder
			log.SetOutput(&writer)
			log.SetFlags(0)

			log.Println(test.input)

			got := writer.String()
			if got != test.want {
				t.Errorf("%#v != %#v", got, test.want)
			}
		})
	}
}

func TestPanic(t *testing.T) {
	tests := []struct {
		description string
		input       string
		want        string
	}{
		{
			input: "text",
			want:  "text",
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					switch got := r.(type) {
					case string:
						if got != test.want {
							t.Errorf("%#v != %#v", got, test.want)
						}
					default:
						t.Errorf("invalid type: %T", got)
					}
				}
			}()

			var writer strings.Builder
			log.SetOutput(&writer)
			log.SetFlags(0)

			log.Panic(test.input)

			got := writer.String()
			if got != test.want {
				t.Errorf("%#v != %#v", got, test.want)
			}
		})
	}
}

func TestPanicf(t *testing.T) {
	tests := []struct {
		description string
		input       string
		want        string
	}{
		{
			input: "text",
			want:  "text",
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					switch got := r.(type) {
					case string:
						if got != test.want {
							t.Errorf("%#v != %#v", got, test.want)
						}
					default:
						t.Errorf("invalid type: %T", got)
					}
				}
			}()

			var writer strings.Builder
			log.SetOutput(&writer)
			log.SetFlags(0)

			log.Panicf("%v", test.input)

			got := writer.String()
			if got != test.want {
				t.Errorf("%#v != %#v", got, test.want)
			}
		})
	}
}

func TestPanicln(t *testing.T) {
	tests := []struct {
		description string
		input       string
		want        string
	}{
		{
			input: "text",
			want:  "text\n",
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					switch got := r.(type) {
					case string:
						if got != test.want {
							t.Errorf("%#v != %#v", got, test.want)
						}
					default:
						t.Errorf("invalid type: %T", got)
					}
				}
			}()

			var writer strings.Builder
			log.SetOutput(&writer)
			log.SetFlags(0)

			log.Panicln(test.input)

			got := writer.String()
			if got != test.want {
				t.Errorf("%#v != %#v", got, test.want)
			}
		})
	}
}

func TestOutput(t *testing.T) {
	tests := []struct {
		description string
		input       string
		want        string
	}{
		{
			input: "text",
			want:  "text\n",
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			var writer strings.Builder
			log.SetOutput(&writer)
			log.SetFlags(0)

			_ = log.Output(0, test.input)

			got := writer.String()
			if got != test.want {
				t.Errorf("%#v != %#v", got, test.want)
			}
		})
	}
}

func TestFlags(t *testing.T) {
	tests := []struct {
		description string
		input       int
		want        int
	}{
		{input: log.Ldate | log.Ltime, want: 3},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			log.SetFlags(test.input)
			got := log.Flags()

			if got != test.want {
				t.Errorf("%#v != %#v", got, test.want)
			}
		})
	}
}

func TestPrefix(t *testing.T) {
	tests := []struct {
		description string
		input       string
		want        string
	}{
		{input: "prefix", want: "prefix"},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			log.SetPrefix(test.input)

			got := log.Prefix()

			// reset standard logger
			log.SetPrefix("")

			if got != test.want {
				t.Errorf("%#v != %#v", got, test.want)
			}
		})
	}
}

func TestCurrentLevel(t *testing.T) {
	tests := []struct {
		description string
		input       log.Level
		want        log.Level
	}{
		{input: log.LevelError, want: log.LevelError},
		{input: log.LevelWarn, want: log.LevelWarn},
		{input: log.LevelInfo, want: log.LevelInfo},
		{input: log.LevelDebug, want: log.LevelDebug},
		{input: log.LevelTrace, want: log.LevelTrace},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			log.SetLevel(test.input)

			got := log.CurrentLevel()

			// reset standard logger
			log.SetLevel(log.LevelError)

			if got != test.want {
				t.Errorf("%#v != %#v", got, test.want)
			}
		})
	}
}
