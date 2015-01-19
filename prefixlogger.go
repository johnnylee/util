package util

import (
	"fmt"
	"log"
	"os"
)

type PrefixLogger struct {
	logger *log.Logger
}

func NewPrefixLogger(prefix string) *PrefixLogger {
	pl := new(PrefixLogger)
	pl.logger = log.New(os.Stderr, "["+prefix+"] ", log.Ldate|log.Ltime)
	return pl
}

// Msg: Just link Printf, but adds an additional newline to the end of the
// message.
func (pl *PrefixLogger) Msg(format string, v ...interface{}) {
	pl.logger.Printf(format+"\n", v...)
}

// Err: Prints the error with the given message. Appends a colon and a newline
// to the message.
func (pl *PrefixLogger) Err(err error, format string, v ...interface{}) {
	msg := fmt.Sprintf(format, v...)
	pl.logger.Printf("ERROR: %v:\n    %v\n", msg, err)
}
