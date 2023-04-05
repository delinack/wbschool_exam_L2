package main

import (
	"bufio"
	"os"
	"reflect"
	"strings"
	"testing"
)

func TestSortFile(t *testing.T) {
	testCases := []struct {
		name       string
		inputFile  string
		outputFile string
		flag       *flags
		expected   string
	}{
		{
			name:       "test basic alphabetical sort",
			inputFile:  "input1.txt",
			outputFile: "output1.txt",
			flag: &flags{
				key:    0,
				unique: false,
			},
			expected: "apple\nbanana\nkiwi\norange\n",
		},
		{
			name:       "test reverse alphabetical sort",
			inputFile:  "input1.txt",
			outputFile: "output2.txt",
			flag: &flags{
				key:     0,
				reverse: true,
			},
			expected: "orange\nkiwi\nbanana\napple\n",
		},
		{
			name:       "test unique alphabetical sort",
			inputFile:  "input2.txt",
			outputFile: "output3.txt",
			flag: &flags{
				key:    0,
				unique: true,
			},
			expected: "apple\nbanana\nkiwi\norange\n",
		},
		{
			name:       "test numeric sort",
			inputFile:  "input3.txt",
			outputFile: "output4.txt",
			flag: &flags{
				key:     0,
				numeric: true,
			},
			expected: "1\n2\n3\n4\n5\n7\n",
		},
		{
			name:       "test column sort",
			inputFile:  "input4.txt",
			outputFile: "output5.txt",
			flag: &flags{
				key:     1,
				numeric: true,
			},
			expected: "a 1\nc 1\nc 2\nb 2\na 2\nb 3\n",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			f := tc.flag
			err := f.sortFile(tc.inputFile, tc.outputFile)
			if err != nil {
				t.Fatalf("failed to sort file: %v", err)
			}

			file, err := os.Open(tc.outputFile)
			if err != nil {
				t.Fatalf("failed to open output file: %v", err)
			}
			defer file.Close()

			scanner := bufio.NewScanner(file)
			var actual strings.Builder
			for scanner.Scan() {
				actual.WriteString(scanner.Text() + "\n")
			}

			if actual.String() != tc.expected {
				t.Fatalf("expected sorted data to be:\n%v\nBut got:\n%v", tc.expected, actual.String())
			}
		})
	}
}

func TestGetValue(t *testing.T) {
	input := "qwe asd zxc"

	col0 := getValue(input, 0)
	if col0 != "qwe" {
		t.Errorf("expected %s, but got %s", "qwe", col0)
	}

	col1 := getValue(input, 1)
	if col1 != "asd" {
		t.Errorf("expected %s, but got %s", "asd", col1)
	}

	col2 := getValue(input, 2)
	if col2 != "zxc" {
		t.Errorf("expected %s, but got %s", "zxc", col2)
	}

	col3 := getValue(input, 3)
	if col3 != "" {
		t.Errorf("expected empty string, but got %s", col3)
	}
}

func TestUniqueStrings(t *testing.T) {
	input := []string{"qwe", "asd", "qwe", "zxc", "asd"}
	expected := []string{"qwe", "asd", "zxc"}
	output := uniqueStrings(input)

	if !reflect.DeepEqual(output, expected) {
		t.Errorf("expected %v, but got %v", expected, output)
	}
}
