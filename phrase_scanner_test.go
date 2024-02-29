package go_phrase_scanner

import (
	"testing"
)

func DoStringSearch(t *testing.T, phrases []string, text string, expected []string) {
	n := NewTrie(phrases)
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
		[]string{"hello", "goodbye"})
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
		[]string{"once upon a time", "time out", "a long time ago", "time bomb"})
}
