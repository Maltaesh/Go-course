package main

import utils "example/project/mypackage"

type MyConstraint interface {
	int | float64
}

func getSumGen[T MyConstraint](x T, y T) T {
	return x + y
}

func main() {
	utils.Pl("5 + 4 =", getSumGen(5, 4))
	utils.Pl("5.5 + 4.2 =", getSumGen(5.5, 4.2))
}
