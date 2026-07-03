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
	got := Hello("Murillo")
	want := "Hello, Murillo"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
