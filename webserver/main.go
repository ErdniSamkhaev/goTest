package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"net/http"
)

func main() {
	connStr := "postgres://postgres:secret@127.0.0.1:5432/checker?sslmode=disable"

	pool, err := pgxpool.New(context.Background(), connStr)

	if err != nil {
		fmt.Println("не удалось подключиться к базе:", err)
		return
	}

	defer pool.Close()

	db = pool // сохраняем в глобальную переменную

	if err := initDB(); err != nil {
		fmt.Println("не удалось создать таблицу:", err)
		return
	}
	fmt.Println("База готова, таблица на месте")

	fmt.Println("База подключена")

	http.HandleFunc("/", home)
	http.HandleFunc("/status", status)

	http.HandleFunc("/targets", listTargets)
	http.HandleFunc("/targets/add", addTarget)

	http.HandleFunc("/targets/delete", deleteTarget)

	fmt.Println("Сервер запущен на: http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
