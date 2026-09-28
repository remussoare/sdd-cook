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
echo "downloading sdd ($OS/$ARCH) to $BIN_DIR/sdd ..."
curl -fsSL "$URL" -o "$BIN_DIR/sdd"
chmod +x "$BIN_DIR/sdd"

case ":$PATH:" in
  *":$BIN_DIR:"*) ;;
  *) echo "note: $BIN_DIR is not on PATH — add it to use sdd directly." >&2 ;;
esac

exec "$BIN_DIR/sdd" install "$@"
