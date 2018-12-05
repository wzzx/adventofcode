package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
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

func reactPolymer(polymer string) int {
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
	return len(polymer)
}

func main() {
	fh, err := os.Open("./input")
	if err != nil {
		fmt.Println(err)
	}
	var polymer_orig string
	var unittypes = make(map[string]string)

	// read the polymer
	scanner := bufio.NewScanner(fh)
	for scanner.Scan() {
		polymer_orig = scanner.Text()
	}
	// map the units in the polymer
	for i := 0; i < len(polymer_orig); i++ {
		s := strings.ToLower(string(polymer_orig[i]))
		if _, ok := unittypes[s]; !ok {
			unittypes[s] = s
		}
	}

	units := reactPolymer(polymer_orig)
	fmt.Printf("Original polymer units: %d\n", units)

	// remove units and check for most compact polymer
	for k, _ := range unittypes {
		polymer := polymer_orig
		regex_str := fmt.Sprintf("(?i)%s", unittypes[k])
		re := regexp.MustCompile(regex_str)
		polymer = re.ReplaceAllString(polymer, "")
		units = reactPolymer(polymer)
		fmt.Printf("Polymer without '%s' units: %d\n", k, units)

	}
}
