// Package sddcook exposes the canonical skill sources embedded in the binary.
// The installer (internal/install) and transpiler (internal/transpile) read
// from SkillsFS, so the sdd binary carries the full skill payload and the
// client never needs to clone the repository or build anything.
package sddcook

import "embed"

// SkillsFS is the embedded canonical skill tree (skills/<name>/...).
//
//go:embed all:skills
var SkillsFS embed.FS

// SkillName is the canonical skill installed by this CLI.
const SkillName = "sdd-cook"
