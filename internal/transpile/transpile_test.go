package transpile

import (
	"io/fs"
	"strings"
	"testing"
	"testing/fstest"

	sddcook "github.com/remussoare/sdd-cook"
)

var testSkill = fstest.MapFS{
	"SKILL.md":                 {Data: []byte("# Skill: test\n\nBody.\n")},
	"commands/init.md":         {Data: []byte("# /sdd:init\n\n## Purpose\n\nStart things.\n")},
	"commands/new.md":          {Data: []byte("# /sdd:new\n\n## Purpose\n\nBegin evolution.\n")},
	"commands/plan.md":         {Data: []byte("# /sdd:plan\n\nNo purpose section.\n")},
	"commands/run.md":          {Data: []byte("# /sdd:run\n\n## Purpose\n\nDo work.\n")},
	"commands/validate.md":     {Data: []byte("# /sdd:validate\n\n## Purpose\n\nCheck work.\n")},
	"commands/change.md":       {Data: []byte("# /sdd:change\n\n## Purpose\n\nEvolve.\n")},
	"commands/status.md":       {Data: []byte("# /sdd:status\n\n## Purpose\n\nShow state.\n")},
	"principles/a.md":          {Data: []byte("principle\n")},
	"nested/dir/file.md":       {Data: []byte("nested\n")},
}

func paths(files []File) []string {
	var out []string
	for _, f := range files {
		out = append(out, f.Path)
	}
	return out
}

func hasPath(files []File, p string) bool {
	for _, f := range files {
		if f.Path == p {
			return true
		}
	}
	return false
}

func contentOf(files []File, p string) string {
	for _, f := range files {
		if f.Path == p {
			return string(f.Data)
		}
	}
	return ""
}

func TestPayloadOpenCodeCopiesTree(t *testing.T) {
	files, err := Payload("opencode", testSkill)
	if err != nil {
		t.Fatal(err)
	}
	for _, want := range []string{
		".agents/skills/sdd-cook/SKILL.md",
		".agents/skills/sdd-cook/commands/init.md",
		".agents/skills/sdd-cook/principles/a.md",
		".agents/skills/sdd-cook/nested/dir/file.md",
	} {
		if !hasPath(files, want) {
			t.Errorf("missing %s (got %v)", want, paths(files))
		}
	}
	if got := contentOf(files, ".agents/skills/sdd-cook/SKILL.md"); !strings.Contains(got, "# Skill: test") {
		t.Errorf("SKILL.md content not copied verbatim")
	}
}

func TestPayloadCursor(t *testing.T) {
	files, err := Payload("cursor", testSkill)
	if err != nil {
		t.Fatal(err)
	}
	agent := contentOf(files, ".cursor/agents/sdd-cook.md")
	if !strings.HasPrefix(agent, "---\nname: sdd-cook") {
		t.Errorf("agent missing frontmatter")
	}
	for _, want := range []string{"`/sdd:init` — Start things.", "`/sdd:status` — Show state.", "`/sdd:plan` — see commands/plan.md"} {
		if !strings.Contains(agent, want) {
			t.Errorf("agent missing %q", want)
		}
	}
	rules := contentOf(files, ".cursor/rules/sdd-cook.mdc")
	if !strings.Contains(rules, "Master decision rule") {
		t.Errorf("rules missing decision rule")
	}
}

func TestPayloadVSCode(t *testing.T) {
	files, err := Payload("vscode", testSkill)
	if err != nil {
		t.Fatal(err)
	}
	sk := contentOf(files, ".github/skills/sdd-cook/SKILL.md")
	if !strings.HasPrefix(sk, "---\nname: sdd-cook") {
		t.Errorf("skill missing frontmatter")
	}
	var agents *File
	for i, f := range files {
		if f.Path == "AGENTS.md" {
			agents = &files[i]
		}
	}
	if agents == nil || !agents.KeepExisting {
		t.Errorf("AGENTS.md must exist with KeepExisting=true")
	}
}

func TestPayloadUnknownHost(t *testing.T) {
	if _, err := Payload("nope", testSkill); err == nil {
		t.Errorf("expected error for unknown host")
	}
}

// TestPurposeOfEmbeddedCommands guards against heading rot in the real
// command docs: every /sdd command must yield an extracted purpose from the
// embedded skill, never the fallback string.
func TestPurposeOfEmbeddedCommands(t *testing.T) {
	sub, err := fs.Sub(sddcook.SkillsFS, "skills/sdd-cook")
	if err != nil {
		t.Fatal(err)
	}
	for _, c := range []string{"init", "new", "plan", "run", "validate", "change", "status"} {
		if p := purposeOf(sub, c); strings.HasPrefix(p, "see commands/") {
			t.Errorf("purposeOf(%s) fell back to %q — check the '## PURPOSE' heading in embedded commands/%s.md", c, p, c)
		}
	}
}
