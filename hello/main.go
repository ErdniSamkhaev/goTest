package main

import (
	"flag"
	"fmt"
	"net"
	// "os"
	"time"
)

// тип для ответа
type Result struct {
	Address  string
	Alive    bool
	Duration time.Duration
}

func check(address string, timeout time.Duration, ch chan Result) {
	start := time.Now()
	// для проверки соединения
	conn, err := net.DialTimeout("tcp", address, timeout)
	elapsed := time.Since(start)

	if err != nil {
		ch <- Result{Address: address, Alive: false, Duration: elapsed}
		return
	}
	defer conn.Close()
	ch <- Result{Address: address, Alive: true, Duration: elapsed}
}

func main() {
	timeout := flag.Int("timeout", 2, "таймаут в секундах")
	flag.Parse()
	start := time.Now()
	addresses := flag.Args() // хосты теперь берем отсюда, не из os.Args
	if len(addresses) == 0 {
		fmt.Println("Использование: checker <хост:порт> [хост:порт ...]")
		fmt.Println("Пример: checker google.com:443 github.com:80")
		return
	}

	ch := make(chan Result)
	// Запускаем все проверки паралельно
	for _, addr := range addresses {
		go check(addr, time.Duration(*timeout)*time.Second, ch)
	}

	alive := 0
	for i := 0; i < len(addresses); i++ {
		result := <-ch
		if result.Alive {
			alive++
		}
		fmt.Printf("%s | alive=%v | %v\n", result.Address, result.Alive, result.Duration)
	}
	fmt.Printf("Доступно %d из %d\n", alive, len(addresses))
	fmt.Println("Всего заняло:", time.Since(start))
}
