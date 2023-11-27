package showLogs

import (
	"fmt"
	"os"
)

// Show reads and displays the contents of the "logs.txt" file
func Show() {
	// Read the contents of the "logs.txt" file
	file, err := os.ReadFile("logs.txt")

	// Check for errors when reading the file
	if err != nil {
		fmt.Println("Something went wrong...")
		fmt.Println(err)
		os.Exit(0)
	}

	// Print the contents of the file to the console
	fmt.Println(string(file))
}
