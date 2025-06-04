start:
	@echo "Starting apps..."
	@go run main.go http
install:
	@echo "Installing ..."
	@go mod tidy && go mod vendor && git config core.hooksPath .git-hooks
	
clean:
	@echo "Cleaning"
	@rm -rf vendor