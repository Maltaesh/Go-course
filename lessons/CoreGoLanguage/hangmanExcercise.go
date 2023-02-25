package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"regexp"
	"strings"
	"time"
)

/*

  +---+
  0   |
 /|\  |
 / \  |
     ===


Secret Word: M_N___

Incorrect Guesses: A

Guess a Letter: Y

Sorry Your Dead! The word is MONKEY
Yes the Secret Word is MONKEY

Please enter only one letter
Please enter a letter
Please a enter you haven't already guessed


*/

var hmArr = [7]string{
	"  +---+\n" +
		"      |\n" +
		"      |\n" +
		"      |\n" +
		"     ===\n",
	"  +---+\n" +
		"  0   |\n" +
		"      |\n" +
		"      |\n" +
		"     ===\n",
	"  +---+\n" +
		"  0   |\n" +
		"  |   |\n" +
		"      |\n" +
		"     ===\n",
	"  +---+\n" +
		"  0   |\n" +
		" /|   |\n" +
		"      |\n" +
		"     ===\n",
	"  +---+\n" +
		"  0   |\n" +
		" /|\\  |\n" +
		"      |\n" +
		"     ===\n",
	"  +---+\n" +
		"  0   |\n" +
		" /|\\  |\n" +
		" /    |\n" +
		"     ===\n",
	"  +---+\n" +
		"  0   |\n" +
		" /|\\  |\n" +
		" / \\  |\n" +
		"     ===\n",
}

var wordArr = [7]string{
	"JAZZ", "ZIGZAG", "ZILCH", "ZIPPER", "ZODIAC", "ZOMBIE", "FLUFF",
}

var randWord string
var guessedLetters string
var correctLetters []string
var wrongGuesses []string

func getRandWord() string {
	seedSecs := time.Now().Unix()
	rand.Seed(seedSecs)
	randWord = wordArr[rand.Intn(7)]
	correctLetters = make([]string, len(randWord))
	return randWord
}

func showBoard() {
	fmt.Println(hmArr[len(wrongGuesses)])
	fmt.Print("Secret Word: ")
	for _, v := range correctLetters {
		if v == "" {
			fmt.Print("_")
		} else {
			fmt.Print(v)
		}
	}
	if len(wrongGuesses) > 0 {
		fmt.Print("\nIncorrect Guesses: ")
		for _, v := range wrongGuesses {
			fmt.Print(v + " ")
		}
	}
	fmt.Println()
}

func getUserLetter() string {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("\nGuess a letter: ")
		guess, err := reader.ReadString('\n')

		if err != nil {
			log.Fatal(err)
		}

		guess = strings.ToUpper(guess)
		guess = strings.TrimSpace(guess)
		var IsLetter = regexp.MustCompile(`^[a-zA-Z]+$`).MatchString

		if len(guess) != 1 {
			fmt.Println("Please enter only one letter")
		} else if !IsLetter(guess) {
			fmt.Println("Please enter only letters A-Z/a-z")
		} else if strings.Contains(guessedLetters, guess) {
			fmt.Println("Please enter a letter you haven't guesses")
		}
		return guess
	}
}

func getAllIndexes(theStr, subStr string) (indices []int) {
	if len(subStr) == 0 || len(theStr) == 0 {
		return indices
	}

	offset := 0
	for {
		i := strings.Index(theStr[offset:], subStr)
		if i == -1 {
			return indices
		}
		offset += i
		indices = append(indices, offset)
		offset += len(subStr)
	}
}

func updateCorrectLetters(letter string) {
	indexMatches := getAllIndexes(randWord, letter)
	for _, v := range indexMatches {
		correctLetters[v] = letter
	}
}

func sliceHasEmptys(theSlice []string) bool {
	for _, v := range theSlice {
		if len(v) == 0 {
			return true
		}
	}

	return false
}

func main() {
	fmt.Println(getRandWord())

	for {
		showBoard()

		guess := getUserLetter()

		if strings.Contains(randWord, guess) {
			updateCorrectLetters(guess)

			if sliceHasEmptys(correctLetters) {
				fmt.Println("More letters to guess")
			} else {
				fmt.Println("Yes, the secret word is: ", randWord)
				break
			}
		} else {
			guessedLetters += guess
			wrongGuesses = append(wrongGuesses, guess)

			if len(wrongGuesses) >= 6 {
				fmt.Println("Sorry, to many trys, you're died! The secret word is: ", randWord)
				break
			}
		}
	}
}
