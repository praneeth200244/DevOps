package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	// Check if the program has received an argument
	if len(os.Args) != 2 {
		log.Fatal("Usage: <program> <day-number>")
	}

	// Read the day number from command line argument
	dayNum := os.Args[1]
	day, err := strconv.Atoi(dayNum)
	if err != nil || day < 1 || day > 7 {
		log.Fatal("Error: Please provide a valid number between 1 and 7.")
	}

	// Define the days of the week
	days := []string{
		"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday",
	}

	// Print the corresponding weekday
	fmt.Printf("The day corresponding to number %d is %s.\n", day, days[day-1])
}

