package main

import "testing"

func TestDictionary(t *testing.T) {
	dictionary := Dictionary{"test": "verify if something works as expected"}

	t.Run("search for a existing word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "verify if something works as expected"

		assertStrings(t, want, got)
	})

	t.Run("search for a non-existing word", func(t *testing.T) {
		_, error := dictionary.Search("inexist")

		if error == nil {
			t.Fatal("expected an error")
		}
		assertError(t, error, ErrWordDefinitionNotFound)
	})
}

func assertStrings(t *testing.T, want, got string) {
	t.Helper()
	if got != want {
		t.Errorf("want %q got %q", want, got)
	}
}

func assertError(t *testing.T, want, got error) {
	t.Helper()
	if got != want {
		t.Errorf("expected error %q but got %q", want, got)
	}
}
