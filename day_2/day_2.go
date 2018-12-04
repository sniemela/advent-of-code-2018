package main

import (
	"bufio"
	"fmt"
	"os"
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
		letterCount := make(map[rune]int)

		for _, letter := range id {
			letterCount[letter]++
		}

		for _, v := range letterCount {
			if v == 2 {
				boxesWithTwoMatches++
				break
			}
		}

		for _, v := range letterCount {
			if v == 3 {
				boxesWithThreeMatches++
				break
			}
		}
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
