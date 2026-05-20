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
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

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
		current_column := 0
		for current_column < len(block[0]) {
			match := false
			for asciVl := 32; asciVl < 127; asciVl++ {
				asciiArt := banner[asciVl-32]
				charWidth := len(asciiArt[0])

				matchCount := 0

				for lineidx := 0; lineidx < 8; lineidx++ {

					if current_column+charWidth > len(block[0]) {
						continue
					}
					//fmt.Println("a")
					if block[lineidx][current_column:current_column+charWidth] == asciiArt[lineidx] {
						matchCount++
					} else {
						break
					}

					if matchCount == 8 {
						match = true
					}
				}
				if match == true {
					b.WriteRune(rune(asciVl))
					match = false
					current_column += charWidth
					break
				}

			}
		}
		if i < len(blocks)-1 {
			b.WriteString("\\n")
		}
	}

	fmt.Println(b.String())
}
