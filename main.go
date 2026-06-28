package main

import (
	"ascii-art/functions"
	"fmt"
	"log"
	"os"
	"strings"
)

// Recieve user input and call other functions to validat and process ascii art
func main() {
	defaultArgs := os.Args[1:]
	// validat program usage
	if len(defaultArgs) == 0 || len(defaultArgs[0]) == 0 {
		if len(defaultArgs) > 1 && len(defaultArgs[0]) == 0 {
			functions.Usage()
		}
		return
	}
	flag, args := functions.ExtractFlags(defaultArgs)
	if len(flag) == 0 && len(args) > 2 {
		functions.Usage()
		os.Exit(2)
	}

	subStr, text, banner := functions.GetArgs(args, flag["color"] != "")

	err := functions.CheckArgument(args, defaultArgs)
	if err != nil {
		log.Fatal(err)
	}

	// load data from banner into [][]string
	asciiTable, err := functions.LoadBanner(banner + ".txt")
	if err != nil {
		fmt.Println("Error.. Loading banner")
		return
	}

	inputArg := strings.Split(text, `\n`)

	functions.ProcessInput(asciiTable, inputArg, flag, subStr)

}
