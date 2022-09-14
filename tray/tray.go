package tray

import (
	"github.com/getlantern/systray"
	"github.com/tiredsosha/warden/tools/logger"
)

func onReady() {

	systray.SetIcon(iconBytes)
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

func TrayStart() {
	systray.Run(onReady, onExit)
}
