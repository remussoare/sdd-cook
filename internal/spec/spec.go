// Package spec parses and updates the .spec/ SDD state directory.
package spec

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

// Status is the lifecycle status recorded in an artifact frontmatter.
type Status string

const (
	Draft    Status = "DRAFT"
	Accepted Status = "ACCEPTED"
	Locked   Status = "LOCKED"
	Invalid  Status = "INVALID"
)

// Chain is the ordered artifact chain of the core workflow.
var Chain = []string{"objective", "domain", "specification", "plan", "tests", "tasks", "validation"}

var phaseDirs = map[string]string{
	"objective":     "objective",
	"domain":        "domain",
	"specification": "specifications",
	"plan":          "plan",
	"tests":         "tests",
	"tasks":         "tasks",
	"validation":    "validation",
}

var fileNames = map[string]string{
	"objective":     "objective.md",
	"domain":        "domain.md",
	"specification": "specification.md",
	"plan":          "plan.md",
	"tests":         "tests.md",
	"tasks":         "tasks.md",
	"validation":    "validation.md",
}

// AuxDirs are supporting directories of .spec/.
var AuxDirs = []string{"changes", "traceability"}

// AllDirs returns the canonical .spec/ subdirectory order.
func AllDirs() []string {
	dirs := make([]string, 0, len(Chain)+len(AuxDirs))
	for _, name := range Chain {
		dirs = append(dirs, phaseDirs[name])
	}
	return append(dirs, AuxDirs...)
}

// Dir returns the directory name of a chain artifact.
func Dir(name string) string { return phaseDirs[name] }

// File returns the canonical file name of a chain artifact.
func File(name string) string { return fileNames[name] }

// Artifact is one SDD artifact of the chain.
type Artifact struct {
	Name      string
	Path      string // empty when the artifact file does not exist
	Component string
	Evolution string
	Status    Status
	Body      string
}

// Missing reports whether the artifact file does not exist.
func (a *Artifact) Missing() bool { return a.Path == "" }

// Accepted reports whether the artifact is accepted or locked.
func (a *Artifact) Accepted() bool { return a.Status == Accepted || a.Status == Locked }

// Spec is the parsed .spec/ state of a project.
type Spec struct {
	Root      string
	Artifacts map[string]*Artifact
}

// Exists reports whether .spec/ is present in root.
func Exists(root string) bool {
	st, err := os.Stat(filepath.Join(root, ".spec"))
	return err == nil && st.IsDir()
}

// Load parses the chain artifacts of the .spec/ directory in root.
func Load(root string) (*Spec, error) {
	s := &Spec{Root: root, Artifacts: map[string]*Artifact{}}
	for _, name := range Chain {
		a := &Artifact{Name: name}
		p := filepath.Join(root, ".spec", phaseDirs[name], fileNames[name])
		if st, err := os.Stat(p); err == nil && !st.IsDir() {
			a.Path = p
			parseArtifact(a)
		}
		s.Artifacts[name] = a
	}
	return s, nil
}

func parseArtifact(a *Artifact) {
	data, err := os.ReadFile(a.Path)
	if err != nil {
		a.Status = Invalid
		return
	}
	text := string(data)
	fm, body, ok := splitFrontmatter(text)
	if !ok {
		a.Status = Invalid
		a.Body = text
		return
	}
	a.Body = body
	for _, line := range fm {
		k, v, ok := strings.Cut(line, ":")
		if !ok {
			continue
		}
		switch strings.ToLower(strings.TrimSpace(k)) {
		case "component":
			a.Component = strings.TrimSpace(v)
		case "evolution":
			a.Evolution = strings.TrimSpace(v)
		case "status":
			switch strings.ToUpper(strings.TrimSpace(v)) {
			case "DRAFT":
				a.Status = Draft
			case "ACCEPTED":
				a.Status = Accepted
			case "LOCKED":
				a.Status = Locked
			default:
				a.Status = Invalid
			}
		}
	}
	if a.Status == "" {
		a.Status = Invalid
	}
}

func splitFrontmatter(text string) (fm []string, body string, ok bool) {
	lines := strings.Split(text, "\n")
	i := 0
	for i < len(lines) && strings.TrimSpace(lines[i]) == "" {
		i++
	}
	if i >= len(lines) || strings.TrimSpace(lines[i]) != "---" {
		return nil, text, false
	}
	i++
	start := i
	for i < len(lines) {
		if strings.TrimSpace(lines[i]) == "---" {
			return lines[start:i], strings.Join(lines[i+1:], "\n"), true
		}
		i++
	}
	return nil, text, false
}

