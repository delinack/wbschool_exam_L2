package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
=== Утилита sort ===

Отсортировать строки (man sort)
Основное

Поддержать ключи

-k — указание колонки для сортировки
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки

Дополнительное

Поддержать ключи

-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учётом суффиксов

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

type flags struct {
	key     int
	numeric bool
	reverse bool
	unique  bool
}

func (f *flags) sortFile(inputFile string, outputFile string) error {
	file, err := os.Open(inputFile) // чтение содержимого файла
	if err != nil {
		return err
	}
	defer file.Close() // отложенный вызов закртытия файла

	scanner := bufio.NewScanner(file)
	var data []string    // слайс для хранения данных
	for scanner.Scan() { // сканируем содержимое файла
		data = append(data, scanner.Text())
	}

	less := func(i, j int) bool { // сортируем данные
		if f.numeric {            // сортируем по числу, если указан ключ -n
			ni, err := strconv.Atoi(getValue(data[i], f.key))
			if err == nil {
				nj, err := strconv.Atoi(getValue(data[j], f.key))
				if err == nil {
					return ni < nj
				}
			}
		}
		return getValue(data[i], f.key) < getValue(data[j], f.key)
	}

	if f.reverse { // сортируем в обратном порядке, если указан ключ -r
		sort.Slice(data, func(i, j int) bool {
			return !less(i, j)
		})
	} else {
		sort.Slice(data, less)
	}

	if f.unique { // исключаем повторы, если указан ключ -u
		data = uniqueStrings(data)
	}

	file, err = os.Create(outputFile) // создаем файл для записи результата
	if err != nil {
		return err
	}
	defer file.Close() // отложенный вызов закрытия файла

	writer := bufio.NewWriter(file)
	for _, str := range data {
		writer.WriteString(str + "\n")
	}
	writer.Flush() // записываем отсортированные данные в файл

	return nil
}

func getValue(s string, col int) string { // функция для сортировки по колонкам
	fields := strings.Fields(s) // сплитуем строки
	if len(fields) <= col {     // проверка на наличие колонки
		return ""
	}
	return fields[col] // возвращаем нужную колонку
}

func uniqueStrings(slice []string) []string { // функция, возвращающая уникальные значения
	unique := make(map[string]bool)
	var result []string
	for _, str := range slice {
		if !unique[str] { // добавляем строку, если она ещё не встречалась
			unique[str] = true
			result = append(result, str)
		}
	}
	return result
}
