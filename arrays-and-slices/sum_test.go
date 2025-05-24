// Learn more about slices by writing more tests about them. https://blog.golang.org/go-slices-usage-and-internals
// Go playground: https://play.golang.org/p/ICCWcRGIO68
// Slicing Examples:
// - https://go.dev/play/p/bTrRmYfNYCp
// - https://go.dev/play/p/Poth8JS28sc
package arraysandslices

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		expected := 15

		if got != expected {
			t.Errorf("expected '%d' but got '%d' given, %v", expected, got, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	t.Run("collection of sums from varying number of slices", func(t *testing.T) {
		got := SumAll([]int{1, 2}, []int{2, 3}, []int{5, 7})
		want := []int{3, 5, 12}

		if !reflect.DeepEqual(got, want) { // Careful when using reflect.DeepEqual(), might still compile and compare things that shouldn't be compared
			t.Errorf("wanted '%v' got '%v'", want, got)
		}
	})
}

func TestSumAllTails(t *testing.T) {
	checkSums := func(t testing.TB, got, want []int) { // Assigning a function to a var. Great for binding a function to local variables. Reduces surface area in api
		t.Helper() // Note: This function is bound to the variable. It cannot be used outside this func.
		if !reflect.DeepEqual(got, want) {
			t.Errorf("wanted %v got %v", want, got)
		}
	}

	t.Run("collection of totals of the tails of each slice.", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{2, 3}, []int{5, 7})
		want := []int{2, 3, 7}

		checkSums(t, got, want)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}

		checkSums(t, got, want)
	})

	// Example bad test. Shows that checkSums throws proper error when given wrong param
	// t.Run("DAVE :D", func(t *testing.T) {
	// 	got := SumAllTails([]int{}, []int{3, 4, 5})
	// 	want := []int{0, 9}

	// 	checkSums(t, got, "dave")
	// })
}
