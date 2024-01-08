#!/bin/sh

set -x

git apply .beagle/v1.7.2-support-loong64.patch

export GOARCH=loong64
make binaries
mkdir -p _output/linux/$GOARCH
mv _output/nerdctl _output/linux/$GOARCH/nerdctl

git apply -R .beagle/v1.7.2-support-loong64.patch