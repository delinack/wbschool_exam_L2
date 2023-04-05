package pattern

import "fmt"

/*
	Реализовать паттерн «комманда».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Command_pattern
*/

type Order interface { // собой интерфейс заказа, который определяет метод Execute()
	Execute()
}

type Cook struct{} // структура повара, который может готовить блюда

func (c *Cook) CookDish(dish string) {
	fmt.Printf("Повар готовит %s\n", dish)
}

type CookDishCommand struct { // структура команды готовки блюда, которая реализует интерфейс Order
	cook *Cook
	dish string
}

func NewCookDishCommand(cook *Cook, dish string) *CookDishCommand {
	return &CookDishCommand{cook: cook, dish: dish}
}

func (c *CookDishCommand) Execute() {
	c.cook.CookDish(c.dish)
}

type Waiter struct { // структура официанта, который принимает заказы и передает их на выполнение
	orders []Order
}

func (w *Waiter) TakeOrder(order Order) {
	w.orders = append(w.orders, order)
}

func (w *Waiter) ProcessOrders() {
	for _, order := range w.orders {
		order.Execute()
	}
	w.orders = nil
}

//func main() {
//	cook := &Cook{}     // создаём экземпляр повара
//	waiter := &Waiter{} // создаём экзмепляр официанта
//
//	// создаём экземпляры заказов через конструктор
//	order1 := NewCookDishCommand(cook, "пиццу")
//	order2 := NewCookDishCommand(cook, "салат")
//
//	// вызываем методы официанта
//	waiter.TakeOrder(order1)
//	waiter.TakeOrder(order2)
//
//	waiter.ProcessOrders()
//}

/*
	Команда является поведенческим паттерном.
	Она превращает запросы в объекты, позоляя передавать их как аргументы при вызове методов,
	ставить запросы в очередь, логировать их, поддерживать отмену операций.

	Применимость:
	Когда нужно составить очередь из операций, выполнять их по расписанию или передавать их по сети.
	Когда нужна поддержка отмены.
	Когда нужно параметризовать объекты выполняемым действием.

	+ убирает прямую зависимость между объектами
	+ позволяет реализовать простую отмену, повтор операций, отложенный запуск команд
	+ позволет собирать сложные команды из простых
	+ реализует принцип открытости/закрытости
	- усложняет код засчет доп. структур
*/
