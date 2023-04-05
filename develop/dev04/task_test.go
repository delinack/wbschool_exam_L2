package main

import (
	"reflect"
	"testing"
)

func TestSortString(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{"empty string", "", ""},
		{"single character", "a", "a"},
		{"multiple characters", "hello", "ehllo"},
		{"mixed case", "AbC", "ACb"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := sortString(tc.input)
			if got != tc.expected {
				t.Errorf("sortString(%q) = %q; expected %q", tc.input, got, tc.expected)
			}
		})
	}
}

func TestSearch(t *testing.T) {
	testCases := []struct {
		name         string
		input        []string
		expected     map[string][]string
		expectedSize int
	}{
		{"empty list", []string{}, map[string][]string{}, 0},
		{"single word", []string{"hello"}, map[string][]string{}, 0},
		{"multiple words, no anagrams", []string{"hello", "world"}, map[string][]string{}, 0},
		{"multiple words, one anagram", []string{"listen", "silent"}, map[string][]string{"eilnst": {"listen", "silent"}}, 1},
		{"multiple words, multiple anagram 1", []string{"listen", "silent", "earth", "heart"}, map[string][]string{"aehrt": {"earth", "heart"}, "eilnst": {"listen", "silent"}}, 2},
		{"multiple words, multiple anagram 2", []string{"пятак", "пятка", "тяпка", "листок", "столик", "слиток"},
			map[string][]string{"акптя": {"пятак", "пятка", "тяпка"}, "иклост": {"листок", "слиток", "столик"}}, 2},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := Search(&tc.input)
			if len(*got) != tc.expectedSize {
				t.Errorf("Search(%q) returned map of size %d; expected size %d", tc.input, len(*got), tc.expectedSize)
			}
			if !reflect.DeepEqual(*got, tc.expected) {
				t.Errorf("Search(%q) = %v; expected %v", tc.input, *got, tc.expected)
			}
		})
	}
}
