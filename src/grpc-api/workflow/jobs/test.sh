#!/bin/bash
set -euo pipefail

MY_LEVEL="info"
MY_ENVIRONMENT="dev"
MY_PARAMS="{}"

while [[ $# -gt 0 ]]; do
  key="$1"
  case $key in
    --level)
      MY_LEVEL="$2"
      shift 2
      ;;
    --env)
      MY_ENVIRONMENT="$2"
      shift 2
      ;;
    --params)
      MY_PARAMS="$2"
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