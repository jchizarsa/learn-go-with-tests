/* Tests must follow these rules
 * 1. File name needs to be <x>_test.go
 * 2. Test functions must start with "Test"
 * 3. Each test function only takes in one argument `t *testing.T`
 * 4. Must `import "testing"` to use the *testing.T type.
 */
package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"

		if got != want {
			assertCorrectMessage(t, got, want)
		}
	})
	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		if got != want {
			assertCorrectMessage(t, got, want)
		}
	})
	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in French", func(t *testing.T) {
		got := Hello("Napoleon", "French")
		want := "Bonjour, Napoleon"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in Tagalog", func(t *testing.T) {
		got := Hello("Josh", "Tagalog")
		want := "Kumusta, Josh"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want) // %q wraps values in quotes. See `go doc fmt` for more.
	}
}
