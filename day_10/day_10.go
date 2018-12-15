package main

import (
	"bufio"
	"fmt"
	"os"
)

const inputFile = "input.txt"
const maxSeconds = 100000

type point struct {
	ppx int
	ppy int
	px  int
	py  int
	vx  int
	vy  int
}

func (p *point) move() {
	p.px = p.px + p.vx
	p.py = p.py + p.vy
}

type canvas struct {
	points []*point
}

func (c *canvas) move() {
	for _, p := range c.points {
		p.move()
	}
}

func (c *canvas) edges() (int, int, int, int) {
	minx := 1000
	miny := 1000
	maxx := -1000
	maxy := -1000

	for _, p := range c.points {
		if p.px < minx {
			minx = p.px
		}
		if p.px > maxx {
			maxx = p.px
		}
		if p.py < miny {
			miny = p.py
		}
		if p.py > maxy {
			maxy = p.py
		}
	}

	return minx, miny, maxx, maxy
}

func (c *canvas) print(second int) {
	minx, miny, maxx, maxy := c.edges()

	if maxy-miny < 30 {
		fmt.Printf("State at %d\n", second+1)

		m := make(map[string]string)
		for _, p := range c.points {
			key := fmt.Sprintf("%d,%d", p.px, p.py)
			m[key] = "#"
		}

		for y := miny; y <= maxy; y++ {
			line := ""
			for x := minx; x <= maxx; x++ {
				key := fmt.Sprintf("%d,%d", x, y)
				_, exists := m[key]
				if exists {
					line += "#"
				} else {
					line += " "
				}
			}
			fmt.Println(line)
		}
	}
}

func main() {
	file, _ := os.Open(inputFile)
	defer file.Close()
	scanner := bufio.NewScanner(file)

	canvas := &canvas{}

	for scanner.Scan() {
		p := &point{}
		fmt.Sscanf(scanner.Text(), "position=<%d,  %d> velocity=<%d,  %d>", &p.px, &p.py, &p.vx, &p.vy)
		canvas.points = append(canvas.points, p)
	}

	for i := 0; i < maxSeconds; i++ {
		canvas.move()
		canvas.print(i)
	}
}
