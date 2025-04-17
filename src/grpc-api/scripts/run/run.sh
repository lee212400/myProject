#!/bin/bash

set -ex

ROOT_DIR=$(cd "$(dirname "$0")/../.." && pwd)

go run "$ROOT_DIR/cmd/main.go"