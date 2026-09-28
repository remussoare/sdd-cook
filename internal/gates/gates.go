// Package gates evaluates SDD gates from .spec/ state.
package gates

import (
	"fmt"

	"github.com/remussoare/sdd-cook/internal/spec"
)

// Result is the outcome of a gate evaluation.
type Result struct {
	Outcome string   // PASS, FAIL, BLOCKED, INCONCLUSIVE
	Details []string
}

func downstreamAccepted(s *spec.Spec, idx int) bool {
	for _, name := range spec.Chain[idx+1:] {
		if s.Artifacts[name].Accepted() {
			return true
		}
	}
	return false
}

// Evaluate checks the artifact chain. It reports drift (FAIL), missing
// prerequisites (BLOCKED) and unreadable state (INCONCLUSIVE).
func Evaluate(s *spec.Spec) Result {
	var r Result
	for i, name := range spec.Chain {
		a := s.Artifacts[name]
		switch {
		case a.Missing():
			if downstreamAccepted(s, i) {
				r.Outcome = "FAIL"
				r.Details = append(r.Details, fmt.Sprintf("%s missing while downstream artifacts are accepted (drift)", name))
				return r
			}
			r.Outcome = "BLOCKED"
			r.Details = append(r.Details, fmt.Sprintf("%s missing", name))
			return r
		case a.Status == spec.Invalid:
			r.Outcome = "INCONCLUSIVE"
			r.Details = append(r.Details, fmt.Sprintf("%s has missing or unreadable status in frontmatter", name))
			return r
		case a.Status == spec.Draft:
			if downstreamAccepted(s, i) {
				r.Outcome = "FAIL"
				r.Details = append(r.Details, fmt.Sprintf("%s is DRAFT while downstream artifacts are accepted (drift)", name))
				return r
			}
			r.Outcome = "BLOCKED"
			r.Details = append(r.Details, fmt.Sprintf("%s is DRAFT (not accepted)", name))
			return r
		}
	}
	r.Outcome = "PASS"
	r.Details = append(r.Details, "artifact chain objective → validation is accepted")
	if spec.ChainAllLocked(s) {
		r.Details = append(r.Details, "state is LOCKED")
	}
	return r
}