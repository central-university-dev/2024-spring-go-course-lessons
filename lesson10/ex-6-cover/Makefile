test:
	go test -v ./...

cover:
	go test -v -cover ./...

test-cover:
	go test -v -coverprofile=coverage.out ./... -coverpkg=./... -covermode=atomic

report-console:
	go tool cover -func=coverage.out

report-html:
	go tool cover -html=coverage.out
