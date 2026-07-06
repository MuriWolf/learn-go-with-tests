package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	t.Run("repeat with string", func(t *testing.T) {
		repeated := Repeat("a", 5, "string")
		expected := "aaaaa"
		assertIsExpected(t, repeated, expected)
	})

	t.Run("repeat with write", func(t *testing.T) {
		repeated := Repeat("a", 5, "write")
		expected := "aaaaa"
		assertIsExpected(t, repeated, expected)
	})

	t.Run("repeat with write string", func(t *testing.T) {
		repeated := Repeat("a", 5, "writeString")
		expected := "aaaaa"
		assertIsExpected(t, repeated, expected)
	})

	t.Run("repeat with undefined method", func(t *testing.T) {
		repeated := Repeat("a", 5, "notDefined")
		expected := ""
		assertIsExpected(t, repeated, expected)

	})
}

func assertIsExpected(t *testing.T, repeated, expected string) {
	if repeated != expected {
		t.Errorf("got %q but expected %q", repeated, expected)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for b.Loop() {
		Repeat("a", 5, "writeString")
	}
}

func ExampleRepeat() {
	repeated := Repeat("a", 3, "writeString")
	fmt.Println(repeated)
	// Output: aaa
}
