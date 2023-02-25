package main

import (
	"fmt"
)

func sayHello() {
	fmt.Println("Hello")
}

func getSum(x int, y int) int {
	return x + y
}

func getQuotient(x float64, y float64) (ans float64, err error) {
	if y == 0 {
		return 0, fmt.Errorf("You can't divide by 0")
	}

	return x / y, nil
}

func getTwo(x int) (int, int) {
	return x + 1, x + 2
}

func getSum2(nums ...int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

func getArraySum(arr []int) int {
	sum := 0
	for _, num := range arr {
		sum += num
	}
	return sum
}

func changeVal(f3 int) int {
	f3 += 1
	return f3
}

func main() {
	//func funcName(parameters) returnType {
	//function body
	//}

	sayHello()
	getSum(2, 5)
	ans, err := getQuotient(5, 0)
	if err == nil {
		fmt.Printf("%f / %f = %f", 5.0, 0.0, ans)
	} else {
		fmt.Println(err)
	}

	ans, err = getQuotient(5, 2)
	if err == nil {
		fmt.Printf("%f / %f = %f", 5.0, 2.0, ans)
	} else {
		fmt.Println(err)
	}

	f1, f2 := getTwo(5)
	fmt.Printf("%d, %d\n", f1, f2)

	resultGetSum2 := getSum2(2, 3, 4, 1, 2, 3, 5, 2, 0)
	fmt.Printf("result of the getSum2: %d\n", resultGetSum2)

	vArr := []int{1, 2, 3, 4, 5, 2, 3, 0}
	resultGetArraySum := getArraySum(vArr)
	fmt.Printf("result of the getArraySum: %d\n", resultGetArraySum)

	//parameters of the functions is by default passed as a value, not refference
	f3 := 5
	fmt.Println("f3 before func changeVal: ", f3)
	changeVal(f3)
	fmt.Println("f3 after func changeVal: ", f3)

}
