package gates

import (
	"os"
	"path/filepath"
	"github.com/remussoare/sdd-cook/internal/lifecycle"
	"github.com/remussoare/sdd-cook/internal/spec"
	"testing"
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

func TestEvaluateEmptySpec(t *testing.T) {
	s, _ := spec.Load(t.TempDir())
	g := Evaluate(s)
	if g.Outcome != "BLOCKED" {
		t.Fatalf("got %q", g.Outcome)
	}
}

func TestEvaluateDraftBlocks(t *testing.T) {
	root := t.TempDir()
	writeArt(t, root, "objective", string(spec.Draft))
	s, _ := spec.Load(root)
	g := Evaluate(s)
	if g.Outcome != "BLOCKED" {
		t.Fatalf("got %q", g.Outcome)
	}
}

func TestEvaluateDriftFails(t *testing.T) {
	root := t.TempDir()
	writeArt(t, root, "objective", string(spec.Draft))
	writeArt(t, root, "domain", string(spec.Accepted))
	s, _ := spec.Load(root)
	g := Evaluate(s)
	if g.Outcome != "FAIL" {
		t.Fatalf("got %q", g.Outcome)
	}
}

func TestEvaluateMissingWithDownstreamFails(t *testing.T) {
	root := t.TempDir()
	writeArt(t, root, "domain", string(spec.Accepted))
	s, _ := spec.Load(root)
	g := Evaluate(s)
	if g.Outcome != "FAIL" {
		t.Fatalf("got %q", g.Outcome)
	}
}

func TestEvaluateInvalidInconclusive(t *testing.T) {
	root := t.TempDir()
	writeArt(t, root, "objective", "ACCEPTED")
	writeArt(t, root, "domain", "WEIRD")
	s, _ := spec.Load(root)
	g := Evaluate(s)
	if g.Outcome != "INCONCLUSIVE" {
		t.Fatalf("got %q", g.Outcome)
	}
}

func TestEvaluatePass(t *testing.T) {
	root := t.TempDir()
	for _, name := range spec.Chain {
		writeArt(t, root, name, string(spec.Accepted))
	}
	s, _ := spec.Load(root)
	g := Evaluate(s)
	if g.Outcome != "PASS" {
		t.Fatalf("got %q", g.Outcome)
	}
}

func TestEvaluatePassLocked(t *testing.T) {
	root := t.TempDir()
	for _, name := range spec.Chain {
		writeArt(t, root, name, string(spec.Locked))
	}
	s, _ := spec.Load(root)
	g := Evaluate(s)
	if g.Outcome != "PASS" {
		t.Fatalf("got %q", g.Outcome)
	}
	if len(g.Details) == 0 || g.Details[len(g.Details)-1] != "state is LOCKED" {
		t.Fatalf("expected locked detail, got %v", g.Details)
	}
}

// cross-check with lifecycle to keep both packages consistent.
func TestConsistencyWithLifecycle(t *testing.T) {
	root := t.TempDir()
	for _, name := range spec.Chain {
		writeArt(t, root, name, string(spec.Accepted))
	}
	s, _ := spec.Load(root)
	g := Evaluate(s)
	if g.Outcome != "PASS" || lifecycle.NextAction(s) != "sdd lock" {
		t.Fatalf("inconsistent: %q / %q", g.Outcome, lifecycle.NextAction(s))
	}
}