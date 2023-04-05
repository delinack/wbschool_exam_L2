package pattern

import "fmt"

/*
	Реализовать паттерн «состояние».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/State_pattern
*/

type State interface { // интерфейс состояния
	DoAction(*Person)
	String() string
}

// Work, Rest и Sleep - это структуры, представляющие различные состояния человека

type Work struct{} // состояние работы

func (w *Work) DoAction(person *Person) {
	fmt.Println("Человек работает")
	person.SetState(w)
}

func (w *Work) String() string {
	return "работа"
}

type Rest struct{} // состояние отдыха

func (r *Rest) DoAction(person *Person) {
	fmt.Println("Человек отдыхает")
	person.SetState(r)
}

func (r *Rest) String() string {
	return "отдых"
}

type Sleep struct{} // состояние сна

func (s *Sleep) DoAction(person *Person) {
	fmt.Println("Человек спит")
	person.SetState(s)
}

func (s *Sleep) String() string {
	return "сон"
}

type Person struct { // структура человека
	state State // поле для хранения текущего состояния
}

func (p *Person) SetState(state State) {
	p.state = state
}

func (p *Person) GetCurrentState() State {
	return p.state
}

//func main() {
//	// создаём экземпляры состояний и человека
//	work := &Work{}
//	rest := &Rest{}
//	sleep := &Sleep{}
//	person := &Person{}
//
//	// вызываем метод DoAction() для каждого состояния, передавая экземпляр человека
//	work.DoAction(person)
//	fmt.Println("Текущее состояние:", person.GetCurrentState())
//
//	rest.DoAction(person)
//	fmt.Println("Текущее состояние:", person.GetCurrentState())
//
//	sleep.DoAction(person)
//	fmt.Println("Текущее состояние:", person.GetCurrentState())
//}

/*
	Состояние является поведенческим паттерном.
	Он позволяет объектам менять поведение в зависимости от своего состояния.
	Извне создаётся впечатление, что изменилась структура объекта.

	Применимость:
	Когда есть объект, поведение которого кардинально меняется в зависимости от внутреннего состояния.
	Когда код содержит множество бусловных операторов, которые выбирают поведение в зависимости от текущих полей структуры.

	+ избавляет от множества условных операторов машины состояний
	+ концентрирует в одном место код, связанныйм определённым состоянием
	+ упрощает код контекста
	- может усложнить код, если состояний мало или они редко меняются
*/
