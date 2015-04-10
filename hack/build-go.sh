#!/bin/bash
# This script sets up a go workspace locally and builds all go components.

set -o errexit
set -o nounset
set -o pipefail

SY_ROOT=$(dirname "${BASH_SOURCE}")/..
source "${SY_ROOT}/hack/lib/init.sh"

strayard::golang::build_binaries "$@"
strayard::golang::place_bins
