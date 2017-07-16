#!/usr/bin/env bash
#
# This script builds the utility targeting the Travis sudo build environment
set -o nounset
set -o errexit -o errtrace

# Golang architecture and operating system documentation
# https://golang.org/doc/install/source#environment

CGO_ENABLED=0
GOOS=linux
GOARCH=amd64

go build -a -installsuffix -v
