package recursive_search

import (
	"io/ioutil"
	"log"
	"os"
	"testing"
)

const needle = "god"

var haystack []string

func TestMain(m *testing.M) {
	for i := 0; i < 2000; i++ {
		haystack = append(haystack, []string{"tac", "esuom", "horse", "mouse", "cat", "dog"}...)
	}
	haystack = append(haystack, []string{"frog", "god", "esroh", "elephant", "tiger"}...)
	haystack = append(haystack, []string{"frog", "ged", "esroh", "elephant", "tiger"}...)

	log.SetOutput(ioutil.Discard)
	os.Exit(m.Run())
}

func TestStringRecursiveSearch(t *testing.T) {
	result := StartStrRecursiveSearch(haystack, needle)
	if result == -1 {
		t.Errorf("needle '%s' not found in haystack: %q\n", needle, haystack)
	} else {
		t.Logf("Target string: '%s' found at index %d\n", needle, result)
	}
}

func BenchmarkStartStrRecursiveSearch(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		StartStrRecursiveSearch(haystack, needle)
	}
}

func BenchmarkLinearSearch(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		for i := 0; i < len(haystack); i++ {
			if haystack[i] == needle {
				break
			}
		}
	}
}
