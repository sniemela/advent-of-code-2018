package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

const mapWidth = 1000

type inchOwner struct {
	claimID string
	free    bool
}

func main() {
	claims := readClaims("input.txt")
	regex, _ := regexp.Compile("#(.+)\\s@\\s(\\d+),(\\d+):\\s(\\d+)x(\\d+)")
	fabricMap := make([]int, mapWidth*mapWidth, mapWidth*mapWidth)

	overlaps := 0
	inchOwners := make(map[int]*inchOwner)

	for _, claim := range claims {
		matches := regex.FindAllStringSubmatch(claim, -1)
		id := matches[0][1]
		x, _ := strconv.Atoi(matches[0][2])
		y, _ := strconv.Atoi(matches[0][3])
		width, _ := strconv.Atoi(matches[0][4])
		height, _ := strconv.Atoi(matches[0][5])

		owner := &inchOwner{claimID: id, free: true}

		targetX := x + width
		targetY := y + height

		for i := x; i < targetX; i++ {
			for j := y; j < targetY; j++ {
				position := mapWidth*i + j
				inch := fabricMap[position]

				o := inchOwners[position]

				if o == nil {
					inchOwners[position] = owner
				} else {
					o.free = false
					owner.free = false
				}

				if inch == 1 {
					overlaps++
				}

				fabricMap[position]++
			}
		}
	}

	fmt.Printf("Part 1 (overlaps): %d\n", overlaps)

	for _, v := range inchOwners {
		if v.free {
			fmt.Printf("Part 2 (free claim): %s\n", v.claimID)
			break
		}
	}
}

func readClaims(filename string) []string {
	file, _ := os.Open(filename)
	scanner := bufio.NewScanner(file)
	var claims []string

	for scanner.Scan() {
		claims = append(claims, scanner.Text())
	}

	return claims
}
