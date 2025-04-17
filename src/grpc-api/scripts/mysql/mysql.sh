#!/bin/bash

ROOT_DIR=$(cd "$(dirname "$0")/../.." && pwd)

docker compose -f "$ROOT_DIR/docker-compose/mysql/docker-compose.yaml" down -v
docker compose -f "$ROOT_DIR/docker-compose/mysql/docker-compose.yaml" up -d

MAX_RETRIES=20
RETRY_COUNT=0

until docker exec mysql-server mysql -u root -ppassword -e "SHOW DATABASES LIKE 'mydb';" ; do
  RETRY_COUNT=$((RETRY_COUNT+1))
  
  if [ "$RETRY_COUNT" -ge "$MAX_RETRIES" ]; then
    echo "MySQL did not become ready after $MAX_RETRIES attempts. Exiting."
    exit 1
  fi
  
  echo "Attempt $RETRY_COUNT/$MAX_RETRIES: MySQL not ready yet, retrying in 5 seconds..."
  sleep 5
done

echo "MySQL is ready and accepting connections!"

go run "$ROOT_DIR/cmd/migration/main.go"