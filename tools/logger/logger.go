package logger

import (
	"log"
	"os"
)

var (
	WarningLogger *log.Logger
	InfoLogger    *log.Logger
	ErrorLogger   *log.Logger
)

func LogInit() {
	file, err := os.OpenFile("warden.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.SetOutput(os.Stdout)
		log.SetOutput(os.Stderr)
		log.Println("can't open/make a warden.log. Logging in console")

	} else {
		log.SetOutput(file)
	}
	InfoLogger = log.New(file, "INFO: ", log.Ldate|log.Ltime)
	WarningLogger = log.New(file, "WARNING: ", log.Ldate|log.Ltime)
	ErrorLogger = log.New(file, "ERROR: ", log.Ldate|log.Ltime)
}