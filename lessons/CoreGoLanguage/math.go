package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("5 + 4 = ", 5+4)
	fmt.Println("5 - 4 = ", 5-4)
	fmt.Println("5 * 4 = ", 5*4)
	fmt.Println("5 / 4 = ", 5/4)
	fmt.Println("5 % 4 = ", 5%4)

	mInt := 1
	fmt.Println(mInt)
	mInt += 1
	fmt.Println(mInt)
	mInt = mInt + 1
	fmt.Println(mInt)
	mInt++
	fmt.Println(mInt)

	fmt.Println("Float Precision =", 0.111111111111111111111111111111+0.2222222222222222222222222222222)

	//Random values
	seedSecs := time.Now().Unix()
	rand.Seed(seedSecs)
	randNum := rand.Intn(50) + 1
	fmt.Println(randNum)

	fmt.Println("Absolute -10 =", math.Abs(-10))
	fmt.Println("4 to the power of 2 =", math.Pow(4, 2))
	fmt.Println("Square root of 16 =", math.Sqrt(16))
	fmt.Println("Cube root of 16 =", math.Cbrt(16))
	fmt.Println("Ceil 4.4 =", math.Ceil(4.4))
	fmt.Println("Floor 4.4 =", math.Floor(4.4))
	fmt.Println("Rounding 4.4 =", math.Round(4.4))
	fmt.Println("Binary Logarithm of 8 =", math.Log2(8))
	fmt.Println("Decimal Logarithm of 100 =", math.Log10(100))
	fmt.Println("Natural Logarithm of 7.389 =", math.Log(7.389))
	fmt.Println("Maximum of 5 and 4 =", math.Max(5, 4))
	fmt.Println("Minimum of 5 and 4 =", math.Min(5, 4))

	//Radians
	r90 := 90 * math.Pi / 180
	d90 := r90 * (180 / math.Pi)
	fmt.Printf("%f radians = %f degrees\n", r90, d90)

}
