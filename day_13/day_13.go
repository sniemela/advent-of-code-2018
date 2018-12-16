package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

const fileName = "input.txt"

type cart struct {
	x            int
	y            int
	dx           int
	dy           int
	intersection int
	crashed      bool
}

func (c *cart) move() {
	c.x += c.dx
	c.y += c.dy
}

func (c *cart) turnLeft() {
	dx := c.dx
	c.dx = c.dy
	c.dy = dx * -1
}

func (c *cart) turnRight() {
	dx := c.dx
	c.dx = c.dy * -1
	c.dy = dx
}

func (c *cart) icon() rune {
	if c.dy == 1 {
		return 'v'
	} else if c.dy == -1 {
		return '^'
	} else if c.dx == -1 {
		return '<'
	} else if c.dx == 1 {
		return '>'
	}
	return '?'
}

func (c *cart) turn() {
	switch c.intersection {
	case 0:
		c.turnLeft()
		c.intersection = 1
		break
	case 1:
		// straight
		// do nothing
		c.intersection = 2
		break
	case 2:
		// right
		c.turnRight()
		c.intersection = 0
		break
	}
}

type carts []*cart

func (a carts) Len() int { return len(a) }
func (a carts) Less(i, j int) bool {
	l := a[i]
	r := a[j]
	return l.x*l.x+l.y*l.y < r.x*r.x+r.y*r.y
}
func (a carts) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

func main() {
	file, _ := os.Open(fileName)
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var m [][]string
	var carts carts

	for scanner.Scan() {
		track := strings.Split(scanner.Text(), "")
		m = append(m, track)

		for i, v := range track {
			var c *cart
			if v == ">" {
				c = &cart{dx: 1, dy: 0}
				m[len(m)-1][i] = "-"
			} else if v == "<" {
				c = &cart{dx: -1, dy: 0}
				m[len(m)-1][i] = "-"
			} else if v == "v" {
				c = &cart{dx: 0, dy: 1}
				m[len(m)-1][i] = "|"
			} else if v == "^" {
				c = &cart{dx: 0, dy: -1}
				m[len(m)-1][i] = "|"
			}
			if c != nil {
				c.x = i
				c.y = len(m) - 1
				carts = append(carts, c)
			}
		}
	}

	ticks := 100000
	collision := false
	oneLeft := false

	for i := 1; i <= ticks && !oneLeft; i++ {
		sort.Sort(carts)

		for j := 0; j < len(carts); j++ {
			c := carts[j]
			moveCart(c, m[c.y][c.x])

			// check collision at this point
			for _, otherCart := range carts {
				if c != otherCart && !c.crashed && !otherCart.crashed && c.x == otherCart.x && c.y == otherCart.y {
					if !collision {
						fmt.Printf("Part 1 = %d,%d\n", c.x, c.y)
					}
					c.crashed = true
					otherCart.crashed = true
					collision = true
				}
			}

			var livingCarts []*cart

			for _, c := range carts {
				if !c.crashed {
					livingCarts = append(livingCarts, c)
				}
			}

			if len(livingCarts) == 1 {
				lastCart := livingCarts[0]
				// Simulate to next tick
				moveCart(lastCart, m[lastCart.y][lastCart.x])

				fmt.Printf("Part 2 = %d,%d\n", lastCart.x, lastCart.y)
				oneLeft = true
				break
			}
		}
	}
}

func moveCart(c *cart, tile string) {
	if tile == "+" {
		c.turn()
	} else if tile == "\\" {
		if c.dy == -1 || c.dy == 1 {
			c.turnLeft()
		} else {
			c.turnRight()
		}
	} else if tile == "/" {
		if c.dy == -1 || c.dy == 1 {
			c.turnRight()
		} else {
			c.turnLeft()
		}
	}

	c.move()
}

func printMap(m [][]string, carts []*cart) {
	for y, line := range m {
		l := strings.Join(line, "")

		for _, c := range carts {
			if c.y == y {
				runes := []rune(l)
				runes[c.x] = c.icon()
				l = string(runes)
			}
		}

		fmt.Println(l)
	}
}
