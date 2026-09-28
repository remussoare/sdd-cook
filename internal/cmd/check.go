package cmd

import (
	"fmt"

	"github.com/remussoare/sdd-cook/internal/gates"
	"github.com/remussoare/sdd-cook/internal/lifecycle"
)

// Check evaluates the artifact-chain gate.
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
	return nil
}