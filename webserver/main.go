package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"net/http"
	"sync"
)

var (
	targets []Result
	mu      sync.Mutex
	db      *pgxpool.Pool
)

type Status struct {
	Service string `json:"service"`
	Version string `json:"version"`
	Healthy bool   `json:"healthy"`
}

type Result struct {
	Address string `json:"address"`
	Alive   bool   `json:"alive"`
}

func results(w http.ResponseWriter, r *http.Request) {
	data := []Result{
		{Address: "google.com:443", Alive: true},
		{Address: "github.com:443", Alive: true},
		{Address: "192.0.2.1:80", Alive: false},
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func status(w http.ResponseWriter, r *http.Request) {
	s := Status{
		Service: "checker-api",
		Version: "1.0",
		Healthy: true,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(s)
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Ты запросил путь: %s\n", r.URL.Path)
}

func addTarget(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "только POST", http.StatusMethodNotAllowed)
		return
	}

	var target Result
	if err := json.NewDecoder(r.Body).Decode(&target); err != nil {
		http.Error(w, "Плохой JSON", http.StatusBadRequest)
		return
	}

	_, err := db.Exec(context.Background(),
		"INSERT INTO targets (address, alive) VALUES ($1, $2)",
		target.Address, target.Alive)
	if err != nil {
		http.Error(w, "ошибка записи в базу", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(target)
}

func listTargets(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query(context.Background(), "SELECT address, alive FROM targets")
	if err != nil {
		http.Error(w, "ошибка чтения из базы", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	results := []Result{}
	for rows.Next() {
		var t Result
		rows.Scan(&t.Address, &t.Alive)
		results = append(results, t)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}

func deleteTarget(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "только DELETE", http.StatusMethodNotAllowed)
		return
	}

	address := r.URL.Query().Get("address")
	if address == "" {
		http.Error(w, "укажи address", http.StatusBadRequest)
		return
	}

	result, err := db.Exec(context.Background(), "DELETE FROM targets WHERE address=$1", address)
	if err != nil {
		http.Error(w, "ошибка удаления", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if result.RowsAffected() > 0 {
		fmt.Fprintf(w, "Удалён: %s\n", address)
	} else {
		fmt.Fprintf(w, "Не найден: %s\n", address)
	}
}

func main() {
	connStr := "postgres://postgres:secret@127.0.0.1:5432/checker?sslmode=disable"

	pool, err := pgxpool.New(context.Background(), connStr)
	if err != nil {
		fmt.Println("не удалось подключиться к базе:", err)
		return
	}
	defer pool.Close()
	db = pool // сохраняем в глобальную переменную

	fmt.Println("База подключена")

	http.HandleFunc("/", home)
	http.HandleFunc("/status", status)
	http.HandleFunc("/results", results)

	http.HandleFunc("/targets", listTargets)
	http.HandleFunc("/targets/add", addTarget)

	http.HandleFunc("/targets/delete", deleteTarget)

	fmt.Println("Сервер запущен на: http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
