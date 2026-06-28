package functions

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// rendering text into art of []string of length 8 
func RenderLine(text string, banner [][]string) []string {
	var result []string
	var b strings.Builder

	for i := 0; i < 8; i++ {
		for _, c := range text {
			line := banner[c-32]
			b.WriteString(line[i])
		}
		result = append(result, b.String())
		b.Reset()
	}
	return result
}

// writing art to file
func saveTofile(art, fileName string) {

	ext := filepath.Ext(fileName)
	if ext == "" {
		fileName += ".txt"
	}
	if ext != ".txt" && ext != "" {
		log.Fatal("Error file must have txt as an extention: <filename.txt>, ", ext)
	}

	file, err := os.Create(fileName)
	if err != nil {
		log.Fatal("Error creating file")
	}
	defer file.Close()

	_, err = file.Write([]byte(art))
	if err != nil {
		log.Fatal("Error... writing to file", err)
	}

}
// printing art to the terminal
func PrintArt(s string) {
	fmt.Printf("%s", s)
}

// helper function to print to terminal or to file
func SaveOrPrintart(s string, flag map[string]string) {
	file := flag["output"]
	if file != "" {
		saveTofile(s, file)
	} else {
		PrintArt(s)
	}
}
