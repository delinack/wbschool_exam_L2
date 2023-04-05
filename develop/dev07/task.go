package main

import (
	"sync"
)

/*
=== Or channel ===

Реализовать функцию, которая будет объединять один или более done каналов в single канал если один из его составляющих каналов закроется.
Одним из вариантов было бы очевидно написать выражение при помощи select, которое бы реализовывало эту связь,
однако иногда неизестно общее число done каналов, с которыми вы работаете в рантайме.
В этом случае удобнее использовать вызов единственной функции, которая, приняв на вход один или более or каналов, реализовывала весь функционал.

Определение функции:
var or func(channels ...<- chan interface{}) <- chan interface{}

Пример использования функции:
sig := func(after time.Duration) <- chan interface{} {
	c := make(chan interface{})
	go func() {
		defer close(c)
		time.Sleep(after)
}()
return c
}

start := time.Now()
<-or (
	sig(2*time.Hour),
	sig(5*time.Minute),
	sig(1*time.Second),
	sig(1*time.Hour),
	sig(1*time.Minute),
)

fmt.Printf(“fone after %v”, time.Since(start))
*/

// orChannel объединяет один или более done каналов в один single канал,
// который закрывается, когда один из составляющих каналов закрывается.
func orChannel(channels ...<-chan interface{}) <-chan interface{} {
	// если нет каналов, возвращаем nil
	if len(channels) == 0 {
		return nil
	}

	// если только один канал, возвращаем его без изменений
	if len(channels) == 1 {
		return channels[0]
	}

	orDone := make(chan interface{})
	go func() {
		defer close(orDone) // закроем orDone канал после завершения работы

		var wg *sync.WaitGroup

		wg.Add(len(channels))
		// пройдёмся по каждому каналу и дождёмся его закрытия
		for _, ch := range channels {
			go func(c <-chan interface{}) {
				defer wg.Done() // завершим ожидание после закрытия входного канала

				select {
				case <-c: // ждём закрытия входного канала
				case <-orDone: // ждём закрытия orDone канала
				}
			}(ch)
		}

		wg.Wait() // ожидаем завершения всех горутин
	}()

	return orDone
}
