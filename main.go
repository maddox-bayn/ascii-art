package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	input := os.Args[1]

	fmt.Println(input)

	inputArg := strings.Split(input, "\n")
	fmt.Println(inputArg)

	for i := 0; i < len(inputArg); i++ {
		for j := 0; j <= 7; j++ {
			for k := 0; k < len(inputArg[i]); k++ {
				fmt.Print(string(inputArg[i][k]))
			}
		}
	}

}
