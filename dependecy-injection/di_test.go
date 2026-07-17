package main

import (
	"bytes"
	"testing"
)

func TestPrint(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "mademoiselle")

	got := buffer.String()
	want := "Bonjour, mademoiselle"

	if got != want {
		t.Errorf("want %q got %q", want, got)
	}
}
