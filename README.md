# Tic-Tac-Toe in Go

This is a simple implementation of the classic game Tic-Tac-Toe in Go. It features an unbeatable AI opponent that uses the minimax algorithm. The project also includes OpenTelemetry integration for debugging and tracing.

## Codebase Structure

The codebase is organized into the following packages:

-   `main`: The entry point of the application.
-   `board`: Contains the game board logic, including functions for checking for a win or draw condition.
-   `computer`: Implements the AI opponent's logic, including the minimax algorithm.
-   `player`: Handles the human player's input and interaction with the game.

## Getting Started

### Prerequisites

-   Go 1.23 or later.

### Installation

1.  Clone the repository:

    ```bash
    git clone https://github.com/sreconsulting/tic-tac-toe.git
    cd tic-tac-toe
    ```

2.  Install the dependencies:

    ```bash
    go mod download
    ```

### Running the Game

To run the game, execute the following command:

```bash
go run main.go
```

### Building the Executable

To build the executable, run the following command:

```bash
go build -o ttt main.go
```

This will create an executable file named `ttt` in the root directory of the project.

### Running Tests

To run the tests, execute the following commands:

```bash
go test ttt/board
go test ttt/computer
```

## Debugging with OpenTelemetry

This project uses OpenTelemetry for debugging and tracing. To use it, you'll need a free [Honeycomb](https://www.honeycomb.io) account.

Once you have an account, set the following environment variables in your terminal, replacing `<YOUR_API_KEY_HERE>` with your Honeycomb API key:

```bash
export OTEL_SERVICE_NAME="tic-tac-toe"
export OTEL_EXPORTER_OTLP_PROTOCOL="http/protobuf"
export OTEL_EXPORTER_OTLP_ENDPOINT="https://api.honeycomb.io"
export OTEL_EXPORTER_OTLP_HEADERS="x-honeycomb-team=<YOUR_API_KEY_HERE>"
```

If you don't set these variables, you may encounter errors.

## Contributing

Contributions are welcome! Please feel free to open an issue or submit a pull request.

## Contact

For questions or comments, please email me at schuyler@sreconsulting.io.
