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
	// load data from banner into [][]string
	asciiTable, err := functions.LoadBanner("standard.txt")
	if err != nil {
		fmt.Println("Error.. Loading banner")
		return
	}

	if len(os.Args) < 2 {
		fmt.Println("Error.. RUN: go run . \"input-text\"")
		return
	}

	input := os.Args[1]

	if input == "" {
		return
	}

	err = verifyInput(input)
	if err != nil {
		log.Fatal(err)
	}
	// for _, char := range input {
	// 	if char < 32 || char > 126 {
	// 		fmt.Println("Error: Wrong input Try: PRINTABLE ASCII CHARACTER")
	// 		return
	// 	}
	// }

	inputArg := strings.Split(input, `\n`)

	for i, word := range inputArg {
		if word == "" {
			if len(inputArg)-1 == i {
				continue
			}
			fmt.Println()
			continue
		}
		for lineChar := 0; lineChar < 8; lineChar++ {
			for _, char := range word {
				fmt.Print(asciiTable[char][lineChar])
			}
			fmt.Println()
		}
	}
}
