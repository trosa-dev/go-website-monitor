package main

import "fmt"

func main() {
    // Print a message indicating the start of the application
    fmt.Println("Starting application...")

    // Display a menu for the user to choose an option
    fmt.Println("Choose an option:")
    fmt.Println("1 - Start monitoring")
    fmt.Println("2 - Display logs")
    fmt.Println("3 - Exit")

    // Declare a variable to store the user's selected option
    var selectedOption int

    // Prompt the user to input their choice and store it in the selectedOption variable
    fmt.Scan(&selectedOption)

    // Print the selected option to the console
    fmt.Println("Selected option: ", selectedOption)

    // Use a switch statement to perform actions based on the user's choice
    switch selectedOption {
    case 1:
        fmt.Println("Monitoring....")
    case 2:
        fmt.Println("Displaying logs")
    case 3:
        // Exit the program for option 3
        break
    default:
        fmt.Println("Selected option is invalid", selectedOption)
    }
}
