package main

import "testing"

func TestDictionary(t *testing.T) {
	dictionary := Dictionary{"test": "verify if something works as expected"}

	t.Run("search for an existing word", func(t *testing.T) {
		word := "test"
		definition := "verify if something works as expected"

		assertDefinition(t, dictionary, word, definition)
	})

	t.Run("search for a non-existing word", func(t *testing.T) {
		_, error := dictionary.Search("inexist")

		if error == nil {
			t.Fatal("expected an error")
		}
		assertError(t, error, NotFound)
	})

	t.Run("add word definition", func(t *testing.T) {
		newDictionary := Dictionary{}
		word := "test"
		definition := "verify if something works as expected"

		error := newDictionary.Add(word, definition)

		assertError(t, nil, error)
		assertDefinition(t, newDictionary, word, definition)
	})

	t.Run("prevent from adding already defined word", func(t *testing.T) {
		word := "test"
		definition := "another definition"

		error := dictionary.Add(word, definition)

		assertError(t, AlreadyDefined, error)
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

func assertDefinition(t *testing.T, dictionary Dictionary, word, definition string) {
	got, error := dictionary.Search(word)

	if error != nil {
		t.Fatal("didn't expected error")
	}

	assertStrings(t, definition, got)
}
