#!/bin/bash

set -e

ROOT_DIR=$(cd "$(dirname "$0")/.." && pwd)

docker build -t go-grpc-server -f "$ROOT_DIR/docker/grpc/dockerfile2" "$ROOT_DIR"

docker run --rm -p 50051:50051 go-grpc-server