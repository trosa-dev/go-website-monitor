package saveLog

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

// Save writes a log entry with the current date and time, site, and status to a file named "logs.txt"
func Save(site string, status bool) int {
	// Declare a variable to store the user's selected option
	var selectedOption int

	// Open or create the "logs.txt" file for reading, writing, and appending
	file, err := os.OpenFile("logs.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	// Check for errors when opening or creating the file
	if err != nil {
		fmt.Println("Something went wrong...")
		fmt.Println(err)
		os.Exit(0)
	}

	// Write a log entry with the current date and time, site, and status to the file
	// https://pkg.go.dev/time#Time.Format
	file.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " status: " + strconv.FormatBool(status) + "\n")

	// Close the file after writing the log entry
	file.Close()

	// Return the selected option (which is not used in the current implementation)
	return selectedOption
}
