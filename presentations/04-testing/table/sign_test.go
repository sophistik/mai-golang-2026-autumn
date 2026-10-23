package tabledemo

import "testing"

func TestSign(t *testing.T) {
	tests := []struct {
		name string
		in   int
		want string
	}{
		{"zero", 0, "ноль"},
		{"positive", 5, "плюс"},
		{"negative", -3, "минус"},
	}
	for _, tt := range tests {
		if got := Sign(tt.in); got != tt.want {
			t.Errorf("%s: Sign(%d) = %q, want %q", tt.name, tt.in, got, tt.want)
		}
	}
}
