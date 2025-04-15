#!/bin/bash

ROOT_DIR=$(cd "$(dirname "$0")/.." && pwd)

docker compose -f "$ROOT_DIR/docker-compose/docker-compose.yaml build grpc-server"

docker compose -f "$ROOT_DIR/docker-compose/docker-compose.yaml" up -d