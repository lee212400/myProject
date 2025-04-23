#!/bin/bash

set -ex

NAMESPACE="my-api"
SECRET_NAME="myapp-secret"

kubectl create secret generic $SECRET_NAME \
    --from-literal=KEY_NAME=TEST_KEY_NAME \
    --from-literal=APP_SECRET=TEST_APP_SECRET \
    -n $NAMESPACE \
    --dry-run=client -o yaml | kubectl apply -f -