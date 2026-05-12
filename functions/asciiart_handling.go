package functions

import (
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
			art.WriteString("\n")
			continue
		}
		needColor := GetColorMusk(word, subStr)

		for lineChar := 0; lineChar < 8; lineChar++ {
			for i, char := range word {
				//	Stage Change: should turn color on
				if needColor[i] && (i == 0 || !needColor[i-1]) {
					art.WriteString(ColorMap[color])
				}

				// Stage Change: should turn off?
				if !needColor[i] && (i > 0 || needColor[i-1]) {
					art.WriteString(Reset)
				}

				//Action: Always draw the character art
				art.WriteString(asciiTable[char-32][lineChar])
			}

			// Safety: if the last character of the word was colored, rest it at the end of the line
			if needColor[len(word)-1] {
				art.WriteString(Reset)
			}
			art.WriteString("\n")
		}

	}
	return art.String()
}
