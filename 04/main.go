package main

import (
	"bufio"
	"fmt"
	"os"
	_ "strconv"
)

type Guard struct {
	id         int
	totalsleep int
}

func NewGuard() *Guard {
	return &Guard{
		id:         -1,
		totalsleep: 0}
}

func isWhitespace(r rune) bool {
	return r == ' ' || r == '\t' || r == '\n'
}

func isControlChar(r rune) bool {
	return r == '[' || r == ']' || r == '#'
}

func parseLine(line string, current_g *Guard, guards map[int]*Guard) {
	// I could fucking do this with splits on space and get done with it but I
	// decided that it would be more fun if I built a fucking parser... pain in
	// the arse...
	var (
		word        string
		isDateStart = false
		isDateEnd   = false
		isTime      = false
		isId        = false

		id     int
		action string
		date   string
		time   string
	)
	for i, r := range line {
		// parse line and get guard id
		if !isWhitespace(r) {
			switch r {
			case '[':
				isHeaderStart = true
			case ']':
				isHeaderEnd = true
			case '#':
				isId = true
			}

			// don't store control characters
			if !isControlChar(r) {
				word += string(r)
			}
		} else {
			// we reached space therefore need to evaluate what's our word
			if isHeaderStart && !isHeaderEnd {
				// we have finished with the date and we are starting the time
				date = word
				isTime = true
			}
			if isHeaderStart && isHeaderEnd && isTime {
				time = word
				isTime = false
			}

			word = strings.ToLower(word)
			switch word {
			case "guard":
				isNewSequence = true
				var number string

				i = i + 2 // skip the space and #
				for !isWhitespace(line[i]) {
					number += line[i]
				}
				id, err := strconv.Atoi(number)

				// create guard if doesn't exist
				if _, ok := guards[id]; ok {
					fmt.Printf("New guard!\n")
				}
			case "falls":

			case "wakes":
			}
			// get ready for next word
			word = ""
		}

	}
}

func main() {
	fh, err := os.Open("./sorted_input")
	if err != nil {
		fmt.Println(err)
	}
	defer fh.Close()

	var guards = make(map[int]*Guard)

	scanner := bufio.NewScanner(fh)
	for scanner.Scan() {
		parseLine(scanner.Text(), guards)
	}
}
