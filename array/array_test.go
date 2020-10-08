package array

import "testing"

func TestArray(t *testing.T) {
	arr := [5]int{1, 2, 3, 4, 5}
	sum := array(arr)
	want := 15
	if sum != want {
		t.Errorf("got %v want %v", sum, want)
	}
}
