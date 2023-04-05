package main

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

/*
=== Задача на распаковку ===

Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
	- "a4bc2d5e" => "aaaabccddddde"
	- "abcd" => "abcd"
	- "45" => "" (некорректная строка)
	- "" => ""
Дополнительное задание: поддержка escape - последовательностей
	- qwe\4\5 => qwe45 (*)
	- qwe\45 => qwe44444 (*)
	- qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.

Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func Unpack(input string) (string, error) {
	var output strings.Builder
	escaped := false // меняем на true, если встретим один /
	repeat := false  // меняем на true, если встретили "букву"

	for _, r := range input { // итерация по символам строки
		if unicode.IsDigit(r) {
			if !repeat && !escaped { // если символ является цифрой,
				return "", errors.New("invalid string") // но "повторять" нечего, то возвращаем ошибку
			}
			if escaped {
				output.WriteRune(r) // записываем символ, если перед ним \
				escaped = false     // сбрасываем флаг
			} else {
				repeatCount, _ := strconv.Atoi(string(r))                           // сколько раз нужно повторить
				prevRune := output.String()[output.Len()-1]                         // достаём предыдущий символ
				output.WriteString(strings.Repeat(string(prevRune), repeatCount-1)) // записываем символы в строку
				repeat = false                                                      // сбрасываем флаг
			}
		} else if string(r) == `\` {
			if escaped {
				output.WriteRune(r) // записываем символ, если перед ним \
				escaped = false     // сбрасываем флаг
			} else {
				escaped = true // переводим флаг в true, чтобы знать, что мы встреитли один \
			}
		} else { // если символ не является цифрой и обратным слешем
			if escaped {
				return "", errors.New("invalid escape sequence")
			}
			output.WriteRune(r)
			repeat = true
		}
	}

	if escaped { // если на конце остался \, строка неккоректная
		return "", errors.New("invalid escape sequence")
	}

	return output.String(), nil
}

//func main() {
//	testCases := []string{
//		"a4bc2d5e", // "aaaabccddddde"
//		"abcd",     // "abcd"
//		"45",       // (некорректная строка)
//		"",         // ""
//		`qwe\4\5`,  // qwe45
//		`qwe\45`,   // qwe44444
//		`qwe\\5`,   // qwe\\\\\
//		`qwe\\5\`,  // (некорректная строка)
//	}
//
//	for _, s := range testCases {
//		result, err := Unpack(s)
//		if err != nil {
//			fmt.Printf("Error: %v for input: %s\n\n", err, s)
//		} else {
//			fmt.Printf("Unpacked string for %s: %s\n\n", s, result)
//		}
//	}
//}
