#!/usr/bin/env bash

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
export GOROOT="$SCRIPT_DIR/.local/go"
export GOPATH="$SCRIPT_DIR/.local/gopath"
export GOCACHE="$SCRIPT_DIR/.local/gocache"
export PATH="$GOROOT/bin:$PATH"
