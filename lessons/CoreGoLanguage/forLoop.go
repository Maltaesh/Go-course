package main

import "fmt"

func main() {
	// for itialzation; condition; postStatement {BODY of the loop}

	fmt.Println("For loop 1-5:")
	for x := 1; x <= 5; x++ {
		fmt.Println(x)
	}

	fmt.Println("For loop 5-1:")
	for x := 5; x >= 1; x-- {
		fmt.Println(x)
	}

	fmt.Println("While loop equivalent 0-5:")
	fX := 0
	for fX <= 5 {
		fmt.Println(fX)
		fX++
	}

	fmt.Println("Looping over arrays:")
	aNums := []int{1, 2, 3}
	for _, num := range aNums {
		fmt.Println(num)
	}

	fmt.Println("Do while loop 1-5:")
	xVal := 1
	for true {
		if xVal == 5 {
			break
		}

		fmt.Println(xVal)
		xVal++
	}
}
