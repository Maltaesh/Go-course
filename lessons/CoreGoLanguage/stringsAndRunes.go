package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	// Strings
	sV1 := "A word"
	replacer := strings.NewReplacer("A", "Another")
	sV2 := replacer.Replace(sV1)
	fmt.Println(sV2)

	fmt.Println("Length: ", len(sV2))
	fmt.Println("Contains Another: ", strings.Contains(sV2, "Another"))
	fmt.Println("o index: ", strings.Index(sV2, "o"))
	fmt.Println("Replace: ", strings.Replace(sV2, "o", "0", -1))

	sV3 := "\nSome Words \n"
	fmt.Println(sV3)
	sV3 = strings.TrimSpace(sV3)
	fmt.Println(sV3)

	fmt.Println("Split: ", strings.Split("a-r-c-y-m-a-g", "-"))
	fmt.Println("Lower: ", strings.ToLower(sV2))
	fmt.Println("Upper: ", strings.ToUpper(sV2))
	fmt.Println("Prefix: ", strings.HasPrefix("tacocat", "taco"))
	fmt.Println("Suffix: ", strings.HasSuffix("tacocat", "taco"))

	// Runes
	rStr := "abcdefg"
	fmt.Println("Rune count: ", utf8.RuneCountInString(rStr))

	for i, runeVal := range rStr {
		fmt.Printf("%d : %#U : %c\n", i, runeVal, runeVal)
	}
}
