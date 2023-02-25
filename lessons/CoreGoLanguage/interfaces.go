package main

import (
	utils "example/project/mypackage"
)

type Animal interface {
	AngrySound()
	HappySound()
}

type Cat string

func (c Cat) Attack() {
	utils.Pl("Cat Attack its Prey")
}

func (c Cat) Name() string {
	return string(c)
}

func (c Cat) AngrySound() {
	utils.Pl("Cat says Hisssss")
}

func (c Cat) HappySound() {
	utils.Pl("Cat says Mrrrr")
}

func main() {
	var kitty Animal
	kitty = Cat("Kitty")
	kitty.AngrySound()
	kitty.HappySound()

	var kitty2 = kitty.(Cat)
	utils.Pl(kitty2.Name())

	kitku := Cat("Kitku")
	utils.Pl(kitku.Name())

}
