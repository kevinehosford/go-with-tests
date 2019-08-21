package arrays

// Sum adds all of the numbers in the given array
func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}

	return sum
}