// SetStatus rewrites the status field in the artifact frontmatter.
func (a *Artifact) SetStatus(s Status) error {
	if a.Missing() {
		return fmt.Errorf("artifact %s does not exist", a.Name)
	}
	data, err := os.ReadFile(a.Path)
	if err != nil {
		return err
	}
	lines := strings.Split(string(data), "\n")
	inFM := false
	for i, l := range lines {
		t := strings.TrimSpace(l)
		if t == "---" {
			if !inFM {
				inFM = true
				continue
			}
			break
		}
		if inFM {
			if k, _, ok := strings.Cut(t, ":"); ok && strings.ToLower(strings.TrimSpace(k)) == "status" {
				lines[i] = "status: " + string(s)
				return os.WriteFile(a.Path, []byte(strings.Join(lines, "\n")), 0o644)
			}
		}
	}
	return fmt.Errorf("artifact %s has no status field in frontmatter", a.Name)
}

// ChainAllLocked reports whether every chain artifact is LOCKED.
func ChainAllLocked(s *Spec) bool {
	for _, name := range Chain {
		if s.Artifacts[name].Status != Locked {
			return false
		}
	}
	return true
}

// NextEvolution returns the next evolution label (v1 -> v2).
func NextEvolution(cur string) string {
	if len(cur) > 1 && (cur[0] == 'v' || cur[0] == 'V') {
		if n, err := strconv.Atoi(cur[1:]); err == nil {
			return fmt.Sprintf("v%d", n+1)
		}
	}
	return "v1"
}

// CreateObjective writes a new objective draft. It never overwrites.
func CreateObjective(root, component, evolution string) error {
	p := filepath.Join(root, ".spec", "objective", "objective.md")
	body := fmt.Sprintf(`---
artifact: objective
component: %s
evolution: %s
status: DRAFT
---

# Objective

PURPOSE:
SCOPE:
CONSTRAINTS:
ACCEPTANCE CRITERIA:
`, component, evolution)
	return writeFileNew(p, body)
}

// CreateChangeRequest writes a new change request draft in .spec/changes/.
// It never overwrites and preserves the previous accepted evolution.
func CreateChangeRequest(root, component, evolution, title string) error {
	p := filepath.Join(root, ".spec", "changes", fmt.Sprintf("change-%s-%s.md", time.Now().Format("20060102-150405"), slug(title)))
	body := fmt.Sprintf(`---
artifact: change-request
component: %s
evolution: %s
status: PROPOSED
---

# Change Request: %s

REQUESTED CHANGE:
IMPACT ANALYSIS:
DECISION:
AFFECTED ARTIFACTS:
NEXT EVOLUTION:
`, component, evolution, title)
	return writeFileNew(p, body)
}

// ActiveChangeRequest returns the path of a change request in a
// non-terminal state, if any. Terminal states are ACCEPTED, REJECTED,
// CANCELLED and BLOCKED.
func ActiveChangeRequest(root string) (string, bool) {
	dir := filepath.Join(root, ".spec", "changes")
	entries, err := os.ReadDir(dir)
	if err != nil {
		return "", false
	}
	for _, e := range entries {
		if e.IsDir() || !strings.HasSuffix(e.Name(), ".md") {
			continue
		}
		p := filepath.Join(dir, e.Name())
		data, err := os.ReadFile(p)
		if err != nil {
			return p, true
		}
		fm, _, ok := splitFrontmatter(string(data))
		if !ok {
			return p, true
		}
		status := ""
		for _, line := range fm {
			if k, v, ok := strings.Cut(line, ":"); ok && strings.ToLower(strings.TrimSpace(k)) == "status" {
				status = strings.ToUpper(strings.TrimSpace(v))
			}
		}
		switch status {
		case "ACCEPTED", "REJECTED", "CANCELLED", "BLOCKED":
		default:
			return p, true
		}
	}
	return "", false
}

func writeFileNew(p, content string) error {
	if _, err := os.Stat(p); err == nil {
		return fmt.Errorf("refusing to overwrite %s", p)
	}
	if err := os.MkdirAll(filepath.Dir(p), 0o755); err != nil {
		return err
	}
	return os.WriteFile(p, []byte(content), 0o644)
}

func slug(s string) string {
	var b strings.Builder
	for _, r := range strings.ToLower(strings.TrimSpace(s)) {
		if (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') {
			b.WriteRune(r)
		} else if b.Len() > 0 && !strings.HasSuffix(b.String(), "-") {
			b.WriteByte('-')
		}
	}
	return strings.TrimSuffix(b.String(), "-")
}