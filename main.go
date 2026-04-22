package main

import (
	"ascii-art/functions"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

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

	input := os.Args[1]
	banner := "standard.txt"
	if len(os.Args) > 2 {
		if os.Args[2] == "standard.txt" || os.Args[2] == "shadow.txt" || os.Args[2] == "thinkertoy.txt" {
			banner = os.Args[2]
		}
	}


	if input == "" {
		return
	}

	err := verifyInput(input)
	if err != nil {
		log.Fatal(err)
	}

	// load data from banner into [][]string
	asciiTable, err := functions.LoadBanner(banner)
	if err != nil {
		fmt.Println("Error.. Loading banner")
		return
	}

	inputArg := strings.Split(input, `\n`)

	// if args are  only newlines
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
		for lineChar := 0; lineChar < 8; lineChar++ {
			for _, char := range word {
				fmt.Print(asciiTable[char-32][lineChar])
			}
			fmt.Println()
		}
	}
}
