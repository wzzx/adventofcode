package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Claim struct {
	id      int64
	x       int64
	y       int64
	size_x  int64
	size_y  int64
	overlap bool
}

func ParseClaim(claim string) *Claim {
	// equivalent of Split on space but it also squeezes spaces
	claim_parts := strings.Fields(claim)

	id, _ := strconv.ParseInt(strings.Replace(claim_parts[0], "#", "", -1), 10, 32)
	pos_str := claim_parts[2]
	pos := strings.Split(pos_str, ",")
	x, _ := strconv.ParseInt(pos[0], 10, 32)
	y, _ := strconv.ParseInt(strings.Replace(pos[1], ":", "", -1), 10, 32)
	size_str := claim_parts[3]
	size := strings.Split(size_str, "x")
	size_x, _ := strconv.ParseInt(size[0], 10, 32)
	size_y, _ := strconv.ParseInt(size[1], 10, 32)

	//fmt.Printf("Claim id: %d\n\tPosition: (%d,%d)\n\tSize: %dx%d\n", id, x, y, size_x, size_y)

	// fmt.Printf("From %d,%d for %d,%d\n", x, y, size_x, size_y)
	return &Claim{
		id:      id,
		x:       x,
		y:       y,
		size_x:  size_x,
		size_y:  size_y,
		overlap: false}
}

func UpdateFabric(c *Claim, claims map[int]*Claim, fabric *[1000][1000][]int) {
	// mark the fabric fields as used or not
	for i := int64(0); i < c.size_x; i++ {
		for j := int64(0); j < c.size_y; j++ {
			// fmt.Printf("\tPrinting id %d in pos: (%d, %d)\n", c.id, c.x+i, c.y+j)

			// if the fabric square has already some entries mark the current claim and the others as overlapping
			if fabric[c.x+i][c.y+j] != nil {
				c.overlap = true

				// detect if the square was already marked for overlaps and mark claims as overlapping
				found := false
				for s := 0; s < len(fabric[c.x+i][c.y+j]); s++ {
					claimid := fabric[c.x+i][c.y+j][s]

					if claimid == -1 {
						found = true
					} else {
						claims[claimid].overlap = true
					}
				}
				// if it's the first overlap, mark also only once the fabric square as used with -1
				if !found {
					fabric[c.x+i][c.y+j] = append(fabric[c.x+i][c.y+j], -1)
				}
			}

			// write the id of the current claim to the square
			fabric[c.x+i][c.y+j] = append(fabric[c.x+i][c.y+j], int(c.id))
		}
	}
}

func main() {
	fh, err := os.Open("./input")
	if err != nil {
		fmt.Println(err)
	}
	defer fh.Close()

	var claims = make(map[int]*Claim)
	fabric := [1000][1000][]int{}

	// read the file line by line and build the fabric as well as the map of claims
	scanner := bufio.NewScanner(fh)
	for scanner.Scan() {
		claimstr := scanner.Text()
		c := ParseClaim(claimstr)
		claims[int(c.id)] = c

		UpdateFabric(c, claims, &fabric)
	}

	total_overlap := 0
	for i := 0; i < len(fabric); i++ {
		for j := 0; j < len(fabric[i]); j++ {
			for s := 0; s < len(fabric[i][j]); s++ {
				if fabric[i][j][s] == -1 {
					total_overlap += 1
					break
				}
			}
		}
	}
	fmt.Printf("Total overlapping: %d\n", total_overlap)

	for k, _ := range claims {
		if claims[k].overlap == false {
			fmt.Printf("Non-overlapping claim is the one with ID: %d\n", k)
			break
		}
	}
}
