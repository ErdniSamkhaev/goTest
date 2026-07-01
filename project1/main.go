package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"net/http"
)

func main() {
	connStr := "postgres://postgres:secret@127.0.0.1:5433/links-db"

	// 1. Создаём пул соединений
	pool, err := pgxpool.New(context.Background(), connStr)
	if err != nil {
		fmt.Println("не удалось подключиться к бд:", err)
		return
	}

	// 2. Регистрируем отложенный вызов (но НЕ выполняем его сейчас!)
	defer pool.Close() // ← Go запоминает: "перед выходом из main вызвать pool.Close()"
	// 3. Присваиваем pool переменную db
	db = pool

	if err := initDB(); err != nil {
		fmt.Println("не удалось создать таблицу:", err)
		return
	}

	fmt.Println("База подключена")

	http.HandleFunc("/shorten", shorten)
	http.HandleFunc("/links", getLinks)
	fmt.Println("Сервер запущен на: http://localhost:8080")
	// чтобы сервак не падал, только принудительно
	http.ListenAndServe(":8080", nil)
}
