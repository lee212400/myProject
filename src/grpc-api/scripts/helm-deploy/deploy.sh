#!/bin/bash

set -ex

ROOT_DIR=$(cd "$(dirname "$0")/../.." && pwd)


NAMESPACE="my-api"
IMAGE_NAME="myapp:$(date +'%Y-%m-%d_%H%M%S')"
MY_DOMAIN="https://example.com"

CHART_DIR="$ROOT_DIR/app-deployment"

RELEASE_NAME="myapp"

eval $(minikube -p minikube docker-env)

# Docker image bulid
docker build -t $IMAGE_NAME -f "$ROOT_DIR/dockerfile/app/dockerfile" "$ROOT_DIR"

kubectl get namespace "$NAMESPACE" >/dev/null 2>&1 || kubectl create namespace "$NAMESPACE"

helm upgrade --install "$RELEASE_NAME" "$CHART_DIR" \
  --namespace "$NAMESPACE" \
  --set imageName="$IMAGE_NAME" \
  --set myDomain="$MY_DOMAIN" \
  --set namespace="$NAMESPACE"