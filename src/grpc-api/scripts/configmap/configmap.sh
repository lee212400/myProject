#!/bin/bash

set -ex

#ROOT_DIR=$(cd "$(dirname "$0")/../.." && pwd)

#kubectl apply -f "$ROOT_DIR/config-map/configmap.yaml"

#kubectl apply -f "$ROOT_DIR/secret/secret.yaml"

NAMESPACE="my-api"
CONFIGMAP_NAME="myapp-config"

kubectl create configmap $CONFIGMAP_NAME \
    --from-literal=LOG_LEVEL=debeg \
    --from-literal=TIMEOUT=30 \
    --from-literal=MYSQL_HOST=mysql.default.svc.cluster.local \
    --from-literal=MYSQL_USER=root \
    --from-literal=MYSQL_PASSWORD=password \
    --from-literal=MYSQL_PORT=3306 \
    --from-literal=MYSQL_DATABASES=mydb \
    --from-literal=MY_CONFIG=my_config \
    -n $NAMESPACE \
    --dry-run=client -o yaml | kubectl apply -f -