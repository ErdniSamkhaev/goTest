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

func check(address string) Result {
	start := time.Now()
	// для проверки соединения
	conn, err := net.DialTimeout("tcp", address, 2*time.Second)
	elapsed := time.Since(start)

	if err != nil {
		return Result{Address: address, Alive: false, Duration: elapsed}
	}
	defer conn.Close()
	return Result{Address: address, Alive: true, Duration: elapsed}
}

func main() {
	addresses := []string{
		"google.com:443",
		"github.com:443",
		"192.0.2.1:80", // мертвый
		"google.com:9999", //закрытый
	}

	for _, addr := range addresses {
		result := check(addr)
		fmt.Printf("%s | alive=%v | %v\n", result.Address, result.Alive, result.Duration)
	}
}
