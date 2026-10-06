package cmd

import (
	"fmt"

	"github.com/remussoare/sdd-cook/internal/gates"
	"github.com/remussoare/sdd-cook/internal/lifecycle"
)

// Check evaluates the artifact-chain gate. It exits non-zero when the gate
// outcome is FAIL or INCONCLUSIVE so scripts and CI can rely on it. BLOCKED
// (chain not finished yet) still exits 0.
func Check() error {
	s, err := load()
	if err != nil {
		return err
	}
	g := gates.Evaluate(s)
	fmt.Println("SDD CHECK")
	fmt.Printf("- Gate: %s\n", g.Outcome)
	for _, d := range g.Details {
		fmt.Printf("  - %s\n", d)
	}
	fmt.Printf("- Next action: %s\n", lifecycle.NextAction(s))
	if g.Outcome == "FAIL" || g.Outcome == "INCONCLUSIVE" {
		return fmt.Errorf("gate %s", g.Outcome)
	}
	return nil
}