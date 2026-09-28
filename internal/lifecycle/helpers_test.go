package lifecycle

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/remussoare/sdd-cook/internal/spec"
)

func writeArt(t *testing.T, root, name, status string) {
	t.Helper()
	dir := filepath.Join(root, ".spec", spec.Dir(name))
	if err := os.MkdirAll(dir, 0o755); err != nil {
		t.Fatal(err)
	}
	content := "---\nartifact: " + name + "\ncomponent: demo\nevolution: v1\nstatus: " + status + "\n---\n\nbody\n"
	if err := os.WriteFile(filepath.Join(dir, spec.File(name)), []byte(content), 0o644); err != nil {
		t.Fatal(err)
	}
}

// fullAccepted writes a complete chain of accepted artifacts.
func fullAccepted(t *testing.T, root string) {
	t.Helper()
	for _, name := range spec.Chain {
		writeArt(t, root, name, string(spec.Accepted))
	}
}

func loadedSpec(t *testing.T, root string) *spec.Spec {
	t.Helper()
	s, err := spec.Load(root)
	if err != nil {
		t.Fatal(err)
	}
	return s
}

func emptySpec(t *testing.T) *spec.Spec {
	t.Helper()
	return loadedSpec(t, t.TempDir())
}