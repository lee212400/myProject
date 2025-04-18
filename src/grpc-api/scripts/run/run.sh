#!/bin/bash

set -ex

ROOT_DIR=$(cd "$(dirname "$0")/../.." && pwd)

source "$ROOT_DIR/conf/config.conf"

bash "$ROOT_DIR/scripts/mysql/mysql.sh"

bash "$ROOT_DIR/scripts/proto/proto.sh"

go run "$ROOT_DIR/cmd/gateway/main.go" & \
docker compose -f "$ROOT_DIR/docker-compose/app/docker-compose.yaml" up --build