package utils

import (
	"fmt"
	"strconv"
)

var Name string = "My utils package"

var Pl = fmt.Println
var F = fmt.Printf

func IntArrToStrArr(intArr []int) []string {
	var strArr []string
	for _, i := range intArr {
		strArr = append(strArr, strconv.Itoa(i))
	}
	return strArr
}
