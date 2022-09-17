package logger

import (
	"io"
	"log"
	"os"
	"syscall"
	"time"
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
		deleteLog := logCreation()
		if deleteLog {
			os.Remove("warden.log")
		}
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

func DebugLog(version string, debug, state bool, hostname, broker, username, password string) {
	Debug.Println("---------------------------")
	Debug.Println("common data:")
	Debug.Printf("\t\tversion  - %s\n", version)
	Debug.Println("- - - - - - - - - - - - - -")
	Debug.Println("logging data:")
	Debug.Printf("\t\tdebug    - %t\n", debug)
	Debug.Printf("\t\tcli conf - %t\n", state)
	Debug.Println("- - - - - - - - - - - - - -")
	Debug.Println("сonnection data:")
	Debug.Printf("\t\thostname - %s\n", hostname)
	Debug.Printf("\t\tbroker   - %s\n", broker)
	Debug.Printf("\t\tusername - %s\n", username)
	Debug.Printf("\t\tpassword - %s\n", password)
	Debug.Println("---------------------------")
}

func logCreation() bool {
	var deleteLog bool = false

	log, _ := os.Stat("warden.log")
	createTime := time.Unix(0, log.Sys().(*syscall.Win32FileAttributeData).CreationTime.Nanoseconds())
	currTime := time.Now()
	diff := currTime.Sub(createTime).Milliseconds()

	if diff > 604800000 {
		deleteLog = true
	}

	return deleteLog
}
