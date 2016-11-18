run:
	@echo "Fetching dependencies..."
	@go get github.com/joho/godotenv
	@echo "Running application..."
	@go run main.go
.PHONY: run
