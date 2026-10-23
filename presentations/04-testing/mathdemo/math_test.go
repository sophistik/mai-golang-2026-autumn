package mathdemo

import "testing"

func TestDiv(t *testing.T) {
	got := Div(10, 2)
	if got != 5 {
		t.Errorf("Div(10, 2) = %d, want 5", got)
	}
}
