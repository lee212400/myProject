#!/bin/bash
set -euo pipefail

echo "Installing golangci-lint..."
export PATH=$HOME/.local/bin:$PATH
curl -sSfL https://github.com/golangci/golangci-lint/releases/download/v2.1.2/golangci-lint-2.1.2-linux-amd64.tar.gz | tar -xvzf - -C $HOME/.local/bin

echo "golangci-lint version:"
golangci-lint --versio

echo "Running golangci-lint..."
golangci-lint run ./... -v