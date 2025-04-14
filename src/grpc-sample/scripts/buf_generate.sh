#!/bin/bash

set -e

PROTO_DIR="proto"
BUF_CONFIG_DIR="buf"
OUT_DIR="rpc"
ROOT_DIR=$(cd "$(dirname "$0")/.." && pwd)

# rpcファイル削除
rm -f "$ROOT_DIR/$OUT_DIR"/*.pb.go "$ROOT_DIR/$OUT_DIR/*.pb.gw.go" "$ROOT_DIR/$OUT_DIR/*.grpc.pb.go"

# bufファイルをprotoにコピー
cp "$ROOT_DIR/$BUF_CONFIG_DIR/buf.yaml" "$ROOT_DIR/$PROTO_DIR/"
cp "$ROOT_DIR/$BUF_CONFIG_DIR/buf.gen.yaml" "$ROOT_DIR/$PROTO_DIR/"

cd "$PROTO_DIR"

# proto dirでbuf 実行
cd "$ROOT_DIR/$PROTO_DIR"
buf dep update
buf generate

# 依存関係ファイルをbufにコピー
cp "$ROOT_DIR/$PROTO_DIR/buf.lock" "$ROOT_DIR/$BUF_CONFIG_DIR/"

# proto/yamlとlockファイルは削除する
rm "$ROOT_DIR/$PROTO_DIR/buf.yaml" "$ROOT_DIR/$PROTO_DIR/buf.gen.yaml" "$ROOT_DIR/$PROTO_DIR/buf.lock"

echo "OK"