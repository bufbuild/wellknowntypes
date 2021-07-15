#!/usr/bin/env bash

set -eo pipefail

DIR="$(cd "$(dirname "${0}")/../.." && pwd)"
cd "${DIR}"

# args: error message
fail() {
  echo "error: $@" >&2
  exit 1
}

# arg1: command
check_command_installed() {
	if ! command -v "${1}" >/dev/null 2>/dev/null; then
    fail "command ${1} must be installed"
  fi
}

check_command_installed curl
check_command_installed unzip

CURL_FLAGS=()
if [ -n "${GITHUB_TOKEN}" ]; then
  CURL_FLAGS+=("-H" "Authorization: token ${GITHUB_TOKEN}")
fi

TMP="$(mktemp -d)"
trap 'rm -rf "${TMP}"' EXIT

find "${DIR}" -type d -name 'v3*' | xargs rm -rf
pushd "${TMP}" >/dev/null
for tag in $(list-proto3-stable-release-tags-with-wkt); do
  mkdir -p "${tag}"
  pushd "${tag}" >/dev/null
	curl -sSL "${CURL_FLAGS[@]}" \
    "https://github.com/protocolbuffers/protobuf/releases/download/${tag}/protoc-$(echo ${tag} | sed 's/^v//')-linux-x86_64.zip" \
    -o protoc.zip
  unzip protoc.zip
  mkdir -p "${DIR}/${tag}"
  mv include/google "${DIR}/${tag}/google"
  popd >/dev/null
done
popd >/dev/null
