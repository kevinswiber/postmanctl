#!/bin/sh

set -e

if [ $# -lt 1 ]; then
  echo "Usage: $0 <collection-id>"
  exit 1
fi

dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"

postmanctl get co -o go-template-file="${dir}/jmeter.tmpl" $1