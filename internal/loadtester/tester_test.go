package loadtester

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRun(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	requests := 100
	concurrency := 10
	report := Run(server.URL, requests, concurrency)

	if report.TotalRequests != requests {
		t.Errorf("expected %d requests, but got %d", requests, report.TotalRequests)
	}

	if report.StatusCounts[http.StatusOK] != requests {
		t.Errorf("expected %d requests with status 200, but got %d", requests, report.StatusCounts[http.StatusOK])
	}
}
