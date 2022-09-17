package tray

import (
	"time"

	"github.com/getlantern/systray"
	"github.com/tiredsosha/warden/tools/logger"
)

var Conn bool = false
var icon bool

func onReady() {

	systray.SetIcon(grayIcon)
	systray.SetTitle("Warden")

	systray.SetTooltip("Warden")
	menuQuit := systray.AddMenuItem("Quit", "Quit the whole app")

	go func() {
		<-menuQuit.ClickedCh
		systray.Quit()
	}()
	go func() {
		for {
			time.Sleep(3 * time.Second)

			if Conn == icon {
				continue
			}
			if Conn {
				systray.SetIcon(blueIcon)
			} else {
				systray.SetIcon(grayIcon)
			}

			icon = Conn
		}
	}()
}

func onExit() {
	logger.Error.Fatal("EXITING MANUALLY")
}

func TrayStart() {
	systray.Run(onReady, onExit)
}
