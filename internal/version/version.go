// Package version carries the CLI build version.
// Release builds override Version via ldflags:
//
//	go build -ldflags "-X github.com/remussoare/sdd-cook/internal/version.Version=v0.1.0" ./cmd/sdd
package version

// Version is the CLI version. Defaults to "dev" for local builds.
var Version = "dev"

// SourceRepo is the canonical repository of sdd-cook.
const SourceRepo = "github.com/remussoare/sdd-cook"

// SourceURL is the canonical clone URL of sdd-cook.
const SourceURL = "https://github.com/remussoare/sdd-cook.git"
