package loop

import "testing"

func TestLoop(t *testing.T) {
	got := loop("a", 5)
	want := "aaaaa"

	if got != want {
		t.Errorf("got  %v want %v", got, want)
	}
}

func BenchmarkLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		loop("a", 5)
	}
}
