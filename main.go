package main

import (
	"fmt"
	"os"
	"website-monitor/internal/menu"
	"website-monitor/internal/readInput"
	"website-monitor/internal/showLogs"
	"website-monitor/internal/startMonitoring"
)

func main() {
	// Print a message indicating the start of the application
	fmt.Println("Starting application...")

	// Show the menu to the user
	menu.ShowMenu()

	// Read the user's choice from the input
	selectedOption := readInput.Read()
	fmt.Println("")

	// Use a switch statement to perform actions based on the user's choice
	switch selectedOption {
	case 1:
		// If the user chooses option 1, start the monitoring process
		startMonitoring.Start()
	case 2:
		// If the user chooses option 2, display logs
		fmt.Println("Displaying logs")
		showLogs.Show()
	case 3:
		// If the user chooses option 3, exit the application
		fmt.Println("Bye! See you soon...")
		os.Exit(0)
	default:
		// If the user chooses an invalid option, display an error message
		fmt.Println("Selected option is invalid", selectedOption)
	}
}
