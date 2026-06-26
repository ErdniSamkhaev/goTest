package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

var (
	targets []Result
	mu      sync.Mutex
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
	}

	mu.Lock()
	targets = append(targets, target)
	mu.Unlock()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(target)
}

func listTargets(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(targets)
}

func deleteTarget(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "только DELETE", http.StatusMethodNotAllowed)
		return
	}

	address := r.URL.Query().Get("address")
	if address == "" {
		http.Error(w, "укажи address", http.StatusBadRequest)
	}

	mu.Lock()
	defer mu.Unlock()

	newTargets := []Result{}
	found := false
	for _, t := range targets {
		if t.Address == address {
			found = true
			continue
		}
		newTargets = append(newTargets, t)
	}
	targets = newTargets

	w.Header().Set("Content-Type", "application/json")
	if found {
		fmt.Fprintf(w, "Удален: %s\n", address)
	} else {
		fmt.Fprintf(w, "Не найден: %s\n", address)
	}
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/status", status)
	http.HandleFunc("/results", results)
	http.HandleFunc("/addTarget", addTarget)

	http.HandleFunc("/targets", listTargets)
	http.HandleFunc("/targets/add", addTarget)

	http.HandleFunc("/targets/delete", deleteTarget)

	fmt.Println("Сервер запущен на: http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
