package install

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// prompter reads interactive answers. In tests Stdin can be any reader;
// interactivity additionally requires a real terminal on os.Stdin.
type prompter struct {
	stdin  io.Reader
	stdout io.Writer
	reader *bufio.Reader
}

func newPrompter(stdin io.Reader, stdout io.Writer) *prompter {
	if stdin == nil {
		stdin = os.Stdin
	}
	if stdout == nil {
		stdout = os.Stdout
	}
	return &prompter{stdin: stdin, stdout: stdout, reader: bufio.NewReader(stdin)}
}

// interactive reports whether we may ask the user questions.
func (p *prompter) interactive() bool {
	if os.Getenv("CI") != "" {
		return false
	}
	fi, err := os.Stdin.Stat()
	if err != nil {
		return false
	}
	return fi.Mode()&os.ModeCharDevice != 0
}

func (p *prompter) readLine() string {
	// One shared buffered reader: recreating it per call discards input
	// buffered past the first newline (e.g. pasted answers).
	line, _ := p.reader.ReadString('\n')
	return strings.TrimSpace(line)
}

func (p *prompter) askScope(repoDir, homeDir string) Scope {
	fmt.Fprintf(p.stdout, "Install scope:\n  1) This repository (%s)\n  2) Home directory (%s)\nChoose [1]: ", repoDir, homeDir)
	switch p.readLine() {
	case "", "1":
		return ScopeRepo
	case "2":
		return ScopeHome
	default:
		fmt.Fprintln(p.stdout, "Unknown choice, using this repository.")
		return ScopeRepo
	}
}

func (p *prompter) askHosts(tools []Tool) []string {
	fmt.Fprintln(p.stdout, "Install for hosts (comma-separated numbers, or 'all'):")
	for i, t := range tools {
		mark := ""
		if t.Detected {
			mark = " [detected]"
		}
		fmt.Fprintf(p.stdout, "  %d) %s — %s%s\n", i+1, t.Label, t.Hint, mark)
	}
	fmt.Fprint(p.stdout, "Choose [all]: ")
	line := strings.ToLower(strings.TrimSpace(p.readLine()))
	if line == "" || line == "all" {
		return allHostIDs(tools)
	}
	var out []string
	for _, part := range strings.Split(line, ",") {
		part = strings.TrimSpace(part)
		for i, t := range tools {
			if part == fmt.Sprint(i+1) || part == t.ID {
				out = append(out, t.ID)
			}
		}
	}
	if len(out) == 0 {
		fmt.Fprintln(p.stdout, "No valid choice, using all hosts.")
		return allHostIDs(tools)
	}
	return dedupe(out)
}

func (p *prompter) askConfirm(summary string) bool {
	fmt.Fprintf(p.stdout, "%s\nApply? [y/N]: ", summary)
	line := strings.ToLower(strings.TrimSpace(p.readLine()))
	return line == "y" || line == "yes"
}

func allHostIDs(tools []Tool) []string {
	var out []string
	for _, t := range tools {
		out = append(out, t.ID)
	}
	return out
}

func detectedHostIDs(tools []Tool) []string {
	var out []string
	for _, t := range tools {
		if t.Detected {
			out = append(out, t.ID)
		}
	}
	return out
}

func dedupe(in []string) []string {
	seen := map[string]bool{}
	var out []string
	for _, s := range in {
		if !seen[s] {
			seen[s] = true
			out = append(out, s)
		}
	}
	return out
}
