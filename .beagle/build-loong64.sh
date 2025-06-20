#!/bin/sh

set -x

git config --global --add safe.directory $PWD
# git apply .beagle/v2.0.3-images-prune.patch

export VERSION="${BUILD_VERSION:-v2.0.0-rc.2}-beagle-$(git rev-parse --short HEAD 2>/dev/null || true)"

export GOARCH=loong64
make binaries
mkdir -p _output/linux/$GOARCH
mv _output/nerdctl _output/linux/$GOARCH/nerdctl
