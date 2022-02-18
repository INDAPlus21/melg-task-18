// go build -o C:\Users\Marcus\Documents\adk\labb1\filegenerator\out
// out/filegenerator.go.exe
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Generating indexfile.txt and magicfile.txt, this will take a while")
	createfiles()
	fmt.Println("Successfully generated files")
}

func createfiles() {
	tokenfile, _ := os.Open("../token.txt")
	indexfile, _ := os.Create("../indexfile.txt")
	magicfile, _ := os.Create("../magicfile.txt")

	// Close files when finished
	defer tokenfile.Close()
	defer indexfile.Close()
	defer magicfile.Close()

	scanner := bufio.NewScanner(tokenfile)

	lastword := ""
	lasthash := -1

	// Every word in file
	for scanner.Scan() {
		word := strings.Split(scanner.Text(), " ")

		// New word = new line
		if lastword != word[0] {
			if lastword != "" {
				indexfile.WriteString("\n")
			}

			currenthash := hash(word[0])
			// New element in magic file
			if currenthash != lasthash {
				stat, _ := indexfile.Stat()

				// Assign all from last to this index
				for i := lasthash + 1; i <= currenthash; i++ {
					// Let every index have 10 bytes in magicfile, print index of element in indexfile
					bytearray := []byte(strconv.FormatInt(stat.Size(), 10))
					magicfile.WriteAt(bytearray, int64(i*10))
				}

				lasthash = currenthash
			}

			indexfile.WriteString(word[0])
			lastword = word[0]
		}

		indexfile.WriteString(" " + word[1]) // Write byte-index
	}

	// Add last index to magicfile for +1 to always work
	stat, _ := indexfile.Stat()
	bytearray := []byte(strconv.FormatInt(stat.Size(), 10))
	magicfile.WriteAt(bytearray, int64(lasthash+1)*10)
}

// Code in both projects
func hash(input string) int {
	input = strings.ToLower(input) // Lower and upper chars should be treated as same
	one := getindexorempty(input, 0)
	two := getindexorempty(input, 1)
	three := getindexorempty(input, 2)

	return chartoint(one)*900 + chartoint(two)*30 + chartoint(three) - 900
}

func getindexorempty(input string, pos int) rune {
	runes := []rune(input)
	if len(runes) > pos {
		return runes[pos] // Support non-ascii
	} else {
		return ' '
	}
}

var charmap = map[rune]int{'_': 0, 'a': 1, 'b': 2, 'c': 3, 'd': 4, 'e': 5, 'f': 6, 'g': 7, 'h': 8, 'i': 9, 'j': 10, 'k': 11, 'l': 12, 'm': 13, 'n': 14, 'o': 15, 'p': 16, 'q': 17, 'r': 18, 's': 19, 't': 20, 'u': 21, 'v': 22, 'w': 23, 'x': 24, 'y': 25, 'z': 26, 'ä': 27, 'å': 28, 'ö': 29, '�': 29}

func chartoint(char rune) int {
	return charmap[char] // convert a to 1, b to 2 and so on
}
