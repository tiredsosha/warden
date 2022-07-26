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
		log.Println("can't reboot pc, error in Windows API")
		log.Println(err)
	}
}

func Shutdown() {
	cmd := "shutdown /f"

	exe := exec.Command("cmd", "/C", cmd)
	err := exe.Run()

	if err != nil {
		log.Println("can't shutdown pc, error in Windows API")
		log.Println(err)
	}
}
