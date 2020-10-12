package structs

import "testing"

func TestBasicStruct(t *testing.T) {

	data := StructData{}
	data.a = 10
	data.b = 20

	got := StructSum(data)
	want := 30

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}

}

func TestStructR(t *testing.T) {
	data := StructData{}
	data.a = 10
	data.b = 20
	got := data.StructSumR()
	want := 30

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
