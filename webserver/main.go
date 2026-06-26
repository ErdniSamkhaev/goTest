package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	// "strconv"
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
	err := json.NewDecoder(r.Body).Decode(&target)
	if err != nil {
		http.Error(w, "Плохой JSON", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Println(w, "Принял хост: %s (alive=%v)\n", target.Address, target.Alive)
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/status", status)
	http.HandleFunc("/results", results)
	http.HandleFunc("/addTarget", addTarget)

	fmt.Println("Сервер запущен на: http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
