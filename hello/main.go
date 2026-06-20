package main

import "fmt"

// Структуры
// Поле Hobbies — это срез внутри структуры.
// К нему работает append
// type User struct {
// 	Name    string
// 	Hobbies []string
// }

//	func main() {
//		u := User{Name: "Erdni", Hobbies: []string{"boxing", "go"}}
//		u.Hobbies = append(u.Hobbies, "dogs")
//		// для отладки "%+v\n"
//		fmt.Printf("%+v\n", u)
//	}
//
// -------------------------
// Структура внутри структуры
// type Address struct {
// 	City    string
// 	Country string
// }
// // Address — отдельный тип, и он поле внутри User
// type User struct {
// 	Name    string
// 	Address Address
// }

//	func main() {
//		u := User{
//			Name:    "Erdni",
//			Address: Address{City: "Riga", Country: "Latvia"},
//		}
//		fmt.Println(u.Address.City)
//	}
//
// ------------------------------------
// Структура с картой внутри
// но здесь проблема
// Я создал Account, но не задал поле Balances.
// Значит Balances встал в своё zero value.
//  А zero value карты — это nil
// в nil-срез append делать можно
// а в nil-карту писать — паника.
// Карта требует, чтобы её сначала создали, срез — нет.
// type Account struct {
// 	Owner    string
// 	Balances map[string]int
// }

//	func main() {
//		a := Account{Owner: "Erdni"}
//		a.Balances["USD"] = 100
//		fmt.Println(a.Balances)
//	}
//
// -------------------
// Создаем нормально с инициализированной картой
// type Account struct {
// 	Owner    string
// 	Balances map[string]int
// }

// func NewAccount(owner string) Account {
// 	return Account{
// 		Owner:    owner,
// 		Balances: map[string]int{}, //создаем пустую карту, не nil
// 	}
// }

//	func main() {
//		a := NewAccount("Erdni")
//		a.Balances["USD"] = 100 //теперь не паникует
//		fmt.Println(a.Balances)
//	}
//
// ----------------------------
// конструктор, возвращающий указатель
// type User struct {
// 	Name string
// 	Age  int
// }

// func NewUser(name string) *User {
// 	// &User{...} — создаёт структуру и сразу даёт её адрес. Тип возврата — *User
// 	return &User{Name: name}
// }

//	func main() {
//		u := NewUser("Erdni")
//		fmt.Println(u.Name)
//	}
//
// ----------------------
// конструктор с проверкой.
// type User struct {
// 	Name string
// 	Age  int
// }
// // Возврат (*User, error) — указатель И ошибка, тот же мультивозврат, что у strconv.Atoi
// func NewUser(name string, age int) (*User, error) {
// 	if age < 0 {
// 		// Плохой ввод → nil (zero value указателя) + заполненная ошибка. Хороший → адрес + nil
// 		return nil, fmt.Errorf("Возврат не может быть отрицательным: %d", age)
// 	}
// 	return &User{Name: name, Age: age}, nil
// }

//	func main() {
//		u, err := NewUser("Erdni", -5)
//		if err != nil {
//			fmt.Println("ошибка", err)
//			return
//		}
//		fmt.Println(u.Name, u.Age)
//	}
//
// ----------------------
// Задача
type Wallet struct {
	Owner string
	// Карта — это «коробка пар ключ→значение
	Coins map[string]int
}

// Инициализируем карту, иначе она nil, а запись в nil-карту = паника
func NewWallet(owner string) *Wallet {
	return &Wallet{
		Owner: owner,
		Coins: map[string]int{},
	}
}

// метод привязанный к Wallet
func (w *Wallet) Add(coin string, amount int) {
	w.Coins[coin] += amount
}

func main() {
	w := NewWallet("Erdni")
	w.Add("BTC", 1)
	w.Add("BTC", 2)
	fmt.Printf("%+v\n", w)
}
