# Concurrency in Go - Race Condition Example

This project demonstrates the use of concurrency in Go, specifically showcasing the differences between handling race conditions with and without proper synchronization mechanisms such as WaitGroup and Mutex. The project includes an API to illustrate these concepts.


## Getting Started

### Prerequisites

- Go installed on your system

### Running the Project

1. **Clone the repository:**

2. **Navigate to the project directory:**

    ```sh
    cd go_race_condition
    ```

3. **Run the server:**

    ```sh
    go run main.go
    ```

The server will run on port 8080.

### API Endpoints

- **GET /get-stock:** Retrieves the current stock value.

    Example:

    ```sh
    curl http://localhost:8080/stock
    ```

- **POST /update-stock-wg:** Updates the stock using WaitGroup (prone to race conditions).

    Example:

    ```sh
    curl -X POST http://localhost:8080/update-stock-wg
    ```

- **POST /update-stock-mutex:** Updates the stock using Mutex (safe from race conditions).

    Example:

    ```sh
    curl -X POST http://localhost:8080/update-stock-mutex
    ```

# Race Condition Case
clone the test repo "go_race_condition_test" and run the program