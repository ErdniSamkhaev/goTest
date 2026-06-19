package main

import "fmt"

// func main() {
// 	age := 20
// 	// скобок в if как в JS нет
// 	if age >= 18
// 		fmt.Println("совершенолетний")
// 	else {
// 		fmt.Println("нет")
// 	}
// }

// --------------------------------
//
//	func main() {
//		age := 20
//		// Переменная doubled создается внутри условия if
//		if doubled := age * 2; doubled > 30 {
//			fmt.Println(doubled, "больше 30")
//		}
//	}
//
// --------------------------------------
// тернарки как в JS неь
// status := age > 18 ? "взрослый" : "ребенок"
// в Go так задумано: меньше способов написать одно и то же = меньше холиваров и проще читать чужой код
// func main() {
// 	age := 20
// 	var status string
// 	if age > 18 {
// 		status = "взрослый"
// 	} else {
// 		status = "ребенок"
// 	}
// 	fmt.Println(status)
// }

// while нет , только for
//
//	func main() {
//		for i := 0; i < 3; i++ {
//			fmt.Println(i)
//		}
//	}
//
// ------------------------------------------
// в Go нет while и do-while. Нужен цикл «пока условие истинно» — это тот же for, просто с одним условием:
//
//	func main() {
//		n := 0
//		for n < 3 {
//			fmt.Println(n)
//			n++
//		}
//	}
//
// ------------------------------------------------
// бесконечный for + break.
//
//	func main() {
//		count := 0
//		for {
//			if count == 3 {
//				break
//			}
//			fmt.Println(count)
//			count++
//		}
//	}
//
// ------------------------------------------------
// тут нет break после каждого case
// В Go switch по умолчанию не проваливается в следующий case — отработал свой и вышел сам.
//
//	func main() {
//		day := 3
//		switch day {
//		case 1:
//			fmt.Println("ПОнедельгтк")
//		case 3:
//			fmt.Println("Среда")
//		default:
//			fmt.Println("Другой")
//		}
//	}
//
// Можно несколько значений в одном case
//
//	func main() {
//		day := 3
//		switch day {
//		case 6, 7:
//			fmt.Println("выходной")
//		default:
//			fmt.Println("будний")
//		}
//	}
//
// --------------------------------
func main() {
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println("Четное:", i)
		} else {
			fmt.Println("Нечетное:", i)
		}
	}
}
