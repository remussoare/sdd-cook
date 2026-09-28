package install

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// LatestReleaseTag returns the latest release tag (e.g. "v0.1.0") of sdd-cook
// on GitHub. Network failures are returned as errors; callers should treat
// them as non-fatal (offline installs still work).
func LatestReleaseTag() (string, error) {
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Get("https://api.github.com/repos/remussoare/sdd-cook/releases/latest")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("github api: %s", resp.Status)
	}
	var body struct {
		TagName string `json:"tag_name"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
		return "", err
	}
	if body.TagName == "" {
		return "", fmt.Errorf("no releases found")
	}
	return body.TagName, nil
}

// NewerThan reports whether latest is a newer version than current.
// Versions are compared as dot-separated numbers after stripping a leading
// "v"; "dev" is always considered older than any release.
func NewerThan(current, latest string) bool {
	if current == "" || current == "dev" {
		return true
	}
	return compareVersions(norm(current), norm(latest)) < 0
}

func norm(v string) []int {
	v = strings.TrimPrefix(strings.TrimSpace(v), "v")
	var out []int
	for _, p := range strings.Split(v, ".") {
		n, err := strconv.Atoi(strings.TrimSpace(p))
		if err != nil {
			return out
		}
		out = append(out, n)
	}
	return out
}

func compareVersions(a, b []int) int {
	for i := 0; i < len(a) || i < len(b); i++ {
		var x, y int
		if i < len(a) {
			x = a[i]
		}
		if i < len(b) {
			y = b[i]
		}
		if x != y {
			if x < y {
				return -1
			}
			return 1
		}
	}
	return 0
}
