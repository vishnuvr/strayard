#!/bin/bash

# This script sets up a go workspace locally and builds all go components.

set -o errexit
set -o nounset
set -o pipefail

SY_ROOT=$(dirname "${BASH_SOURCE}")/..
source "${SY_ROOT}/hack/common.sh"

sy::build::build_binaries "$@"
sy::build::place_bins
sy::build::make_strayard_binary_symlinks
