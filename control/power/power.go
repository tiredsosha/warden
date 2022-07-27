package power

import (
	"os/exec"

	"github.com/tiredsosha/warden/tools/logger"
)

func Reboot() {
	cmd := "shutdown /f /r"

	exe := exec.Command("cmd", "/C", cmd)
	err := exe.Run()

	if err != nil {
		logger.Warn.Println("can't reboot pc, error in Windows API")
		logger.Warn.Println(err)
	}
}

func Shutdown() {
	cmd := "shutdown /f"

	exe := exec.Command("cmd", "/C", cmd)
	err := exe.Run()

	if err != nil {
		logger.Warn.Println("can't shutdown pc, error in Windows API")
		logger.Warn.Println(err)
	}
}
