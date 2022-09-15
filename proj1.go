package main

import (
	"fmt"
)

type gender int64

type manatee struct {
	id int
	age int
	weight int
	mf gender
}

func newManatee(identifier int, age int, weight int, mf gender) manatee {
	p := manatee{id: identifier, age: age, weight: weight, mf: mf}
	return p
}

func main() {

	const(
		male gender = iota
		female
	)

	const numManatees int = 4
	var i, j, k int = 0, 0, 1
	var stack []manatee
	var input = [numManatees][numManatees]int{
		{3, 2, 1, 2} ,
		{2, 3, 4, 3} ,
		{2, 1, 2, 1} ,
		{2, 2, 1, 3} }

	for i = 0; i < numManatees; i++ {
		for j = 0; j < numManatees; j++ {
			// need to create the manatees, and figure exactly how to step through input array.
			k++
			fmt.Print(input[i][j])
		}
	}


	manatee1 := newManatee(1, 12, 2, male)
	manatee2 := newManatee(2, 13, 3, female)

	stack = append(stack, manatee2)
	stack = append(stack, manatee1)

	for len(stack) > 0 {
		n := len(stack) - 1 // size of the stack is 2, therefore n is 1
		fmt.Print(stack[n]) // print element at 1

		stack = stack[:n] // slice the stack (in effect removing element at 1) now stack has size 1, n is 0.
	}

}
