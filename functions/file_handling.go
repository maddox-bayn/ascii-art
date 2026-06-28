package functions

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// function helps to reverse ascii representation back to text format
func ToReverse(fileName string, banner [][]string) {
	var b strings.Builder
	// open and read file line by line 
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Group the lines into blocks of 8 lines.
	var blocks [][]string
	var block []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if len(scanner.Text()) == 0 {
			blocks = append(blocks, []string{})
			continue
		}
		block = append(block, scanner.Text())

		if len(block) == 8 {
			blocks = append(blocks, block)
			block = nil
		}
	}

	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for i, block := range blocks {
		if len(block) == 0 {
			b.WriteString("\\n")
			continue
		}
		// Start at column 0 
		current_column := 0
		// loop runs as long as column is less than art width
		for current_column < len(block[0]) {
			match := false
			for asciVl := 32; asciVl < 127; asciVl++ {
				asciiArt := banner[asciVl-32]
				charWidth := len(asciiArt[0])

				matchCount := 0
				// Safety guard: Don't slice past the end of the input line!
				if current_column+charWidth > len(block[0]) {
					continue
				}

				for lineidx := 0; lineidx < 8; lineidx++ {

					// slice out the correct column window from this line
					if block[lineidx][current_column:current_column+charWidth] == asciiArt[lineidx] {
						matchCount++
					} else {
						break
					}

					if matchCount == 8 {
						match = true
					}
				}
				// if match found write to the string builder
				if match == true {
					b.WriteRune(rune(asciVl))
					match = false
					current_column += charWidth
					break // stop searching the data base
				}

			}
		}
		
		if i < len(blocks)-1 {
			b.WriteString("\\n")
		}
	}

	fmt.Println(b.String())
}
