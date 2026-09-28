package cmd

import (
	"fmt"

	"github.com/remussoare/sdd-cook/internal/spec"
)

// Lock locks the accepted validated state.
func Lock() error {
	s, err := load()
	if err != nil {
		return err
	}
	for _, name := range spec.Chain {
		a := s.Artifacts[name]
		if a.Missing() {
			return fmt.Errorf("BLOCKED: %s missing — cannot lock", name)
		}
		switch a.Status {
		case spec.Draft:
			return fmt.Errorf("BLOCKED: %s is DRAFT — run: sdd accept %s", name, name)
		case spec.Invalid:
			return fmt.Errorf("BLOCKED: %s has an invalid frontmatter status", name)
		}
	}
	if spec.ChainAllLocked(s) {
		fmt.Println("State is already LOCKED.")
		return nil
	}
	for _, name := range spec.Chain {
		a := s.Artifacts[name]
		if a.Status == spec.Accepted {
			if err := a.SetStatus(spec.Locked); err != nil {
				return err
			}
			fmt.Printf("%s: ACCEPTED → LOCKED\n", name)
		}
	}
	fmt.Println("Lifecycle: LOCKED")
	fmt.Println("Next action: /sdd:change (to start a new evolution)")
	return nil
}