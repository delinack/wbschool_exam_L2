package pattern

import "fmt"

/*
	Реализовать паттерн «стратегия».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Strategy_pattern
*/

type RouteStrategy interface { // интерфейс стратегии маршрута
	BuildRoute(from, to string) string
}

// FastestRouteStrategy и ShortestRouteStrategy - это структуры,
// представляющие различные стратегии построения маршрутов

type FastestRouteStrategy struct{} // стратегия быстрейшего маршрута

func (frs *FastestRouteStrategy) BuildRoute(from, to string) string {
	return fmt.Sprintf("Быстрейший маршрут из %s в %s", from, to)
}

type ShortestRouteStrategy struct{} // стратегия кратчайшего маршрута

func (srs *ShortestRouteStrategy) BuildRoute(from, to string) string {
	return fmt.Sprintf("Кратчайший маршрут из %s в %s", from, to)
}

type Navigator struct { // структура навигатора
	routeStrategy RouteStrategy // поле для хранения текущей стратегии маршрута
}

func (n *Navigator) SetRouteStrategy(strategy RouteStrategy) {
	n.routeStrategy = strategy
}

func (n *Navigator) Navigate(from, to string) {
	route := n.routeStrategy.BuildRoute(from, to)
	fmt.Println(route)
}

//func main() {
//	// создаём экземпляры стратегий маршрутов и навигатора
//	fastestStrategy := &FastestRouteStrategy{}
//	shortestStrategy := &ShortestRouteStrategy{}
//	navigator := &Navigator{}
//
//	// устанавливаем различные стратегии маршрута в навигаторе и вызывает метод Navigate() для каждой из них
//	navigator.SetRouteStrategy(fastestStrategy)
//	navigator.Navigate("A", "B")
//
//	navigator.SetRouteStrategy(shortestStrategy)
//	navigator.Navigate("A", "B")
//}

/*
	Стратегия является поведенческим паттерном.
	Она определяет семейство схожих алгоритмов и помещает каждый в собственную структуру.
	После этого алгоритмы можно взаимозаменять по ходу программы.

	Применимость:
	Когда нужно использовать разные вариации какого-то алгоритма внутри одного объекта.
	Когда похожие подструктуры отличают только некоторым поведением.
	Когда нужно скрыть детали реализации алгоритмов от другоих структур.
	Когда различне вариации алгоритмов реализованы в виде развесистого условного оператора.

	+ "горячая" замена алгоритмов "налету"
	+ изолирует код и данные алгоритмов от остального кода
	+ уход от наследования к делегированию
	+ реализует пирнцип открытости/закрытости
	- усложняет код засчёт доп. структур
	- клиент должен знать разницу между стратегиями, чтобы выбрать подходящую
*/
