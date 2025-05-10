#!/bin/bash
set -ex
ROOT_DIR=$(cd "$(dirname "$0")/.." && pwd)

bash "$ROOT_DIR/scripts/mysql.sh"

docker compose -f "$ROOT_DIR/docker-compose/protoc/docker-compose.yaml" \
    -f "$ROOT_DIR/docker-compose/app/docker-compose.yaml" \
    up --build