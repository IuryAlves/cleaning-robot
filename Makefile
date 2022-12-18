.PHONY: test
test:
	go test ./...

.PHONY: fmt
fmt:
	gofmt -w robot client.go client_test.go

.PHONY: lint
lint:
	test $(shell gofmt -l robot client.go client_test.go | wc -l) = 0 || exit 1
