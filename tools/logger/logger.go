package logger

import (
	"io"
	"log"
	"os"
)

var (
	Warn  *log.Logger
	Debug *log.Logger
	Info  *log.Logger
	Error *log.Logger
)

func LogInit(debug bool) {
	var out interface{}
	out = io.Discard

	if debug {
		file, err := os.OpenFile("warden.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		if err == nil {
			out = file
		}
	}

	Debug = log.New(out.(io.Writer), "DEBUG: ", log.Ldate|log.Ltime)
	Info = log.New(out.(io.Writer), "INFO:  ", log.Ldate|log.Ltime)
	Warn = log.New(out.(io.Writer), "WARN:  ", log.Ldate|log.Ltime)
	Error = log.New(out.(io.Writer), "ERROR: ", log.Ldate|log.Ltime)
}
