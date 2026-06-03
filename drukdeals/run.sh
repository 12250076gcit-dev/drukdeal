#!/bin/bash
# Run Go app with the working Go version (1.23.5)
cd /workspaces/drukdeals/drukdeals
GOROOT=/tmp/go /tmp/go/bin/go run main.go "$@"