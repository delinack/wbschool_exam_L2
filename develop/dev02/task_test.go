package main

import (
	"errors"
	"testing"
)

func TestUnpack(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
		err      error
	}{
		{"a4bc2d5e", "aaaabccddddde", nil},
		{"abcd", "abcd", nil},
		{"45", "", errors.New("invalid string")},
		{"", "", nil},
		{`qwe\4\5`, "qwe45", nil},
		{`qwe\45`, "qwe44444", nil},
		{`qwe\\5`, `qwe\\\\\`, nil},
		{`qwe\\5\`, "", errors.New("invalid escape sequence")},
	}

	for _, tc := range testCases {
		result, err := Unpack(tc.input)
		if result != tc.expected || (err != nil && tc.err == nil) || (err == nil && tc.err != nil) || (err != nil && tc.err != nil && err.Error() != tc.err.Error()) {
			t.Errorf("Unpack(%q) = %q, %v, expected %q, %v", tc.input, result, err, tc.expected, tc.err)
		}
	}
}
