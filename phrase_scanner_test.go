package go_phrase_scanner

import (
	"strings"
	"testing"
)

func DoStringSearch(t *testing.T, phrases []string, text string, expected []string) {
	n, _ := NewTrie(phrases)
	i := 0
	for s := range n.ScanString(text) {
		if i == len(expected) {
			t.Errorf("Expected end of values, got \"%s\"", s)
			return
		}
		if s != expected[i] {
			t.Errorf("Expected \"%s\", got \"%s\"", expected[i], s)
			return
		}
		i++
	}
	if i < len(expected) {
		t.Errorf("Expected \"%s\", got end of values", expected[i])
		return
	}
}

func TestSimple(t *testing.T) {
	DoStringSearch(t,
		[]string{"hello", "goodbye"},
		"hello goodbye",
		[]string{"hello", "goodbye"})
}

func TestVariedCase(t *testing.T) {
	DoStringSearch(t,
		[]string{"hello", "goodbye"},
		"hellO GOOdbYe",
		[]string{"hellO", "GOOdbYe"})
}

func TestEmptyResults(t *testing.T) {
	DoStringSearch(t,
		[]string{"hello", "goodbye"},
		"helo goodby hel lo good bye ello helllo goodbya",
		[]string{})
}

func TestOverlappingResults(t *testing.T) {
	DoStringSearch(t,
		[]string{"hello", "low", "water", "waters"},
		"hello hellow slow low lowaters lowater waters",
		[]string{"hello", "hello", "low", "low", "low", "low", "water", "waters", "low", "water", "water", "waters"})
}

func TestPhrases(t *testing.T) {
	DoStringSearch(t,
		[]string{"a long time ago", "once upon a time", "time out", "time bomb"},
		"Once upon a time outside of a city, a long time ago, there was a time bomb that had a timeout.",
		[]string{"Once upon a time", "time out", "a long time ago", "time bomb"})
}

func FuzzScanReaderAgainstScanString(f *testing.F) {
	testcases := []string{
		"Hello world hello there",
		"Goodbye test case case good test",
	}
	for _, tc := range testcases {
		f.Add(tc) // Use f.Add to provide a seed corpus
	}
	f.Fuzz(func(t *testing.T, orig string) {
		strs_ := strings.Split(orig, " ")
		strs := make([]string, 0, len(strs_))
		for _, str := range strs_ {
			if str == "" {
				continue
			}
			strs = append(strs, str)
		}
		if len(strs) < 2 {
			return
		}
		n, _ := NewTrie(strs[1:])
		r_iter := n.ScanReader(strings.NewReader(strs[0]))
		for s := range n.ScanString(strs[0]) {
			rstring, ok := <-r_iter
			if !ok {
				t.Errorf("Expected \"%s\", got end of values", s)
				return
			}
			if s != rstring {
				t.Errorf("Expected \"%s\", got \"%s\"", rstring, s)
				return
			}
		}
		rstring, ok := <-r_iter
		if ok {
			t.Errorf("Expected end of values, got \"%s\"", rstring)
			return
		}
	})
}
