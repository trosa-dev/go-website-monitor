package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

const monitoring = 10
const delay = 10

func main() {
	// Print a message indicating the start of the application
	fmt.Println("Starting application...")

	showMenu()

	selectedOption := readInput()
	fmt.Println("")

	// Use a switch statement to perform actions based on the user's choice
	switch selectedOption {
	case 1:
		startMonitoring()
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

func showMenu() {
	// Display a menu for the user to choose an option
	fmt.Println("Choose an option:")
	fmt.Println("1 - Start monitoring")
	fmt.Println("2 - Display logs")
	fmt.Println("3 - Exit")
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

func startMonitoring() {
	fmt.Println("Monitoring....")

	sitesUlr := []string{"https://www.google.com.br", "https://www.netflix.com"}

	for i := 0; i < monitoring; i++ {
		for i, urlSite := range sitesUlr {
			fmt.Println("Testing site ", i+1)
			testSite(urlSite)
		}
		time.Sleep(delay * time.Second)
		fmt.Println("")
	}
}

func testSite(urlSite string) {
	res, err := http.Get(urlSite)

	if err != nil {
		fmt.Println("Something went wrong...")
		fmt.Println(err)
		os.Exit(0)
	}

	statusCode := res.StatusCode

	switch res.StatusCode {
	case 200:
		fmt.Println("Site", urlSite, "is online!")
	default:
		fmt.Println("Something went wrong! site:", urlSite, ". status code: ", statusCode)
	}
}
