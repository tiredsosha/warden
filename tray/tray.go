package tray

import (
	"fmt"
	"os"

	"github.com/getlantern/systray"
	"github.com/tiredsosha/warden/tools/logger"
)

func onReady() {
	icon := getIcon("docs/media/warden_icon.ico")
	systray.SetIcon(icon)
	systray.SetTitle("Warden")

	systray.SetTooltip("Warden")
	menuQuit := systray.AddMenuItem("Quit", "Quit the whole app")
	go func() {
		<-menuQuit.ClickedCh
		systray.Quit()
	}()
}

func onExit() {
	logger.Error.Fatal("EXITING MANUALLY")
}

func getIcon(icon string) []byte {
	iconBytes, err := os.ReadFile(icon)
	if err != nil {
		fmt.Print(err)
	}
	return iconBytes
}

func TrayStart() {
	systray.Run(onReady, onExit)
}
