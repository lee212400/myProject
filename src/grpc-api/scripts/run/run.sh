#!/bin/bash

set -ex

ROOT_DIR=$(cd "$(dirname "$0")/../.." && pwd)

bash "$ROOT_DIR/scripts/mysql/mysql.sh"

go run "$ROOT_DIR/cmd/main.go"