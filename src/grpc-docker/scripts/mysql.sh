#!/bin/bash

ROOT_DIR=$(cd "$(dirname "$0")/.." && pwd)

docker compose -f "$ROOT_DIR/docker-compose/mysql/docker-compose.yaml" up -d
