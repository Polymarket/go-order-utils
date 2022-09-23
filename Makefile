PRIVATE_REPOS=github.com/polymarket/*

.PHONY: lint
lint:
	@echo "Linting code..."
	gofmt -s -w ./.
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.43.0
	golangci-lint run ./... -E gofmt
	go mod tidy
	@echo "Linting complete!"

.PHONY: test
test:
	@echo "Running tests..."
	GOPRIVATE=${PRIVATE_REPOS} go test ./... -v
	@echo "Tests complete!"
