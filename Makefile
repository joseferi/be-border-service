start:
	@echo "Starting apps..."
	@go run main.go http
install:
	@echo "Installing ..."
	@go mod tidy && go mod vendor && git config core.hooksPath .git-hooks
	
clean:
	@echo "Cleaning"
	@rm -rf vendor

test:
	@go test internal/... -v -cover -coverprofile=coverage.out
view-test:
	@go tool cover -html=coverage.out -o coverage.html