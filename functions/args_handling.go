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
