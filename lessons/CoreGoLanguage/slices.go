package main

import (
	"fmt"
)

func main() {
	// var name []datatype

	sl1 := make([]string, 6)
	sl1[0] = "Society"
	sl1[1] = "of"
	sl1[2] = "the"
	sl1[3] = "Simulated"
	sl1[4] = "Universe"

	fmt.Println("Slice size: ", len(sl1))
	for i := 0; i < len(sl1); i++ {
		fmt.Println(sl1[i])
	}

	for _, v := range sl1 {
		fmt.Println(v)
	}

	sl2 := []int{12, 21, 1974}
	fmt.Println(sl2)

	//this is an array, on with slices will operate
	sArr := [5]int{1, 2, 3, 4, 5}
	//extracting 3 elements counted from the begening, using a slice
	sl3 := sArr[:3]
	fmt.Println(sl3)
	//extracting last 3 elements counted from the end, using a slice
	sl4 := sArr[2:]
	fmt.Println(sl4)

	//slices care if you change the array that they operate on
	sArr[0] = 10
	fmt.Println(sl3)

	sl3 = append(sl3, 12)
	fmt.Println(sl3)
	fmt.Println(sArr)

	sl5 := make([]string, 6)
	fmt.Println(sl5)
	fmt.Println(sl5[0])

}
