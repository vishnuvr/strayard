#!/bin/bash

# Script to generate docs from the latest swagger spec.

set -o errexit
set -o nounset
set -o pipefail

SY_ROOT=$(dirname "${BASH_SOURCE}")/..
source "${SY_ROOT}/hack/util.sh"

pushd "${SY_ROOT}/hack/swagger-doc" > /dev/null
  gradle gendocs --info
popd > /dev/null

echo "SUCCESS"
