.PHONY: lint
default: run

COVERAGE_FILE ?= coverage.out

.PHONY: lint
lint:
	@golangci-lint --version && echo "golangci-lint -v run --fix ./..." || echo "golangci-lint not found"
	@golangci-lint -v run --fix ./...

.PHONY: run
run:
	go run ./cmd/server

.PHONY: test
test:
	@go test -v -coverpkg=./... -covermode=atomic -race -coverprofile='$(COVERAGE_FILE)' ./...
	@go tool cover -func='$(COVERAGE_FILE)' | grep ^total | tr -s '\t'

.PHONY: test-short
test-short:
	@go test -short -v ./...

.PHONY: mockgen
mockgen:
	@go generate ./...


API_ENDPOINT = http://localhost:3001
URL :=

.PHONY: shorten
shorten:
	curl -X POST -d '{"url": "$(URL)"}' $(API_ENDPOINT)
