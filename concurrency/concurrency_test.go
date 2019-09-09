package concurrency

import (
	"reflect"
	"testing"
	"time"
)

func mockCheckWebsite(url string) bool {
	return url != "waaat://"
}

func TestCheckwebsites(t *testing.T) {
	urls := []string{
		"http://foo.com",
		"http://bar.com",
		"waaat://",
	}

	got := CheckWebsites(mockCheckWebsite, urls)
	want := map[string]bool{
		"http://foo.com": true,
		"http://bar.com": true,
		"waaat://":       false,
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("%v is not %v", got, want)
	}
}

func mockSlowCheckWebsite(url string) bool {
	time.Sleep(20 * time.Millisecond)

	return true
}

func BenchmarkCheckwebsites(b *testing.B) {
	urls := make([]string, 100)

	for i := range urls {
		urls[i] = "a url"
	}

	for i := 0; i < b.N; i++ {
		CheckWebsites(mockSlowCheckWebsite, urls)
	}
}
