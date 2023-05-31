package main

import (
	"flag"
	"io"
	"os"

	"git.sr.ht/~spc/go-log"
)

func main() {
	level := flag.String("level", "error", "set log level to `STRING`")
	stdout := flag.Bool("stdout", true, "log to stdout")
	file := flag.String("file", "", "log to `FILE`")
	flag.Parse()

	logLevel, err := log.ParseLevel(*level)
	if err != nil {
		log.Fatalf("cannot parse level: %v", err)
	}
	log.SetLevel(logLevel)

	var data []byte
	var r io.Reader
	if flag.Arg(0) == "-" {
		r = os.Stdin
	} else {
		r, err = os.Open(flag.Arg(0))
	}
	if err != nil {
		log.Fatalf("cannot open file for reading: %v", err)
	}
	data, err = io.ReadAll(r)
	if err != nil {
		log.Fatalf("cannot read data: %v", err)
	}

	writers := []io.Writer{}

	if *stdout {
		writers = append(writers, os.Stdout)
	}

	if *file != "" {
		f, err := os.Create(*file)
		if err != nil {
			log.Fatalf("cannot open file for writing: %v", err)
		}
		writers = append(writers, f)
	}

	mw := io.MultiWriter(writers...)
	log.SetOutput(mw)

	switch log.CurrentLevel() {
	case log.LevelError:
		log.Error(string(data))
	case log.LevelWarn:
		log.Warn(string(data))
	case log.LevelInfo:
		log.Info(string(data))
	case log.LevelDebug:
		log.Debug(string(data))
	case log.LevelTrace:
		log.Trace(string(data))
	default:
		log.Print(string(data))
	}
}
