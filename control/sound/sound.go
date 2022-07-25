package sound

import (
	"errors"
	"log"

	"github.com/go-ole/go-ole"
	"github.com/moutend/go-wca/pkg/wca"
)

func GetVolume() (uint8, error) {
	vol, err := invoke(func(aev *wca.IAudioEndpointVolume) (interface{}, error) {
		var level float32
		err := aev.GetMasterVolumeLevelScalar(&level)
		volume := uint8(level*100.0 + 0.5)
		return volume, err
	})

	if err != nil || vol == nil {
		log.Println("Can't get volume level - Windows API doesn't responce")
		err = errors.New("WIN API")
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
		log.Println("Can't set volume level - Windows API doesn't responce")
	}
}

func Mute(state bool) {
	_, err := invoke(func(aev *wca.IAudioEndpointVolume) (interface{}, error) {
		err := aev.SetMute(state, nil)
		return nil, err
	})

	if err != nil {
		log.Println("Can't set mute - Windows API doesn't responce")
	}
}

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
