package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strings"
	"testing"
)

func TestParseFields(t *testing.T) {
	tests := []struct {
		input    string
		expected []int
	}{
		{"1,2", []int{0, 1}},
		{"3,4,5", []int{2, 3, 4}},
	}

	for _, test := range tests {
		result, err := parseFields(test.input)
		if err != nil {
			t.Errorf("Ошибка разбора полей: %v", err)
		}
		if !equalSlices(result, test.expected) {
			t.Errorf("Ожидалось %v, получено %v", test.expected, result)
		}
	}
}

func equalSlices(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func runCut(in io.Reader, out io.Writer, fieldsFlag, delimiterFlag string, separatedFlag bool) error {
	// Преобразование флага полей в срез индексов
	fields, err := parseFields(fieldsFlag)
	if err != nil {
		return fmt.Errorf("Ошибка: %v", err)
	}

	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		line := scanner.Text()
		if separatedFlag && !strings.Contains(line, delimiterFlag) {
			continue
		}

		parts := strings.Split(line, delimiterFlag)
		var selected []string
		for _, index := range fields {
			if index < len(parts) {
				selected = append(selected, parts[index])
			}
		}
		fmt.Fprintln(out, strings.Join(selected, delimiterFlag))
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("Ошибка чтения: %v", err)
	}

	return nil
}

func TestCut(t *testing.T) {
	input := "a\tb\tc\nd\te\tf\n"
	expected := "b\tc\ne\tf\n"

	in := bytes.NewBufferString(input)
	out := &bytes.Buffer{}

	err := runCut(in, out, "2,3", "\t", false)
	if err != nil {
		t.Errorf("Ошибка выполнения runCut: %v", err)
	}

	result := out.String()
	if result != expected {
		t.Errorf("Ожидалось %q, получено %q", expected, result)
	}
}
