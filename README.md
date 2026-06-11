A tic-tac-toe CLI built as an apprenticeship project. Implements the minimax algorithm to make the computer theoretically unbeatable, and uses OpenTelemetry for tracing (streamed to Honeycomb).

## Installation

Requires Go 1.23+. See [Golang install](https://go.dev/doc/install) for all platforms, or on macOS:

```bash
brew install golang
```

Dependencies are managed with Go modules — run `go mod download` to fetch them.

## Usage

```bash
make run      # run the game
make build    # build ./ttt executable
make test     # run all tests
make clean    # remove ./ttt
```

Or directly:

```bash
go run main.go
go test ./...
```

## Tracing (Honeycomb)

Tracing is wired to Honeycomb via OTLP. Create a free account at [honeycomb.io](https://www.honeycomb.io) and set these env vars before running (replace the API key):

```bash
export OTEL_SERVICE_NAME="tic-tac-toe"
export OTEL_EXPORTER_OTLP_PROTOCOL="http/protobuf"
export OTEL_EXPORTER_OTLP_ENDPOINT="https://api.honeycomb.io"
export OTEL_EXPORTER_OTLP_HEADERS="x-honeycomb-team=<YOUR_API_KEY_HERE>"
```

If these are not set you may see errors on startup, but the game still runs.

---

Comments, questions? Email: schuyler@sreconsulting.io
