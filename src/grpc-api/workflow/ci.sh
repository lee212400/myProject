#!/bin/bash
set -euo pipefail

echo "Running golangci-lint..."
golangci-lint run ./... -v