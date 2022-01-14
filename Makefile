PRIVATE_REPOS=github.com/polymarket/*

.PHONY : clean
clean:
	@echo "Cleaning env..."
	go clean -cache
	go clean -testcache
	@echo "Cleaned env!"

.PHONY: lint
lint:
	@echo "Linting code..."
	gofmt -s -w ./.
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.43.0
	golangci-lint run ./... -E gofmt
	go mod tidy
	@echo "Linting complete!"

.PHONY: test
test: clean
	@echo "Running tests..."
	GOPRIVATE=${PRIVATE_REPOS} go test ./... -v
	@echo "Tests complete!"

.PHONY: test-coverage
test-coverage:
	@echo "Running tests with coverage..."
	GOPRIVATE=${PRIVATE_REPOS} go test ./... -v -cover
	@echo "Tests complete!"

.PHONY: test-racing
test-racing:
	@echo "Running race condition tests..."
	GOPRIVATE=${PRIVATE_REPOS} go test ./... -v -race
	@echo "Tests complete!"
