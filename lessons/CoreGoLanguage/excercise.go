package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var pl = fmt.Println

func main() {
	// Receive customer data (their age)
	// Google hot to trim whitespace from input
	// Age < 5 to young for school
	// age == 5 go to kindergarden
	// age > 5 or age <= 17 go to grade
	//default go to college

	fmt.Print("What is your name? ")
	reader := bufio.NewReader(os.Stdin)
	name, err := reader.ReadString('\n')
	if err == nil {
		pl("Hello, ", strings.TrimSpace(name))
	} else {
		log.Fatal("Oops! Something went wrong!")
	}

	fmt.Print("How old are you? ")
	reader = bufio.NewReader(os.Stdin)
	age, err := reader.ReadString('\n')
	age = strings.TrimSpace(age)
	if err == nil {
		ageInt, err := strconv.Atoi(age)

		if err == nil {
			if ageInt < 5 {
				pl("You're to young for school")
			} else if ageInt == 5 {
				pl("You're going to kindergarden!")
			} else {
				if ageInt <= 13 {
					pl("You're going to elementary school")
				} else if ageInt <= 17 {
					pl("You're going to high school")
				} else {
					pl("You're going to the college")
				}
			}
		} else {
			log.Fatal("Something went wrong with check age section")
		}
	} else {
		log.Fatal("Something went wrong with age question")
	}

}
