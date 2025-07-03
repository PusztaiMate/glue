#!/bin/bash
echo "Building Go Test Runner..."
go build -o glue .
if [ $? -eq 0 ]; then
    echo "✓ Build successful!"
    echo "Usage: ./gotest-runner [directory]"
    echo "Example: ./gotest-runner examples/"
else
    echo "✗ Build failed!"
    exit 1
fi
