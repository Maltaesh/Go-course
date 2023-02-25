package main

import (
	utils "example/project/mypackage"
	"reflect"
)

func main() {
	utils.Pl(utils.Name)
	intArr := []int{2, 3, 4, 5, 11}
	strArr := utils.IntArrToStrArr(intArr)
	utils.Pl(reflect.TypeOf(intArr))
	utils.Pl(reflect.TypeOf(strArr))
}
