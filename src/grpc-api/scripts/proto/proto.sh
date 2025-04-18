#!/bin/bash

set -ex

ROOT_DIR=$(cd "$(dirname "$0")/../.." && pwd)

rm -rf "$ROOT_DIR/rpc"

mkdir "$ROOT_DIR/rpc"

(cd "$ROOT_DIR/buf" && buf dep update)

(cd "$ROOT_DIR/buf" && buf generate)
