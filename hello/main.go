package main

import "fmt"

// Просто два типа
// type Dog struct{ Name string }
// type Cat struct{ Name string }

// func (d Dog) Sound() string { return "Гав" }
// func (c Cat) Sound() string { return "Мяу" }

//	func main() {
//		d := Dog{Name: "Курама"}
//		c := Cat{Name: "Мося"}
//		fmt.Println(d.Sound())
//		fmt.Println(c.Sound())
//	}
//
// -------------------
// а если я хочу функцию, которая примет любого, кто умеет издавать звук — и собаку, и кошку, и кого угодно ещё
// добавляем интерфейс
// type Dog struct{ Name string }
// type Cat struct{ Name string }

// func (d Dog) Sound() string { return "Гав" }
// func (c Cat) Sound() string { return "Мяу" }

// // интерфейс
// type Sounder interface {
// 	Sound() string
// }

// // функция принимающая интерефейс
// func describe(s Sounder) {
// 	fmt.Println("Звук:", s.Sound())
// }

//	func main() {
//		d := Dog{Name: "Курама"}
//		c := Cat{Name: "Мося"}
//		describe(d)
//		describe(c)
//	}
//
// -------------------------------
// Делаем все тоже самое но с добавлением типа без метода
// и будет ошибка implement Sounder
// type Dog struct{ Name string }
// type Cat struct{ Name string }
// type Rock struct{}

// func (d Dog) Sound() string { return "Гав" }
// func (c Cat) Sound() string { return "Мяу" }

// // интерфейс
// type Sounder interface {
// 	Sound() string
// }

// // функция принимающая интерефейс
// func describe(s Sounder) {
// 	fmt.Println("Звук:", s.Sound())
// }

//	func main() {
//		d := Dog{Name: "Курама"}
//		c := Cat{Name: "Мося"}
//		r := Rock{}
//		describe(d)
//		describe(c)
//		describe(r)
//	}
//
// -----------------------------
// Теперь добавляем третий тип так, чтобы не трогать describe(функицю принимающая интерфейс)
// type Dog struct{ Name string }
// type Cat struct{ Name string }
// type Bird struct{ Name string }

// func (d Dog) Sound() string  { return "Гав" }
// func (c Cat) Sound() string  { return "Мяу" }
// func (b Bird) Sound() string { return "Кря" }

// // интерфейс
// type Sounder interface {
// 	Sound() string
// }

// // функция принимающая интерефейс
// func describe(s Sounder) {
// 	fmt.Println("Звук:", s.Sound())
// }

// func main() {
// 	d := Dog{Name: "Курама"}
// 	c := Cat{Name: "Мося"}
// 	describe(d)
// 	describe(c)
// 	describe(Bird{Name: "Кеша"})
// }
// ---------------------------------------
// теперь error
// type Dog struct{ Name string }
// type Cat struct{ Name string }
// type Bird struct{ Name string }

// func (d Dog) Sound() string  { return "Гав" }
// func (c Cat) Sound() string  { return "Мяу" }
// func (b Bird) Sound() string { return "Кря" }

// // интерфейс
// type Sounder interface {
// 	Sound() string
// }
// // интрефейс ошибки
// type error interface{
// 	Error() string
// }

// // функция принимающая интерефейс
// func describe(s Sounder) {
// 	fmt.Println("Звук:", s.Sound())
// }

//	func main() {
//		d := Dog{Name: "Курама"}
//		c := Cat{Name: "Мося"}
//		describe(d)
//		describe(c)
//		describe(Bird{Name: "Кеша"})
//	}
//
// -------------------------------
// Задача
// Создаем типы простые
type English struct{}
type Russian struct{}
// даем типам метод Greet
func (e English) Greet() string { return "Hello" }
func (r Russian) Greet() string { return "Привет" }
// интерфейс
type Greeter interface {
	Greet() string
}
// функция принимающая интерфейс
func sayHi(g Greeter) {
	fmt.Println(g.Greet())
}

func main() {
	sayHi(English{})
	sayHi(Russian{})
}
