PROJECTNAME := $(shell basename "$(PWD)")

## run: run to start server
.PHONY: run
run:
	@echo "Server Starting"
	go run cmd/*.go

.PHONY: test
test:
	@echo "Running Tests"
	go test ./...

.PHONY: coverage
coverage:
	@echo "Generating Test Coverage & See"
	go test ./... -coverprofile=coverage.out
	go tool cover -func=coverage.out

.PHONY: mock
mock:
	@echo "Generating Mocks"
	mockery --all