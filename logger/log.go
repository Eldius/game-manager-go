package logger

import (
	"log"
	"strings"
)

/*
LogWriter wrapper for log to use at os.exec.Cmd
*/
type LogWriter struct {
    logger *log.Logger
}

/*
NewLogWriter creates a new LogWriter
*/
func NewLogWriter(l *log.Logger) *LogWriter {
    lw := &LogWriter{}
    lw.logger = l
    return lw
}

/*
DefaultLogger returns a default logger
*/
func DefaultLogger() *log.Logger {
	return log.New(
		log.Writer(),
		log.Prefix() + "> ",
		log.Flags(),
	)
}

/*
Write write logs to log (doh)
*/
func (lw LogWriter) Write (p []byte) (n int, err error) {
    lw.logger.Println(strings.ReplaceAll(string(p), "\n", "\n> "))
    return len(p), nil
}