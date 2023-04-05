Что выведет программа? Объяснить вывод программы. Объяснить как работают defer’ы и их порядок вызовов.

```go
package main

import (
	"fmt"
)


func test() (x int) {
	defer func() {
		x++
	}()
	x = 1
	return
}


func anotherTest() int {
	var x int
	defer func() {
		x++
	}()
	x = 1
	return x
}


func main() {
	fmt.Println(test())
	fmt.Println(anotherTest())
}
```

Ответ:
```

Программа выведет:
2
1

В функции test() defer функция увеличивает значение x, которое 
было установлено в 1, а затем возвращаемое значение увеличивается на 1.
 
В функции anotherTest() defer функция увеличивает значение x, 
которое еще не было установлено возвращаемым значением, 
поэтому она не влияет на результат.

```
