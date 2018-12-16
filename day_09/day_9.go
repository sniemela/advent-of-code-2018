package main

import (
	"container/ring"
	"fmt"
	"io/ioutil"
)

func main() {
	var players, marbles int
	bytes, _ := ioutil.ReadFile("input.txt")
	fmt.Sscanf(string(bytes), "%d players; last marble is worth %d points", &players, &marbles)

	fmt.Printf("Part 1 = %d\n", play(marbles, players))
	fmt.Printf("Part 2 = %d\n", play(marbles*100, players))
}

func play(marbles, players int) int {
	r := ring.New(1)
	r.Value = 0

	scores := map[int]int{}

	for marble := 1; marble <= marbles; marble++ {
		if marble%23 == 0 {
			r = r.Move(-8)
			removed := r.Unlink(1)
			scores[marble%players] += removed.Value.(int) + marble
			r = r.Next()
		} else {
			tmp := ring.New(1)
			tmp.Value = marble
			r = r.Move(1)
			r = r.Link(tmp)
			r = r.Prev()
		}
	}

	var highScore int
	for _, score := range scores {
		if score > highScore {
			highScore = score
		}
	}

	return highScore
}
