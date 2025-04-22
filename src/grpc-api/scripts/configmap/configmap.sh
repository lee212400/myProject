#!/bin/bash

set -ex

ROOT_DIR=$(cd "$(dirname "$0")/../.." && pwd)

kubectl apply -f "$ROOT_DIR/config-map/configmap.yaml"