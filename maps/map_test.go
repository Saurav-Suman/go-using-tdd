package maps

import "testing"

func TestMap(t *testing.T) {

	mapdata := map[string]string{"name": "saurav"}
	got := Search(mapdata, "name")
	want := "saurav"

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestMapR(t *testing.T) {

	dict := CustomMap{"name": "saurav"}
	got := dict.Search("name")
	want := "saurav"

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
