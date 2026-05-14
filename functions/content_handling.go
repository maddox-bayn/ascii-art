package functions

import (
	"fmt"
	"strings"
)

func RenderLine(text string, banner [][]string) []string {
	var result []string
	var b strings.Builder

	for i := 0; i < 8; i++ {
		for _, c := range text {
			line := banner[c-32]
			b.WriteString(line[i])
		}
		result = append(result, b.String())
		b.Reset()
	}
	return result
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
	space := strings.Repeat(" ", padd)
	return space
}
