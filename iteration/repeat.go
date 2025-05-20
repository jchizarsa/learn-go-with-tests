// Test with `go test -bench=. -benchmem`
package iteration

import "strings"

// Example for loops: https://gobyexample.com/for

func Repeat(character string, repeatCount int) string {
	var repeated strings.Builder
	for i := 0; i < repeatCount; i++ {
		repeated.WriteString(character)
	}
	return repeated.String()
}
