package main

import "fmt"

//	func main() {
//		// []int{...} — срез целых чисел (аналог массива в JS, но с нюансами)
//		nums := []int{10, 20, 30}
//		// fmt.Println(nums)
//		// fmt.Println(len(nums))
//		// обращение по индексу
//		fmt.Println(nums[0])
//		// обращение за пределы длины
//		// fmt.Println(nums[5])
//		// добавление append и ловушка
//		nums = append(nums, 40)
//		// append не меняет срез на месте, он возвращает новый — поэтому результат надо присвоить обратно: nums = append(...).
//		// append(nums, 40) — компилятор даже ругнётся, что результат не используется
//		fmt.Println(nums)
//	}
//
// -------------------------------------------
//
//	func main() {
//		nums := []int{10, 20, 30}
//		// Перебор через range
//		// for i, v := range nums {
//		// 	fmt.Println(i, v)
//		// }
//		// Если индекс не нужен — for _, v := range nums
//		for _, v := range nums {
//			fmt.Println(v)
//		}
//	}
//
// ------------------------------------------
// nil-срез и zero value.
//
//	func main() {
//		// Объявил срез без значения → его zero value это nil, длина 0.
//		// к nil-срезу можно сразу делать append, он сам становится нормальным.
//		// Снова zero value спасает: не надо отдельно «создавать пустой массив», как в JS. Выведет true, 0, [1].
//		var s []int
//		fmt.Println(s == nil)
//		fmt.Println(len(s))
//		s = append(s, 1)
//		fmt.Println(s)
//	}
//
// -----------------------------
// (maps)
//
//	func main() {
//		// map[string]int — карта, где ключ строка, значение число (аналог объекта/словаря в JS). Пишешь m[ключ] = значение, читаешь m[ключ]
//		ages := map[string]int{}
//		ages["Erdni"] = 30
//		ages["Lyuba"] = 27
//		fmt.Println(ages["Erdni"])
//	}
//
// ---------------------------------
// несуществующий ключ.
//
//	func main() {
//		// В JS obj["нетТакого"] дал бы undefined. В Go — 0. Снова zero value!
//		ages := map[string]int{"Erdni": 30}
//		fmt.Println(ages["НетТакого"])
//	}
//
// ------------------------------
// Паттерн «запятая, ok
//
//	func main() {
//		//  (n, err :=) — два возврата, второй говорит «всё ок или нет»
//		ages := map[string]int{"Erdni": 30}
//		v, ok := ages["Erdni"]
//		fmt.Println(v, ok)
//		v2, ok2 := ages["Нет такого"]
//		fmt.Println(v2, ok2)
//	}
//
// ---------------------------------
//
//	удаление и перебор.
//
//	func main() {
//		ages := map[string]int{"Erdni": 30, "Lyuba": 27, "Неизвестный": 222}
//		delete(ages, "Lyuba")
//		for k, v := range ages {
//			fmt.Println(k, v)
//		}
//	}
//
// ---------------------
// Возвращаем сумму среза
func sumSlice(num []int) int {
	sum := 0
	for _, v := range num {
		sum += v
	}
	return sum
}

func main() {
	nums := []int{10, 20, 30, 40, 5}
	result := sumSlice(nums)
	fmt.Println(result)
}
