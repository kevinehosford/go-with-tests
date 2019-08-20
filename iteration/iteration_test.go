package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	got := Repeat("a", 10)
	want := "aaaaaaaaaa"

	if got != want {
		t.Errorf("got %q but wanted %q", got, want)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 10)
	}
}

func ExampleRepeat() {
	got := Repeat("z", 5)
	fmt.Println(got)

	// Output: zzzzz
}
