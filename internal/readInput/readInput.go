package readInput

import "fmt"

// Read reads the user's input for an integer and returns the selected option
func Read() int {
	// Declare a variable to store the user's selected option
	var selectedOption int

	// Prompt the user to input their choice and store it in the selectedOption variable
	fmt.Scan(&selectedOption)

	// Print the selected option to the console
	fmt.Println("Selected option: ", selectedOption)

	// Return the selected option
	return selectedOption
}
