package main

import (
	"fmt"
	"os"
)

func main() {
	// Print a message indicating the start of the application
	fmt.Println("Starting application...")

	showMenu()

	selectedOption := readInput()

	// Use a switch statement to perform actions based on the user's choice
	switch selectedOption {
	case 1:
		fmt.Println("Monitoring....")
	case 2:
		fmt.Println("Displaying logs")
	case 3:
		fmt.Println("Bye! See you soon...")
		os.Exit(0)
		break
	default:
		fmt.Println("Selected option is invalid", selectedOption)
	}
}

func readInput() int {
	// Declare a variable to store the user's selected option
	var selectedOption int

	// Prompt the user to input their choice and store it in the selectedOption variable
	fmt.Scan(&selectedOption)

	// Print the selected option to the console
	fmt.Println("Selected option: ", selectedOption)

	return selectedOption
}

func showMenu() {
	// Display a menu for the user to choose an option
	fmt.Println("Choose an option:")
	fmt.Println("1 - Start monitoring")
	fmt.Println("2 - Display logs")
	fmt.Println("3 - Exit")
}
