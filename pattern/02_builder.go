package pattern

import "fmt"

/*
	Реализовать паттерн «строитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Builder_pattern
*/

type House struct { // структура дома, представляющая основной продукт, который будет "строиться"
	Windows int
	Doors   int
	Floors  int
}

func (h *House) String() string {
	return fmt.Sprintf("Дом с %d окнами, %d дверьми и %d этажами", h.Windows, h.Doors, h.Floors)
}

type HouseBuilder interface { // интерфейс строителя: определяет методы для строительства дома
	SetWindows(windows int) HouseBuilder
	SetDoors(doors int) HouseBuilder
	SetFloors(floors int) HouseBuilder
	Build() *House
}

type ConcreteHouseBuilder struct { // конкретная реализация строителя
	House *House
}

func NewConcreteHouseBuilder() *ConcreteHouseBuilder { // конструктор конкретного строителя
	return &ConcreteHouseBuilder{
		House: &House{},
	}
}

func (c *ConcreteHouseBuilder) SetWindows(windows int) HouseBuilder {
	c.House.Windows = windows
	return c
}

func (c *ConcreteHouseBuilder) SetDoors(doors int) HouseBuilder {
	c.House.Doors = doors
	return c
}

func (c *ConcreteHouseBuilder) SetFloors(floors int) HouseBuilder {
	c.House.Floors = floors
	return c
}

func (c *ConcreteHouseBuilder) Build() *House {
	return c.House
}

//func main() {
//	builder := NewConcreteHouseBuilder() // создаём экземпляр строителя
//	house := builder.SetWindows(10).SetDoors(4).SetFloors(2).Build()
//	fmt.Println(house)
//}

/*
	Строитель является порождающим паттерном.
	Он позволяет создавать сложные объекты пошагово.
	Он даёт возможность использовать один и тот же код
	строительства для построения разных представлений объектов.

	Применимость:
	Когда нужно создавать разные представления какого-то объекта.
	Когда нужно собирать сложные составные объекты.

	+ позволяет создавать продукты пошагово
	+ позволяет использовать один и тот же код для создания различных продуктов
	+ изолирует сложный код сборки продукта от бизнес-логики
	- усложняет код засчёт доп. структур
	- клиент привязывается к конкретным структурам строителей
*/
