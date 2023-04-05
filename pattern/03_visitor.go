package pattern

import "fmt"

/*
	Реализовать паттерн «посетитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Visitor_pattern
*/

type InsuranceElement interface { // интерфейс для элементов страхования
	Accept(InsuranceAgentVisitor)
}

type CarInsurance struct { // структура страхования автомобиля
	cost int
}

func (c *CarInsurance) Accept(visitor InsuranceAgentVisitor) {
	visitor.VisitCarInsurance(c)
}

type LifeInsurance struct { // структура страхования жизни
	cost int
}

func (l *LifeInsurance) Accept(visitor InsuranceAgentVisitor) {
	visitor.VisitLifeInsurance(l)
}

type InsuranceAgentVisitor interface { // интерфейс посетителя страхового агента
	VisitCarInsurance(*CarInsurance)
	VisitLifeInsurance(*LifeInsurance)
}

type InsuranceCalculator struct{} // структура для расчета стоимости страховки
// это конкретная реализация посетителя, которая реализует методы интерфейса InsuranceAgentVisitor.

func (ic *InsuranceCalculator) VisitCarInsurance(ci *CarInsurance) {
	tax := 50
	fmt.Printf("Расчет стоимости страхования автомобиля, включая налог: %d\n", ci.cost+tax)
}

func (ic *InsuranceCalculator) VisitLifeInsurance(li *LifeInsurance) {
	tax := 100
	fmt.Printf("Расчет стоимости страхования жизни, включая налог: %d\n", li.cost+tax)
}

//func main() {
//	// создаём экземпляры различных видов страхования
//	carInsurance := &CarInsurance{cost: 5000}    // страхование машины
//	lifeInsurance := &LifeInsurance{cost: 10000} // страхование жизни
//
//	insuranceCalculator := &InsuranceCalculator{} // экземпляр посетителя, считающего стоимость страховки
//
//	// на каждом объекте страхования вызовем метод Accept с аргументом в виде посетителя-калькулятора
//	carInsurance.Accept(insuranceCalculator)
//	lifeInsurance.Accept(insuranceCalculator)
//}

/*
	Посетитель является поведенческим паттерном.
	Он позволяет создавать новые операции, не меняя
	структуры объектов, над которыми эти операции могут выполняться.

	Применимость:
	Когда нужно выполнить операцию над всеми элементами сложной структуры объектов:
	посетитель позволяет применять одну и ту же операцию к объектам различных структур.

	+ упрощает добавление новых операций над всей связанной структурой объектов
	+ объединяет родственные операции в одной структуре
	+ может накапливать состояние при обходе структуры компонентов
	- применение неоправдано, если иерархия компонентов часто меняется
	- может привести к нарушению инкапсуляции компонентов
*/
