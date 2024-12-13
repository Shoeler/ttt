This is a repo I'm putting together as the entrypoint to an apprenticeship.  I haven't written much code since I moved into leadership, so this is a new way to enhance my skills and experience.
It implements the minimax algorithm to make the computer theoretically unbeatable and also uses open telemetry for debugging.

**Golang Installation - all other platforms**

See:  [Golang install](https://go.dev/doc/install)

**Golang Installation - macOS with homebrew**
```
brew install golang
cd <repo directory>
go get github.com/honeycombio/otel-config-go/otelconfig
go get go.opentelemetry.io/otel
go get go.opentelemetry.io/otel/attribute
go get go.opentelemetry.io/otel/codes
go get go.opentelemetry.io/otel/trace
```

**Running the game**
```
go run main.go
```
**Building the executable**
```
go build 
```

**Running tests**
```
go test ttt/board
go test ttt/computer
```

**Debugging**

I put all of my debugging code into open telemetry, which is setup to stream to honeycomb.  Create a free account with [honeycomb](https://www.honeycomb.io)  and then the following environment variables must be set in your terminal.  Make sure you change the `x-honeycomb-team` to your API key.
```
export OTEL_SERVICE_NAME="tic-tac-toe"
export OTEL_EXPORTER_OTLP_PROTOCOL="http/protobuf"
export OTEL_EXPORTER_OTLP_ENDPOINT="https://api.honeycomb.io"
export OTEL_EXPORTER_OTLP_HEADERS="x-honeycomb-team=<YOUR_API_KEY_HERE>"
```

If you don't set these variables, you may get odd errors.


Comments, questions?  E-mail me:  schuyler@sreconsulting.io