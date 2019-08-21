package arrays

// Sum adds all of the numbers in the given array
func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}

	return sum
}

// SumAll returns an array containing each sum of argument
func SumAll(slices ...[]int) []int {
	lengthOfSlices := len(slices)

	sums := make([]int, lengthOfSlices)

	for idx, slice := range slices {
		sums[idx] = Sum(slice)
	}

	return sums
}
