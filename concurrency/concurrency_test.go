package main

import (
	"maps"
	"testing"
	"time"
)

func mockWebsiteChecker(url string) bool {
	return url != "waat://sdfsdfsd.dfd"
}

func TestCheckWebsites(t *testing.T) {
	websites := []string{
		"http://google.com",
		"http://youtube.com",
		"waat://sdfsdfsd.dfd",
	}

	want := map[string]bool{
		"http://google.com":   true,
		"http://youtube.com":  true,
		"waat://sdfsdfsd.dfd": false,
	}

	got := CheckWebsites(mockWebsiteChecker, websites)

	if !maps.Equal(want, got) {
		t.Fatalf("wanted %v got %v", want, got)
	}
}

func slowStubWebsiteChecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func BenchmarkCheckWebsites(b *testing.B) {
	urls := make([]string, 100)
	for i := 0; i < len(urls); i++ {
		urls[i] = "a url"
	}
	b.ResetTimer() // resets the timer of the test before it runs
	for i := 0; i < b.N; i++ {
		CheckWebsites(slowStubWebsiteChecker, urls)
	}
}
