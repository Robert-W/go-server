fmt/list:

audit:
	go mod tidy -diff
	go mod verify
	gofmt -s -l -e ./internal
	go vet ./...
	go run golang.org/x/vuln/cmd/govulncheck@latest ./...

test/unit:
	go test -race -cover -coverprofile=./tmp/coverage.out ./internal/...

tag = latest
docker/build:
	docker build -t go-server:$(tag) .
