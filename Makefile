SERVICE_NAME = "potato"
SERVICE_VERSION = "v1"

test:
	@echo "Running tests..."
	$(MAKE) unit-test

unit-test:
	@echo "Running unit tests..."
	go test -v ./...

coverage:
	@echo "Calculating coverage..."
	@scripts/coverage.sh

profile:
	@echo "Generating coverage report..."
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out

format:
	@echo "Formatting..."
	gofmt -s -w .

lint:
	@echo "Applying lint..."
	go vet ./...

build:
	@echo "Building app..."
	CGO_ENABLED=0 GOOS=linux GARCH=amd64 go build -o app

run:
	@echo "Running app..."
	go run main.go -config config/config