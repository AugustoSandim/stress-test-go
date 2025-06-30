package loadtester

import (
	"net/http"
	"sync"
	"time"
)

type Result struct {
	StatusCode int
	Duration   time.Duration
}

type Report struct {
	TotalTime     time.Duration
	TotalRequests int
	StatusCounts  map[int]int
}

func Run(url string, requests int, concurrency int) *Report {
	var wg sync.WaitGroup
	results := make(chan Result, requests)
	jobs := make(chan struct{}, concurrency)

	startTime := time.Now()

	for i := 0; i < requests; i++ {
		wg.Add(1)
		jobs <- struct{}{}
		go func() {
			defer wg.Done()
			defer func() { <-jobs }()

			startReqTime := time.Now()
			resp, err := http.Get(url)
			duration := time.Since(startReqTime)

			if err != nil {
				results <- Result{StatusCode: 0, Duration: duration}
				return
			}
			defer resp.Body.Close()

			results <- Result{StatusCode: resp.StatusCode, Duration: duration}
		}()
	}

	wg.Wait()
	close(results)

	totalTime := time.Since(startTime)
	statusCounts := make(map[int]int)
	var totalRequests int

	for result := range results {
		totalRequests++
		statusCounts[result.StatusCode]++
	}

	return &Report{
		TotalTime:     totalTime,
		TotalRequests: totalRequests,
		StatusCounts:  statusCounts,
	}
}
