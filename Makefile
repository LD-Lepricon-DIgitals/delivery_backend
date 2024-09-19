BINARY_NAME = delivery_service
BUILD_DIR = build

up:
	@echo "Starting Docker images..."
	docker-compose up -d
	@echo "Docker images started!"

up_build:
	@echo "Stopping docker images (if running...)"
	docker-compose down
	@echo "Building (when required) and starting docker images..."
	docker-compose up --build -d
	@echo "Docker images built and started!"

down:
	@echo "Stopping docker-compose..."
	docker-compose down
	@echo "DONE!"

build_service:
	@echo "Building service binary..."
	cd cmd/app && set GOOS=linux&& set CGO_ENABLED=0 && go build -o ../../$(BUILD_DIR)/$(BINARY_NAME) .
	@echo "Binary built: $(BUILD_DIR)/$(BINARY_NAME)"