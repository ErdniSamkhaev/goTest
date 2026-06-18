package main

import "fmt"

//	func main() {
//		var age int = 32
//		var name string = "Erdni"
//		fmt.Println(name, age)
//	}
//
// ------------------------------------------------------------
//
//	func main() {
//		var city = "Riga"   // Тип выведен сам string
//		// := работает только внутри функций
//		country := "Latvia" // Короткое объявление, тоже выводит тип
//		fmt.Println(city, country)
//	}
//
// --------------------------------------------------------------
//
//	func main() {
//		// zero value
//		//  Выведет 0 0  false
//		// В Go у каждого типа есть нулевое значение: число → 0,
//		// строка → "" (пустая, поэтому между 0 и false ты видишь двойной пробел — там пустая строка), bool → false.
//		//  Неинициализированных «мусорных» переменных в Go не бывает.
//		var count int
//		var price float64
//		var label string
//		var ready bool
//		fmt.Println(count, price, label, ready)
//	}
//
// ---------------------------------------------------
//
//	func main () {
//		var a int = 5
//		var b float64 = 2.5
//		// fmt.Println(a + b) так низя складывать как в JS
//		fmt.Println(float64(a) + b)
//	}
//
// ---------------------------------------------------
//
//	func main () {
//		const pi = 3.14159
//		const greeting = "Привет"
//		fmt.Println(pi, greeting)
//		// константу низя менять
//		pi = 4
//		fmt.Println(pi)
//	}
//
// ----------------------------------------------
func main() {
	var weight float64 = 78
	var reps int = 12
	fmt.Println(weight + (float64(reps)))
}
