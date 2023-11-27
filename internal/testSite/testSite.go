package testSite

import (
	"fmt"
	"net/http"
	"os"
	"website-monitor/internal/saveLog"
)

// Test performs a simple HTTP GET request to check the status of a website
func Test(urlSite string) {
	// Perform an HTTP GET request to the specified URL
	res, err := http.Get(urlSite)

	// Check for errors during the HTTP request
	if err != nil {
		fmt.Println("Something went wrong...")
		fmt.Println(err)
		os.Exit(0)
	}

	// Get the HTTP status code from the response
	statusCode := res.StatusCode

	// Use a switch statement to perform actions based on the HTTP status code
	switch res.StatusCode {
	case 200:
		// If the status code is 200, the site is online
		fmt.Println("Site", urlSite, "is online!")
		// Save a log entry indicating a successful test
		saveLog.Save(urlSite, true)
	default:
		// If the status code is different from 200, something went wrong
		fmt.Println("Something went wrong! site:", urlSite, ". status code: ", statusCode)
		// Save a log entry indicating a failed test
		saveLog.Save(urlSite, false)
	}
}
