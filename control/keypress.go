package control

import "syscall"

const (
	key  = 58
	vkey = key + 0x80
)

var dll = syscall.NewLazyDLL("User32.dll")
var procKey = dll.NewProc("keypress")

func KeyPress() {
	flag := 0
	procKey.Call(uintptr(key), uintptr(vkey), uintptr(flag), 0)
	flag = 0x0002
	procKey.Call(uintptr(key), uintptr(vkey), uintptr(flag), 0)
}
