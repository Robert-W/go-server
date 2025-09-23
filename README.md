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

### Running integration tests
This is currently a pain in the ass and there is a lot here to be desired, but
I'll get to that later.

```bash
# Start the database
docker compose -f docker-compose.test.yaml -d

# Run the server, make sure the database url matches whats set in compose
DATABASE_URL=postgres://username:password@0.0.0.0:5432/sweet_potato go run cmd/api/main.go

# Run the tests
go test ./test/...
```

## Building with Docker
You can build the image by running

```bash
docker build -t <tag-name> .
```

and then you can verify it by running the image locally

```bash
# In one terminal
docker compose up

# In another terminal
docker run --env-file .env.docker -p 3000:3000 <tag-name>
```
