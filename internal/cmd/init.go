package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/remussoare/sdd-cook/internal/lifecycle"
	"github.com/remussoare/sdd-cook/internal/spec"
)

// Init scaffolds the minimum .spec/ structure. It never overwrites.
func Init() error {
	root := "."
	specDir := filepath.Join(root, ".spec")
	existed := spec.Exists(root)

	var created []string
	if err := os.MkdirAll(specDir, 0o755); err != nil {
		return err
	}
	for _, d := range spec.AllDirs() {
		p := filepath.Join(specDir, d)
		if _, err := os.Stat(p); os.IsNotExist(err) {
			if err := os.MkdirAll(p, 0o755); err != nil {
				return err
			}
			if err := os.WriteFile(filepath.Join(p, ".gitkeep"), nil, 0o644); err != nil {
				return err
			}
			created = append(created, ".spec/"+d)
		}
	}

	var discovery []string
	for _, f := range []string{"AGENTS.md", "CLAUDE.md", "CONSTITUTION.md", "README.md"} {
		if _, err := os.Stat(filepath.Join(root, f)); err == nil {
			discovery = append(discovery, f)
		}
	}

	s, err := spec.Load(root)
	if err != nil {
		return err
	}

	fmt.Println("SDD INIT")
	fmt.Printf("- Project type: %s\n", projectType(root))
	fmt.Printf("- Discovery: %s\n", joinOr(discovery, "nothing relevant found"))
	fmt.Printf("- Existing SDD artifacts: %s\n", joinOr(existingArtifacts(root), "none"))
	if len(created) > 0 {
		fmt.Printf("- Created: %s\n", joinOr(created, "none"))
	} else {
		fmt.Println("- Created: nothing (structure already present)")
	}
	fmt.Println("- Overwritten: nothing")
	fmt.Println("- Conflicts/drift: none detected by init")
	fmt.Println("- Initialization: PASS")
	next := "sdd new <name>"
	if existed && !s.Artifacts["objective"].Missing() {
		next = lifecycle.NextAction(s)
	}
	fmt.Printf("- Next action: %s\n", next)
	return nil
}