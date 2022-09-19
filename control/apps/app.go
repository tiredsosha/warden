package app

import (
	"os/exec"

	"github.com/tiredsosha/warden/tools/logger"
)

func Quit(app string) {

	cmd := "taskkill /im " + app + " /t /f"

	exe := exec.Command("cmd", "/C", cmd)
	err := exe.Run()

	if err != nil {
		logger.Warn.Printf("can't close the app - %s\n", err)
	}
}
