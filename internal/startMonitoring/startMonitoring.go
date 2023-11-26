package startMonitoring

import (
	"fmt"
	"time"
	"website-monitor/internal/internal/readFile"
	"website-monitor/internal/internal/testSite"
)

const monitoring = 10
const delay = 10

func Start() {
	fmt.Println("Monitoring....")

	sitesUlr := readFile.Read()

	for i := 0; i < monitoring; i++ {
		for i, urlSite := range sitesUlr {
			fmt.Println("Testing site ", i+1)
			testSite.Test(urlSite)
		}
		time.Sleep(delay * time.Second)
		fmt.Println("")
	}
}
