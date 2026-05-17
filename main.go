package main

import (
	"ascii-art/functions"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	defaultArgs := os.Args[1:]

	if len(defaultArgs) == 0 || len(defaultArgs[0]) == 0 {
		if len(defaultArgs) > 1 && len(defaultArgs[0]) == 0 {
			functions.Usage()
		}
		return
	}
	flag, args := functions.ExtractFlags(defaultArgs)

	subStr, text, banner := functions.GetArgs(args, flag["color"] != "")

	err := functions.CheckArgument(args)
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
