package main

import (
	"fmt"
	"os"

	"github.com/remussoare/sdd-cook/internal/cmd"
)

const usage = `sdd — Specification-Driven Development state manager (sdd-cook)

Usage:

  sdd init              Create the minimum .spec/ structure (never overwrites)
  sdd new <name>        Start a new component or (after LOCKED) an evolution
  sdd status            Show the real SDD state of the project
  sdd check             Evaluate the artifact-chain gate
  sdd show <artifact>   Print an artifact (objective, domain, ...)
  sdd accept <artifact> Record acceptance of a DRAFT artifact
  sdd lock              Lock the accepted validated state
  sdd install [flags]   Install the sdd-cook skill for AI tools (scope, hosts)
  sdd update [flags]    Re-apply skill payloads to tracked installs
  sdd version           Print version and platform

  install flags: --scope repo|home, --hosts opencode,cursor,vscode|all,
                 --target <dir>, --dry-run, --yes

AI work (/sdd:plan, /sdd:run, /sdd:validate, /sdd:change) stays with the agent;
this CLI manages and enforces the deterministic state.`

func main() {
	if len(os.Args) < 2 {
		fmt.Println(usage)
		os.Exit(2)
	}
	var err error
	switch os.Args[1] {
	case "init":
		err = cmd.Init()
	case "new":
		err = cmd.New(os.Args[2:])
	case "status":
		err = cmd.Status()
	case "check":
		err = cmd.Check()
	case "show":
		err = cmd.Show(os.Args[2:])
	case "accept":
		err = cmd.Accept(os.Args[2:])
	case "lock":
		err = cmd.Lock()
	case "install":
		err = cmd.Install(os.Args[2:])
	case "update":
		err = cmd.Update(os.Args[2:])
	case "version":
		err = cmd.Version()
	case "help", "-h", "--help":
		fmt.Println(usage)
		return
	default:
		fmt.Fprintf(os.Stderr, "unknown command %q\n\n%s\n", os.Args[1], usage)
		os.Exit(2)
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}