#!/bin/bash

ROOT_DIR=$(cd "$(dirname "$0")/../.." && pwd)

docker compose -f "$ROOT_DIR/docker-compose/postgre/docker-compose.yaml" down -v
docker compose -f "$ROOT_DIR/docker-compose/postgre/docker-compose.yaml" up -d