package main

import "fmt"

func changeVal2(myPointer *int) {
	*myPointer = 12
}

func dblArrVals(arr *[4]int) {
	for x := 0; x < 4; x++ {
		arr[x] *= 2
	}
}

func getAverage(nums ...float64) float64 {
	var sum float64 = 0.0
	var numSize float64 = float64(len(nums))
	for _, val := range nums {
		sum += val
	}

	return (sum / numSize)
}

func main() {
	f4 := 10

	fmt.Println("f4: ", f4)
	fmt.Println("f4 address: ", &f4)
	var f4Ptr *int = &f4
	fmt.Println("f4 address:", f4Ptr)
	fmt.Println("f4 value: ", *f4Ptr)

	*f4Ptr = 11

	fmt.Println("f4 value: ", *f4Ptr)

	fmt.Println("f4 before function with pointer argument: ", f4)
	changeVal2(&f4)
	fmt.Println("f4 after function with pointer argument: ", f4)

	pArr := [4]int{1, 2, 3, 4}
	dblArrVals(&pArr)
	fmt.Println(pArr)

	iSlice := []float64{11, 2, 3, 43, 2, 3}
	fmt.Printf("Average is: %3f\n", getAverage(iSlice...))

}
