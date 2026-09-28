package install

import "testing"

func TestNewerThan(t *testing.T) {
	cases := []struct {
		current, latest string
		want            bool
	}{
		{"dev", "v0.1.0", true},
		{"", "v0.1.0", true},
		{"v0.1.0", "v0.1.0", false},
		{"v0.1.0", "v0.2.0", true},
		{"v0.2.0", "v0.1.9", false},
		{"v0.1.9", "v0.1.10", true},
		{"0.1.0", "v0.1.1", true},
		{"v1.0.0", "v0.9.9", false},
	}
	for _, c := range cases {
		if got := NewerThan(c.current, c.latest); got != c.want {
			t.Errorf("NewerThan(%q, %q) = %v, want %v", c.current, c.latest, got, c.want)
		}
	}
}
