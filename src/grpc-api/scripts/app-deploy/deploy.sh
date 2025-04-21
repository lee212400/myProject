#!/bin/bash

set -ex

ROOT_DIR=$(cd "$(dirname "$0")/../.." && pwd)

NAMESPACE="my-api"
IMAGE_NAME="myapp:local"

eval $(minikube -p minikube docker-env)

# Docker image bulid
docker build -t $IMAGE_NAME -f "$ROOT_DIR/dockerfile/app/dockerfile" "$ROOT_DIR"

# create namespace
kubectl get namespace $NAMESPACE || kubectl create namespace $NAMESPACE

# Minikube image deploy
export NAMESPACE IMAGE_NAME
envsubst < "$ROOT_DIR/docker-compose/app/deployment/deployment.yaml.tpl" | kubectl apply -f -