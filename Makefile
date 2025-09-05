fmt/list:

audit:
	go mod tidy -diff
	go mod verify
	gofmt -s -l -e ./internal
	go vet ./...
	go run golang.org/x/vuln/cmd/govulncheck@latest ./...

test/unit:
	go test -race -shuffle=on -v -cover -coverprofile=./coverage.out ./internal/...

# make migrate/new name=migration-name
migrate/new:
	goose create $(name) go -dir ./migrations

# make docker/build tag=latest
tag = latest
docker/build:
	docker build -t go-server:$(tag) .
