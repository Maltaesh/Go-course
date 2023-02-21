package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("First number: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')

	if err != nil {
		log.Fatal("Something went wrong!")
	}

	input = strings.TrimSpace(input)
	number, err := strconv.ParseFloat(input, 64)

	fmt.Print("Second number: ")
	reader = bufio.NewReader(os.Stdin)
	input2, err := reader.ReadString('\n')

	if err != nil {
		log.Fatal("Something went wrong!")
	}

	input2 = strings.TrimSpace(input2)
	number2, err := strconv.ParseFloat(input, 64)

	fmt.Println(number, "+", number2, "=", number+number2)
}
