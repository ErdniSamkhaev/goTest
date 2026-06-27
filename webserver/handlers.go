package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Ты запросил путь: %s\n", r.URL.Path)
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
