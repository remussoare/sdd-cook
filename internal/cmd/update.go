package cmd

import (
	"flag"

	"github.com/remussoare/sdd-cook/internal/install"
)

// Update runs `sdd update`: re-applies embedded payloads to every tracked
// installation and reports newer releases.
func Update(args []string) error {
	fs := flag.NewFlagSet("update", flag.ContinueOnError)
	dryRun := fs.Bool("dry-run", false, "print plan, apply nothing")
	yes := fs.Bool("yes", false, "skip confirmation")
	if err := fs.Parse(args); err != nil {
		return err
	}
	return install.Update(install.Options{DryRun: *dryRun, Yes: *yes})
}
