Что выведет программа? Объяснить вывод программы. Рассказать про внутреннее устройство слайсов и что происходит при передачи их в качестве аргументов функции.

```go
package main

import (
	"fmt"
)

func main() {
	var s = []string{"1", "2", "3"}
	modifySlice(s)
	fmt.Println(s)
}

func modifySlice(i []string) {
	i[0] = "3"
	i = append(i, "4")
	i[1] = "5"
	i = append(i, "6")
}
```

Ответ:
```
Программа выведет:

[3 2 3]

Когда слайс передается в функцию как аргумент, создается новый слайс, 
который является копией исходного, но с тем же указателем на массив.
При изменении элементов внутри функции, изменения отображаются на исходном слайсе,
так как он по-прежнему ссылается на тот же массив, что и исходный слайс. 

Но при добавлении элементов внутри функции, создается новый массив с большей ёмкостью, 
который не является частью исходного слайса. Поэтому изменения, внесенные в новый массив, 
не будут отображаться в исходном слайсе.
```