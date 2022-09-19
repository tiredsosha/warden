package power

import (
	"os/exec"

	"github.com/lxn/win"
	"github.com/tiredsosha/warden/tools/logger"
)

func Reboot() {
	cmd := "shutdown /f /r /t 0"

	exe := exec.Command("cmd", "/C", cmd)
	err := exe.Run()

	if err != nil {
		logger.Warn.Printf("can't reboot pc - %s\n", err)
	}
}

func Shutdown() {
	cmd := "shutdown /f /s /t 0"

	exe := exec.Command("cmd", "/C", cmd)
	err := exe.Run()

	if err != nil {
		logger.Warn.Printf("can't shutdown pc - %s\n", err)
	}
}

func Sleep(state bool) {
	switch state {
	case false:
		win.SendMessage(0xFFFF, 0x0112, 0xF170, 2)
	}
}
