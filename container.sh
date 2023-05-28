#!/bin/bash

SCRIPT_DIRECTORY="$(dirname -- "$(readlink -f -- "${BASH_SOURCE[0]}")")"
cd "${SCRIPT_DIRECTORY}" || exit

docker run \
  --interactive \
  --tty \
  --user "$(id -u)":"$(id -g)" \
  --volume "${PWD}":/codewars/go120 \
  --volume "${PWD}/.cache":/.cache \
  andrerademacher/codewars-go120 "$@"
