package main

import (
	"fmt"
)

func main() {
	assertEqual(-1, chop(3, []int{}))
	assertEqual(-1, chop(3, []int{1}))
	assertEqual(0, chop(1, []int{1}))

	assertEqual(0, chop(1, []int{1, 3, 5}))
	assertEqual(1, chop(3, []int{1, 3, 5}))
	assertEqual(2, chop(5, []int{1, 3, 5}))
	assertEqual(-1, chop(0, []int{1, 3, 5}))
	assertEqual(-1, chop(2, []int{1, 3, 5}))
	assertEqual(-1, chop(4, []int{1, 3, 5}))
	assertEqual(-1, chop(6, []int{1, 3, 5}))

	assertEqual(0, chop(1, []int{1, 3, 5, 7}))
	assertEqual(1, chop(3, []int{1, 3, 5, 7}))
	assertEqual(2, chop(5, []int{1, 3, 5, 7}))
	assertEqual(3, chop(7, []int{1, 3, 5, 7}))
	assertEqual(-1, chop(0, []int{1, 3, 5, 7}))
	assertEqual(-1, chop(2, []int{1, 3, 5, 7}))
	assertEqual(-1, chop(4, []int{1, 3, 5, 7}))
	assertEqual(-1, chop(6, []int{1, 3, 5, 7}))
	assertEqual(-1, chop(8, []int{1, 3, 5, 7}))
}

/*
I ported a simple ruby binary chop, and added the goroutines bit
here is the ruby solution (i used the recursive one obviously)

http://softwareramblings.com/2009/11/codekata-kata-two-solution.html
*/

func chop(target int, list []int) int {
	// 0 items
	if len(list) == 0 {
		return -1
	}

	// 1 item
	if len(list) == 1 {
		if target == list[0] {
			return 0
		} else {
			return -1
		}
	}

	// more than 1 items
	half := len(list) / 2
	top := NewChunk(list[0:half])
	bottom := NewChunk(list[half:])
	done := make(chan bool)

	// it chops the chunks or it gets the hose...
	chopping := func(chunk *Chunk, done chan bool) {
		chunk.result = chop(target, chunk.list)
		done <- true
	}

	go chopping(top, done)
	go chopping(bottom, done)

	// wait for both goroutines
	<-done
	<-done

	if top.hasValue() {
		return top.result
	}

	if bottom.hasValue() {
		return bottom.result + half
	}

	return -1
}

// just for fun... Chunks!
type Chunk struct {
	list   []int
	result int
}

func NewChunk(list []int) *Chunk {
	return &Chunk{list, -1}
}

func (c *Chunk) hasValue() bool {
	return c.result > -1
}

func assertEqual(expected, actual int) {
	if expected == actual {
		fmt.Printf("pass\n")
	} else {
		fmt.Printf("fail. expected %v to be %v\n", actual, expected)
	}
}

// synchronous version:
func simple_chop(target int, list []int) int {
	// 0 items
	if len(list) == 0 {
		return -1
	}

	// 1 item
	if len(list) == 1 {
		if target == list[0] {
			return 0
		} else {
			return -1
		}
	}

	half := len(list) / 2

	pos := chop(target, list[0:half])
	if pos != -1 {
		return pos
	}

	pos = chop(target, list[half:])
	if pos != -1 {
		return pos + (half)
	}

	// not found
	return -1
}
