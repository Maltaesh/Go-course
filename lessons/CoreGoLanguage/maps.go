package main

import (
	utils "example/project/mypackage"
)

func main() {
	//maps are collection of key - value pairs
	//key can be any data type, that can be compared with the == sign
	//var myMap map[keyType]valueType

	heroes := make(map[string]string)
	villians := make(map[string]string)

	heroes["Batman"] = "Bruce Wayne"
	heroes["Superman"] = "Clark Kent"
	heroes["The Flash"] = "Barry Allen"

	villians["Lex Luther"] = "Lex Luther"

	superPets := map[int]string{1: "Krypto", 2: "Bat Hound"}

	utils.F("Batman is %v\n", heroes["Batman"])

	utils.Pl("Chip :", superPets[3])
	_, ok := superPets[3]
	utils.Pl("Is there a 3rd pet: ", ok)

	for k, v := range heroes {
		utils.F("%s is %s\n", k, v)
	}

	utils.Pl("Deleting insane hero - Flash. Now heroes map has:")
	delete(heroes, "The Flash")
	for k, v := range heroes {
		utils.F("%s is %s\n", k, v)
	}

}
