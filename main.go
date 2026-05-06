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

	defaultArgs := os.Args[1:]

	if len(defaultArgs) == 0 || len(defaultArgs[0]) == 0 {
		if len(defaultArgs) > 1 && len(defaultArgs[0]) == 0 {
		functions.Usage()
		}
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

	functions.PrintArt(asciiTable, inputArg, color, subStr)

}
