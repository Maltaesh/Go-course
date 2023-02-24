package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	now := time.Now()

	fmt.Println(now.Year(), now.Month(), now.Day())
	fmt.Println(now.Hour(), now.Minute(), now.Second())

	loc, err := time.LoadLocation("America/New_York")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Time in New York %s\n", now.In(loc))

	loc, _ = time.LoadLocation("Asia/Shanghai")

	fmt.Printf("Time in Shanghai %s\n", now.In(loc))

	locEST, _ := time.LoadLocation("EST")
	locUTC, _ := time.LoadLocation("UTC")
	locMST, _ := time.LoadLocation("MST")

	fmt.Printf("EST: %s\n", now.In(locEST))
	fmt.Printf("UTC: %s\n", now.In(locUTC))
	fmt.Printf("MST: %s\n", now.In(locMST))

	birthDate := time.Date(1991, time.December, 12, 1, 00, 00, 00, time.Local)

	diff := now.Sub(birthDate)
	fmt.Printf("Days Alive: %d days\n", int(diff.Hours()/24))
	fmt.Printf("Hours Alive: %d hours\n", int(diff.Hours()))
}
