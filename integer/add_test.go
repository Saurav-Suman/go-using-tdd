package integer

import "testing"

func TestAdd(t *testing.T) {
	got := add(2, 6)
	want := 8
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
