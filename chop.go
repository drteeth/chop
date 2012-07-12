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
  top := make(chan int)
  bottom := make(chan int)

  go func() {
    top <- chop(target, list[0:half])
  }

  go func() {
    bottom <- chop(target, list[half:])
  }

  // TODO use a select to return the first result
  // wait for the top half
  t := <- top
  if t > -1 {
    return t
  }

  //
  b := <- bottom

  if top > -1 {
    return top
  }

  if bottom > -1 {
    return bottom + half
  }

  return -1
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
