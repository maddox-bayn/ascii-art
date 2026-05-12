package functions

import (
	"errors"
	"fmt"
	"os"
)

func CheckArgument(args []string) error {
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
	return nil
}
func ValidateInput(str string) {
	for _, char := range str {
		if char < 32 || char > 126 {
			//fmt.Println("Error: Wrong input Try: PRINTABLE ASCII CHARACTER")
			fmt.Fprint(os.Stdout, "Error: Wrong input Try: PRINTABLE ASCII CHARACTER\n")
			os.Exit(2)
		}
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
