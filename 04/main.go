package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Activity struct {
	action string
	date   string
	time   string
}

type Guard struct {
	id         int
	totalsleep int
	sleepmap   [60]int
	activity   []*Activity
}

func NewGuard() *Guard {
	return &Guard{
		id:         -1,
		totalsleep: 0}
}

func NewActivity() *Activity {
	return &Activity{
		action: "",
		date:   "",
		time:   ""}
}

func isWhitespace(r byte) bool {
	return r == ' ' || r == '\t' || r == '\n'
}

func isControlChar(r byte) bool {
	return r == '[' || r == ']' || r == '#'
}

func readWord(line string, start_pos int) (int, string) {
	// reads until hits space or end of line
	var word string
	var i int
	for i = start_pos; i < len(line); i++ {
		if isWhitespace(line[i]) {
			break
		}
		if !isControlChar(line[i]) {
			word += string(line[i])
		}
	}
	return i, word
}

func parseLine(line string, guards map[int]*Guard) (int, *Activity) {
	// I could fucking do this with splits on space and get done with it but I
	// decided that it would be more fun if I built a fucking parser... pain in
	// the arse...
	var (
		word string

		id     int
		action string
		date   string
		time   string
		gid    int       = -1
		a      *Activity = NewActivity()
	)

	for i := 0; i < len(line); i++ {
		switch line[i] {
		case '[':
			i, date = readWord(line, 1)
			i += 1
			i, time = readWord(line, i)
		default:
			i, word = readWord(line, i)
			word = strings.ToLower(word)
			switch word {
			case "guard":
				var number string
				i += 1
				i, number = readWord(line, i)
				id, _ = strconv.Atoi(number)
				break
			case "falls":
				action = "falls"
				break
			case "wakes":
				action = "wakes"
				break
			}
		}
	}

	if action != "" {
		a = &Activity{
			action: action,
			date:   date,
			time:   time}
	} else {
		gid = id
	}
	return gid, a
}

func main() {
	fh, err := os.Open("./sorted_input")
	if err != nil {
		fmt.Println(err)
	}
	defer fh.Close()

	var guards = make(map[int]*Guard)
	var current_g *Guard
	var start_count = false
	var start_minutes int
	var end_minutes int

	scanner := bufio.NewScanner(fh)
	for scanner.Scan() {
		gid, a := parseLine(scanner.Text(), guards)
		// update current guard
		if gid != -1 {
			// create guard if doesn't exist
			if _, ok := guards[gid]; !ok {
				g := NewGuard()
				g.id = gid
				guards[gid] = g
			}

			current_g = guards[gid]
		} else {
			current_g.activity = append(current_g.activity, a)
			if a.time[0:2] == "00" && a.action == "falls" {
				start_count = true
				sm, _ := strconv.ParseInt(a.time[3:], 10, 32)
				start_minutes = int(sm)
			}
			if start_count && a.action == "wakes" {
				start_count = false
				em, _ := strconv.ParseInt(a.time[3:], 10, 32)
				end_minutes = int(em)
				sleep_minutes := end_minutes - start_minutes

				// update the total register
				current_g.totalsleep += sleep_minutes

				// update heatmap
				for i := start_minutes; i < end_minutes; i++ {
					current_g.sleepmap[i] += 1
				}

				// reset counters
				start_minutes = 0
				end_minutes = 0
			}

		}
	}

	// present data and calculate answers
	total_sleep := 0
	total_min := 0
	sleeps_more_gid := 0
	sleeps_more_min := 0
	// -- part 2 --
	tot_abs_min := 0
	absolute_gid := 0
	absolute_min := 0

	fmt.Printf("------------ ")
	for i := 0; i < 60; i++ {
		fmt.Printf("%02d ", i)
	}
	fmt.Println()
	for k, v := range guards {
		fmt.Printf("[%4d] (%3d) ", k, v.totalsleep)
		if v.totalsleep > total_sleep {
			total_sleep = v.totalsleep
			sleeps_more_gid = k
			sleeps_more_min = 0
		}
		for i := 0; i < len(v.sleepmap); i++ {
			fmt.Printf("%02d ", v.sleepmap[i])
			if sleeps_more_gid == k && v.sleepmap[i] > total_min {
				sleeps_more_min = i
				total_min = v.sleepmap[i]
			}
			if v.sleepmap[i] > tot_abs_min {
				tot_abs_min = v.sleepmap[i]
				absolute_min = i
				absolute_gid = k
			}
		}
		fmt.Println()
	}

	fmt.Println()
	fmt.Printf("[%d] Sleeps more than anyone else, specially on minute %d. Answer: %d\n", sleeps_more_gid, sleeps_more_min, sleeps_more_min*sleeps_more_gid)
	fmt.Printf("Of all guards, [%d] tends to often fall asleep on minute %d more than anyone else. Answer %d\n", absolute_gid, absolute_min, absolute_gid*absolute_min)
}
