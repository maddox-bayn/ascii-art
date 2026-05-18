package functions

import (
	"bufio"
	"log"
	"os"
	"strings"
)

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
		block = append(block, scanner.Text())

		if len(block) == 8 {
			blocks = append(blocks, block)
			block = nil
		}
	}

	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}
	for _, block := range blocks {
		current_column := 0

		// Safety check: make sure the block actually has lines
		if len(block) < 8 {
			continue
		}

		// Loop across the width of the ASCII art line
		for current_column < len(block[0]) {
			matched := false

			// Test every printable character in your database
			for asciiVal := 32; asciiVal < 127; asciiVal++ {
				charArt := banner[asciiVal-32]
				charWidth := len(charArt[0])

				// Safety guard: Don't slice past the end of the input line!
				if current_column+charWidth > len(block[0]) {
					continue
				}

				// Check if the 8-line window matches charArt
				matchCount := 0
				for lineIdx := 0; lineIdx < 8; lineIdx++ {
					// Slice out the exact column window for this line
					inputSlice := block[lineIdx][current_column : current_column+charWidth]

					if inputSlice == charArt[lineIdx] {
						matchCount++
					} else {
						break // Wrong character, stop checking lines
					}
				}

				// If all 8 lines matched perfectly
				if matchCount == 8 {
					b.WriteRune(rune(asciiVal))
					current_column += charWidth // Move cursor past this character
					matched = true
					break // Stop searching the database, move to next character position
				}
			}

			// Edge Case Safety: If no character matched, prevent an infinite loop
			if !matched {
				current_column++
			}
		}
		b.WriteString("\n") // Newline after finishing an 8-line row
	}
	PrintArt(b.String())
}
