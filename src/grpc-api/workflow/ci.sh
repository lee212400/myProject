#!/bin/bash
set -euo pipefail

export PATH="$(go env GOPATH)/bin:${PATH}"
mkdir -p "$(go env GOPATH)"/bin

curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/HEAD/install.sh \
  | sh -s -- -b "$(go env GOPATH)/bin" v2.1.2

echo "golangci-lint version: $(golangci-lint --version)"

cd src/grpc-api

echo "Running golangci-lint..."
golangci-lint run ./... -v