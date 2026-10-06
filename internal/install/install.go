// Package install implements `sdd install` / `sdd update`: it writes the
// embedded skill payloads into a scope root (project dir or home dir) for
// the selected hosts. Re-running is idempotent: unchanged files skip, local
// edits are overwritten (this repository is the version store), and files
// removed upstream are deleted.
package install

import (
	"encoding/json"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"time"

	sddcook "github.com/remussoare/sdd-cook"
	"github.com/remussoare/sdd-cook/internal/transpile"
	"github.com/remussoare/sdd-cook/internal/version"
)

// Scope is the installation scope.
type Scope string

const (
	// ScopeRepo installs into the current project directory.
	ScopeRepo Scope = "repo"
	// ScopeHome installs into the user home directory (global).
	ScopeHome Scope = "home"
)

// DefaultSource is the canonical source recorded on install.
const DefaultSource = version.SourceURL

// Options controls an install run.
type Options struct {
	Scope  Scope    // "" = decide (interactive, else default)
	Hosts  []string // nil = decide (interactive, else detected-or-all)
	Target string   // override scope root (scripts, tests)
	DryRun bool     // print plan, apply nothing
	Yes    bool     // skip confirmation
	Source string   // recorded source URL (default DefaultSource)
	Stdin  io.Reader
	Stdout io.Writer
	Stderr io.Writer
}

// Op is a single planned file operation.
type Op struct {
	Action string // create | update | skip | keep | delete
	Path   string // relative to scope root, slashes
}

// Plan is the computed install plan for one scope root.
type Plan struct {
	Root  string
	Scope Scope
	Hosts []string
	Ops   []Op
}

// Record tracks one applied installation for future updates.
type Record struct {
	Root        string   `json:"root"`
	Scope       string   `json:"scope"`
	Hosts       []string `json:"hosts"`
	Source      string   `json:"source"`
	Version     string   `json:"version"`
	Files       []string `json:"files"`
	InstalledAt string   `json:"installed_at"`
}

type store struct {
	Installs []Record `json:"installs"`
}

func (o *Options) out() io.Writer {
	if o.Stdout != nil {
		return o.Stdout
	}
	return os.Stdout
}

func (o *Options) errOut() io.Writer {
	if o.Stderr != nil {
		return o.Stderr
	}
	return os.Stderr
}

// skillFS returns the embedded canonical skill tree.
func skillFS() (fs.FS, error) {
	return fs.Sub(sddcook.SkillsFS, "skills/"+sddcook.SkillName)
}

// Run executes the full install flow: resolve, plan, confirm, apply, record.
func Run(o Options) error {
	if o.Source == "" {
		o.Source = DefaultSource
	}
	p := newPrompter(o.Stdin, o.out())
	repoDir, _ := os.Getwd()
	homeDir, _ := os.UserHomeDir()

	scope := o.Scope
	if scope == "" {
		if v := strings.ToLower(os.Getenv("INSTALL_SCOPE")); v == "global" || v == "home" {
			scope = ScopeHome
		} else if v == "project" || v == "repo" {
			scope = ScopeRepo
		}
	}
	hosts := o.Hosts
	if hosts == nil {
		if v := os.Getenv("INSTALL_TARGET"); v != "" {
			if strings.ToLower(v) == "all" {
				hosts = transpile.HostIDs()
			} else {
				parts := strings.Split(v, ",")
				for i, h := range parts {
					parts[i] = strings.TrimSpace(h)
				}
				hosts = parts
			}
		}
	}

	tools := DetectTools(repoDir)
	interactive := p.interactive() && !o.Yes

	if scope == "" {
		if interactive {
			scope = p.askScope(repoDir, homeDir)
		} else {
			scope = ScopeRepo
		}
	}
	if hosts == nil {
		if interactive {
			hosts = p.askHosts(tools)
		} else if d := detectedHostIDs(tools); len(d) > 0 {
			hosts = d
		} else {
			hosts = transpile.HostIDs()
		}
	}
	for _, h := range hosts {
		if !transpile.ValidHost(strings.TrimSpace(h)) {
			return fmt.Errorf("unknown host %q (valid: opencode, cursor, vscode)", h)
		}
	}

	root := o.Target
	if root == "" {
		if scope == ScopeHome {
			root = homeDir
		} else {
			root = repoDir
		}
	}
	if root == "" {
		return fmt.Errorf("cannot resolve install root")
	}

	plan, err := buildPlan(root, scope, hosts)
	if err != nil {
		return err
	}
	printPlan(o.out(), plan)

	if o.DryRun {
		fmt.Fprintln(o.out(), "dry-run: nothing applied.")
		return nil
	}
	confirmed := o.Yes
	if !confirmed {
		if !p.interactive() {
			return fmt.Errorf("refusing to apply in non-interactive mode without --yes — pass --yes to confirm, or run interactively")
		}
		confirmed = p.askConfirm(summarize(plan))
	}
	if !confirmed {
		fmt.Fprintln(o.out(), "aborted.")
		return nil
	}
	if err := apply(plan); err != nil {
		return err
	}
	rec := Record{
		Root:        root,
		Scope:       string(scope),
		Hosts:       hosts,
		Source:      o.Source,
		Version:     version.Version,
		Files:       manifest(plan),
		InstalledAt: time.Now().UTC().Format(time.RFC3339),
	}
	if err := saveRecord(rec); err != nil {
		fmt.Fprintln(o.errOut(), "warning: could not record installation:", err)
	}
	fmt.Fprintf(o.out(), "installed sdd-cook %s for %s (scope %s) in %s\n",
		version.Version, strings.Join(hosts, ", "), scope, root)
	fmt.Fprintln(o.out(), "next: /sdd:init in your project, or run: sdd status")
	return nil
}

