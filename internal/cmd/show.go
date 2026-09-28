package cmd

import (
	"fmt"
	"os"
)

// Show prints an artifact file.
func Show(args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("usage: sdd show <objective|domain|specification|plan|tests|tasks|validation>")
	}
	s, err := load()
	if err != nil {
		return err
	}
	a, ok := s.Artifacts[args[0]]
	if !ok {
		return fmt.Errorf("unknown artifact %q", args[0])
	}
	if a.Missing() {
		return fmt.Errorf("artifact %q is MISSING", args[0])
	}
	data, err := os.ReadFile(a.Path)
	if err != nil {
		return err
	}
	fmt.Print(string(data))
	return nil
}