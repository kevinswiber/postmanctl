#!/bin/sh

set -e

if [ $# -lt 2 ]; then
  echo "Usage: $0 <api_name> <api_version> -- <spectral_args: optional>"
  echo "Examples: "
  echo "  $0 Cosmos v1.0.0"
  echo "  $0 Users v1.3.0 -- -q --skip-rule=\"operation-tags\""
  exit 1
fi

api_name="$1"
api_version_name="$2"

shift
shift

spectral_params=""
if [ "$1" = "--" ]; then
  shift
  spectral_params=$@
fi

api_id=$(postmanctl get apis \
  -o jsonpath="{[?(.name == '$api_name')].id}")

api_version_id=$(postmanctl get api-versions \
  --for-api "$api_id" \
  -o jsonpath="{[?(.name == '$api_version_name')].id}")

schema=$(postmanctl get schema \
  --for-api "$api_id" \
  --for-api-version "$api_version_id" \
  -o jsonpath="{.schema}")

echo "$schema" | sh -c "spectral lint $spectral_params"
