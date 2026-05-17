package functions

import (
	"fmt"
	"log"
	"strings"
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
	var ws = &winsize{}

	r, _, errno := syscall.Syscall(syscall.SYS_IOCTL,
		uintptr(syscall.Stdout),
		uintptr(syscall.TIOCGWINSZ),
		uintptr(unsafe.Pointer(ws)))
	if errno != 0 {
		log.Fatal(errno)
	}
	if int(r) == -1 {
		return 80
	}
	return int(ws.Col)
}

func Padding(flag map[string]string, tWidth, ascii_len int) int {
	//	var alignv = []string{"right", "left", "center"}
	align := flag["align"]
	fmt.Println(align)
	spacetoAdd := tWidth - ascii_len
	if align != "" {
		if align == "center" {
			spacetoAdd = spacetoAdd / 2
		}
	}
	return spacetoAdd
}

func Addpadding(padd int) string {
	if padd < 1 {
		log.Fatal("Error... with window size.. try resize your window")
	}
	space := strings.Repeat(" ", padd)
	return space
}
