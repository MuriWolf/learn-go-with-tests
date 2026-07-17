package main

import (
	"testing"
)

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "verify if something works as expected"}

	t.Run("search defined word", func(t *testing.T) {
		word := "test"
		definition := "verify if something works as expected"

		assertDefinition(t, dictionary, word, definition)
	})

	t.Run("search undefined word", func(t *testing.T) {
		_, error := dictionary.Search("inexist")

		if error == nil {
			t.Fatal("expected an error")
		}
		assertError(t, error, NotFound)
	})
}

func TestAdd(t *testing.T) {
	dictionary := Dictionary{"test": "verify if something works as expected"}

	t.Run("add word", func(t *testing.T) {
		newDictionary := Dictionary{}
		word := "test"
		definition := "verify if something works as expected"

		error := newDictionary.Add(word, definition)

		assertError(t, nil, error)
		assertDefinition(t, newDictionary, word, definition)
	})

	t.Run("add defined word", func(t *testing.T) {
		word := "test"
		definition := "another definition"

		error := dictionary.Add(word, definition)

		assertError(t, AlreadyDefined, error)
	})
}

func TestUpdate(t *testing.T) {
	dictionary := Dictionary{"test": "verify if something works as expected"}

	t.Run("update defined word", func(t *testing.T) {
		word := "test"
		newDefinition := "another definition"

		error := dictionary.Update(word, newDefinition)

		assertError(t, nil, error)
		assertDefinition(t, dictionary, word, newDefinition)
	})

	t.Run("update undefined word", func(t *testing.T) {
		word := "nonexist"
		newDefinition := "i simply don't exist"

		error := dictionary.Update(word, newDefinition)

		assertError(t, NotDefined, error)
	})
}

func TestDelete(t *testing.T) {
	dictionary := Dictionary{"test": "verify if something works as expected"}

	t.Run("delete word", func(t *testing.T) {
		word := "test"

		error := dictionary.Delete(word)

		assertError(t, nil, error)

		_, error = dictionary.Search(word)

		assertError(t, NotFound, error)
	})

	t.Run("delete undefined word", func(t *testing.T) {
		word := "nonexist"

		error := dictionary.Delete(word)

		assertError(t, NotDefined, error)
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
	t.Helper()
	got, error := dictionary.Search(word)

	if error != nil {
		t.Fatal("didn't expected error")
	}

	assertStrings(t, definition, got)
}
