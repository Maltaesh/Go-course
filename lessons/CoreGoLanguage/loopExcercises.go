package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {

	//Guess a number between 0 and 50.
	//Ex. Random = 25
	//User: 20
	//Higher
	//User: 40
	//Lover

	numOfTries := 0
	maxNumOfTries := 5
	seedSecs := time.Now().Unix()
	rand.Seed(seedSecs)
	randNum := rand.Intn(50) + 1

	for {
		if numOfTries >= maxNumOfTries {
			fmt.Println("To many tries!")
		}

		fmt.Print("Guess number between 1-50: ")
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')

		if err != nil {
			log.Fatal("Something went wrong")
		}

		input = strings.TrimSpace(input)

		inputNumber, err := strconv.Atoi(input)

		if err != nil {
			log.Fatal("Something went wrong")
		}

		if randNum != inputNumber {
			numOfTries++
			fmt.Println("Left chances: ", maxNumOfTries-numOfTries)
			if inputNumber > randNum {
				fmt.Println("Lower!")
			} else {
				fmt.Println("Higher!")
			}
		} else {
			fmt.Println("Good Guess! You win!")
			break
		}
	}
}
