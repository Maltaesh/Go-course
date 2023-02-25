package main

import utils "example/project/mypackage"

func factorial(num uint64) uint64 {
	if num == 0 {
		return 1
	}

	return num * factorial(num-1)
}

func main() {
	// Recursion
	utils.Pl("Factorial 5:", factorial(60))
}
