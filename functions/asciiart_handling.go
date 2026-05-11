package functions

import (
	"fmt"
	"strings"
)

func PrintArt(asciiTable [][]string, inputArg []string, color string, subStr string) string {

	var art strings.Builder
	isOnlyNewline := true

	for _, word := range inputArg {
		if word != "" {
			isOnlyNewline = false
			break
		}
	}

	if isOnlyNewline {
		for i := 0; i < len(inputArg)-1; i++ {
			art.WriteString("\n")

		}
		return art.String()
	}

	for _, word := range inputArg {
		if word == "" {
			fmt.Println()
			continue
		}
		needColor := GetColorMusk(word, subStr)

		for lineChar := 0; lineChar < 8; lineChar++ {
			for i, char := range word {
				//	needColor := i >= stIdx || i <= stIdx+len(subStr)-1

				if color != "" && needColor[i] {
					art.WriteString(ColorMap[color] + asciiTable[char-32][lineChar] + "\033[0m")
					//
				} else {
					art.WriteString(asciiTable[char-32][lineChar])
				}
			}
			art.WriteString("\n")
		}
	}
	return art.String()
}
