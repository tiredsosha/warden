package sound

import (
	"errors"

	"github.com/go-ole/go-ole"
	"github.com/moutend/go-wca/pkg/wca"
	"github.com/tiredsosha/warden/tools/logger"
)

func GetVolume() (uint8, error) {
	vol, err := invoke(func(aev *wca.IAudioEndpointVolume) (interface{}, error) {
		var level float32
		err := aev.GetMasterVolumeLevelScalar(&level)
		volume := uint8(level*100.0 + 0.5)
		return volume, err
	})

	if err != nil || vol == nil {
		logger.Warn.Printf("can't get volume level - %s\n", err)
		logger.Debug.Printf("volume:%s\n", vol)
		err = errors.New("WINAPI")
		return 0, err
	}

	return vol.(uint8), err
}

func SetVolume(volume int) {
	_, err := invoke(func(aev *wca.IAudioEndpointVolume) (interface{}, error) {
		err := aev.SetMasterVolumeLevelScalar(float32(volume)/100.0, nil)
		return nil, err
	})

	if err != nil {
		logger.Warn.Printf("can't set volume level - %s\n", err)
	}
}

func GetMute() (bool, error) {
	mute, err := invoke(func(aev *wca.IAudioEndpointVolume) (interface{}, error) {
		var muted bool
		err := aev.GetMute(&muted)
		return muted, err
	})

	if err != nil {
		logger.Warn.Printf("can't get mute status - %s\n", err)
	}

	if mute == nil {
		return false, err
	}

	return mute.(bool), err
}

func Mute(state bool) {
	_, err := invoke(func(aev *wca.IAudioEndpointVolume) (interface{}, error) {
		err := aev.SetMute(state, nil)
		return nil, err
	})

	if err != nil {
		logger.Warn.Printf("can't set mute - %s\n", err)
	}
}

// CONNECTION TO WINDOWS API //
func invoke(f func(aev *wca.IAudioEndpointVolume) (interface{}, error)) (ret interface{}, err error) {
	if err = ole.CoInitializeEx(0, ole.COINIT_APARTMENTTHREADED); err != nil {
		return
	}
	defer ole.CoUninitialize()

	var mmde *wca.IMMDeviceEnumerator
	if err = wca.CoCreateInstance(wca.CLSID_MMDeviceEnumerator, 0, wca.CLSCTX_ALL, wca.IID_IMMDeviceEnumerator, &mmde); err != nil {
		return
	}
	defer mmde.Release()

	var mmd *wca.IMMDevice
	if err = mmde.GetDefaultAudioEndpoint(wca.ERender, wca.EConsole, &mmd); err != nil {
		return
	}
	defer mmd.Release()

	var ps *wca.IPropertyStore
	if err = mmd.OpenPropertyStore(wca.STGM_READ, &ps); err != nil {
		return
	}
	defer ps.Release()

	var pv wca.PROPVARIANT
	if err = ps.GetValue(&wca.PKEY_Device_FriendlyName, &pv); err != nil {
		return
	}

	var aev *wca.IAudioEndpointVolume
	if err = mmd.Activate(wca.IID_IAudioEndpointVolume, wca.CLSCTX_ALL, nil, &aev); err != nil {
		return
	}
	defer aev.Release()

	ret, err = f(aev)
	return
}
