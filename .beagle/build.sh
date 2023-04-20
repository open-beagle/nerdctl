#!/bin/sh

set -x

export GOARCH=amd64
make binaries
mkdir -p _output/linux/$GOARCH
mv _output/nerdctl _output/linux/$GOARCH/nerdctl

export GOARCH=arm64
make binaries
mkdir -p _output/linux/$GOARCH
mv _output/nerdctl _output/linux/$GOARCH/nerdctl

export GOARCH=ppc64le
make binaries
mkdir -p _output/linux/$GOARCH
mv _output/nerdctl _output/linux/$GOARCH/nerdctl

export GOARCH=mips64le
make binaries
mkdir -p _output/linux/$GOARCH
mv _output/nerdctl _output/linux/$GOARCH/nerdctl

export GOARCH=loong64
make binaries
mkdir -p _output/linux/$GOARCH
mv _output/nerdctl _output/linux/$GOARCH/nerdctl
