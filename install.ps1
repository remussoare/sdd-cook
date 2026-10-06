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
$Tmp = New-TemporaryFile
Invoke-WebRequest -UseBasicParsing -Uri $Url -OutFile $Tmp

# Verify the checksum against the release SHA256SUMS when available.
$SumsUrl = $Url.Substring(0, $Url.LastIndexOf('/')) + "/SHA256SUMS"
$Verified = $false
try {
  $Sums = (Invoke-WebRequest -UseBasicParsing -Uri $SumsUrl).Content
  $Line = ($Sums -split "`n") | Where-Object { $_ -match [regex]::Escape($Asset) } | Select-Object -First 1
  if ($Line) {
    $Expected = ($Line -split '\s+')[0].Trim()
    $Actual = (Get-FileHash -Algorithm SHA256 $Tmp).Hash.ToLower()
    if ($Actual -ne $Expected.ToLower()) {
      Write-Error "checksum mismatch (expected $Expected, got $Actual)"
      exit 1
    }
    Write-Host "checksum ok"
    $Verified = $true
  }
} catch {
  Write-Warning "SHA256SUMS unavailable for this release — skipping checksum verification"
}
if (-not $Verified) {
  Write-Warning "binary downloaded WITHOUT checksum verification"
}

Move-Item -Force $Tmp $Dest

$PathParts = $env:PATH -split ";"
if ($PathParts -notcontains $BinDir) {
  Write-Warning "$BinDir is not on PATH — add it to use sdd directly."
}

& $Dest install @args
