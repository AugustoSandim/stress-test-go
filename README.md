# Go Stress Test

This is a command-line interface (CLI) tool for load testing web services, built with Go.

## Features

- Perform a specified number of HTTP requests to a target URL.
- Control the level of concurrency for simultaneous requests.
- Generate a report with key metrics, including total time, total requests, and status code distribution.

## Prerequisites

- [Docker](https://www.docker.com/get-started)

## Getting Started

This project is designed to be run as a Docker container. It uses Go for the implementation and provides a simple CLI interface for load testing.

## Cloning the Repository

To get started, clone the repository:

```sh
git clone git@github.com:AugustoSandim/stress-test-go.git
cd stress-test-go
```

## How to Run

1.  **Build the Docker image:**

    ```sh
    docker build -t stress-test-go .
    ```

2.  **Run the load test:**

    Use the `docker run` command with the desired flags to start the test.

    ```sh
    docker run stress-test-go --url <TARGET_URL> --requests <TOTAL_REQUESTS> --concurrency <CONCURRENT_REQUESTS>
    ```

### Example

```sh
docker run stress-test-go --url=http://google.com --requests=1000 --concurrency=10
# OR
docker run stress-test-go --url http://google.com --requests 1000 --concurrency 10
```

## Command-Line Flags

- `--url`: The URL of the service to be tested. (Required)
- `--requests`: The total number of requests to be made. (Required)
- `--concurrency`: The number of simultaneous requests. (Default: 1)

## Project Structure

- `cmd/stress-test`: Contains the entry point of the application (`main.go`).
- `internal/loadtester`: Contains the core logic for running the load test and generating reports.
- `Dockerfile`: Defines the steps to build the Docker image.
- `go.mod`, `go.sum`: Manage the project's dependencies.
- `README.md`: Provides instructions and documentation for the project.
