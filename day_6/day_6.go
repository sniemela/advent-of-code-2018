package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

type point struct {
	x        int
	y        int
	infinite bool
	area     int
}

type xy struct {
	x int
	y int
}

func main() {
	bytes, _ := ioutil.ReadFile("input.txt")
	lines := strings.Split(string(bytes), "\n")

	top := 0
	bottom := math.MaxInt32
	right := 0
	left := math.MaxInt32

	var points []*point

	for _, record := range lines {
		pointData := strings.Split(record, ",")
		x, _ := strconv.Atoi(strings.TrimSpace(pointData[0]))
		y, _ := strconv.Atoi(strings.TrimSpace(pointData[1]))
		point := &point{x: x, y: y}
		points = append(points, point)
		top = max(top, y)
		bottom = min(bottom, y)
		left = min(left, x)
		right = max(right, x)
	}

	distances := map[xy]int{}

	for y := bottom; y <= top; y++ {
		for x := left; x <= right; x++ {
			xy := xy{x: x, y: y}
			minDist := math.MaxInt32
			var minPoint *point
			minPointInf := false

			for _, p := range points {
				dist := distance(x, y, p)
				distances[xy] += dist

				if dist == minDist {
					minPoint = p
					minPointInf = true
				} else if dist < minDist {
					minDist = dist
					minPoint = p
					minPointInf = false
				}
			}

			if !minPointInf {
				minPoint.area++

				if y == top || y == bottom || x == left || x == right {
					minPoint.infinite = true
				}
			}
		}
	}

	var maxArea int
	var bestPoint *point
	for _, point := range points {
		if !point.infinite && point.area > maxArea {
			maxArea = point.area
			bestPoint = point
		}
	}

	fmt.Printf("Part 1, best area: %d (%d,%d)\n", maxArea, bestPoint.x, bestPoint.y)

	var regions int
	for _, totalDist := range distances {
		if totalDist < 10000 {
			regions++
		}
	}

	fmt.Printf("Part 2, regions: %d\n", regions)
}

func distance(x, y int, p *point) int {
	x1, y1 := float64(x), float64(y)
	x2, y2 := float64(p.x), float64(p.y)
	return int(math.Abs(x1-x2) + math.Abs(y1-y2))
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
