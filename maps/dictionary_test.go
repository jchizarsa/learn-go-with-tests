package maps

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is a test"}

	got := dictionary.Search("test")
	want := "this is a test"

	AssertStrings(t, got, want)
}

func AssertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
