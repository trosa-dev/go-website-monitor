package saveLog

import (
	"fmt"
	"os"
	"strconv"
)

func Save(site string, status bool) int {
	// Declare a variable to store the user's selected option
	var selectedOption int

	file, err := os.OpenFile("logs.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("Something went wrong...")
		fmt.Println(err)
		os.Exit(0)
	}

	file.WriteString(site + " status: " + strconv.FormatBool(status) + "\n")

	file.Close()

	return selectedOption
}
