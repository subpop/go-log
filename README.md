# package log

Package log implements a simple level logging package that maintains API
compatibility with the standard library `log` package. It extends the standard
library `log.Logger` type with a `Level` type that can be used to define output
verbosity. It adds additional methods that will be printed only if the logger
is configured at or below a given level. For example:

```
package main

import "git.sr.ht/~spc/go-log"

func main() {
    log.SetLevel(log.LevelWarn)
    log.Traceln("Messages at LevelTrace will not print if the Level is Warn")
    log.Warnln("Messages at LevelTrace will print with the Level set to Warn")
}
```
