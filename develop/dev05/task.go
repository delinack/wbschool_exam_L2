package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strings"
)

/*
=== Утилита grep ===

Реализовать утилиту фильтрации (man grep)

Поддержать флаги:
-A - "after" печатать +N строк после совпадения
-B - "before" печатать +N строк до совпадения
-C - "context" (A+B) печатать ±N строк вокруг совпадения
-c - "count" (количество строк)
-i - "ignore-case" (игнорировать регистр)
-v - "invert" (вместо совпадения, исключать)
-F - "fixed", точное совпадение со строкой, не паттерн
-n - "line num", печатать номер строки

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

type flags struct {
	after      int
	before     int
	context    int
	count      bool
	ignoreCase bool
	invert     bool
	fixed      bool
	lineNum    bool
	pattern    string
}

func grep(r io.Reader, w io.Writer, f *flags) error {
	// создаём сканер для чтения и буфер для записи
	scanner := bufio.NewScanner(r)
	writer := bufio.NewWriter(w)
	defer writer.Flush() // отложенный вызов записи данных

	if f.ignoreCase { // если указан флаг игнорированиия регистра, приводим всё к нижнему
		f.pattern = strings.ToLower(f.pattern)
	}

	var re *regexp.Regexp
	var err error
	// если не установлен флаг точного совпадения, компилируем регулярку
	// для сопоставления с текстом
	if !f.fixed {
		re, err = regexp.Compile(f.pattern)
		if err != nil {
			return err
		}
	}

	var (
		lineNum     int      // для хранения номера строки
		matchBuffer []string // для буфера совпадений
		matchCount  int      // для количества найденных совпадений
	)

	for scanner.Scan() { // читаем строки
		line := scanner.Text()
		lineNum++ // считаем количество строк для флага lineNum

		match := false
		// проверяем совпадения в зависимости от указанных флагов
		if f.fixed {
			if f.ignoreCase {
				match = strings.Contains(strings.ToLower(line), f.pattern)
			} else {
				match = strings.Contains(line, f.pattern)
			}
		} else {
			if f.ignoreCase {
				match = re.MatchString(strings.ToLower(line))
			} else {
				match = re.MatchString(line)
			}
		}

		// если указан флаг invert, инвертировуем совпадения
		if f.invert {
			match = !match
		}

		// обрабатываем совпадения и выводим результат
		if match {
			matchCount++
			for i := 0; i < len(matchBuffer); i++ { // выведем строки перед совпадением
				fmt.Fprintln(writer, matchBuffer[i])
			}
			matchBuffer = matchBuffer[:0]

			// если указан флаг lineNum, выведем номер строки
			if f.lineNum {
				fmt.Fprintf(writer, "%d:", lineNum)
			}
			fmt.Fprintln(writer, line)
		} else { // если совпадение не найдено, то добавляем текущую строку в буфер контекстных строк
			if f.before > 0 || f.context > 0 {
				matchBuffer = append(matchBuffer, line)
				if len(matchBuffer) > f.before+f.context {
					matchBuffer = matchBuffer[1:]
				}
			}
		}
	}

	// если установлен флаг count, выведем количество найденных совпадений
	if f.count {
		fmt.Fprintln(writer, matchCount)
	}

	return scanner.Err() // вернём ошибку, если произошла ошибка сканирования и nil, если всё в порядке
}

func main() {
	f := &flags{}

	// парсим флаги
	flag.IntVar(&f.after, "A", 0, "Print +N lines after match")
	flag.IntVar(&f.before, "B", 0, "Print +N lines before match")
	flag.IntVar(&f.context, "C", 0, "Print ±N lines around match")
	flag.BoolVar(&f.count, "c", false, "Print count of matching lines")
	flag.BoolVar(&f.ignoreCase, "i", false, "Ignore case")
	flag.BoolVar(&f.invert, "v", false, "Invert match")
	flag.BoolVar(&f.fixed, "F", false, "Fixed match, not pattern")
	flag.BoolVar(&f.lineNum, "n", false, "Print line number")

	flag.Parse()

	// если аргумент не один, то завершаем программу
	if flag.NArg() != 1 {
		log.Fatal("Invalid arguments count")
	}

	f.pattern = flag.Arg(0) // записывам первый аргумент в структуру

	if err := grep(os.Stdin, os.Stdout, f); err != nil {
		log.Fatal(err)
	}
}
