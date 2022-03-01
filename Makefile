default:
	@echo "Building Go app..."
	go build -o network-monitor-macos
	@echo "Running test coverage report..."
	go test -v -cover ./...

github:
	@echo "Committing changes to Github..."
	git add -A
	git commit -m "$m"
	git push

arm-build:
	env GOOS=linux GOARCH=arm GOARM=5 go build -o network-monitor-arm