package pattern

import "fmt"

/*
	Реализовать паттерн «цепочка вызовов».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Chain-of-responsibility_pattern
*/

type SupportTicket struct { // структура заявки в техподдержку
	level       int
	description string
}

type SupportHandler interface { // интерфейс обработчика заявки
	SetNext(SupportHandler)
	HandleTicket(*SupportTicket)
}

type BaseSupportHandler struct { // базовая структура обработчика заявки
	next SupportHandler // поле для связи между уровнями
}

func (b *BaseSupportHandler) SetNext(next SupportHandler) {
	b.next = next
}

// Level1SupportHandler,
// Level2SupportHandler и
// Level3SupportHandler -
// это структуры обработчиков заявок разных уровней,
// каждый из которых наследует базовый обработчик и
// реализует метод HandleTicket() для обработки заявок своего уровня.

// Level1SupportHandler - обработчик заявок 1-го уровня
type Level1SupportHandler struct {
	BaseSupportHandler
}

func (l1 *Level1SupportHandler) HandleTicket(ticket *SupportTicket) {
	if ticket.level == 1 {
		fmt.Printf("Заявка уровня 1: %s\n", ticket.description)
	} else if l1.next != nil {
		l1.next.HandleTicket(ticket)
	}
}

// Level2SupportHandler - обработчик заявок 2-го уровня
type Level2SupportHandler struct {
	BaseSupportHandler
}

func (l2 *Level2SupportHandler) HandleTicket(ticket *SupportTicket) {
	if ticket.level == 2 {
		fmt.Printf("Заявка уровня 2: %s\n", ticket.description)
	} else if l2.next != nil {
		l2.next.HandleTicket(ticket)
	}
}

// Level3SupportHandler - обработчик заявок 3-го уровня
type Level3SupportHandler struct {
	BaseSupportHandler
}

func (l3 *Level3SupportHandler) HandleTicket(ticket *SupportTicket) {
	if ticket.level == 3 {
		fmt.Printf("Заявка уровня 3: %s\n", ticket.description)
	} else if l3.next != nil {
		l3.next.HandleTicket(ticket)
	}
}

//func main() {
//	// создаём обработчики заявок разных уровней
//	l1 := &Level1SupportHandler{}
//	l2 := &Level2SupportHandler{}
//	l3 := &Level3SupportHandler{}
//
//	// связываем обработчики в цепочку обязанностей
//	l1.SetNext(l2)
//	l2.SetNext(l3)
//
//	// создаём набор заявок на техподдержку и передает их цепочке обязанностей, начиная с первого
//	tickets := []*SupportTicket{
//		{level: 1, description: "Проблема с интернетом"},
//		{level: 2, description: "Сервер не отвечает"},
//		{level: 3, description: "Проблема с электронной почтой"},
//	}
//
//	for _, ticket := range tickets {
//		l1.HandleTicket(ticket)
//	}
//}

/*
	Цепочка обязаностей является поведенческим паттерном.
	Она позволяет передавать запросы последовательно по цепочке обработчиков.
	Каждый последующий обработчик решает, может ли он обработать запрос сам и
	стоит ли передавать его дальше.

	Применимость:
	Когда программа содержит несколько объектов, способных обрабатывать тот или иной запрос,
	но заранее неизвестно какой запрос придёт и какой обработчик понадобится.
	Когда важно, чтобы обработчики выполнялись один за другим в строгом порядке.
	Когда набор объектов, способных обработать запрос, должен задаваться динамически.

	+ уменьшает зависимость между клиентом и обработчиком
	+ реализует принцип единственной обязанности
	+ реализует принцип открытости/закрытости
	- запрос может остаться никем не обработанный
*/
