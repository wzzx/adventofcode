package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fh, err := os.Open("./input")
	if err != nil {
		fmt.Println(err)
	}
	defer fh.Close()

	var total int64 = 0
	var firstrepfreq int64 = 0
	var found = false
	var m = make(map[int64]int)

	for !found {
		scanner := bufio.NewScanner(fh)
		for scanner.Scan() {
			if curr, err := strconv.ParseInt(scanner.Text(), 10, 32); err == nil {
				// fmt.Printf("%d + %d = ", total, curr)
				total += curr

				m[total] += 1
				if m[total] > 1 && !found {
					// fmt.Printf("\t*** %d\n", total)
					firstrepfreq = total
					found = true
				}

				// fmt.Printf("%d | Map: %d \n", total, m[total])
			}
		}
		if err := scanner.Err(); err != nil {
			panic(err)
		}

		fmt.Printf("Total: %d. Repeated frequency: %d\n", total, firstrepfreq)
		fh.Seek(0, 0)

	}
}
