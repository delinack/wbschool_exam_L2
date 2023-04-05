Что выведет программа? Объяснить вывод программы.

```go
package main

type customError struct {
	msg string
}

func (e *customError) Error() string {
	return e.msg
}

func test() *customError {
	{
		// do something
	}
	return nil
}

func main() {
	var err error
	err = test()
	if err != nil {
		println("error")
		return
	}
	println("ok")
}
```

Ответ:
```
Программа выведет:
error

В err лежит результат выполнения функции test(), которая возвращает нам
указатель на customError.
err != nil выполняетя, т.к. значение любого интерфейса является nil в случае,
когда значение и тип интерфейса это nil.

То есть сравниваются два разных типа - *main.customError != nill
```
