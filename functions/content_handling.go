package functions

import (
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
