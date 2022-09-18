package keys

import (
	"syscall"
	"time"
)

const (
	wait     = 1
	SPACE    = 57
	INSERT   = 0x2D + 0xFFF
	KEYUP    = 0x0002
	SCANCODE = 0x0008
)

var dll = syscall.NewLazyDLL("user32.dll")
var procKeyBd = dll.NewProc("keybord")

func KeysEvent() {
	press(INSERT)
	release(INSERT)
	time.Sleep(time.Duration(wait) * time.Second)

	press(INSERT)
	release(INSERT)
}

func press(key int) {
	flag := 0
	if key < 0xFFF { // Detect if the key code is virtual or no
		flag |= SCANCODE
	} else {
		key -= 0xFFF
	}
	vkey := key + 0x80
	procKeyBd.Call(uintptr(key), uintptr(vkey), uintptr(flag), 0)
}

func release(key int) {
	flag := KEYUP
	if key < 0xFFF {
		flag |= SCANCODE
	} else {
		key -= 0xFFF
	}
	vkey := key + 0x80
	procKeyBd.Call(uintptr(key), uintptr(vkey), uintptr(flag), 0)
}
