.PHONY: test
test:
	go test ./robot

.PHONY: integration-test
integration-test:
	bash -c "trap 'docker-compose stop' EXIT; docker compose up database -d"

	DATABASE_HOST=localhost \
    DATABASE_PORT="5432" \
    DATABASE_USER=test \
    DATABASE_PASSWORD=test \
    DATABASE_NAME=test go test ./app/

.PHONY: fmt
fmt:
	 gofmt -w $(shell find . -iname '*.go' -not -path "./vendor/*" | xargs)

.PHONY: lint
lint:
	test $(shell gofmt -l robot app main.go | tee /dev/stderr | wc -l) = 0 || exit 1
