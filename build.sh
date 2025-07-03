#!/bin/bash
echo "Building glue..."
go build -o glue .
if [ $? -eq 0 ]; then
    echo "✓ Build successful!"
    echo "Usage: ./glue [directory]"
    echo "Example: ./glue examples/"
else
    echo "✗ Build failed!"
    exit 1
fi
