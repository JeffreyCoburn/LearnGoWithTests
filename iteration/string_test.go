package iteration

import (
	"strings"
	"testing"
)

func TestCompare(t *testing.T) {
	var a, b string
	var want int

	assertCorrectMessage := func(t *testing.T, got, want int) {
		t.Helper()
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("compare strings that are equal", func(t *testing.T) {
		a = "aaa"
		b = "aaa"
		want = 0
		got := strings.Compare(a, b)
		assertCorrectMessage(t, got, want)

	})

	t.Run("compare strings where a>b", func(t *testing.T) {
		a = "aaa"
		b = "AAA"
		want = 1
		got := strings.Compare(a, b)
		assertCorrectMessage(t, got, want)
	})

	t.Run("compare strings where a<b", func(t *testing.T) {
		a = "AAA"
		b = "aaa"
		want = -1
		got := strings.Compare(a, b)
		assertCorrectMessage(t, got, want)
	})

	t.Run("compare empty strings", func(t *testing.T) {
		a = ""
		b = ""
		want = 0
		got := strings.Compare(a, b)
		assertCorrectMessage(t, got, want)
	})
}
