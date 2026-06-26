package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Ты запросил путь: %s\n", r.URL.Path)
}

func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "О сервере: чекер проект, версия 1.0")
}

func ping(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Метод запроса: %s\n", r.Method)
}

func greet(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "незнакомец"
	}
	fmt.Fprintf(w, "Привет, %s\n", name)
}

func sum(w http.ResponseWriter, r *http.Request) {
	aStr := r.URL.Query().Get("a")
	bStr := r.URL.Query().Get("b")

	a, err := strconv.Atoi(aStr)
	if err != nil {
		fmt.Fprintf(w, "Ошибка: '%s' не число\n", aStr)
		return
	}

	b, err := strconv.Atoi(bStr)
	if err != nil {
		fmt.Fprintf(w, "Ошибка: '%s' не число\n", bStr)
		return
	}

	fmt.Fprintf(w, "Сумма: %d\n", a+b)
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/about", about)
	http.HandleFunc("/ping", ping)
	http.HandleFunc("/greet", greet)
	http.HandleFunc("/sum", sum)

	fmt.Println("Сервер на http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
