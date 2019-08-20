package iteration

// Repeat repeats a string count times
func Repeat(s string, count int) string {
	var word string
	for i := 0; i < count; i++ {
		word += s
	}

	return word
}
