# go-server
Practice project building an API server in Go

## Pre-requisites
- Install [Go](https://go.dev/doc/install)
- Install [Docker](https://docs.docker.com/get-docker)

## Running the app
In one terminal session: `docker compose up`
In another terminal: `go run cmd/api/main.go`

## Running tests
To generate coverage and see it in the browser, run the following commands:

```bash
go test -cover -coverprofile=coverage.out ./internal/...
go tool cover -html=coverage.out
```

## Building with Docker
You can build the image by running

```bash
docker build -t <tag-name> .
```

and then you can verify it by running the image locally

```bash
docker run -p 3000:3000 <tag-name>
```
