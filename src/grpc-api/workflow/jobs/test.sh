#!/bin/bash
set -euo pipefail

LEVEL="info"
ENVIRONMENT="dev"
PARAMS="{}"

while [[ $# -gt 0 ]]; do
  key="$1"
  case $key in
    --level)
      LEVEL="$2"
      shift 2
      ;;
    --env)
      ENVIRONMENT="$2"
      shift 2
      ;;
    --params)
      PARAMS="$2"
      shift 2
      ;;
    *)
      echo "Unknown option: $1"
      exit 1
      ;;
  esac
done

echo "LEVEL: $MY_LEVEL"
echo "ENVIRONMENT: $MY_ENVIRONMENT"
echo "PARAMS: $MY_PARAMS"