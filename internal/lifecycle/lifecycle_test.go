package lifecycle

import (
	"github.com/remussoare/sdd-cook/internal/spec"
	"testing"
)

func TestCurrentPhaseEmpty(t *testing.T) {
	s := emptySpec(t)
	if got := CurrentPhase(s); got != "OBJECTIVE (not started)" {
		t.Fatalf("got %q", got)
	}
	if got := NextAction(s); got != "sdd new <name>" {
		t.Fatalf("got %q", got)
	}
}

func TestCurrentPhaseObjectiveDraft(t *testing.T) {
	root := t.TempDir()
	writeArt(t, root, "objective", string(spec.Draft))
	s := loadedSpec(t, root)
	if got := CurrentPhase(s); got != "OBJECTIVE (draft)" {
		t.Fatalf("got %q", got)
	}
	if got := NextAction(s); got != "/sdd:new" {
		t.Fatalf("got %q", got)
	}
}

func TestCurrentPhasePlanPhase(t *testing.T) {
	root := t.TempDir()
	writeArt(t, root, "objective", string(spec.Locked))
	writeArt(t, root, "domain", string(spec.Locked))
	writeArt(t, root, "specification", string(spec.Locked))
	s := loadedSpec(t, root)
	if got := CurrentPhase(s); got != "PLAN (not started)" {
		t.Fatalf("got %q", got)
	}
	if got := NextAction(s); got != "/sdd:plan" {
		t.Fatalf("got %q", got)
	}
}

func TestCurrentPhaseImplementation(t *testing.T) {
	root := t.TempDir()
	for _, name := range []string{"objective", "domain", "specification", "plan", "tests", "tasks"} {
		writeArt(t, root, name, string(spec.Accepted))
	}
	s := loadedSpec(t, root)
	if got := CurrentPhase(s); got != "IMPLEMENTATION" {
		t.Fatalf("got %q", got)
	}
	if got := NextAction(s); got != "/sdd:run" {
		t.Fatalf("got %q", got)
	}
	if got := ImplementationStatus(s); got != "IN PROGRESS" {
		t.Fatalf("implementation: got %q", got)
	}
}

func TestCurrentPhaseValidationDraft(t *testing.T) {
	root := t.TempDir()
	for _, name := range []string{"objective", "domain", "specification", "plan", "tests", "tasks"} {
		writeArt(t, root, name, string(spec.Accepted))
	}
	writeArt(t, root, "validation", string(spec.Draft))
	s := loadedSpec(t, root)
	if got := CurrentPhase(s); got != "VALIDATION (draft)" {
		t.Fatalf("got %q", got)
	}
	if got := NextAction(s); got != "/sdd:validate" {
		t.Fatalf("got %q", got)
	}
}

func TestCurrentPhasePendingLock(t *testing.T) {
	root := t.TempDir()
	fullAccepted(t, root)
	s := loadedSpec(t, root)
	if got := CurrentPhase(s); got != "LOCKED (pending lock)" {
		t.Fatalf("got %q", got)
	}
	if got := NextAction(s); got != "sdd lock" {
		t.Fatalf("got %q", got)
	}
	if got := ImplementationStatus(s); got != "DONE" {
		t.Fatalf("implementation: got %q", got)
	}
}

func TestCurrentPhaseLocked(t *testing.T) {
	root := t.TempDir()
	for _, name := range spec.Chain {
		writeArt(t, root, name, string(spec.Locked))
	}
	s := loadedSpec(t, root)
	if got := CurrentPhase(s); got != "LOCKED" {
		t.Fatalf("got %q", got)
	}
	if got := NextAction(s); got != "/sdd:change" {
		t.Fatalf("got %q", got)
	}
}

func TestCanAccept(t *testing.T) {
	root := t.TempDir()
	writeArt(t, root, "objective", string(spec.Draft))
	writeArt(t, root, "domain", string(spec.Draft))
	s := loadedSpec(t, root)

	if err := CanAccept(s, "domain"); err == nil {
		t.Fatal("expected BLOCKED when upstream objective is not accepted")
	}
	if err := CanAccept(s, "nonexistent"); err == nil {
		t.Fatal("expected error for unknown artifact")
	}
	writeArt(t, root, "objective", string(spec.Accepted))
	s = loadedSpec(t, root)
	if err := CanAccept(s, "domain"); err != nil {
		t.Fatalf("expected accept allowed, got %v", err)
	}
	writeArt(t, root, "domain", string(spec.Accepted))
	s = loadedSpec(t, root)
	if err := CanAccept(s, "domain"); err == nil {
		t.Fatal("expected error accepting non-draft artifact")
	}
}