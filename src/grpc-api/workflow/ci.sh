#!/bin/bash
set -euo pipefail

export PATH="$(go env GOPATH)/bin:${PATH}"
mkdir -p "$(go env GOPATH)"/bin

# 2. Install golangci-lint into $GOPATH/bin
curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/HEAD/install.sh \
  | sh -s -- -b "$(go env GOPATH)/bin" v2.1.2

# 3. Verify installation
echo "golangci-lint version: $(golangci-lint --version)"

go --verison

echo "Running golangci-lint..."
golangci-lint run ./... -v