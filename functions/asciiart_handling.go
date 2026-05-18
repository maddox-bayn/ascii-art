package functions

import (
	"strings"
)

func hancleSpace(SliceInput []string) []string {
	for i, input := range SliceInput {
		sortinput := strings.Fields(input)

		SliceInput[i] = strings.Join(sortinput, " ")
	}
	return SliceInput
}
func ProcessInput(asciiTable [][]string, inputArg []string, flag map[string]string, subStr string) {

	var art strings.Builder
	isOnlyNewline := true
	color := flag["color"]
	align := flag["align"]
	fileName := flag["reverse"]
	if fileName != "" {
		ToReverse(fileName, asciiTable)
		return
	}
	terminalWidth := 0

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
		SaveOrPrintart(art.String(), flag)
	}

	if align == "justify" {
		inputArg = hancleSpace(inputArg)
	}

	for _, word := range inputArg {
		if word == "" {
			art.WriteString("\n")
			continue
		}
		needColor := GetColorMusk(word, subStr, color)
		rendLines := RenderLine(word, asciiTable)
		if align != "" {
			terminalWidth = GetTerminalWidth()
		}

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
					} else if align == "justify" && char == 32 {
						art.WriteString("{space}")
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
	if align == "justify" {
		artSlice := strings.Split(art.String(), "\n")
		for i, line := range artSlice {
			linelen := len(strings.ReplaceAll(line, "{space}", ""))
			spaceCoutbetweenWord := strings.Count(line, "{space}")

			padd := GetTerminalWidth() - linelen
			if spaceCoutbetweenWord > 0 {
				padd = padd / spaceCoutbetweenWord
			}
			space := Addpadding(padd)
			artSlice[i] = strings.ReplaceAll(line, "{space}", space)
		}
	}
	SaveOrPrintart(art.String(), flag)
}
