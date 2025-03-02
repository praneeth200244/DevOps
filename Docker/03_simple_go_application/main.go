package main

import (
	"fmt"
)

func main() {
	// Get user's name
	var name string
	fmt.Print("Enter your name: ")
	fmt.Scanln(&name)

	// Display a crazy welcome message
	fmt.Printf("🎉🌪️🌟🚀 Welcome to Docker, %s! 🐳💥✨ It's time to spin up containers and make some magic happen! 🔥🌍\n", name)
}
