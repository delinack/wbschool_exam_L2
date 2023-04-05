package pattern

import "fmt"

/*
	Реализовать паттерн «фасад».
Объяснить применимость паттерна, его плюсы и минусы,а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Facade_pattern
*/

type SubsystemItem struct {
	name  string
	price string
}

func (si *SubsystemItem) AddItem() {
	fmt.Println("item", si.name, si.price, "added")
}

type SubsystemPayment struct {
	paid bool
}

func (sp *SubsystemPayment) GetPayment() {
	if sp.paid {
		fmt.Println("payment received")
	} else {
		fmt.Println("payment is not received")
	}
}

type SubsystemDelivery struct {
	delivered bool
}

func (sd *SubsystemDelivery) DeliveryItem() {
	if sd.delivered {
		fmt.Println("order delivered")
	} else {
		fmt.Println("order is not delivered")
	}
}

type OrderFacade struct { // фасад, ссылающийся на элементы подсистемы
	item     *SubsystemItem
	payment  *SubsystemPayment
	delivery *SubsystemDelivery
}

func NewFacade(itemName, itemPrice string, isPaid, isDelivered bool) *OrderFacade { // конструктор фасада
	return &OrderFacade{
		&SubsystemItem{name: itemName, price: itemPrice},
		&SubsystemPayment{paid: isPaid},
		&SubsystemDelivery{delivered: isDelivered},
	}
}

func (of *OrderFacade) OrderInfo() { // метод фасада, показывающий основную информацию о заказе
	of.item.AddItem()       // теперь вместо того, чтобы поотдельности дергать методы подсистемы,
	of.payment.GetPayment() // достаточно вызвать один метод фасада
	of.delivery.DeliveryItem()
}

//func main() {
//	facade := NewFacade("Shoes", "2500.00", true, false)
//	facade.OrderInfo()
//}

/*
	Фасад является структурным паттерном.
	Он предоставляет простой интерфейс к сложной подсистеме.

	Применимость:
	Когда нужен простой интерфейс к сложной подсистеме.
	Когда нужно разложить подсистему на отдельные слои.

	+ изолирует клиентов от компонентов системы
	+ уменьшает зависимость между подсистемой и клиентом
	- есть риск создания божественного объекта (хранящий в себе или делающий "слишком много"),
	  но в таком случае следует ввести дополнительный фасад.
*/
