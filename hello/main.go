package main

import "fmt"

// Интерфейсы под разные типы

// type Sounder interface {
// 	Sound() string
// }

// type Dog struct{ Name string }
// type Cat struct{ Name string }

// func (d Dog) Sound() string { return "Гав" }
// func (c Cat) Sound() string { return "Мяу" }

//	func main() {
//		// []Sounder — это срез, в который сложены разные конкретные типы (Dog и Cat вперемешку), но снаружи они все «звучащие».
//		animals := []Sounder{
//			Dog{Name: "Курама"},
//			Cat{Name: "Мося"},
//			Dog{Name: "Rex"},
//		}
//		for _, a := range animals {
//			fmt.Println(a.Sound())
//		}
//	}
//
// ----------------------------------
// функция работающая со всем зоопарком (расширяемость)
// type Sounder interface {
// 	Sound() string
// }

// type Dog struct{ Name string }
// type Cat struct{ Name string }

// func (d Dog) Sound() string { return "Гав" }
// func (c Cat) Sound() string { return "Мяу" }

// func concert(sounders []Sounder) {
// 	for _, s := range sounders {
// 		fmt.Println("♪", s.Sound())
// 	}
// }

//	func main() {
//		animals := []Sounder{
//			Dog{Name: "Курама"},
//			Cat{Name: "Мося"},
//		}
//		// принимает срез интерфейса
//		concert(animals)
//	}
//
// --------------------
// Пустой интерфейс any
// any (раньше писали interface{}) — это интерфейс с пустым списком методов. А раз методов ноль — то любой тип его удовлетворяет (у каждого типа «есть все ноль методов»).
// Поэтому в any можно положить что угодно: число, строку, bool, структуру. Это Go-аналог «динамической переменной», как let x в JS, куда можно пихать всё подряд.
// func main() {
// 	var x any
// 	x = 42
// 	fmt.Println(x)

// 	x = "Привет"
// 	fmt.Println(x)

//		x = true
//		fmt.Println(x)
//	}
//
// -------------------------
// Но в any есть ограничение
//
//	func main() {
//		// Не скомпилируется. Хотя ты знаешь, что там 42, Go видит только тип any — а с any нельзя делать +, потому что компилятор не знает, число там или строка.
//		// Вот плата за «any»: ты теряешь типовую информацию, и обычные операции становятся недоступны.
//		var x any = 42
//		fmt.Println(x + 1)
//	}
//
// ---------------------------
// Как достать типо обратно: type assertion
//
//	func main() {
//		// x.(int) — это type assertion: «я утверждаю, что внутри x лежит int, дай мне его как int». После этого n — нормальный int, с ним + 1 работае
//		var x any = 42
//		n := x.(int)
//		fmt.Println(n + 1)
//	}
//
// ------------------------------
// А если ошибся типом?
//
//	func main() {
//		// .(int) тут уронила бы панику. А форма с двумя значениями n, ok := — безопасная: если тип не совпал, ok = false, а n = 0 (zero value).
//		// Это тот же value, ok паттерн, что у карт (v, ok := m[key]) — два возврата, второй говорит «получилось или нет».
//		var x any = "строка, а не число"
//		n, ok := x.(int)
//		fmt.Println(n, ok)
//		// any только когда тип реально заранее неизвестен (универсальные контейнеры, парсинг произвольного JSON)
//	}
//
// ----------------------------------
// Задача
// интерфейс
type Shape interface {
	Area() float64
}

// типы
type Circle struct{ Radius float64 }
type Square struct{ Side float64 }

// методы
func (c Circle) Area() float64 { return 3.14159 * c.Radius * c.Radius }
func (s Square) Area() float64 { return s.Side * s.Side }

func main() {
	result := []Shape{
		Circle{Radius: 10},
		Square{Side: 10},
	}
	for _, v := range result {
		fmt.Println(v.Area())
	}
}
