// Package transpile generates per-host install payloads from the canonical
// skill sources. One origin (skills/sdd-cook), N targets (opencode, cursor,
// vscode). Adding a host means adding one generator here.
package transpile

import (
	"fmt"
	"io/fs"
	"path"
	"sort"
	"strings"
)

// Host describes an installation target.
type Host struct {
	ID          string
	Label       string
	Description string
}

// Hosts lists the supported installation targets.
var Hosts = []Host{
	{ID: "opencode", Label: "OpenCode", Description: "skill in .agents/skills/sdd-cook/ (+ .opencode/) — primary host"},
	{ID: "cursor", Label: "Cursor", Description: "agent in .cursor/agents/ + rules in .cursor/rules/"},
	{ID: "vscode", Label: "VS Code / Copilot", Description: "skill in .github/skills/ + AGENTS.md"},
}

// HostIDs returns the IDs of all supported hosts.
func HostIDs() []string {
	out := make([]string, 0, len(Hosts))
	for _, h := range Hosts {
		out = append(out, h.ID)
	}
	return out
}

// ValidHost reports whether id is a supported host.
func ValidHost(id string) bool {
	for _, h := range Hosts {
		if h.ID == id {
			return true
		}
	}
	return false
}

// File is a generated file. Path uses slashes and is relative to the scope
// root (project dir for scope=repo, home dir for scope=home).
type File struct {
	Path         string
	Data         []byte
	Mode         uint32
	KeepExisting bool // if true, never overwrite an existing file
}

func (f File) mode() fs.FileMode {
	if f.Mode == 0 {
		return 0o644
	}
	return fs.FileMode(f.Mode)
}

// Command IDs in canonical order.
var commands = []string{"init", "new", "plan", "run", "validate", "change", "status"}

// Payload generates the install files for hostID from skill, an fs.FS rooted
// at the canonical skill directory (SKILL.md at its top).
func Payload(hostID string, skill fs.FS) ([]File, error) {
	switch hostID {
	case "opencode":
		return payloadOpenCode(skill)
	case "cursor":
		return payloadCursor(skill)
	case "vscode":
		return payloadVSCode(skill)
	default:
		return nil, fmt.Errorf("unknown host %q", hostID)
	}
}

// payloadOpenCode copies the canonical skill tree verbatim.
func payloadOpenCode(skill fs.FS) ([]File, error) {
	var out []File
	err := fs.WalkDir(skill, ".", func(p string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		data, err := fs.ReadFile(skill, p)
		if err != nil {
			return err
		}
		out = append(out, File{
			Path: path.Join(".agents/skills/sdd-cook", p),
			Data: data,
		})
		return nil
	})
	if err != nil {
		return nil, err
	}
	sort.Slice(out, func(i, j int) bool { return out[i].Path < out[j].Path })
	return out, nil
}

// payloadCursor builds a single agent file plus a rules file.
func payloadCursor(skill fs.FS) ([]File, error) {
	sk, err := fs.ReadFile(skill, "SKILL.md")
	if err != nil {
		return nil, err
	}
	index, err := commandIndex(skill)
	if err != nil {
		return nil, err
	}
	agent := "---\n" +
		"name: sdd-cook\n" +
		"description: Specification-Driven AI Engineering — SDD commands (init, new, plan, run, validate, change, status) with gates and change control\n" +
		"---\n\n" +
		"# SDD Cook (sdd-cook)\n\n" +
		"Follow this workflow for all SDD work. Canonical reference:\n" +
		"https://github.com/remussoare/sdd-cook\n\n" +
		string(sk) + "\n\n" +
		"## Commands\n\n" + index

	rules := "---\n" +
		"description: SDD Cook decision rules and gates — apply when doing SDD work\n" +
		"globs: [\"**/*\"]\n" +
		"alwaysApply: false\n" +
		"---\n\n" +
		"# SDD Cook rules\n\n" +
		"## Master decision rule\n\n" +
		"If a decision is derivable from accepted artifacts, decide and continue. " +
		"If it is not derivable, STOP and request a decision. " +
		"Never convert uncertainty into behavior.\n\n" +
		"## Gates\n\n" +
		"Possible outcomes: PASS, FAIL, BLOCKED, INCONCLUSIVE. " +
		"PASS → continue. FAIL → diagnose and route to earliest affected phase. " +
		"BLOCKED → stop and request the missing decision or artifact. " +
		"INCONCLUSIVE → investigate only with a new verifiable hypothesis.\n\n" +
		"## Anti-loop\n\n" +
		"Retry a correction only when a new verifiable hypothesis justifies it. " +
		"Repeating the same failed action without new evidence is forbidden.\n\n" +
		"## Change control\n\n" +
		"Any change affecting behavior, scope, architecture or an accepted decision " +
		"requires Change Management. Never silently overwrite accepted state.\n"

	return []File{
		{Path: ".cursor/agents/sdd-cook.md", Data: []byte(agent)},
		{Path: ".cursor/rules/sdd-cook.mdc", Data: []byte(rules)},
	}, nil
}

