package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	// Read first number
	fmt.Print("Enter first number: ")
	var num1 string
	fmt.Scanln(&num1)
	// Convert to float64
	n1, err := strconv.ParseFloat(num1, 64)
	if err != nil {
		log.Fatal(err)
	}

	// Read operator
	fmt.Print("Enter operator (+, -, *, /): ")
	var operator string
	fmt.Scanln(&operator)

	// Read second number
	fmt.Print("Enter second number: ")
	var num2 string
	fmt.Scanln(&num2)
	// Convert to float64
	n2, err := strconv.ParseFloat(num2, 64)
	if err != nil {
		log.Fatal(err)
	}

	// Perform calculation based on the operator
	var result float64
	switch operator {
	case "+":
		result = n1 + n2
	case "-":
		result = n1 - n2
	case "*":
		result = n1 * n2
	case "/":
		if n2 == 0 {
			fmt.Println("Error: Division by zero is not allowed.")
			os.Exit(1)
		}
		result = n1 / n2
	default:
		fmt.Println("Error: Invalid operator.")
		os.Exit(1)
	}

	// Output the result
	fmt.Printf("Result: %.2f\n", result)
}


