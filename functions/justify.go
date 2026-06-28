package functions

import (
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)
// get the current termianl window size
func GetTerminalWidth() int {
	// running a stty command
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin

	output, err := cmd.Output()
	if err != nil {
		return -1
	}

	width, err := strconv.Atoi(strings.Fields(string(output))[1])
	if err != nil {
		return -1
	}
	return width
}
// helper function to get the the number of space to add
func Padding(flag map[string]string, tWidth, ascii_len int) int {
	//	var alignv = []string{"right", "left", "center"}
	align := flag["align"]
	spacetoAdd := tWidth - ascii_len
	if align != "" {
		if align == "center" {
			spacetoAdd = spacetoAdd / 2
		}
	}
	return spacetoAdd
}
// helper function to add space
func Addpadding(padd int) string {
	if padd < 1 {
		log.Fatal("Error... with window size.. try resize your window")
	}
	space := strings.Repeat(" ", padd)
	return space
}
