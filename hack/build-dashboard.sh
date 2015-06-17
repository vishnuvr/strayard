#!/bin/bash

set -o errexit
set -o nounset
set -o pipefail

SY_ROOT=$(dirname "${BASH_SOURCE}")/..
source "${SY_ROOT}/hack/common.sh"

pushd "${SY_ROOT}/dashboard" > /dev/null
  grunt build
  ../Godeps/_workspace/bin/go-bindata-assetfs -nocompress -ignore "\\.gitignore" dist/...
popd > /dev/null

#pushd "${SY_ROOT}" > /dev/null
  #Godeps/_workspace/bin/go-bindata -nocompress -prefix "dashboard/dist" -pkg "dashboard" -o "pkg/dashboard/bindata.go" -ignore "\\.gitignore" dashboard/dist/...
  #Godeps/_workspace/bin/go-bindata-assetfs -nocompress -prefix "dashboard/dist" -pkg "dashboard" -ignore "\\.gitignore" dashboard/dist/...
  #mv "${SY_ROOT}/bindata_assetfs.go" "${SY_ROOT}/pkg/dashboard/bindata_assetfs.go"
  #Godeps/_workspace/bin/go-bindata-assetfs -nocompress -ignore "\\.gitignore" dashboard/dist/...
#popd > /dev/null
