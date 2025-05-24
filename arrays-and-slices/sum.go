package arraysandslices

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers { // use `range` to iterate
		sum += number // `_` ignores index because it's not needed.
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	var sums []int
	// sums := make([]int, lengthOfNumbers) // make() lets us make a slice with a starting capacity

	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}

	return sums
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int

	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:] // Slicing a slice. Similar to Python arrays
			sums = append(sums, Sum(tail))
		}
	}

	return sums
}
