package main

import (
	"fmt"
)

// структура описывается на уровне пакета
// type User struct {
// 	Name string
// 	Age  int
// }

// func main() {
// 	u := User{Name: "Erdni", Age: 32}
// 	fmt.Println(u)
// 	fmt.Println(u.Name)
// }
// --------------------
// меняем поля
// type User struct {
// 	Name string
// 	Age  int
// }

// func main() {
// 	u := User{Name: "Erdni", Age: 32}
// 	u.Age = 33
// 	fmt.Println(u.Age)
// }
// ----------------------
// снова zero value
// объявил структуру без значений → все поля сами встали в свои zero values: Name → "", Age → 0.
// Никакого null или undefined объекта, как в JS. Структура всегда существует целиком, с занулёнными полями
// type User struct {
// 	Name string
// 	Age  int
// }

//	func main() {
//		var u User
//		fmt.Println(u)
//		fmt.Println(u.Name == "", u.Age)
//	}
//
// -------------------------------
// выведет {Name:Erdni Age:30}, с именами полей.
// Обычный Println показывает только значения {Erdni 30}, а %+v — ещё и имена
// type User struct {
// 	Name string
// 	Age  int
// }

//	func main() {
//		u := User{Name: "Erdni", Age: 30}
//		fmt.Printf("%+v\n", u)
//	}
//
// ----------------------------------
//
//	ломаем: поле, которого нет.
//
//	type User struct {
//		Name string
//		Age  int
//	}
//
// // Не скомпилируется — u.Email undefined. И вот ключевое отличие от JS: в JS ты можешь налепить на объект любое свойство на лету (obj.email = ... — и оно появится).
// // В Go структура имеет фиксированный набор полей, заданный при объявлении типа. Хочешь Email — добавь его в type User struct, иначе никак.
//
//	func main() {
//		u := User{Name: "Erdni", Age: 30}
//		u.Email = "test@mail.ru"
//		fmt.Println(u)
//	}
//
// ------------------------------
// МЕТОДЫ
// в скобках перед именем — (u User) — называется получатель (receiver)
// Она и делает Greet методом «структуры User».
// Внутри метода u — это и есть тот экземпляр, на котором его вызвали. Вызываешь как u.Greet().
// описал структуру, а потом отдельно привязал к ней функции через receiver. Никакого class
// type User struct {
// 	Name string
// 	Age  int
// }

// func (u User) Greet() string {
// 	return "Привет, " + u.Name
// }

//	func main() {
//		u := User{Name: "Erdni", Age: 32}
//		fmt.Println(u.Greet())
//	}
//
// ---------------------------------------
// ловушка: метод, который пытается изменить структуру.
// (u User) — метод получает копию структуры. Меняет копию, а оригинал не трогает.
// type User struct {
// 	Name string
// 	Age  int
// }

// func (u User) Brithday() {
// 	u.Age = u.Age + 1
// }

// func main() {
// 	u := User{Name: "Erdni", Age: 32}
// 	u.Brithday()
// 	fmt.Println(u.Age)
// }
// --------------------------------
// receiver-указатель
// *User означает «указатель на User» — метод работает с самим оригиналом, а не с его копией
// type User struct {
// 	Name string
// 	Age  int
// }

// func (u *User) Brithday() {
// 	u.Age = u.Age + 1
// }

//	func main() {
//		u := User{Name: "Erdni", Age: 32}
//		u.Brithday()
//		fmt.Println(u.Age)
//	}
//
// -------------------------
// Задача (площадь прямоугольника)
type Rectangle struct {
	Width  float64
	Height float64
}
// без u *Rectangle (метод только читает и ничего не меняет)
// но вообще есть стайл-гайды, которые по умолчанию вообще всё делают на указателях для единообразия
func (r Rectangle) Area() float64 {
	result := r.Width * r.Height
	return result
}

func main() {
	r := Rectangle{Width: 3, Height: 4}
	result := r.Area()
	fmt.Println(result)
}
