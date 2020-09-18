package log_test

import (
	"os"

	"git.sr.ht/~spc/go-log"
)

func ExampleLogger() {
	// Set the output to Stdout so it is captured by the example test harness.
	// Remember the standard library logger output is Stderr by default.
	log.SetOutput(os.Stdout)

	// Clear the prefix flags so we aren't dancing around timestamps.
	log.SetFlags(0)

	// Log!
	log.Info("Info messages are always logged")
	log.Debug("Debug messages are sometimes logged... but not this one...")

	// Raise the log level
	log.SetLevel(log.LevelDebug)

	// Log again!
	log.Debug("Now that the log level is set to Debug, this message is logged!")
	log.Trace("But we still have one more level that is not logged at the Debug level.")

	// Output:
	// Info messages are always logged
	// Now that the log level is set to Debug, this message is logged!
}
