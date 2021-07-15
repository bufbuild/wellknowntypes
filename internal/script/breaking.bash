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

check_command_installed buf

ls | grep ^v3 | sort-semver-tags | list-pairs | xargs  printf -- 'breaking --against %s %s\n' | while read line; do
  IFS=' ' read -r -a args <<< "${line}"
  echo buf "${args[@]}"
  buf "${args[@]}" || true
done
