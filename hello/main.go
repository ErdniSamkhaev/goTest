package main
// strconv.Atoi — это конвертер строки в число
import (
	"fmt"
	// "strconv"
)

//	func add(a int, b int) int {
//		return a + b
//	}
//
// Тоже add, только сократим
// func add(a, b int) int {
// 	return a + b
// }

//	func main() {
//		result := add(2, 3)
//		fmt.Println(result)
//	}
//
// ----------------------------
// Возврат двух значений
// В JS так нельзя без массива или объекта — в Go это родное.
// func divmod(a, b int) (int, int) {
// 	return a / b, a % b
// }

//	func main() {
//		q, r := divmod(17, 5)
//		fmt.Println(q, r)
//	}
//
// ------------------------
// ошибка из stdlib
// strconv.Atoi превращает строку в число
// она возвращает два значения: само число И ошибку
// Тут "42" — валидное число, поэтому ошибки нет, выведет 42 <nil>.
// <nil> значит «ошибки нет» (это nil, аналог «пустоты» для ошибки).
//
//	func main() {
//		// если указать вместо 42 что тодругое например "abc"
//		//  0 — это zero value, будет ошибка синтаксическая, то есть =>
//		// Go сообщает о провале — не исключением, а возвращённым значением ошибки. Никаких throw, никаких try/catch.
//		n, err := strconv.Atoi("42")
//		fmt.Println(n, err)
//	}
//
// --------------------------------
// главный паттерн который будем писать 1000 раз
// «вызвал → проверил err → работаешь дальше».
//
//	func main() {
//		n, err := strconv.Atoi("abc")
//		if err != nil {
//			fmt.Println("ошибка", err)
//			return
//		}
//		fmt.Println("число", n)
//	}
//
// ------------------------------------
// тоже самое но с инициализацией в if
//
//	func main() {
//		if n, err := strconv.Atoi("abc"); err != nil {
//			fmt.Println("ошибка", err)
//		} else {
//			fmt.Println("число", n)
//		}
//	}
//
// ----------------------------------
// _: когда значение не нужно.
//
//	func main() {
//		_, err := strconv.Atoi("99")
//		if err != nil {
//			fmt.Println("не число")
//		} else {
//			fmt.Println("это число")
//		}
//	}
//
// _: когда значение не нужно с инициализацией.
// func main() {
// 	if _, err := strconv.Atoi("12"); err != nil {
// 		fmt.Println("Ошибка", err)
// 	} else {
// 		fmt.Println("Число")
// 	}
// }
// ______________________________________
// func minmax(a, b int) (int, int) {
// 	if a < b {
// 		return a, b
// 	} else {
// 		return b, a
// 	}
// }
// тоже самое без else
func minmax(a, b int) (int, int) {
	if a < b {
		return a, b
	}
	return b, a
}


func main() {
	a, b := minmax(11, 9)
	fmt.Println(a, b)
}