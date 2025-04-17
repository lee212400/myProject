#!/bin/bash

ROOT_DIR=$(cd "$(dirname "$0")/../.." && pwd)

docker compose -f "$ROOT_DIR/docker-compose/mysql/docker-compose.yaml" down -v
docker compose -f "$ROOT_DIR/docker-compose/mysql/docker-compose.yaml" up -d

MAX_RETRIES=10
RETRY_COUNT=0

until docker ps -q -f name=mysql-server;
do
  RETRY_COUNT=$((RETRY_COUNT+1))
  if [ "$RETRY_COUNT" -ge "$MAX_RETRIES" ]; then
    echo "MySQL did not become ready after $MAX_RETRIES attempts. Exiting."
    cleanup
    exit 1
  fi
  echo "Attempt $RETRY_COUNT/$MAX_RETRIES: MySQL not ready yet, retrying in 2s..."
  sleep 2
done