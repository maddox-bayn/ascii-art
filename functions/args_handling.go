package functions

import "strings"

func ExtractFlags(defaultAgs []string) (map[string]string, []string) {
	flags := make(map[string]string)
	var filterdArgs []string

	for _, args := range defaultAgs {
		if strings.HasPrefix(args, "--color=") {
			flags["color"] = args[8:]
		} else if strings.HasPrefix(args, "--output=") {
			flags["ouput"] = args[9:]
		} else {
			filterdArgs = append(filterdArgs, args)
		}

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
