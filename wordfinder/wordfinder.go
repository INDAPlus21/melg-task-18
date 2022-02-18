// go build -o C:\Users\Marcus\Documents\adk\labb1\wordfinder\out
// out/wordfinder.go.exe
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Enter word to search for")
	var word = ""
	fmt.Scan(&word)
	currenthash := hash(word)

	// Read magic file and get magic index
	magicfile, _ := os.Open("../magicfile.txt")
	bytearray := make([]byte, 10)
	magicfile.ReadAt(bytearray, int64(currenthash)*10)
	trimmedstring := strings.Trim(string(bytearray), "\x00")
	magicindex, _ := strconv.Atoi(trimmedstring)

	// Read index file and get index
	indexfile, _ := os.Open("../indexfile.txt")
	indexreader := bufio.NewReaderSize(indexfile, 100000000) // Create huge buffer
	indexreader.Discard(magicindex)                          // Jump to correct position
	var indexes []string

	// Linear search to get correct value
	// 10000 tries
	for i := 0; i < 10000; i++ {
		line, _, _ := indexreader.ReadLine()
		indexes = strings.Split(strings.Trim(string(line), "\x00"), " ")

		// Found word
		if indexes[0] == word {
			break
		}
	}

	korpusfile, _ := os.Open("../korpus.txt")

	for i := 1; i < len(indexes); i++ {
		index, _ := strconv.ParseInt(indexes[i], 10, 64)
		bytearray = make([]byte, 60+len(indexes[0]))
		korpusfile.ReadAt(bytearray, max(0, index-30))
		fmt.Println(strings.ReplaceAll(string(bytearray), "\n", ""))
	}

	// Close all files to prevent permission errors or memory leaks
	korpusfile.Close()
	indexfile.Close()
	magicfile.Close()
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
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
