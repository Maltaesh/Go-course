package main

import (
	"fmt"
	"reflect"
	"strconv"
)

var pl = fmt.Println

func main() {
	// int, float64, bool, string, rune
	// Default types 0, 0.0, false, ""
	pl(reflect.TypeOf(25))
	cv1 := 1.5
	cv2 := int(cv1)
	pl(cv2)

	cv3 := "50000000"
	cv4, err := strconv.Atoi(cv3)
	pl(cv4, err, reflect.TypeOf(cv4))

	cv5 := 500000
	cv6 := strconv.Itoa(cv5)
	pl(cv6, reflect.TypeOf(cv6))

	cv7 := "3.14"
	if cv8, err := strconv.ParseFloat(cv7, 64); err == nil {
		pl(cv8)
	}

	cv9 := fmt.Sprintf("%f", 3.14)
	pl(cv9)
}
