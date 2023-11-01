// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

// ---------------------------------------------------------
// EXERCISE: Days in a Month
//
//	Print the number of days in a given month.
//
// RESTRICTIONS
//
//  1. On a leap year, february should print 29. Otherwise, 28.
//
//     Set your computer clock to 2020 to see whether it works.
//
//  2. It should work case-insensitive. See below.
//
//     Search on Google: golang pkg strings ToLower
//
//  3. Get the current year using the time.Now()
//
//     Search on Google: golang pkg time now year
//
// EXPECTED OUTPUT
//
//	-----------------------------------------
//	Your solution should not accept invalid months
//	-----------------------------------------
//	go run main.go
//	  Give me a month name
//
//	go run main.go sheep
//	  "sheep" is not a month.
//
//	-----------------------------------------
//	Your solution should handle the leap years
//	-----------------------------------------
//	go run main.go january
//	  "january" has 31 days.
//
//	go run main.go february
//	  "february" has 28 days.
//
//	go run main.go march
//	  "march" has 31 days.
//
//	go run main.go april
//	  "april" has 30 days.
//
//	go run main.go may
//	  "may" has 31 days.
//
//	go run main.go june
//	  "june" has 30 days.
//
//	go run main.go july
//	  "july" has 31 days.
//
//	go run main.go august
//	  "august" has 31 days.
//
//	go run main.go september
//	  "september" has 30 days.
//
//	go run main.go october
//	  "october" has 31 days.
//
//	go run main.go november
//	  "november" has 30 days.
//
//	go run main.go december
//	  "december" has 31 days.
//
//	-----------------------------------------
//	Your solution should be case insensitive
//	-----------------------------------------
//	go run main.go DECEMBER
//	  "DECEMBER" has 31 days.
//
//	go run main.go dEcEmBeR
//	  "dEcEmBeR" has 31 days.
//
// ---------------------------------------------------------

// Returns number of days in given month of current year as int. Format month as in "January"
func daysInMonth(m string) int {
	// change user month into generic time
	mTime, err := time.Parse("January", strings.ToTitle(strings.ToLower(m)))
	if err != nil {
		log.Fatal("Cannot parse '", m, "' as month\n")
	}
	// format generic time into time with the month of current year, on the 1st
	t := time.Date(time.Now().Year(), mTime.Month(), 1, 0, 0, 0, 0, time.UTC)
	// add month to that, to calculate length of month
	t2 := time.Date(time.Now().Year(), mTime.AddDate(0, 1, 0).Month(), 1, 0, 0, 0, 0, time.UTC)
	// calc diff between the 2 dates
	diff := t2.Sub(t)
	// divide by 24 to get days from hours, as int
	days := int(diff.Hours() / 24)
	return days
}

func main() {
	if len(os.Args) == 2 {
		num := daysInMonth(os.Args[1])
		fmt.Printf("%q has %d days.\n", os.Args[1], num)
	} else {
		fmt.Println("Give me a month name")
	}
}
