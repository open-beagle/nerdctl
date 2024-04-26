#!/bin/sh

set -x

export GOARCH=loong64
make binaries
mkdir -p _output/linux/$GOARCH
mv _output/nerdctl _output/linux/$GOARCH/nerdctl
