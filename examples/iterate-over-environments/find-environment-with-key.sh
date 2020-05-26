#!/bin/sh

set -e

if [ $# -lt 1 ]; then
  echo "Usage: $0 <search-key>"
  echo "Example: "
  echo "  $0 apiKey"
  exit 1
fi

SEARCH_KEY="$1"
USERID=$(postmanctl get user -o jsonpath="{.id}")

postmanctl get env -o jsonpath="{[?(@.owner=='$USERID')].id}" | \
  xargs postmanctl get env -o json | \
  jq '.[] | select(has("values")) | select(.values[].key=="'"$SEARCH_KEY"'") | {id: .id, name: .name}'
