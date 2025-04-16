#!/bin/bash

set -e

# 必要な.protoダウンロード
ROOT_DIR=$(cd "$(dirname "$0")/.." && pwd)
mkdir -p "$ROOT_DIR/proto/google/api"
curl -o "$ROOT_DIR/proto/google/api/annotations.proto" https://raw.githubusercontent.com/googleapis/googleapis/master/google/api/annotations.proto
curl -o "$ROOT_DIR/proto/google/api/http.proto" https://raw.githubusercontent.com/googleapis/googleapis/master/google/api/http.proto

# rpcファイル削除
rm -f rpc/*.pb.go rpc/*.grpc.pb.go

# rpcファイル生成
protoc \
  --proto_path=proto \
  --proto_path=proto/google/api \
  --go_out=rpc --go_opt=paths=source_relative \
  --go-grpc_out=rpc --go-grpc_opt=paths=source_relative \
  --grpc-gateway_out=rpc --grpc-gateway_opt=paths=source_relative \
  $(find proto -name "*.proto")

# ダウンロードしてproto削除
rm -rf "$ROOT_DIR/proto/google" "$ROOT_DIR/rpc/google"