package spec

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func writeArtifact(t *testing.T, root, name, status, component, evolution string) {
	t.Helper()
	dir := filepath.Join(root, ".spec", Dir(name))
	if err := os.MkdirAll(dir, 0o755); err != nil {
		t.Fatal(err)
	}
	content := "---\nartifact: " + name + "\ncomponent: " + component + "\nevolution: " + evolution + "\nstatus: " + status + "\n---\n\n# Body\n"
	if err := os.WriteFile(filepath.Join(dir, File(name)), []byte(content), 0o644); err != nil {
		t.Fatal(err)
	}
}

func writeRawArtifact(t *testing.T, root, name, content string) {
	t.Helper()
	dir := filepath.Join(root, ".spec", Dir(name))
	if err := os.MkdirAll(dir, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(dir, File(name)), []byte(content), 0o644); err != nil {
		t.Fatal(err)
	}
}

func testRoot(t *testing.T) string {
	t.Helper()
	return t.TempDir()
}

func TestParseArtifact(t *testing.T) {
	root := testRoot(t)
	writeArtifact(t, root, "objective", "DRAFT", "demo", "v1")
	s, err := Load(root)
	if err != nil {
		t.Fatal(err)
	}
	a := s.Artifacts["objective"]
	if a.Missing() {
		t.Fatal("expected objective artifact")
	}
	if a.Status != Draft || a.Component != "demo" || a.Evolution != "v1" {
		t.Fatalf("unexpected parse: %+v", a)
	}
	if a.Body == "" {
		t.Fatal("expected body to be parsed")
	}
}

func TestMissingArtifact(t *testing.T) {
	s, err := Load(testRoot(t))
	if err != nil {
		t.Fatal(err)
	}
	for _, name := range Chain {
		if !s.Artifacts[name].Missing() {
			t.Fatalf("%s should be missing", name)
		}
	}
}

func TestInvalidFrontmatter(t *testing.T) {
	root := testRoot(t)
	writeRawArtifact(t, root, "plan", "# Plan without frontmatter\n")
	s, err := Load(root)
	if err != nil {
		t.Fatal(err)
	}
	if got := s.Artifacts["plan"].Status; got != Invalid {
		t.Fatalf("expected INVALID, got %s", got)
	}
}

func TestSetStatus(t *testing.T) {
	root := testRoot(t)
	writeArtifact(t, root, "objective", "DRAFT", "demo", "v1")
	s, _ := Load(root)
	a := s.Artifacts["objective"]
	if err := a.SetStatus(Accepted); err != nil {
		t.Fatal(err)
	}
	s2, err := Load(root)
	if err != nil {
		t.Fatal(err)
	}
	if got := s2.Artifacts["objective"].Status; got != Accepted {
		t.Fatalf("expected ACCEPTED after SetStatus, got %s", got)
	}
}

func TestNextEvolution(t *testing.T) {
	if got := NextEvolution("v1"); got != "v2" {
		t.Fatalf("expected v2, got %s", got)
	}
	if got := NextEvolution("v9"); got != "v10" {
		t.Fatalf("expected v10, got %s", got)
	}
	if got := NextEvolution(""); got != "v1" {
		t.Fatalf("expected v1, got %s", got)
	}
}

func TestCreateObjectiveNeverOverwrites(t *testing.T) {
	root := testRoot(t)
	if err := CreateObjective(root, "demo", "v1"); err != nil {
		t.Fatal(err)
	}
	if err := CreateObjective(root, "demo", "v1"); err == nil {
		t.Fatal("expected error on second create")
	}
}

func TestCreateChangeRequest(t *testing.T) {
	root := testRoot(t)
	if err := CreateChangeRequest(root, "demo", "v2", "add exports"); err != nil {
		t.Fatal(err)
	}
	entries, err := os.ReadDir(filepath.Join(root, ".spec", "changes"))
	if err != nil || len(entries) != 1 {
		t.Fatalf("expected one change request, err=%v entries=%d", err, len(entries))
	}
}

func TestActiveChangeRequest(t *testing.T) {
	root := testRoot(t)
	if _, active := ActiveChangeRequest(root); active {
		t.Fatal("no .spec should mean no active change request")
	}
	if err := CreateChangeRequest(root, "demo", "v2", "add exports"); err != nil {
		t.Fatal(err)
	}
	p, active := ActiveChangeRequest(root)
	if !active {
		t.Fatal("PROPOSED change request should be active")
	}
	if p == "" {
		t.Fatal("expected path of active change request")
	}
	entries, _ := os.ReadDir(filepath.Join(root, ".spec", "changes"))
	data, _ := os.ReadFile(filepath.Join(root, ".spec", "changes", entries[0].Name()))
	rewritten := strings.Replace(string(data), "status: PROPOSED", "status: REJECTED", 1)
	if err := os.WriteFile(filepath.Join(root, ".spec", "changes", entries[0].Name()), []byte(rewritten), 0o644); err != nil {
		t.Fatal(err)
	}
	if _, active := ActiveChangeRequest(root); active {
		t.Fatal("REJECTED change request should not be active")
	}
}