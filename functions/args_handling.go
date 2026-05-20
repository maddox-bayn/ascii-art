package functions

import (
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
			flags["output"] = args[9:]
			continue
		}
		if strings.HasPrefix(args, "--align=") {
			flags["align"] = args[8:]
			continue
		}
		if strings.HasPrefix(args, "--reverse=") {
			flags["reverse"] = args[10:]
			continue
		}

		filterdArgs = append(filterdArgs, args)
	}
	return flags, filterdArgs
}

func GetArgs(args []string, hascolor bool) (substr, text, banner string) {
	banner = "standard"
	if len(args) > 1 {
		last := args[len(args)-1]

		if last == "standard" || last == "shadow" || last == "thinkertoy" {
			banner = last
			args = args[:len(args)-1]
		}
	}
	// if has color get the substr and text
	if hascolor && len(args) >= 2 {
		substr = args[0]
		text = args[1]
		// if no substr get only text
	} else if len(args) >= 1 {
		text = args[0]
	}
	return substr, text, banner
}
