package cmd

import "testing"

func TestValidLabel(t *testing.T) {
	cases := []struct {
		in   string
		want bool
	}{
		{"add exports", true},
		{"component-1", true},
		{"", false},
		{"bad\nname", false},
		{"bad:name", false},
		{"bad\rname", false},
		{"tab\tname", false},
		{string([]byte{0x00}), false},
		{"ok", true},
	}
	for _, c := range cases {
		if got := validLabel(c.in); got != c.want {
			t.Errorf("validLabel(%q) = %v, want %v", c.in, got, c.want)
		}
	}
	long := make([]byte, 201)
	for i := range long {
		long[i] = 'a'
	}
	if validLabel(string(long)) {
		t.Error("expected names over 200 bytes to be rejected")
	}
}
