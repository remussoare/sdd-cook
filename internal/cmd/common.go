// Package cmd implements the sdd CLI commands.
package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/remussoare/sdd-cook/internal/spec"
)

// load parses .spec/ in the current directory or fails with guidance.
func load() (*spec.Spec, error) {
	if !spec.Exists(".") {
		return nil, fmt.Errorf(".spec/ missing — run: sdd init")
	}
	return spec.Load(".")
}

func projectType(root string) string {
	for _, c := range []struct{ file, kind string }{
		{"go.mod", "Go"},
		{"package.json", "Node.js"},
		{"pyproject.toml", "Python"},
		{"requirements.txt", "Python"},
		{"Cargo.toml", "Rust"},
		{"pom.xml", "Java (Maven)"},
		{"build.gradle", "Java (Gradle)"},
		{"composer.json", "PHP"},
		{"Gemfile", "Ruby"},
	} {
		if _, err := os.Stat(filepath.Join(root, c.file)); err == nil {
			return c.kind
		}
	}
	return "unknown"
}

func existingArtifacts(root string) []string {
	var found []string
	for _, name := range spec.Chain {
		p := filepath.Join(root, ".spec", spec.Dir(name), spec.File(name))
		if _, err := os.Stat(p); err == nil {
			found = append(found, name)
		}
	}
	return found
}

func joinOr(items []string, or string) string {
	if len(items) == 0 {
		return or
	}
	return strings.Join(items, ", ")
}

func orDefault(s, def string) string {
	if s == "" {
		return def
	}
	return s
}

func statusLine(a *spec.Artifact) string {
	if a.Missing() {
		return "MISSING"
	}
	return string(a.Status)
}

func baseName(p string) string {
	if p == "." || p == "" {
		if wd, err := os.Getwd(); err == nil {
			p = wd
		}
	}
	return filepath.Base(p)
}