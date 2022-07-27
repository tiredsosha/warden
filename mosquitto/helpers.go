package mosquitto

import (
	"strconv"

	"github.com/tiredsosha/warden/control/sound"
)

func VolStatus() (string, error) {
	volume, err := sound.GetVolume()
	status := strconv.Itoa(int(volume))
	return status, err
}

func PcStatus() (string, error) {
	var err error
	status := "true"
	return status, err
}
