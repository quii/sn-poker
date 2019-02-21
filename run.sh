#!/usr/bin/env bash
set -euo pipefail

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"

${DIR}/build.sh

(
  cd ${DIR}/cmd/poker-app
  ./sn-poker
)
