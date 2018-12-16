package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

const inputFile = "input.txt"

func main() {
	history := readShiftHistory(inputFile)
	fmt.Printf("Part 1 = %d\n", history.part1())
	fmt.Printf("Part 2 = %d\n", history.part2())
}

type shiftHistory struct {
	guards []*guard
}

func (h *shiftHistory) part1() int {
	var sleepiestGuard *guard
	maxMinutes := -1
	sleepiestMinute := -1

	for _, guard := range h.guards {
		var maxMinuteIndex, localTotalMinutes int

		for minute, total := range guard.data {
			localTotalMinutes += total

			if total > guard.data[maxMinuteIndex] {
				maxMinuteIndex = minute
			}
		}

		if localTotalMinutes > maxMinutes {
			sleepiestGuard = guard
			sleepiestMinute = maxMinuteIndex
			maxMinutes = localTotalMinutes
		}
	}

	return sleepiestGuard.ID * sleepiestMinute
}

func (h *shiftHistory) part2() int {
	var sleepiestGuard *guard
	maxTotalMinutes := -1
	sleepiestMinute := -1

	for _, guard := range h.guards {
		var maxMinuteIndex int

		for minute, total := range guard.data {
			if total > guard.data[maxMinuteIndex] {
				maxMinuteIndex = minute
			}
		}

		if guard.data[maxMinuteIndex] > maxTotalMinutes {
			sleepiestGuard = guard
			sleepiestMinute = maxMinuteIndex
			maxTotalMinutes = guard.data[maxMinuteIndex]
		}
	}

	return sleepiestGuard.ID * sleepiestMinute
}

type guard struct {
	ID   int
	data []int
}

func readShiftHistory(filename string) *shiftHistory {
	file, _ := os.Open(filename)
	scanner := bufio.NewScanner(file)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	sort.Strings(lines)

	shiftHistory := &shiftHistory{}
	guards := make(map[int]*guard)
	var currentGuard *guard

	var sleepFrom, sleepTo int

	for _, record := range lines {
		description := record[19:]

		if strings.HasPrefix(description, "Guard") {
			guardID, _ := strconv.Atoi(strings.TrimSpace(description[7:10]))
			currentGuard = guards[guardID]

			if currentGuard == nil {
				currentGuard = &guard{
					ID:   guardID,
					data: make([]int, 60, 60),
				}
				guards[guardID] = currentGuard
				shiftHistory.guards = append(shiftHistory.guards, currentGuard)
			}
		} else if strings.HasPrefix(description, "falls") {
			sleepFrom, _ = strconv.Atoi(record[15:17])
		} else {
			sleepTo, _ = strconv.Atoi(record[15:17])

			for i := sleepFrom; i < sleepTo; i++ {
				currentGuard.data[i]++
			}
		}
	}

	return shiftHistory
}
