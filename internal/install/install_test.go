package install

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func testOptions(t *testing.T, target string) Options {
	t.Helper()
	cfg := t.TempDir()
	t.Setenv("SDD_CONFIG_DIR", cfg)
	t.Setenv("CI", "true") // force non-interactive
	return Options{
		Scope:  ScopeRepo,
		Hosts:  []string{"cursor", "vscode"},
		Target: target,
		Yes:    true,
		Stdout: &bytes.Buffer{},
		Stderr: &bytes.Buffer{},
	}
}

func TestRunCreatesAndIsIdempotent(t *testing.T) {
	target := t.TempDir()
	if err := Run(testOptions(t, target)); err != nil {
		t.Fatal(err)
	}
	for _, p := range []string{
		".cursor/agents/sdd-cook.md",
		".cursor/rules/sdd-cook.mdc",
		".github/skills/sdd-cook/SKILL.md",
		"AGENTS.md",
	} {
		if _, err := os.Stat(filepath.Join(target, filepath.FromSlash(p))); err != nil {
			t.Errorf("missing installed file %s: %v", p, err)
		}
	}
	// Opencode host must NOT be installed (not requested).
	if _, err := os.Stat(filepath.Join(target, ".agents")); !os.IsNotExist(err) {
		t.Errorf("unexpected .agents dir for unrequested host")
	}

	// Second run: everything skips.
	var out bytes.Buffer
	o := testOptions(t, target)
	o.Stdout = &out
	if err := Run(o); err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(out.String(), "0 create, 0 update") {
		t.Errorf("second run not idempotent:\n%s", out.String())
	}
}

func TestUpdateAppliesLocalEditsAndKeepsAgents(t *testing.T) {
	target := t.TempDir()
	o := testOptions(t, target)
	if err := Run(o); err != nil {
		t.Fatal(err)
	}
	rules := filepath.Join(target, ".cursor", "rules", "sdd-cook.mdc")
	data, _ := os.ReadFile(rules)
	if err := os.WriteFile(rules, append(data, []byte("\n# local\n")...), 0o644); err != nil {
		t.Fatal(err)
	}
	agents := filepath.Join(target, "AGENTS.md")
	adata, _ := os.ReadFile(agents)
	if err := os.WriteFile(agents, append(adata, []byte("\n# local\n")...), 0o644); err != nil {
		t.Fatal(err)
	}

	var out bytes.Buffer
	o.Stdout = &out
	o.Yes = true
	if err := Update(o); err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(out.String(), "update .cursor/rules/sdd-cook.mdc") {
		t.Errorf("expected rules update:\n%s", out.String())
	}
	if !strings.Contains(out.String(), "keep   AGENTS.md") {
		t.Errorf("expected AGENTS.md keep:\n%s", out.String())
	}
	kept, _ := os.ReadFile(agents)
	if !strings.Contains(string(kept), "# local") {
		t.Errorf("AGENTS.md local edit was overwritten")
	}
}

func TestUpdateWithoutRecordsFails(t *testing.T) {
	t.Setenv("SDD_CONFIG_DIR", t.TempDir())
	t.Setenv("CI", "true")
	if err := Update(Options{Yes: true, Stdout: &bytes.Buffer{}}); err == nil {
		t.Errorf("expected error when nothing is installed")
	}
}

func TestDryRunAppliesNothing(t *testing.T) {
	target := t.TempDir()
	o := testOptions(t, target)
	o.DryRun = true
	if err := Run(o); err != nil {
		t.Fatal(err)
	}
	entries, err := os.ReadDir(target)
	if err != nil {
		t.Fatal(err)
	}
	if len(entries) != 0 {
		t.Errorf("dry-run created files: %v", entries)
	}
}