// buildPlan computes file operations for root+hosts against previous records.
func buildPlan(root string, scope Scope, hosts []string) (*Plan, error) {
	skill, err := skillFS()
	if err != nil {
		return nil, err
	}
	plan := &Plan{Root: root, Scope: scope, Hosts: hosts}
	wanted := map[string]transpile.File{}
	for _, h := range hosts {
		files, err := transpile.Payload(h, skill)
		if err != nil {
			return nil, err
		}
		for _, f := range files {
			wanted[f.Path] = f
		}
	}
	for _, p := range sortedKeys(wanted) {
		f := wanted[p]
		dest := filepath.Join(root, filepath.FromSlash(p))
		data, err := os.ReadFile(dest)
		switch {
		case os.IsNotExist(err):
			plan.Ops = append(plan.Ops, Op{Action: "create", Path: p})
		case err != nil:
			return nil, err
		case string(data) == string(f.Data):
			plan.Ops = append(plan.Ops, Op{Action: "skip", Path: p})
		case f.KeepExisting:
			plan.Ops = append(plan.Ops, Op{Action: "keep", Path: p})
		default:
			plan.Ops = append(plan.Ops, Op{Action: "update", Path: p})
		}
	}
	// Deletions: files installed before (same root) that upstream no longer ships.
	if prev := findRecord(root); prev != nil {
		for _, p := range prev.Files {
			if _, ok := wanted[p]; !ok {
				if _, err := os.Stat(filepath.Join(root, filepath.FromSlash(p))); err == nil {
					plan.Ops = append(plan.Ops, Op{Action: "delete", Path: p})
				}
			}
		}
	}
	return plan, nil
}

// apply executes a plan.
func apply(plan *Plan) error {
	skill, err := skillFS()
	if err != nil {
		return err
	}
	wanted := map[string][]byte{}
	for _, h := range plan.Hosts {
		files, err := transpile.Payload(h, skill)
		if err != nil {
			return err
		}
		for _, f := range files {
			wanted[f.Path] = f.Data
		}
	}
	for _, op := range plan.Ops {
		dest := filepath.Join(plan.Root, filepath.FromSlash(op.Path))
		switch op.Action {
		case "create", "update":
			data, ok := wanted[op.Path]
			if !ok {
				return fmt.Errorf("payload missing for %s", op.Path)
			}
			if err := os.MkdirAll(filepath.Dir(dest), 0o755); err != nil {
				return err
			}
			if err := os.WriteFile(dest, data, 0o644); err != nil {
				return err
			}
		case "delete":
			if err := os.Remove(dest); err != nil && !os.IsNotExist(err) {
				return err
			}
		case "skip", "keep":
			// nothing to do
		}
	}
	return nil
}

func manifest(plan *Plan) []string {
	var out []string
	for _, op := range plan.Ops {
		if op.Action != "delete" {
			out = append(out, op.Path)
		}
	}
	return out
}

func sortedKeys(m map[string]transpile.File) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	for i := 1; i < len(keys); i++ {
		for j := i; j > 0 && keys[j] < keys[j-1]; j-- {
			keys[j], keys[j-1] = keys[j-1], keys[j]
		}
	}
	return keys
}

func summarize(plan *Plan) string {
	counts := map[string]int{}
	for _, op := range plan.Ops {
		counts[op.Action]++
	}
	return fmt.Sprintf("plan for %s: %d create, %d update, %d skip, %d keep, %d delete",
		plan.Root, counts["create"], counts["update"], counts["skip"], counts["keep"], counts["delete"])
}

func printPlan(w io.Writer, plan *Plan) {
	fmt.Fprintf(w, "sdd install plan (scope %s, hosts %s, root %s):\n",
		plan.Scope, strings.Join(plan.Hosts, ","), plan.Root)
	for _, op := range plan.Ops {
		fmt.Fprintf(w, "  %-6s %s\n", op.Action, op.Path)
	}
	fmt.Fprintln(w, summarize(plan))
}

// recordsPath is the user-level install registry.
// SDD_CONFIG_DIR overrides the base config dir (tests, CI, scripts).
func recordsPath() (string, error) {
	if dir := os.Getenv("SDD_CONFIG_DIR"); dir != "" {
		return filepath.Join(dir, "sdd-cook", "installs.json"), nil
	}
	dir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(dir, "sdd-cook", "installs.json"), nil
}

func loadStore() store {
	var st store
	p, err := recordsPath()
	if err != nil {
		return st
	}
	data, err := os.ReadFile(p)
	if err != nil {
		return st
	}
	_ = json.Unmarshal(data, &st)
	return st
}

func findRecord(root string) *Record {
	st := loadStore()
	for i, r := range st.Installs {
		if samePath(r.Root, root) {
			rec := st.Installs[i]
			return &rec
		}
	}
	return nil
}

func samePath(a, b string) bool {
	abs := func(p string) string {
		if ap, err := filepath.Abs(p); err == nil {
			return ap
		}
		return p
	}
	return abs(a) == abs(b)
}

func saveRecord(rec Record) error {
	p, err := recordsPath()
	if err != nil {
		return err
	}
	st := loadStore()
	kept := st.Installs[:0]
	for _, r := range st.Installs {
		if !samePath(r.Root, rec.Root) {
			kept = append(kept, r)
		}
	}
	st.Installs = append(kept, rec)
	data, err := json.MarshalIndent(st, "", "  ")
	if err != nil {
		return err
	}
	if err := os.MkdirAll(filepath.Dir(p), 0o755); err != nil {
		return err
	}
	return os.WriteFile(p, data, 0o644)
}

// Records returns all tracked installations (for `sdd update`).
func Records() []Record {
	return loadStore().Installs
}
