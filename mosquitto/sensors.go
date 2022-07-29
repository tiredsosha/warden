package mosquitto

import (
	"strconv"

	"github.com/tiredsosha/warden/control/sound"
)

func VolumeStatus() (string, error) {
	volume, err := sound.GetVolume()
	status := strconv.Itoa(int(volume))
	return status, err
}

func ConnectStatus() (string, error) {
	var err error
	status := "true"
	return status, err
}
