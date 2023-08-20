#!/usr/bin/env bash

SCRIPT_DIRECTORY="$(dirname -- "$(readlink -f -- "${BASH_SOURCE[0]}")")"
cd "${SCRIPT_DIRECTORY}" || exit

# contains the container ID for referring to the container after running it
CID_FILE=container.cid

# run the container
docker run \
    --cidfile "${CID_FILE}" \
    --interactive \
    --name codewars-go120 \
    --rm \
    --tty \
    --user "$(id -u)":"$(id -g)" \
    --volume "${PWD}":/codewars/go120 \
    --volume "gopkgcache":/home/codewars/go/pkg \
    andrerademacher/codewars-go120 "$@"

rm "${CID_FILE}"
