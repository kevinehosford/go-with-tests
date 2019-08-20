package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	got := Add(2, 2)
	want := 4

	if got != want {
		t.Errorf("got %d but wanted %d", got, want)
	}
}

func ExampleAdd() {
	sum := Add(2, 2)
	fmt.Println(sum)
	// Output: 4
}
