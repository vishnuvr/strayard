#!/bin/bash

set -o errexit
set -o nounset
set -o pipefail

SY_ROOT=$(dirname "${BASH_SOURCE}")/..
source "${SY_ROOT}/hack/common.sh"

pushd "${SY_ROOT}/dashboard" > /dev/null
  grunt build
popd > /dev/null

pushd "${SY_ROOT}" > /dev/null
  Godeps/_workspace/bin/go-bindata -nocompress -prefix "dashboard/dist" -pkg "dashboard" -o "pkg/dashboard/bindata.go" -ignore "\\.gitignore" dashboard/dist/...
popd > /dev/null
