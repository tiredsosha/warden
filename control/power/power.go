package power

import (
	"log"
	"os/exec"
)

func Reboot() {
	var cmd string
	cmd = "shutdown /f /r"

	exe := exec.Command("cmd", "/C", cmd)
	err := exe.Run()

	if err != nil {
		log.Fatal(err)
	}
}

func Shutdown() {
	var cmd string
	cmd = "shutdown /f"

	exe := exec.Command("cmd", "/C", cmd)
	err := exe.Run()

	if err != nil {
		log.Fatal(err)
	}
}
