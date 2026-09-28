package cmd

import (
	"fmt"
	"runtime"

	"github.com/remussoare/sdd-cook/internal/version"
)

// Version prints the CLI version and platform.
func Version() error {
	fmt.Printf("sdd %s (%s/%s)\n", version.Version, runtime.GOOS, runtime.GOARCH)
	return nil
}
