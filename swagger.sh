#!/bin/bash

# Exit on error, print commands
set -xe

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

echo "Swagger documentation generated successfully!"
