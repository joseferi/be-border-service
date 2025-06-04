start:
	@echo "Starting apps..."
	@go run main.go http
install:
	@echo "Installing ..."
	@go mod tidy && go mod vendor
	
clean:
	@echo "Cleaning"
	@rm -rf vendor

run-lint:
	@golangci-lint run internal/usecase