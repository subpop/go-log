[![godocs.io](http://godocs.io/github.com/subpop/go-log?status.svg)](http://godocs.io/github.com/subpop/go-log)
[![Go Report Card](https://goreportcard.com/badge/github.com/subpop/go-log)](https://goreportcard.com/report/github.com/subpop/go-log)

# ⚠️ Deprecated: use `log/slog` instead ⚠️

# package log

Package log implements a simple level logging package that maintains API
compatibility with the standard library `log` package. It extends the standard
library `log.Logger` type with a `Level` type that can be used to define output
verbosity. It adds additional methods that will be printed only if the logger
is configured at or below a given level. For example:

```
package main

import "github.com/subpop/go-log"

func main() {
    log.SetLevel(log.LevelWarn)
    log.Traceln("Messages at LevelTrace will not print if the Level is Warn")
    log.Warnln("Messages at LevelTrace will print with the Level set to Warn")
}
```
