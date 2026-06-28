package functions

import (
	"bufio"
	"os"
)
// function to load the banner into a 2D string
func LoadBanner(filename string) ([][]string, error) {

	// open bannner file to read line by line
	bannerf, err := os.Open(filename)

	if err != nil {
		return nil, err
	}
	defer bannerf.Close()

	scanner := bufio.NewScanner(bannerf)

	var fileLines []string
	asciiTable := make([][]string, 95)

	// load line into a slice 
	for scanner.Scan() {
		line := scanner.Text()
		fileLines = append(fileLines, line)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	// load slice into slice of slice 
	for i := 32; i < 127; i++ {
		startIndex := (i-32)*9 + 1
		endIndex := startIndex + 8

		asciiTable[i-32] = fileLines[startIndex:endIndex]

	}
	//fmt.Printf("%#v", asciiTable)
	return asciiTable, nil
}
