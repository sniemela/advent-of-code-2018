package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	offsets, err := frequencyOffsets("input.txt")

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Part 1 = %d\n", part1(0, offsets))
	fmt.Printf("Part 2 = %d\n", part2(0, offsets))
}

func part1(currentFrequency int, offsets []int) int {
	for _, offset := range offsets {
		currentFrequency += offset
	}
	return currentFrequency
}

func part2(currentFrequency int, offsets []int) int {
	var seen []int

	// We have seen the current frequency
	seen = append(seen, currentFrequency)

	// Using a brute force algorithm to find the match
	for {
		for _, offset := range offsets {
			currentFrequency += offset

			for _, seenFreq := range seen {
				if seenFreq == currentFrequency {
					return seenFreq
				}
			}

			seen = append(seen, currentFrequency)
		}
	}
}

func frequencyOffsets(filename string) ([]int, error) {
	file, err := os.Open(filename)

	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	var offsets []int

	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		offsets = append(offsets, n)
	}

	return offsets, scanner.Err()
}
