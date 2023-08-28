#!/usr/bin/env bash

SCRIPT_DIRECTORY="$(dirname -- "$(readlink -f -- "${BASH_SOURCE[0]}")")"
cd "${SCRIPT_DIRECTORY}" || exit

# set options
# see https://linuxcommand.org/lc3_man_pages/seth.html
set -o errexit  # abort if command exits on error
set -o nounset  # abort when a non-existing variable is expanded
set -o pipefail # abort in case of "pipefail" (one command in pipe fails)
set -o xtrace   # log all actions and variables

# fetch dependencies in case of removal
go mod download -x

# run test suite
go test ./...

# run qa analysis
golangci-lint run
