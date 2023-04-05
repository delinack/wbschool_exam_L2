package main

import (
	"sort"
	"strings"
)

/*
=== Поиск анаграмм по словарю ===

Напишите функцию поиска всех множеств анаграмм по словарю.
Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.

Входные данные для функции: ссылка на массив - каждый элемент которого - слово на русском языке в кодировке utf8.
Выходные данные: Ссылка на мапу множеств анаграмм.
Ключ - первое встретившееся в словаре слово из множества
Значение - ссылка на массив, каждый элемент которого, слово из множества. Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func sortString(s string) string { // функция, сортирующая строки
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
}

type sortRunes []rune // тип с методами, сортирующими руны в строке, удолетворяющий интерфейсу пакета sort

func (s sortRunes) Len() int {
	return len(s)
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func Search(words *[]string) *map[string][]string {
	anagramSets := make(map[string][]string)

	for _, word := range *words { // пройдём по каждому слову в словаре
		sortedWord := sortString(strings.ToLower(word))                 // приведём его к нижнему регистру и отсортируем его буквы
		anagramSets[sortedWord] = append(anagramSets[sortedWord], word) // добавим его в сет
	}

	for k, v := range anagramSets { // удаляем сеты с одним элементом
		if len(v) <= 1 {
			delete(anagramSets, k)
		} else {
			sort.Strings(v) // отсортируем сет по возрастанию
			anagramSets[k] = v
		}
	}

	return &anagramSets
}
