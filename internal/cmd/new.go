package cmd

import (
	"fmt"

	"github.com/remussoare/sdd-cook/internal/spec"
)

// New starts a new component or, after LOCKED, an evolution through a
// change request. It never modifies LOCKED artifacts silently.
func New(args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("usage: sdd new <name>")
	}
	name := args[0]
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
		next := spec.NextEvolution(obj.Evolution)
		if err := spec.CreateChangeRequest(".", obj.Component, next, name); err != nil {
			return err
		}
		fmt.Printf("CHANGE REQUEST: PROPOSED created in .spec/changes/ (evolution %s)\n", next)
		fmt.Println("LOCKED state is preserved. Complete impact analysis, obtain the decision, then evolve artifacts via /sdd:change.")
		return nil
	}
	return fmt.Errorf("BLOCKED: an evolution is already in progress and is not LOCKED — finish it before starting a new one")
}