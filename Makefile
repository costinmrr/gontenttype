export GO111MODULE=on

tidy:
	go mod tidy

build:
	go build

test:
	go test -v ./...

benchmark:
	go test -bench=.

fmt:
	go fmt ./...

lint:
	golangci-lint run ./...

govulncheck:
	@go get golang.org/x/vuln/cmd/govulncheck
	@go run golang.org/x/vuln/cmd/govulncheck ./...
