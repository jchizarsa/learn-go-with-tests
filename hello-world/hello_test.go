/* Tests must follow these rules
 * 1. File name needs to be <x>_test.go
 * 2. Test functions must start with "Test"
 * 3. Each test function only takes in one argument `t *testing.T`
 * 4. Must `import "testing"` to use the *testing.T type.
 */
package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Chris")
	want := "Hello, Chris"

	if got != want {
		t.Errorf("got %q want %q", got, want) // %q wraps values in quotes. See `go doc fmt` for more.
	}
}
