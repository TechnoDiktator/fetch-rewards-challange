# Variables
DOCKER_IMAGE_NAME = receipt-processor
CONTAINER_NAME = receipt-processor-container
DOCKER_PORT = 4040

# Run the project locally
run:
	go run cmd/api/main.go

# Test the project
test:
	go test ./... -v

# Format the code
fmt:
	go fmt ./...

# Tidy up dependencies
tidy:
	go mod tidy

# Generate API documentation (optional, if using Swagger)
swagger:
	swag init -g cmd/api/main.go

# Clean up Go build artifacts
clean:
	go clean
	rm -rf ./bin

# Docker: Build the image
docker-build:
	docker build -t $(DOCKER_IMAGE_NAME) .

# Docker: Run the container
docker-run:
	docker run -d --name $(CONTAINER_NAME) -p $(DOCKER_PORT):4040 $(DOCKER_IMAGE_NAME)

# Docker: Stop and remove the container
docker-stop:
	docker stop $(CONTAINER_NAME) || true
	docker rm $(CONTAINER_NAME) || true

# Docker: Rebuild and restart the container
docker-restart: docker-stop docker-build docker-run
	@echo "Docker container restarted successfully."
