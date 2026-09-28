// Package lifecycle derives the current phase and next authorized action
// from .spec/ state. It never invents states for missing artifacts.
package lifecycle

import (
	"fmt"

	"github.com/remussoare/sdd-cook/internal/spec"
)

var phaseLabel = map[string]string{
	"objective":     "OBJECTIVE",
	"domain":        "DOMAIN",
	"specification": "SPECIFICATION",
	"plan":          "PLAN",
	"tests":         "TESTS",
	"tasks":         "TASKS",
	"validation":    "VALIDATION",
}

// Label returns the workflow phase label of a chain artifact.
func Label(name string) string { return phaseLabel[name] }

// CurrentPhase derives the phase the project is really in.
func CurrentPhase(s *spec.Spec) string {
	tasks := s.Artifacts["tasks"]
	for _, name := range spec.Chain {
		a := s.Artifacts[name]
		if a.Missing() {
			if name == "validation" && tasks.Accepted() {
				return "IMPLEMENTATION"
			}
			return phaseLabel[name] + " (not started)"
		}
		if !a.Accepted() {
			return phaseLabel[name] + " (draft)"
		}
	}
	if s.Artifacts["validation"].Status == spec.Accepted {
		return "LOCKED (pending lock)"
	}
	return "LOCKED"
}

// NextAction returns the next authorized action derived only from actual artifacts.
func NextAction(s *spec.Spec) string {
	if s.Artifacts["objective"].Missing() {
		return "sdd new <name>"
	}
	for _, name := range []string{"objective", "domain", "specification"} {
		if !s.Artifacts[name].Accepted() {
			return "/sdd:new"
		}
	}
	for _, name := range []string{"plan", "tests", "tasks"} {
		if !s.Artifacts[name].Accepted() {
			return "/sdd:plan"
		}
	}
	v := s.Artifacts["validation"]
	switch {
	case v.Missing():
		return "/sdd:run"
	case v.Status == spec.Draft || v.Status == spec.Invalid:
		return "/sdd:validate"
	case v.Status == spec.Accepted:
		return "sdd lock"
	}
	return "/sdd:change"
}

// CanAccept verifies that an artifact can be accepted: it exists, is a draft
// and every upstream artifact is accepted.
func CanAccept(s *spec.Spec, name string) error {
	idx := -1
	for i, n := range spec.Chain {
		if n == name {
			idx = i
			break
		}
	}
	if idx < 0 {
		return fmt.Errorf("unknown artifact %q", name)
	}
	for _, prev := range spec.Chain[:idx] {
		if pa := s.Artifacts[prev]; !pa.Accepted() {
			return fmt.Errorf("upstream artifact %q is not accepted", prev)
		}
	}
	a := s.Artifacts[name]
	if a.Missing() {
		return fmt.Errorf("artifact %q does not exist", name)
	}
	if a.Status != spec.Draft {
		return fmt.Errorf("artifact %q is %s, only DRAFT can be accepted", name, a.Status)
	}
	return nil
}

// ImplementationStatus derives the implementation line of sdd status.
func ImplementationStatus(s *spec.Spec) string {
	switch {
	case s.Artifacts["validation"].Accepted():
		return "DONE"
	case s.Artifacts["tasks"].Accepted():
		return "IN PROGRESS"
	default:
		return "PENDING"
	}
}