package cmd

import (
	"fmt"

	"github.com/remussoare/sdd-cook/internal/lifecycle"
	"github.com/remussoare/sdd-cook/internal/spec"
)

// validLabel reports whether a user-supplied label is safe to embed in a
// YAML frontmatter block and a file body (no newlines, colons or control
// characters that could inject fields or corrupt parsing).
func validLabel(s string) bool {
	if s == "" || len(s) > 200 {
		return false
	}
	for _, r := range s {
		if r < 0x20 || r == 0x7f || r == ':' {
			return false
		}
	}
	return true
}

// New starts a new component or, after LOCKED, an evolution through a
// change request. It never modifies LOCKED artifacts silently.
func New(args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("usage: sdd new <name>")
	}
	name := args[0]
	if !validLabel(name) {
		return fmt.Errorf("invalid name %q: must be non-empty, at most 200 bytes, without newlines, colons or control characters", name)
	}
	s, err := load()
	if err != nil {
		return err
	}
	obj := s.Artifacts["objective"]

	if obj.Missing() {
		if err := spec.CreateObjective(".", name, "v1"); err != nil {
			return err
		}
		fmt.Println("OBJECTIVE: DRAFT created at .spec/objective/objective.md")
		fmt.Println("Fill PURPOSE, SCOPE, CONSTRAINTS and ACCEPTANCE CRITERIA, then run: sdd accept objective")
		return nil
	}

	if spec.ChainAllLocked(s) {
		if p, active := spec.ActiveChangeRequest("."); active {
			return fmt.Errorf("BLOCKED: a change request is already in progress (%s) — resolve it before starting a new evolution", p)
		}
		next, err := spec.NextEvolution(obj.Evolution)
		if err != nil {
			return fmt.Errorf("BLOCKED: %v — rename the evolution label in the chain artifacts before starting a new one", err)
		}
		if err := spec.CreateChangeRequest(".", obj.Component, next, name); err != nil {
			return err
		}
		fmt.Printf("CHANGE REQUEST: PROPOSED created in .spec/changes/ (evolution %s)\n", next)
		fmt.Println("LOCKED state is preserved. Complete impact analysis, obtain the decision, then evolve artifacts via /sdd:change.")
		return nil
	}
	return fmt.Errorf("BLOCKED: chain in progress at %s — finish the current evolution before starting a new one", lifecycle.CurrentPhase(s))
}
