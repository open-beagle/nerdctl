#!/bin/sh

set -x

git config --global --add safe.directory $PWD

export VERSION="${BUILD_VERSION:-v1.7.7}-beagle-$(git rev-parse --short HEAD 2>/dev/null || true)"

export GOARCH=amd64
make binaries
mkdir -p _output/linux/$GOARCH
mv _output/nerdctl _output/linux/$GOARCH/nerdctl

export GOARCH=arm64
make binaries
mkdir -p _output/linux/$GOARCH
mv _output/nerdctl _output/linux/$GOARCH/nerdctl
