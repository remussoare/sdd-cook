package cmd

import (
	"fmt"

	"github.com/remussoare/sdd-cook/internal/lifecycle"
	"github.com/remussoare/sdd-cook/internal/spec"
)

// Accept records acceptance of a DRAFT artifact whose upstream chain is accepted.
func Accept(args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("usage: sdd accept <objective|domain|specification|plan|tests|tasks|validation>")
	}
	name := args[0]
	s, err := load()
	if err != nil {
		return err
	}
	if err := lifecycle.CanAccept(s, name); err != nil {
		return fmt.Errorf("BLOCKED: %v", err)
	}
	if err := s.Artifacts[name].SetStatus(spec.Accepted); err != nil {
		return err
	}
	fmt.Printf("%s: DRAFT → ACCEPTED\n", name)
	s2, err := spec.Load(".")
	if err != nil {
		return err
	}
	fmt.Printf("Next action: %s\n", lifecycle.NextAction(s2))
	return nil
}