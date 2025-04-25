#!/bin/bash


ROOT_DIR=$(cd "$(dirname "$0")/../.." && pwd)

COVERAGE_FILE=coverage.out
TIMEOUT=10m

TARGET_DIRS=(
    $ROOT_DIR/lint-test
)


TEST_DIR="${TARGET_DIRS[@]}"

echo "Running go test with timeout=$TIMEOUT and coverage output to $COVERAGE_FILE"
go test -timeout=$TIMEOUT -covermode=count -coverprofile=$COVERAGE_FILE $TEST_DIR

if [ -f $COVERAGE_FILE ]; then
    echo "Generating HTML coverage report..."
    go tool cover -html=$COVERAGE_FILE -o coverage.html
fi