package main

import "testing"

func TestHello(t *testing.T) {

	t.Run("Hello world", func(t *testing.T) {

		got := hello_world()
		want := "hello world"
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}

	})

}
