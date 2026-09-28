# sdd-cook bootstrap installer (Windows).
# Downloads the prebuilt sdd binary from GitHub releases and runs `sdd install`.
# Usage: irm https://raw.githubusercontent.com/remussoare/sdd-cook/main/install.ps1 | iex
# Env: $env:SDD_VERSION (tag, default latest), $env:SDD_BIN_DIR (default ~\.local\bin)
$ErrorActionPreference = "Stop"

$Repo = if ($env:SDD_REPO) { $env:SDD_REPO } else { "remussoare/sdd-cook" }
$Version = if ($env:SDD_VERSION) { $env:SDD_VERSION } else { "latest" }
$BinDir = if ($env:SDD_BIN_DIR) { $env:SDD_BIN_DIR } else { Join-Path $HOME ".local\bin" }

$Arch = switch ($env:PROCESSOR_ARCHITECTURE) {
  "AMD64" { "amd64" }
  "ARM64" { "arm64" }
  default { Write-Error "unsupported architecture: $env:PROCESSOR_ARCHITECTURE"; exit 1 }
}
$Asset = "sdd-windows-$Arch.exe"

if ($Version -eq "latest") {
  $Url = "https://github.com/$Repo/releases/latest/download/$Asset"
} else {
  $Url = "https://github.com/$Repo/releases/download/$Version/$Asset"
}

New-Item -ItemType Directory -Force -Path $BinDir | Out-Null
$Dest = Join-Path $BinDir "sdd.exe"
Write-Host "downloading sdd (windows/$Arch) to $Dest ..."
Invoke-WebRequest -Uri $Url -OutFile $Dest

$PathParts = $env:PATH -split ";"
if ($PathParts -notcontains $BinDir) {
  Write-Warning "$BinDir is not on PATH — add it to use sdd directly."
}

& $Dest install @args
