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

	colorMap := map[string]string{"red": "\033[31m", "green": "\033[32m", "yellow": "\033[34m"}


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
		stIdx := strings.Index(word, subStr)
		for lineChar := 0; lineChar < 8; lineChar++ {
			for i, char := range word {
				if i >= stIdx || i <= stIdx+len(subStr)-1 {
					colorMap[color] + asciiTable[char-32][lineChar] + "\033[0m"
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
		input = os.Args[2]
		if os.Args[3] == "standard" || os.Args[3] == "shadow" || os.Args[3] == "thinkertoy" {
			banner = os.Args[3]
		}
	}

	if os.Args[2] == "standard" || os.Args[2] == "shadow" || os.Args[2] == "thinkertoy" {
		banner = os.Args[2]
	}

	if strings.HasPrefix(input, "--color=") && len(os.Args) == 5 {
		subStr = os.Args[2]
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
