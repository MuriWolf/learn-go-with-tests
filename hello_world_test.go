package main

import (
	"math/rand"
	"testing"
)

func BenchmarkRandInt(b *testing.B) {
	for b.Loop() {
		rand.Int()
	}
}
func TestHello(t *testing.T) {
	t.Run("Saying hello to people", func(t *testing.T) {
		got := Hello("Murillo")
		want := "Hello, Murillo"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("Saying hello with empty string", func(t *testing.T) {
		got := Hello("")
		want := "Hello, world"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
