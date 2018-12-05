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
			//fmt.Printf("\t[%d] There's a reaction!: '%s' '%s'\n", i, string(polymer[i-1]), string(polymer[i]))
			if i+1 < len(polymer) {
				polymer_right = polymer[i+1:]
			} else {
				break
			}

			polymer = polymer_left + polymer_right
			// make sure we don't miss the first character
			polymer_left = ""
			polymer_right = ""
			i = 0
		} else {
			polymer_left += string(polymer[i-1])
		}
		//fmt.Printf("Polymer orig: %s\nPolymer: %s\n\tPolymer left: %s\n\tPolymer right: %s\n", polymer_orig, polymer, polymer_left, polymer_right)
	}

	fmt.Printf("Resulting polymer: %s\n", polymer)
}
