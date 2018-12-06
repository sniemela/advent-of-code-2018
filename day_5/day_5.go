package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strings"
	"unicode"
)

const inputFile = "input.txt"

func main() {
	fmt.Printf("Part 1 = %d\n", part1())
	fmt.Printf("Part 2 = %d\n", part2())
}

func part1() int {
	polymer, _ := ioutil.ReadFile(inputFile)
	return react(polymer)
}

func part2() int {
	polymer, _ := ioutil.ReadFile(inputFile)
	testChars := "abcdefghijklmnopqrstuvxyz"
	bestScore := math.MaxInt32
	var bestChar string

	for i := 0; i < len(testChars); i++ {
		char := string(testChars[i])
		polymerStr := string(polymer)
		polymerStr = strings.Replace(strings.Replace(polymerStr, strings.ToUpper(char), "", -1), char, "", -1)

		score := react([]byte(polymerStr))

		if score < bestScore {
			bestScore = score
			bestChar = char
		}
	}

	fmt.Printf("Part 2, best char: %s\n", bestChar)

	return bestScore
}

func react(polymer []byte) int {
	for i := 0; i+1 < len(polymer); {
		left := rune(polymer[i])
		right := rune(polymer[i+1])

		if unicode.SimpleFold(left) == right || unicode.SimpleFold(right) == left {
			polymer = append(polymer[:i], polymer[i+2:]...)

			if i > 0 {
				i--
			}

			continue
		}

		i++
	}

	return len(polymer)
}
