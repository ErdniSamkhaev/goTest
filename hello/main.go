package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func main() {
	result := add(5, 10)
	fmt.Print(result)
}

// чтобы посмотреть где ошибки в пробелач gofmt -d main.go
// исправление go fmt ./...
// проверка багов go vet./...