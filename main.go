package main

import (
	"ascii-art/functions"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

func PrintArt(asciiTable [][]string, inputArg []string, color string, subStr string) {

	colorMap := map[string]string{"red": "\033[31m", "green": "\033[32m", "yellow": "\033[33m", "black": "\033[30m", "blue": "\033[34m", "white": "\033[37m", "magenta": "\033[35m"}

	isOnlyNewline := true

	for _, word := range inputArg {
		if word != "" {
			isOnlyNewline = false
			break
		}
	}

	if isOnlyNewline {
		for i := 0; i < len(inputArg)-1; i++ {
			fmt.Println()

		}
		return
	}

	for _, word := range inputArg {
		if word == "" {
			fmt.Println()
			continue
		}

		needColor := make([]bool, len(word))

		// Instead of just one strings.Index:
		if subStr != "" {
			for i := 0; i <= len(word)-len(subStr); i++ {
				if word[i:i+len(subStr)] == subStr {
					for j := 0; j < len(subStr); j++ {
						needColor[i+j] = true
					}
					// Optional: i += len(subStr) - 1 to skip overlapping matches
				}
			}
		}
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

func verifyInput(str string) error {
	for _, char := range str {
		if char < 32 || char > 126 {
			//fmt.Println("Error: Wrong input Try: PRINTABLE ASCII CHARACTER")
			return errors.New("Error: Wrong input Try: PRINTABLE ASCII CHARACTER")
		}
	}
	return nil
}

func main() {
	// check user input
	if len(os.Args) < 2 {
		fmt.Println("Error.. RUN: go run . \"input-text\"")
		return
	}

	color := ""
	input := os.Args[1]
	banner := "standard"
	subStr := ""

	if strings.HasPrefix(input, "--color=") && len(os.Args) == 3 {
		color = input[8:]
		input = os.Args[2]
	}
	if strings.HasPrefix(input, "--color=") && len(os.Args) == 4 {
		color = input[8:]
		subStr = os.Args[2] // kit
		input = os.Args[3]  // a king kitten have kit

		if os.Args[3] == "standard" || os.Args[3] == "shadow" || os.Args[3] == "thinkertoy" {
			banner = os.Args[3]
		}
	}
	if len(os.Args) == 3 {

		if os.Args[2] == "standard" || os.Args[2] == "shadow" || os.Args[2] == "thinkertoy" {
			banner = os.Args[2]
		}
	}

	if strings.HasPrefix(input, "--color=") && len(os.Args) == 5 {
		subStr = os.Args[2]
		input = os.Args[3]
		banner = os.Args[4]
		color = os.Args[1][8:]
	}

	if input == "" {
		return
	}

	err := verifyInput(input)
	if err != nil {
		log.Fatal(err)
	}

	// load data from banner into [][]string
	asciiTable, err := functions.LoadBanner(banner + ".txt")
	if err != nil {
		fmt.Println("Error.. Loading banner")
		return
	}

	inputArg := strings.Split(input, `\n`)

	PrintArt(asciiTable, inputArg, color, subStr)

}
