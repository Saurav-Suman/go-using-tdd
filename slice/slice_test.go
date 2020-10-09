package slice

import (
	"reflect"
	"testing"
)

func TestSlice(t *testing.T) {
	slicedata := []int{1, 2, 3}
	sum := slice(slicedata)
	want := 6
	if sum != want {
		t.Errorf("got %v want %v", sum, want)
	}
}

func TestSumAllSlice(t *testing.T) {
	sum := sumAllSlice([]int{1, 2, 3}, []int{1, 2, 3})
	want := []int{6, 6}
	/*
		Go does not let you use equality operators with slices. You could write a function to iterate over each sum and want
		slice and check their values but for convenience sake,
		we can use reflect.DeepEqual which is useful for seeing if any two variables are the same.
	*/
	if !reflect.DeepEqual(sum, want) {
		t.Errorf("got %v want %v", sum, want)
	}
}

func TestSumTail(t *testing.T) {
	sum := sumTail([]int{2, 3, 4}, []int{5, 6, 7})
	want := []int{7, 13}
	if !reflect.DeepEqual(sum, want) {
		t.Errorf("got %v want %v", sum, want)
	}
}
