#!/bin/bash

set -o errexit
set -o nounset
set -o pipefail

# The root of the build/dist directory
SY_ROOT=$(
  unset CDPATH
  sy_root=$(dirname "${BASH_SOURCE}")/..
  cd "${sy_root}"
  pwd
)

SY_OUTPUT_SUBPATH="${SY_OUTPUT_SUBPATH:-_output/local}"
SY_OUTPUT="${SY_ROOT}/${SY_OUTPUT_SUBPATH}"
SY_LOCAL_BINPATH="${SY_ROOT}/_output/local/go/bin"

readonly SY_GO_PACKAGE=github.com/strayard/strayard
readonly SY_GOPATH="${SY_OUTPUT}/go"

# Asks golang what it thinks the host platform is. 
sy::build::host_platform() {
  echo "$(go env GOHOSTOS)/$(go env GOHOSTARCH)"
}

# Build binaries targets specified
#
# Input:
#   $@ - targets and go flags.  If no targets are set then all binaries targets
#     are built.
sy::build::build_binaries() {
  # Create a sub-shell so that we don't pollute the outer environment
  (
    #GOPATH=${SY_GOPATH}
    #export GOPATH
    local targets=("cmd/strayard")
    local binaries="github.com/strayard/strayard/dashboard"
    local platforms=($(sy::build::host_platform))
    local platform
    for platform in "${platforms[@]}"; do
      echo "++ Building go targets for ${platform}:" "${targets[@]}"
      go install "${binaries[@]}"
    done
  )
}

sy::build::place_bins() {
  (   
    local host_platform
    host_platform=$(sy::build::host_platform)
    echo "++ Placing binaries"
    mkdir -p ${SY_LOCAL_BINPATH}
    local platform
    for platform in "${host_platform}"; do
      echo ${platform}
    done
  )
}

# sy::build::make_strayard_binary_symlinks makes symlinks for the strayard
# binary in _output/local/go/bin
sy::build::make_strayard_binary_symlinks() {
  return
}
