#!/bin/bash
set -ex
ROOT_DIR=$(cd "$(dirname "$0")/.." && pwd)

docker compose -f "$ROOT_DIR/docker-compose/protoc/docker-compose.yaml" \
    -f "$ROOT_DIR/docker-compose/app/docker-compose.yaml" \
    up --build