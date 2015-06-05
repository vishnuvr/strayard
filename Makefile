# Old-skool build tools.
#
# Targets (see each target for more information):
#   all: Build code.
#   build: Build code.
#   clean: Clean up.

OUT_DIR = _output
OUT_PKG_DIR = Godeps/_workspace/pkg

# Build code.
#
# Args:
#   WHAT: Directory names to build.  If any of these directories has a 'main'
#     package, the build will produce executable files under $(OUT_DIR)/go/bin.
#     If not specified, "everything" will be built.
#   GOFLAGS: Extra flags to pass to 'go' when building.
#   TESTFLAGS: Extra flags that should only be passed to hack/test-go.sh
#
# Example:
#   make
#   make all
#   make all WHAT=cmd/kubelet GOFLAGS=-v
all build:
	hack/build-go.sh $(WHAT)
.PHONY: all build

# Remove all build artifacts.
#
# Example:
#   make clean
clean:
	rm -rf $(OUT_DIR) $(OUT_PKG_DIR)
.PHONY: clean