// payloadVSCode builds a Copilot agent skill plus AGENTS.md (created only).
func payloadVSCode(skill fs.FS) ([]File, error) {
	sk, err := fs.ReadFile(skill, "SKILL.md")
	if err != nil {
		return nil, err
	}
	index, err := commandIndex(skill)
	if err != nil {
		return nil, err
	}
	skmd := "---\n" +
		"name: sdd-cook\n" +
		"description: Specification-Driven AI Engineering — SDD commands with gates, traceability and change control\n" +
		"---\n\n" +
		"# SDD Cook\n\n" +
		string(sk) + "\n\n" +
		"## Commands\n\n" + index

	agents := "# AGENTS.md\n\n" +
		"This repository uses SDD Cook (Specification-Driven AI Engineering).\n\n" +
		"- Skill: `.github/skills/sdd-cook/SKILL.md`\n" +
		"- CLI (optional, deterministic `.spec/` state): `sdd` — " +
		"https://github.com/remussoare/sdd-cook\n" +
		"- Workflow: OBJECTIVE → DOMAIN → SPECIFICATION → PLAN → TESTS → TASKS → " +
		"IMPLEMENTATION → VALIDATION → LOCKED\n" +
		"- Master decision rule: if a decision is derivable from accepted artifacts, " +
		"decide and continue; otherwise STOP and request a decision. " +
		"Never convert uncertainty into behavior.\n"

	return []File{
		{Path: ".github/skills/sdd-cook/SKILL.md", Data: []byte(skmd)},
		{Path: "AGENTS.md", Data: []byte(agents), KeepExisting: true},
	}, nil
}

// commandIndex builds a "- `/sdd:x` — purpose" list from commands/*.md.
func commandIndex(skill fs.FS) (string, error) {
	var b strings.Builder
	for _, c := range commands {
		p := purposeOf(skill, c)
		fmt.Fprintf(&b, "- `/sdd:%s` — %s\n", c, p)
	}
	return b.String(), nil
}

// purposeOf extracts the "## Purpose" paragraph of commands/<cmd>.md.
func purposeOf(skill fs.FS, cmd string) string {
	data, err := fs.ReadFile(skill, "commands/"+cmd+".md")
	if err != nil {
		return "see commands/" + cmd + ".md"
	}
	lines := strings.Split(string(data), "\n")
	inPurpose := false
	var b strings.Builder
	for _, l := range lines {
		t := strings.TrimSpace(l)
		if strings.HasPrefix(t, "## ") {
			if inPurpose {
				break
			}
			if t == "## Purpose" {
				inPurpose = true
			}
			continue
		}
		if inPurpose && t != "" {
			if b.Len() > 0 {
				b.WriteString(" ")
			}
			b.WriteString(t)
		}
	}
	if b.Len() == 0 {
		return "see commands/" + cmd + ".md"
	}
	return b.String()
}
