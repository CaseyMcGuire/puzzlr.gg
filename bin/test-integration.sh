#!/usr/bin/env bash
set -euo pipefail

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
cd "$ROOT_DIR"

if ! command -v docker >/dev/null 2>&1; then
  echo "Docker is required for integration tests but is not installed."
  exit 1
fi

if ! docker info >/dev/null 2>&1; then
  echo "Docker is required for integration tests but is not running."
  exit 1
fi

RUN_INTEGRATION_TESTS=1 go test -tags=integration ./src/server/graphql/resolvers "$@"
