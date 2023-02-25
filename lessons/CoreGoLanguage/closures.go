package main

import (
	utils "example/project/mypackage"
)

func useFunction(f func(int, int) int, x, y int) {
	utils.Pl("Answer:", (f(x, y)))
}

func sumVals(x, y int) int {
	return x + y
}

func main() {
	intSum := func(x, y int) int {
		return x + y
	}

	utils.Pl("5 + 4 =", intSum(5, 4))

	sample1 := 1
	changeVar := func() { sample1 += 1 }
	changeVar()
	utils.Pl("Sample1: ", sample1)

	useFunction(sumVals, 5, 8)
}
