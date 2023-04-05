package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestGrep(t *testing.T) {
	testCases := []struct {
		name           string
		input          string
		flags          *flags
		expectedOutput string
	}{
		{
			name:  "case_sensitive_matching",
			input: "Hello\nWorld\nhello\nWORLD",
			flags: &flags{
				pattern:    "hello",
				ignoreCase: false,
				fixed:      true,
			},
			expectedOutput: "hello\n",
		},
		{
			name:  "case_insensitive_matching",
			input: "Hello\nWorld\nhello\nWORLD",
			flags: &flags{
				pattern:    "hello",
				ignoreCase: true,
				fixed:      true,
			},
			expectedOutput: "Hello\nhello\n",
		},
		{
			name:  "invert_match",
			input: "Hello\nWorld\nhello\nWORLD",
			flags: &flags{
				pattern:    "hello",
				ignoreCase: false,
				fixed:      true,
				invert:     true,
			},
			expectedOutput: "Hello\nWorld\nWORLD\n",
		},
		{
			name:  "count_matches",
			input: "Hello\nWorld\nhello\nWORLD",
			flags: &flags{
				pattern:    "hello",
				ignoreCase: true,
				fixed:      true,
				count:      true,
			},
			expectedOutput: "Hello\nhello\n2\n",
		},
		{
			name:  "print_line_numbers",
			input: "Hello\nWorld\nhello\nWORLD",
			flags: &flags{
				pattern:    "hello",
				ignoreCase: false,
				fixed:      true,
				lineNum:    true,
			},
			expectedOutput: "3:hello\n",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			inputReader := strings.NewReader(tc.input)
			outputWriter := &bytes.Buffer{}

			err := grep(inputReader, outputWriter, tc.flags)
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
			}

			output := outputWriter.String()
			if output != tc.expectedOutput {
				t.Errorf("Expected output: %q, got: %q", tc.expectedOutput, output)
			}
		})
	}
}
