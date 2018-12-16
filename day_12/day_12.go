package main

import (
	"bufio"
	"fmt"
	"os"
)

const fileName = "input.txt"

func main() {
	file, _ := os.Open(fileName)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var prevState string
	rules := map[string]string{}

	for scanner.Scan() {
		text := scanner.Text()
		fmt.Sscanf(text, "initial state: %s", &prevState)

		var rule, target string
		n, _ := fmt.Sscanf(text, "%s => %s", &rule, &target)

		if n == 2 {
			rules[rule] = target
		}
	}

	intialState := prevState
	initialSize := len(intialState)
	state := prevState
	stateSum := sum(state, initialSize)
	prevDiff := 0
	diffRep := 0
	genStartRep := 0
	genStartRepSum := 0

	fmt.Printf("Initial state: %s\n", state)

	for gen := 1; gen <= 1000; gen++ {
		state = "...." + state + "...."
		next := ""

		for i := 2; i < len(state)-2; i++ {
			llcrr := state[i-2 : i+3]
			target, exists := rules[llcrr]

			if exists {
				next += target
			} else {
				next += "."
			}
		}

		state = next
		currentSum := sum(state, initialSize)
		diff := currentSum - stateSum
		stateSum = currentSum

		if gen == 20 {
			fmt.Printf("Part 1 = %d\n", stateSum)
		}

		if diff == prevDiff {
			diffRep++
			if diffRep == 10 {
				genStartRep = gen - 9
				break
			}
		} else {
			prevDiff = diff
			diffRep = 1
			genStartRepSum = currentSum
		}
	}

	result := prevDiff*(50000000000-genStartRep) + genStartRepSum
	fmt.Printf("Part 2 gen(%d) diff(%d) = %d\n", genStartRep, prevDiff, result)
}

func sum(state string, initialStateSize int) int {
	diff := (len(state) - initialStateSize) / 2
	var sum int
	for i, v := range state {
		if v == '#' {
			sum += (i - diff)
		}
	}
	return sum
}
