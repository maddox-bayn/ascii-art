package functions

import (
	"fmt"
	"os"
	"strings"
)

func ExtractFlags(defaultAgs []string) (map[string]string, []string) {
	flags := make(map[string]string)
	var filterdArgs []string

	for i := 0; i < len(defaultAgs); i++ {
		args := defaultAgs[i]
		if strings.HasPrefix(args, "--color=") {
			flags["color"] = args[8:]
			_, ok := ColorMap[flags["color"]]
			if !ok {
				PrintColor()
			}

			continue
		}
		if strings.HasPrefix(args, "--output=") {
			flags["ouput"] = args[9:]
			continue
		}
		if strings.HasPrefix(args, "--align=") {
			flags["align"] = args[8:]
			continue
		}

		filterdArgs = append(filterdArgs, args)
	}
	fmt.Println(filterdArgs)
	if len(flags) == 0 && len(filterdArgs) > 2 {
		Usage()
		os.Exit(2)
	}
	return flags, filterdArgs
}

func GetArgs(args []string, hascolor bool) (substr, text, banner string) {
	banner = "standard"
	if len(args) > 0 {
		last := args[len(args)-1]

		if last == "standard" || last == "shadow" || last == "thinkertoy" {
			banner = last
			args = args[:len(args)-1]
		}
	}

	if hascolor && len(args) >= 2 {
		substr = args[0]
		text = args[1]
	} else if len(args) >= 1 {
		text = args[0]
	}
	return substr, text, banner
}
