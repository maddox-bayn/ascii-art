package functions

import "os"

import "fmt"

func PrintArt(asciiTable [][]string, inputArg []string, color string, subStr string) {

	//colorMap := map[string]string{"red": "\033[31m", "green": "\033[32m", "yellow": "\033[33m", "black": "\033[30m", "blue": "\033[34m", "white": "\033[37m", "magenta": "\033[35m"}

	isOnlyNewline := true

	for _, word := range inputArg {
		if word != "" {
			isOnlyNewline = false
			break
		}
	}

	if isOnlyNewline {
		for i := 0; i < len(inputArg)-1; i++ {
			fmt.Fprintln(os.Stdout)

		}
		return
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

				if needColor[i] {
					fmt.Print(colorMap[color] + asciiTable[char-32][lineChar] + "\033[0m")
					//
				} else {
					fmt.Print(asciiTable[char-32][lineChar])
				}
			}
			fmt.Println()
		}
	}
}
