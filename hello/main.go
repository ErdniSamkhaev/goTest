package main

import (
	"fmt"
	"net"
	"time"
)

// тип для ответа
type Result struct {
	Address  string
	Alive    bool
	Duration time.Duration
}

func check(address string, ch chan Result) {
	start := time.Now()
	// для проверки соединения
	conn, err := net.DialTimeout("tcp", address, 2*time.Second)
	elapsed := time.Since(start)

	if err != nil {
		ch <- Result{Address: address, Alive: false, Duration: elapsed}
		return
	}
	defer conn.Close()
	ch <- Result{Address: address, Alive: true, Duration: elapsed}
}

func main() {
	start := time.Now()
	addresses := []string{
		"google.com:443",
		"github.com:443",
		"192.0.2.1:80",    // мертвый
		"google.com:9999", //закрытый
	}

	ch := make(chan Result)
	// Запускаем все проверки паралельно
	for _, addr := range addresses {
		go check(addr, ch)
	}

	for i := 0; i < len(addresses); i++ {
		result := <-ch
		fmt.Printf("%s | alive=%v | %v\n", result.Address, result.Alive, result.Duration)
	}
	fmt.Println("Всего заняло:", time.Since(start))
}
