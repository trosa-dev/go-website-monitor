package readFile

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// Read reads the contents of a file named "sites.txt" and returns them as a slice of strings
func Read() []string {
	// Open the file named "sites.txt" for reading
	file, err := os.Open("sites.txt")

	// Check for errors when opening the file
	if err != nil {
		fmt.Println("Something went wrong...")
		fmt.Println(err)
		os.Exit(0)
	}

	// Initialize an empty slice to store the read lines
	var sites []string

	// Create a reader to read the file line by line
	reader := bufio.NewReader(file)

	// Loop through each line in the file
	for {
		// Read a line from the file
		line, err := reader.ReadString('\n')

		// Trim any leading or trailing whitespaces from the line
		line = strings.TrimSpace(line)

		// Append the cleaned line to the sites slice
		sites = append(sites, line)

		// Print the current state of the sites slice (for debugging)
		fmt.Println(sites)

		// Check for the end of file (EOF) and break the loop if reached
		if err == io.EOF {
			break
		}

		// Check for other errors during file reading
		if err != nil {
			fmt.Println("Something went wrong...")
			fmt.Println(err)
			os.Exit(0)
		}
	}

	// Close the file after reading
	file.Close()

	// Return the slice containing the read lines from the file
	return sites
}
