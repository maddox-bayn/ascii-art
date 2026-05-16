package functions

import (
	"strings"
)

func PrintArt(asciiTable [][]string, inputArg []string, flag map[string]string, subStr string) string {

	var art strings.Builder
	isOnlyNewline := true
	color := flag["color"]
	align := flag["align"]

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
		needColor := GetColorMusk(word, subStr, color)
		rendLines := RenderLine(word, asciiTable)
		terminalWidth := GetTerminalWidth()

		AsciiWidth := len(rendLines[0])

		padd := Padding(flag, terminalWidth, AsciiWidth)

		for lineChar := 0; lineChar < 8; lineChar++ {
			for i, char := range word {
				//	Stage Change: should turn color on
				if needColor[i] && (i == 0 || !needColor[i-1]) {
					art.WriteString(ColorMap[color])
				}

				// Stage Change: should turn off?

				if !needColor[i] && (i > 0 && needColor[i-1]) {
					art.WriteString(Reset)
				}

				if align != "" {
					space := Addpadding(padd)
					CheckAlignment(align)
					if i == 0 && align == "right" {
						art.WriteString(space + asciiTable[char-32][lineChar])
					} else if align == "center" && i == 0 {
						art.WriteString(space + asciiTable[char-32][lineChar])
					} else if align == "left" {
						art.WriteString(asciiTable[char-32][lineChar])
					} else {
						art.WriteString(asciiTable[char-32][lineChar])
					}
					continue
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
