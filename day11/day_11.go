package main

import (
	"fmt"
	"math"
	"strconv"
)

const gridSize = 300
const serial = 7139

func main() {
	left, top := part1()
	fmt.Printf("Part 1 = %d,%d\n", left, top)

	l, t, size := part2()
	fmt.Printf("Part 2 = %d,%d,%d\n", l, t, size)
}

func part1() (int, int) {
	max := math.MinInt32
	var top, left int

	for y := 0; y < gridSize-3; y++ {
		for x := 0; x < gridSize-3; x++ {
			var total int

			for i := y; i < y+3; i++ {
				for j := x; j < x+3; j++ {
					total += power(j+1, i+1, serial)
				}
			}

			if total > max {
				max = total
				top = y
				left = x
			}
		}
	}

	return left + 1, top + 1
}

func part2() (int, int, int) {
	powers := make([][]int, gridSize, gridSize)

	for y := 0; y < gridSize; y++ {
		for x := 0; x < gridSize; x++ {
			powers[x] = append(powers[x], power(x+1, y+1, serial))
		}
	}

	// Summed area table
	for y := 1; y < gridSize; y++ {
		for x := 1; x < gridSize; x++ {
			powers[x][y] = powers[x][y] + powers[x][y-1] + powers[x-1][y] - powers[x-1][y-1]
		}
	}

	currentMax := math.MinInt32
	var top, left, currentSize int

	for y := 0; y < gridSize; y++ {
		for x := 0; x < gridSize; x++ {
			for s := 0; s < gridSize-max(y, x); s++ {
				power := powers[x+s][y+s] + powers[x][y] - powers[x+s][y] - powers[x][y+s]

				if power > currentMax {
					currentMax = power
					top = y + 1
					left = x + 1
					currentSize = s
				}
			}
		}
	}

	return left + 1, top + 1, currentSize
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func power(x, y int, s int) int {
	rackID := x + 10
	digits := strconv.Itoa((rackID*y + s) * rackID)
	digit := 0

	if len(digits) > 2 {
		temp, _ := strconv.Atoi(string(digits[len(digits)-3]))
		digit = temp
	}

	return digit - 5
}
