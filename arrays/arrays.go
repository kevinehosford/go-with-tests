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
	var sums []int

	for _, slice := range slices {
		sums = append(sums, Sum(slice))
	}

	return sums
}

// SumAllTails does something
func SumAllTails(slices ...[]int) []int {
	var sums []int
	for _, slice := range slices {
		sums = append(sums, Sum(slice[1:]))
	}

	return sums
}
