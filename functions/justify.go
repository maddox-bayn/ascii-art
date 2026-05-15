package functions

import (
	"fmt"
	"syscall"
	"unsafe"
)

type winsize struct {
	Row    uint16
	Col    uint16
	Xpixel uint16
	Ypixel uint16
}

func GetTerminalWidth() int {
	ws := &winsize{}
	// TIOCGWINSZ is the magic number (ioctl request) to get window size
	retCode, _, _ := syscall.Syscall(syscall.SYS_IOCTL,
		uintptr(syscall.Stdout),
		uintptr(syscall.TIOCGWINSZ),
		uintptr(unsafe.Pointer(ws)))

	if int(retCode) == -1 {
		return 80 // Default fallback if we can't detect it
	}
	 fmt.Println(int(ws.Col))
	return int(ws.Col)
}
