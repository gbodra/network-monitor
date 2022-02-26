default:
	@echo "Building Go app..."
	go build
	@echo "Running test coverage report..."
	go test -v -cover ./...

github:
	@echo "Committing changes to Github..."
	git add -A
	git commit -m "$m"
	git push