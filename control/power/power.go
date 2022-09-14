package power

import (
	"os/exec"

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

// not working for now
// func Display(state bool) {
// 	switch state {
// 	case true:
// 		time.Sleep(time.Duration(10) * time.Millisecond)
// 		robotgo.Move(10, 20)
// 		time.Sleep(time.Duration(5000) * time.Millisecond)
// 		robotgo.KeyTap("space")
// 	case false:
// 		win.SendMessage(0xFFFF, 0x0112, 0xF170, 2)
// 	}

// }
