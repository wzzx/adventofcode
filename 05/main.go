package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func reactPair(a byte, b byte) bool {
	if a == b {
		return false
	}

	a_lower := strings.ToLower(string(a))
	b_lower := strings.ToLower(string(b))

	// same type there could be reaction
	if a_lower == b_lower && a != b {
		return true
	}
	return false
}

func main() {
	fh, err := os.Open("./input")
	if err != nil {
		fmt.Println(err)
	}
	var polymer_orig string
	var polymer string

	scanner := bufio.NewScanner(fh)
	for scanner.Scan() {
		polymer_orig = scanner.Text()
	}
	polymer = polymer_orig

	polymer_left := ""
	polymer_right := ""

	for i := 1; i < len(polymer); i++ {
		reaction := reactPair(polymer[i-1], polymer[i])
		if reaction {
			if i+1 < len(polymer) {
				polymer_right = polymer[i+1:]
			} else {
				polymer_right = ""
			}

			polymer = polymer_left + polymer_right
			// make sure we don't miss the first character
			polymer_left = ""
			polymer_right = ""
			// -2 because it will be incremented in the for loop
			if i-2 < 0 {
				i = 0
			} else {
				i -= 2
			}
		}
		polymer_left = string(polymer[:i])
	}

	fmt.Printf("Polymer: %s\nUnits: %d\n", polymer, len(polymer))
}
