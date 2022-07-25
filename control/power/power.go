package power

import (
	"log"
	"os/exec"
)

func Reboot() {
	cmd := "shutdown /f /r"

	exe := exec.Command("cmd", "/C", cmd)
	err := exe.Run()

	if err != nil {
		log.Println("Can't reboot pc, error in Windows API")
	}
}

func Shutdown() {
	cmd := "shutdown /f"

	exe := exec.Command("cmd", "/C", cmd)
	err := exe.Run()

	if err != nil {
		log.Println("Can't shutdown pc, error in Windows API")
	}
}
