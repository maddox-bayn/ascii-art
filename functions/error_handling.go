package functions

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

// handle error and checking default argument and filtered args
func CheckArgument(args, defaultarg []string) error {
	arglen := len(args)
	if arglen == 3 {
		switch args[arglen-1] {
		case "standard", "shadow", "thinkertoy":
			// is valide banner
		default:
			return fmt.Errorf("not a banner name '%s' \nAvailable banner types are: 'standard' (default), 'shadow', and 'thinkertoy'", args[arglen-1])
		}
	} else if arglen > 3 {
		return errors.New("too many arguments")
	}
	if len(defaultarg) == 2 && !strings.HasPrefix(defaultarg[0], "--") {
		switch defaultarg[arglen-1] {
		case "standard", "shadow", "thinkertoy":
			// is valide banner
		default:
			return fmt.Errorf("not a banner name '%s' \nAvailable banner types are: 'standard' (default), 'shadow', and 'thinkertoy'", args[arglen-1])
		}
	}
	err := ValidateInput(args)
	if err != nil {
		return err
	}
	return nil
}

// validat args for character not in range of ascii char 32 to 126
func ValidateInput(str []string) error {
	for _, word := range str {
		for _, char := range word {
			if char < 32 || char > 126 {
				//fmt.Println("Error: Wrong input Try: PRINTABLE ASCII CHARACTER")
				return fmt.Errorf("Error: Wrong input '%c' Try: PRINTABLE ASCII CHARACTER  \n", char)
			}
		}
	}

	return nil
}

// helper function to check for invalid alignment
func CheckAlignment(s string) {
	switch s {
	case "left", "right", "center", "justify":
	default:
		fmt.Fprintf(os.Stdout, "Error... Invalide alignment '%s' \nvalide alignment are [left, right, center]\n", s)
		os.Exit(1)
	}
}

func Usage() {
	fmt.Fprintf(os.Stdout, "Usage: go run . [OPTOIN] [STRING] [BANNER]\n")
}

func PrintColor() {
	fmt.Println("Invalide color. Please choose one of the following colors")
	for ckey, clor := range ColorMap {
		fmt.Printf("- %s%s%s\n", clor, ckey, Reset)
	}
	os.Exit(1)
}
