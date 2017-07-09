package middleware

import (
	"github.com/urfave/negroni"
	"log"
	"os"
)

// NewLogger custom logger from negroni logger
func NewLogger() *negroni.Logger {
	logger := &negroni.Logger{ALogger: log.New(os.Stdout, "[mini] ", 0)}
	logger.SetDateFormat(negroni.LoggerDefaultDateFormat)
	logger.SetFormat(negroni.LoggerDefaultFormat)
	return logger
}
