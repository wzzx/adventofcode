package main

import (
	"bufio"
	"fmt"
	"os"
	_ "strconv"
)

func s_distance(s1 string, s2 string) bool {
	// Function will XOR two strings (assuming same length) and return true if
	// there's only one byte difference in the same position and false in any
	// other case. Example, these strings should return true, since there's
	// only one character difference in the same position:
	// abcde ^ abcdZ = 3F
	// abcAe ^ abcZe = 1B

	var diff_bytes = 0
	for i := 0; i < len(s1); i++ {
		if s1[i]^s2[i] != 0 {
			diff_bytes += 1
		}
		if diff_bytes > 1 {
			return false
		}
	}

	return true
}

func main() {
	fh, err := os.Open("./input")
	if err != nil {
		fmt.Println(err)
	}
	defer fh.Close()

	var ids []string
	var total_two, total_three int
	total_two, total_three = 0, 0

	scanner := bufio.NewScanner(fh)
	for scanner.Scan() {
		line := scanner.Text()
		ids = append(ids, line)
		fmt.Printf("*** %s\n", line)
		var m = make(map[byte]int)
		var two, three bool
		two, three = false, false

		// parse the current line
		for i := 0; i < len(line); i++ {
			c := line[i]
			m[c] += 1

		}

		// identify whether there are pairs or triplets
		subtottwo := 0
		subtotthree := 0
		for k, v := range m {
			if v == 2 {
				fmt.Printf("\t\t%s = %d\n", string(k), v)
				two = true
				subtottwo++
			} else if v == 3 {
				fmt.Printf("\t\t%s = %d\n", string(k), v)
				three = true
				subtotthree++
			}
			//fmt.Printf("\t%s %d\n", string(k), v)
		}
		fmt.Printf("\ttwo: %t three: %t (subtottwo: %d, subtotthree: %d)\n", two, three, subtottwo, subtotthree)

		// increase global counters accordingly
		if two {
			total_two += 1
		}
		if three {
			total_three += 1
		}
		fmt.Printf("\tTotal two: %d\n\tTotal three: %d\n\tChecksum: %d\n\n", total_two, total_three, total_two*total_three)

	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	fmt.Printf("Groups of two: %d\nGroups of three: %d\nChecksum: %d\n", total_two, total_three, total_two*total_three)

	var found = false
	fmt.Println("------------------------")
	for i := 0; i < len(ids) && !found; i++ {
		for j := 0; j < len(ids) && !found; j++ {
			if ids[i] != ids[j] {
				res := s_distance(ids[i], ids[j])
				if res {
					found = true
					fmt.Printf("%t: '%s' '%s'\n", res, ids[i], ids[j])
				}
			}
		}
	}

}
