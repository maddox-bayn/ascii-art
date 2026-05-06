package functions

import (
	"fmt"
	"os"
)

func Usage()  {
	fmt.Fprintf(os.Stdout,"Usage: go run . [OPTOIN] [STRING] [BANNER]\n\n")
}