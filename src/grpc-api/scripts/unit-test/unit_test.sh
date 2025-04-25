#!/bin/bash


ROOT_DIR=$(cd "$(dirname "$0")/../.." && pwd)

COVERAGE_FILE=coverage.out
TIMEOUT=10m

TARGET_DIRS=(
    $ROOT_DIR/interface/controller
)


TEST_DIR="${TARGET_DIRS[@]}"

echo "Running go test with timeout=$TIMEOUT and coverage output to $COVERAGE_FILE"
go test -timeout=$TIMEOUT -covermode=atomic -coverprofile=$COVERAGE_FILE $TEST_DIR

if [ -f $COVERAGE_FILE ]; then
    echo "coverage report:$covered_functions/$total_functions"
    go tool cover -func=$COVERAGE_FILE | grep -v "no coverage"


    total_coverage=$(go tool cover -func=$COVERAGE_FILE | grep total | awk '{print $3}')
    echo "total coverage: $total_coverage"

fi

echo "Generating HTML report..."
go tool cover -html=$COVERAGE_FILE -o coverage.html