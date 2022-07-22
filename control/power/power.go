package power

import (
	"log"
	"os/exec"
)

func Reboot() {
	var cmd string
	cmd = "shutdown /f /t 5 /r"

	exe := exec.Command(cmd)
	err := exe.Run()

	if err != nil {
		log.Fatal(err)
	}
}

func Shutdown() {
	var cmd string
	cmd = "shutdown /f /t 5"

	exe := exec.Command(cmd)
	err := exe.Run()

	if err != nil {
		log.Fatal(err)
	}

}
