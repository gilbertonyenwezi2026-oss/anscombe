#!/usr/bin/env bash
set -euo pipefail

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"

echo "== Go =="
(
  cd "$ROOT_DIR/go"
  go test ./... -bench=. -benchmem
  go run ./...
)

echo

echo "== Python =="
python3 "$ROOT_DIR/python/anscombe.py"

echo

echo "== R =="
Rscript "$ROOT_DIR/r/anscombe.R"
