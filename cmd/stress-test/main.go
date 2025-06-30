package main

import (
	"flag"
	"fmt"
	"os"
	"stress-test-go/internal/loadtester"
)

func main() {
	url := flag.String("url", "", "URL of the service to be tested")
	requests := flag.Int("requests", 0, "Total number of requests")
	concurrency := flag.Int("concurrency", 1, "Number of simultaneous calls")
	flag.Parse()

	if *url == "" || *requests == 0 {
		fmt.Println("URL and requests are required")
		flag.Usage()
		os.Exit(1)
	}

	report := loadtester.Run(*url, *requests, *concurrency)

	fmt.Printf("Total time spent: %v\n", report.TotalTime)
	fmt.Printf("Total requests made: %d\n", report.TotalRequests)
	fmt.Printf("Requests with status 200: %d\n", report.StatusCounts[200])
	fmt.Println("Distribution of other status codes:")
	for status, count := range report.StatusCounts {
		if status != 200 {
			fmt.Printf("Status %d: %d\n", status, count)
		}
	}
}
