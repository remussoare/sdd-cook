package install

import (
	"fmt"
	"strings"
	"time"

	"github.com/remussoare/sdd-cook/internal/version"
)

// Update re-applies the embedded payloads to every tracked installation and
// reports whether a newer sdd release exists. Payload content always comes
// from this binary, so updating the skill means upgrading sdd first.
func Update(o Options) error {
	recs := Records()
	if len(recs) == 0 {
		return fmt.Errorf("no installations tracked — run: sdd install")
	}
	p := newPrompter(o.Stdin, o.out())
	interactive := p.interactive() && !o.Yes

	var plans []*Plan
	for _, rec := range recs {
		plan, err := buildPlan(rec.Root, Scope(rec.Scope), rec.Hosts)
		if err != nil {
			return err
		}
		plans = append(plans, plan)
	}
	for _, plan := range plans {
		printPlan(o.out(), plan)
	}
	if o.DryRun {
		fmt.Fprintln(o.out(), "dry-run: nothing applied.")
		return nil
	}
	confirmed := o.Yes || !interactive
	if !confirmed {
		var parts []string
		for _, plan := range plans {
			parts = append(parts, summarize(plan))
		}
		confirmed = p.askConfirm(strings.Join(parts, "\n"))
	}
	if !confirmed {
		fmt.Fprintln(o.out(), "aborted.")
		return nil
	}
	for i, plan := range plans {
		if err := apply(plan); err != nil {
			return err
		}
		recs[i].Version = version.Version
		recs[i].Files = manifest(plan)
		recs[i].InstalledAt = time.Now().UTC().Format(time.RFC3339)
		if err := saveRecord(recs[i]); err != nil {
			fmt.Fprintln(o.errOut(), "warning: could not record installation:", err)
		}
		fmt.Fprintf(o.out(), "updated %s (sdd-cook %s)\n", plan.Root, version.Version)
	}

	// Release check is best-effort: offline or API errors never fail update.
	if tag, err := LatestReleaseTag(); err == nil {
		if NewerThan(version.Version, tag) {
			fmt.Fprintf(o.out(), "newer sdd available: %s (installed %s)\n", tag, version.Version)
			fmt.Fprintln(o.out(), "upgrade, then re-run: sdd update")
			fmt.Fprintln(o.out(), "  go install github.com/remussoare/sdd-cook@latest")
		} else {
			fmt.Fprintf(o.out(), "sdd is up to date (%s)\n", version.Version)
		}
	} else {
		fmt.Fprintln(o.out(), "release check skipped (offline or API unavailable).")
	}
	return nil
}
