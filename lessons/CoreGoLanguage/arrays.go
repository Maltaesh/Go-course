package main

import "fmt"

func main() {
	var arr1 [5]int
	fmt.Println("Array of size 5, type int, after initialization: ", arr1)
	arr1[0] = 1
	fmt.Println("Array of size 5, type int, after adding element at index 0: ", arr1)

	arr2 := [5]int{1, 2, 3, 4, 5}
	fmt.Println("Array of size 5, type int, filled with elements during initialization: ", arr2)

	fmt.Println("Array of size 5, type int, index 4: ", arr2[4])

	fmt.Println("Iterating over array:")
	fmt.Println("Classic for loop:")
	for i := 0; i < len(arr2); i++ {
		fmt.Println(arr2[i])
	}

	fmt.Println("For loop with range:")
	for i, v := range arr2 {
		fmt.Printf("%d : %d\n", i, v)
	}

	fmt.Println("Multidimensional array:")
	arr3 := [2][2]int{
		{1, 2},
		{3, 4},
	}

	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			fmt.Println(arr3[i][j])
		}
	}

	aStr1 := "abcde"
	rArr := []rune(aStr1)
	for _, v := range rArr {
		fmt.Printf("Rune Array: %d\n", v)
	}

	byteArr := []byte{'a', 'b', 'c'}
	bStr := string(byteArr[:])
	fmt.Println("I'am a string: ", bStr)
}
