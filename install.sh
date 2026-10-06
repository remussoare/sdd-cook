#!/bin/sh
# sdd-cook bootstrap installer (macOS / Linux).
# Downloads the prebuilt sdd binary from GitHub releases and runs `sdd install`.
# Usage: curl -fsSL https://raw.githubusercontent.com/remussoare/sdd-cook/main/install.sh | sh
#        (pass install flags after -- : ... | sh -s -- --scope repo --hosts all --yes)
# Env: SDD_VERSION (tag, default latest), SDD_BIN_DIR (default ~/.local/bin)
set -eu

REPO="${SDD_REPO:-remussoare/sdd-cook}"
VERSION="${SDD_VERSION:-latest}"
BIN_DIR="${SDD_BIN_DIR:-$HOME/.local/bin}"

OS="$(uname -s | tr '[:upper:]' '[:lower:]')"
ARCH="$(uname -m)"
case "$ARCH" in
  x86_64|amd64) ARCH="amd64" ;;
  arm64|aarch64) ARCH="arm64" ;;
  *) echo "unsupported architecture: $ARCH" >&2; exit 1 ;;
esac
case "$OS" in
  darwin|linux) ;;
  *) echo "unsupported OS: $OS (use install.ps1 on Windows)" >&2; exit 1 ;;
esac

if [ "$VERSION" = "latest" ]; then
  URL="https://github.com/${REPO}/releases/latest/download/sdd-${OS}-${ARCH}"
else
  URL="https://github.com/${REPO}/releases/download/${VERSION}/sdd-${OS}-${ARCH}"
fi

mkdir -p "$BIN_DIR"
TMP="$(mktemp -d)"
trap 'rm -rf "$TMP"' EXIT
echo "downloading sdd ($OS/$ARCH) to $BIN_DIR/sdd ..."
curl --fail --retry 3 -fsSL "$URL" -o "$TMP/sdd"

# Verify the checksum against the release SHA256SUMS when available.
ASSET="$(basename "$URL")"
SUMS_URL="${URL%/*}/SHA256SUMS"
if curl --fail --retry 3 -fsSL "$SUMS_URL" -o "$TMP/SHA256SUMS" 2>/dev/null; then
  expected="$(awk -v a="$ASSET" '$2 == a || $2 == "*"a {print $1}' "$TMP/SHA256SUMS")"
  if [ -z "$expected" ]; then
    echo "error: $ASSET not listed in SHA256SUMS" >&2
    exit 1
  fi
  if command -v sha256sum >/dev/null 2>&1; then
    actual="$(sha256sum "$TMP/sdd" | awk '{print $1}')"
  elif command -v shasum >/dev/null 2>&1; then
    actual="$(shasum -a 256 "$TMP/sdd" | awk '{print $1}')"
  else
    echo "warning: no sha256 tool available — skipping checksum verification" >&2
    actual="$expected"
  fi
  if [ "$actual" != "$expected" ]; then
    echo "error: checksum mismatch (expected $expected, got $actual)" >&2
    exit 1
  fi
  echo "checksum ok"
else
  echo "warning: SHA256SUMS unavailable for this release — skipping checksum verification" >&2
fi

mv "$TMP/sdd" "$BIN_DIR/sdd"
chmod +x "$BIN_DIR/sdd"

case ":$PATH:" in
  *":$BIN_DIR:"*) ;;
  *) echo "note: $BIN_DIR is not on PATH — add it to use sdd directly." >&2 ;;
esac

exec "$BIN_DIR/sdd" install "$@"
