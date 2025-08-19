# go-server
Practice project building an API server in Go

## Running the app
`go run cmd/api/main.go`

## Running tests
To generate coverage and see it in the browser, run the following commands:

```bash
go test -cover -coverprofile=coverage.out ./internal/...
go tool cover -html=coverage.out
```

## Telemetry
If you want to see telemetry while developing, you just need to have a locally
running collector. Run `docker compoose up` to spin one up and then you can see
traces in your console.
