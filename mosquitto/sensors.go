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

func MuteStatus() (string, error) {
	mute, err := sound.GetMute()
	status := strconv.FormatBool(mute)
	return status, err
}

/*
default sensor func shoud look like it:

func ...Status() (string, error) {
	status, err := your function
	return status, err
}
*/
