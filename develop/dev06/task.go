package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

/*
=== Утилита cut ===

Принимает STDIN, разбивает по разделителю (TAB) на колонки, выводит запрошенные

Поддержать флаги:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func parseFields(fields string) ([]int, error) { // преобразуем строку с номерами полей в слайс индексов
	var indexes []int
	parts := strings.Split(fields, ",")
	for _, part := range parts {
		var index int
		_, err := fmt.Sscanf(part, "%d", &index)
		if err != nil {
			return nil, err
		}
		indexes = append(indexes, index-1)
	}
	return indexes, nil
}

func main() {
	// определим флаги
	fieldsFlag := flag.String("f", "", "Выбрать поля (колонки)")
	delimiterFlag := flag.String("d", "\t", "Использовать другой разделитель")
	separatedFlag := flag.Bool("s", false, "Только строки с разделителем")

	flag.Parse()

	// преобразуем флаги полей в слайс индексов
	fields, err := parseFields(*fieldsFlag)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка: %v\n", err)
		os.Exit(1)
	}

	// читаем строки
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if *separatedFlag && !strings.Contains(line, *delimiterFlag) {
			continue
		}

		parts := strings.Split(line, *delimiterFlag) // разделим строки по разделителю
		var selected []string
		for _, index := range fields {
			if index < len(parts) {
				selected = append(selected, parts[index])
			}
		}
		fmt.Println(strings.Join(selected, *delimiterFlag))
	}

	if err = scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка чтения: %v\n", err)
		os.Exit(1)
	}
}
