package readFile

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func Read() []string {
	file, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Something went wrong...")
		fmt.Println(err)
		os.Exit(0)
	}

	var sites []string

	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')

		line = strings.TrimSpace(line)
		sites = append(sites, line)

		fmt.Println(sites)

		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Println("Something went wrong...")
			fmt.Println(err)
			os.Exit(0)
		}
	}

	file.Close()

	return sites
}
