package main

import (
	"fmt"
	"os"
	"website-monitor/internal/menu"
	"website-monitor/internal/readInput"
	"website-monitor/internal/startMonitoring"
)

func main() {
	// Print a message indicating the start of the application
	fmt.Println("Starting application...")

	menu.ShowMenu()

	selectedOption := readInput.Read()
	fmt.Println("")

	// Use a switch statement to perform actions based on the user's choice
	switch selectedOption {
	case 1:
		startMonitoring.Start()
	case 2:
		fmt.Println("Displaying logs")
	case 3:
		fmt.Println("Bye! See you soon...")
		os.Exit(0)
	default:
		fmt.Println("Selected option is invalid", selectedOption)
	}
}
