package install

import (
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

// Tool is a detectable AI coding tool (install target).
type Tool struct {
	ID       string // opencode | cursor | vscode
	Label    string
	Detected bool
	Hint     string
}

// DetectTools reports which tools are present, given a project directory.
// Detection combines binaries on PATH, well-known config locations and
// project-local markers.
func DetectTools(projectDir string) []Tool {
	home, _ := os.UserHomeDir()
	has := func(p string) bool {
		if p == "" {
			return false
		}
		_, err := os.Stat(p)
		return err == nil
	}
	onPath := func(names ...string) bool {
		for _, n := range names {
			if _, err := exec.LookPath(n); err == nil {
				return true
			}
		}
		return false
	}

	opencode := onPath("opencode") ||
		has(filepath.Join(home, ".config", "opencode")) ||
		has(filepath.Join(projectDir, ".opencode")) ||
		has(filepath.Join(projectDir, ".agents", "skills")) ||
		has(filepath.Join(home, ".agents", "skills"))

	cursorCfg := ""
	switch runtime.GOOS {
	case "darwin":
		cursorCfg = filepath.Join(home, "Library", "Application Support", "Cursor")
	case "windows":
		cursorCfg = filepath.Join(home, "AppData", "Roaming", "Cursor")
	default:
		if xdg := os.Getenv("XDG_CONFIG_HOME"); xdg != "" {
			cursorCfg = filepath.Join(xdg, "Cursor")
		} else {
			cursorCfg = filepath.Join(home, ".config", "Cursor")
		}
	}
	cursor := onPath("cursor", "cursor-cli") ||
		has(filepath.Join(projectDir, ".cursor")) ||
		has(cursorCfg)

	vscode := onPath("code", "code-insiders", "copilot") ||
		has(filepath.Join(projectDir, ".vscode")) ||
		has(filepath.Join(projectDir, ".github"))

	return []Tool{
		{ID: "opencode", Label: "OpenCode", Detected: opencode, Hint: "skill in .agents/skills/sdd-cook/"},
		{ID: "cursor", Label: "Cursor", Detected: cursor, Hint: "agent in .cursor/agents/ + rules in .cursor/rules/"},
		{ID: "vscode", Label: "VS Code / Copilot", Detected: vscode, Hint: "skill in .github/skills/ + AGENTS.md"},
	}
}
