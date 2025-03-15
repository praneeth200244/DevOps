package main

import (
        "fmt"
        "log"
        "os"
        "strconv"
)

func main() {
        // Get the day number from the environment variable
        dayNum := os.Getenv("DAY_NUMBER")
        if dayNum == "" {
                log.Fatal("Error: DAY_NUMBER environment variable not set.")
        }

        day, err := strconv.Atoi(dayNum)
        if err != nil || day < 1 || day > 7 {
                log.Fatal("Error: Please provide a valid number between 1 and 7.")
        }

        days := []string{
                "Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday",
        }

        fmt.Printf("The day corresponding to number %d is %s.\n", day, days[day-1])
}

