#!/bin/bash

# Exit on error, print commands
set -xe

# Ensure bin directory exists
mkdir -p bin

# Download dependencies
go mod download

# Install required packages if not in go.mod
go get -u github.com/gin-gonic/gin
go get -u github.com/tpkeeper/gin-dump
go get -u github.com/go-playground/validator/v10

# Install and generate Swagger docs
# Check if GOPATH is set
if [ -z "$GOPATH" ]; then
  export GOPATH="$HOME/go"
  export PATH="$GOPATH/bin:$PATH"
fi

# Check if swag exists in PATH or GOPATH/bin
if ! command -v swag &>/dev/null && [ ! -f "$GOPATH/bin/swag" ]; then
  echo "Installing swag..."
  go install github.com/swaggo/swag/cmd/swag@latest
fi

# Generate swagger docs
"$GOPATH/bin/swag" init

# Build the application
CGO_ENABLED=1 go build -o bin/application main.go

# Make the binary executable
chmod +x bin/application

echo "Build completed successfully!"
