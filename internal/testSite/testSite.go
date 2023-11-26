package testSite

import (
	"fmt"
	"net/http"
	"os"
	"website-monitor/internal/saveLog"
)

func Test(urlSite string) {
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
		saveLog.Save(urlSite, true)
	default:
		fmt.Println("Something went wrong! site:", urlSite, ". status code: ", statusCode)
		saveLog.Save(urlSite, false)
	}
}
