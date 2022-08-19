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
		logger.Warn.Printf("can't reboot pc - %s\n", err)
	}
}

func Shutdown() {
	cmd := "shutdown /f"

	exe := exec.Command("cmd", "/C", cmd)
	err := exe.Run()

	if err != nil {
		logger.Warn.Printf("can't shutdown pc - %s\n", err)
	}
}
