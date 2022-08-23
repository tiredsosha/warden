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

	Debug.Println("")
	Debug.Println("")
	Info.Print("WARDENER STARTED")
}

func DebugLog(debug, state bool, hostname, broker, username, password string) {
	Debug.Println("---------------------------")
	Debug.Println("logging data:")
	Debug.Printf("\tdebug    - %t\n", debug)
	Debug.Printf("\tcli conf - %t\n", state)
	Debug.Println("- - - - - - - - - - - - - -")
	Debug.Println("сonnection data:")
	Debug.Printf("\thostname - %s\n", hostname)
	Debug.Printf("\tbroker   - %s\n", broker)
	Debug.Printf("\tusername - %s\n", username)
	Debug.Printf("\tpassword - %s\n", password)
	Debug.Println("---------------------------")
}
