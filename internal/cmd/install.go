package cmd

import (
	"flag"
	"fmt"
	"strings"

	"github.com/remussoare/sdd-cook/internal/install"
)

// Install runs `sdd install`: writes the embedded skill payloads for the
// selected hosts into the scope root (project dir or home dir).
func Install(args []string) error {
	fs := flag.NewFlagSet("install", flag.ContinueOnError)
	scope := fs.String("scope", "", "install scope: repo (project dir) or home")
	hosts := fs.String("hosts", "", "comma-separated hosts: opencode,cursor,vscode (or all)")
	target := fs.String("target", "", "override scope root directory")
	dryRun := fs.Bool("dry-run", false, "print plan, apply nothing")
	yes := fs.Bool("yes", false, "skip confirmation")
	source := fs.String("source", "", "source URL recorded on install")
	if err := fs.Parse(args); err != nil {
		return err
	}
	var hostList []string
	if *hosts != "" {
		if strings.ToLower(strings.TrimSpace(*hosts)) == "all" {
			hostList = nil // resolved to all in install.Run
			hostList = []string{"opencode", "cursor", "vscode"}
		} else {
			for _, h := range strings.Split(*hosts, ",") {
				if h = strings.TrimSpace(h); h != "" {
					hostList = append(hostList, h)
				}
			}
		}
	}
	var sc install.Scope
	switch strings.ToLower(*scope) {
	case "", "repo", "project":
		if *scope != "" {
			sc = install.ScopeRepo
		}
	case "home", "global":
		sc = install.ScopeHome
	default:
		return fmt.Errorf("unknown scope %q (valid: repo, home)", *scope)
	}
	return install.Run(install.Options{
		Scope:  sc,
		Hosts:  hostList,
		Target: *target,
		DryRun: *dryRun,
		Yes:    *yes,
		Source: *source,
	})
}
