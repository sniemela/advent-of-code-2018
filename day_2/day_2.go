package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	boxIds := readBoxIds("input.txt")

	fmt.Printf("Part 1 (checksum) = %d\n", part1(boxIds))
	fmt.Printf("Part 2 (common letters) = %s\n", part2(boxIds))
}

func part1(ids []string) int {
	boxesWithTwoMatches := 0
	boxesWithThreeMatches := 0

	for _, id := range ids {
		two := 0
		three := 0
		split := strings.Split(id, "")

		for _, letter := range split {
			count := strings.Count(id, letter)
			if count == 2 {
				two = 1
			} else if count == 3 {
				three = 1
			}
		}

		boxesWithTwoMatches += two
		boxesWithThreeMatches += three
	}

	return boxesWithTwoMatches * boxesWithThreeMatches
}

func part2(ids []string) string {
	for i := 0; i < len(ids); i++ {
		thisID := ids[i]

		for j := i + 1; j < len(ids); j++ {
			anotherID := ids[j]
			letterDiffIndex := compare(thisID, anotherID)

			if letterDiffIndex != -1 {
				return thisID[0:letterDiffIndex] + thisID[letterDiffIndex+1:]
			}
		}
	}
	return ""
}

func compare(a, b string) int {
	diffCount := 0
	lastDiffIndex := -1

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			diffCount++
			lastDiffIndex = i
		}
	}

	if diffCount == 1 {
		return lastDiffIndex
	}

	return -1
}

func readBoxIds(filename string) []string {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	var ids []string

	for scanner.Scan() {
		ids = append(ids, scanner.Text())
	}

	return ids
}
