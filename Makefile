run:
	go run cmd/api/main.go


test:
	go test ./... -v

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
