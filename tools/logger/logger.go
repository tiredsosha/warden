package logger

import (
	"log"
	"os"
)

var (
	Warn  *log.Logger
	Info  *log.Logger
	Error *log.Logger
)

func init() {
	file, err := os.OpenFile("warden.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.SetOutput(os.Stdout)
		log.SetOutput(os.Stderr)
		log.Println("can't open/make a warden.log. Logging in console")

	} else {
		log.SetOutput(file)
	}
	Info = log.New(file, "INFO: ", log.Ldate|log.Ltime)
	Warn = log.New(file, "WARNING: ", log.Ldate|log.Ltime)
	Error = log.New(file, "ERROR: ", log.Ldate|log.Ltime)
}
