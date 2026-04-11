package functions

import (
	"bufio"
	"os"
)

func LoadBanner(filename string) ([][]string, error) {

	bannerf, err := os.Open(filename)

	if err != nil {
		return nil, err
	}
	defer bannerf.Close()

	scanner := bufio.NewScanner(bannerf)

	var fileLines []string
	asciiTable := make([][]string, 127)

	for scanner.Scan() {
		line := scanner.Text()
		fileLines = append(fileLines, line)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	for i := 32; i < 127; i++ {
		startIndex := (i-32)*9 + 1
		endIndex := startIndex + 8

		if len(fileLines) <= 126 {
			asciiTable[i] = fileLines[startIndex:endIndex]
		}
	}
	return asciiTable, nil
}
