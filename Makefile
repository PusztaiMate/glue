.PHONY: build clean test run

build:
	go build -o gotest-runner .

clean:
	rm -f gotest-runner

test:
	go test ./examples/...

run: build
	./gotest-runner

install: build
	sudo cp gotest-runner /usr/local/bin/

help:
	@echo "Available targets:"
	@echo "  build   - Build the application"
	@echo "  clean   - Remove built binary"
	@echo "  test    - Run example tests"
	@echo "  run     - Build and run in current directory"
	@echo "  install - Install to /usr/local/bin"
	@echo "  help    - Show this help"
