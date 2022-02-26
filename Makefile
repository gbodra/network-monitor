default:
	@echo "Building Go app..."
	go build
	@echo "Running test coverage report..."
	go test -v -cover ./...