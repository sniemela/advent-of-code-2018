package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	bytes, _ := ioutil.ReadFile("input.txt")
	temp := strings.Split(string(bytes), " ")
	var entries []int
	for _, entry := range temp {
		n, _ := strconv.Atoi(entry)
		entries = append(entries, n)
	}
	resPart1, _ := part1(0, entries, 0)
	fmt.Printf("Part 1: %d\n", resPart1)

	resPart2, _ := part2(0, entries)
	fmt.Printf("Part 2: %d\n", resPart2)
}

func part1(startIndex int, metadataEntries []int, result int) (int, int) {
	children := metadataEntries[startIndex]
	entries := metadataEntries[startIndex+1]
	startIndex += 2

	for i := 0; i < children; i++ {
		res, endIndex := part1(startIndex, metadataEntries, result)
		startIndex = endIndex
		result = res
	}

	endIndex := startIndex + entries

	for i := startIndex; i < endIndex; i++ {
		result += metadataEntries[i]
	}

	return result, endIndex
}

func part2(startIndex int, metadataEntries []int) (int, int) {
	children := metadataEntries[startIndex]
	entries := metadataEntries[startIndex+1]
	startIndex += 2

	var results []int

	for i := 0; i < children; i++ {
		res, endIndex := part2(startIndex, metadataEntries)
		startIndex = endIndex
		results = append(results, res)
	}

	endIndex := startIndex + entries
	var result int

	for i := startIndex; i < endIndex; i++ {
		entry := metadataEntries[i]

		if children > 0 {
			index := entry - 1
			if index >= 0 && index < children {
				result += results[index]
			}
		} else {
			result += entry
		}
	}

	return result, endIndex
}
