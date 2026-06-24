package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	address := "google.com:80"

	conn, err := net.DialTimeout("tcp", address, 2*time.Second)
	if err != nil {
		fmt.Println("Хост не доступен, ошибка:", err)
		return
	}
	// сразу закрываем
	defer conn.Close()
	fmt.Println(address, "доступен")
}
