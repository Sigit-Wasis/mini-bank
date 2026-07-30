APP_NAME=mini-bank
BUILD_DIR=bin

.PHONY: all build clean run test docker-build docker-up docker-down

all: build

build:
	@echo "Building $(APP_NAME)..."
	@mkdir -p $(BUILD_DIR)
	@go build -o $(BUILD_DIR)/$(APP_NAME) ./main.go

run: build
	@echo "Running $(APP_NAME)..."
	@./$(BUILD_DIR)/$(APP_NAME)

test:
	@echo "Running tests..."
	@go test -v -cover ./...

clean:
	@echo "Cleaning up..."
	@rm -rf $(BUILD_DIR)

docker-build:
	@docker build -t $(APP_NAME):latest .

docker-up:
	@docker-compose up -d

docker-down:
	@docker-compose down
