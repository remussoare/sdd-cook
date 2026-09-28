package cmd

import (
	"fmt"
	"strings"

	"github.com/remussoare/sdd-cook/internal/gates"
	"github.com/remussoare/sdd-cook/internal/lifecycle"
	"github.com/remussoare/sdd-cook/internal/spec"
)

// Status shows the real SDD state of the project.
func Status() error {
	s, err := load()
	if err != nil {
		return err
	}
	obj := s.Artifacts["objective"]
	g := gates.Evaluate(s)

	blockers, drift := "-", "-"
	switch g.Outcome {
	case "FAIL":
		drift = strings.Join(g.Details, "; ")
	case "BLOCKED", "INCONCLUSIVE":
		blockers = strings.Join(g.Details, "; ")
	}

	fmt.Println("SDD STATUS")
	fmt.Printf("- Project: %s\n", baseName(s.Root))
	fmt.Printf("- Component: %s\n", orDefault(obj.Component, "-"))
	fmt.Printf("- Evolution: %s\n", orDefault(obj.Evolution, "-"))
	fmt.Printf("- Current phase: %s\n", lifecycle.CurrentPhase(s))
	fmt.Printf("- Objective: %s\n", statusLine(s.Artifacts["objective"]))
	fmt.Printf("- Domain: %s\n", statusLine(s.Artifacts["domain"]))
	fmt.Printf("- Specification: %s\n", statusLine(s.Artifacts["specification"]))
	fmt.Printf("- Plan: %s\n", statusLine(s.Artifacts["plan"]))
	fmt.Printf("- Tests: %s\n", statusLine(s.Artifacts["tests"]))
	fmt.Printf("- Tasks: %s\n", statusLine(s.Artifacts["tasks"]))
	fmt.Printf("- Implementation: %s\n", lifecycle.ImplementationStatus(s))
	fmt.Printf("- Validation: %s\n", statusLine(s.Artifacts["validation"]))
	fmt.Printf("- Lifecycle: %s\n", chainLine(s))
	fmt.Printf("- Blockers: %s\n", blockers)
	fmt.Printf("- Drift: %s\n", drift)
	fmt.Printf("- Next action: %s\n", lifecycle.NextAction(s))
	return nil
}

func chainLine(s *spec.Spec) string {
	parts := make([]string, 0, len(spec.Chain))
	for _, name := range spec.Chain {
		parts = append(parts, lifecycle.Label(name)+":"+statusLine(s.Artifacts[name]))
	}
	return strings.Join(parts, " → ")
}