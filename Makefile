.PHONY: test
test:
	go test ./...

.PHONY: fmt
fmt:
	 gofmt -w $(shell find . -iname '*.go' -not -path "./vendor/*" | xargs)


.PHONY: lint
lint:
	test $(shell gofmt -l robot server main.go | wc -l) = 0 || exit 1
