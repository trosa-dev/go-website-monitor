package startMonitoring

import (
	"fmt"
	"time"
	"website-monitor/internal/readFile"
	"website-monitor/internal/testSite"
)

// monitoring specifies the number of monitoring cycles
const monitoring = 10

// delay specifies the time delay between monitoring cycles in seconds
const delay = 10

// Start initiates the website monitoring process
func Start() {
	// Print a message indicating the start of the monitoring process
	fmt.Println("Monitoring....")

	// Read the list of sites to monitor from the "sites.txt" file
	sitesUlr := readFile.Read()

	// Loop through the monitoring cycles
	for i := 0; i < monitoring; i++ {
		// Loop through each site and test its status
		for i, urlSite := range sitesUlr {
			fmt.Println("Testing site ", i+1)
			testSite.Test(urlSite)
		}

		// Pause for the specified delay between monitoring cycles
		time.Sleep(delay * time.Second)
		fmt.Println("")
	}
}
