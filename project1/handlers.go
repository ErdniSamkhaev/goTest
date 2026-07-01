package main

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func checkMethod(w http.ResponseWriter, r *http.Request, allowedMethod string) bool {
	if r.Method != allowedMethod {
		http.Error(w, fmt.Sprintf("Метод не поддерживается. Используй %s", allowedMethod), http.StatusMethodNotAllowed)
		return false
	}
	return true
}

func shorten(w http.ResponseWriter, r *http.Request) {
	// Проверяем метод
	if !checkMethod(w, r, http.MethodPost) {
		return
	}

	// парсим запрос
	var req ShortenRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Неверный формат JSON", http.StatusBadRequest)
		return
	}

	// Валидация
	if req.URL == "" {
		http.Error(w, "Поле url не может быть пустым", http.StatusBadRequest)
		return
	}

	// подставляем короткую ссылку
	shortCode, err := generateSymbols()
	if err != nil {
		log.Printf("ошибка генерации кода: %v", err)
		http.Error(w, "Ошибка сервера", http.StatusInternalServerError)
		return
	}

	// сохраняем в бд (переименовал на links Таблицу)
	query := `INSERT INTO links (short_code, original_url) VALUES ($1, $2)`
	_, err = db.Exec(context.Background(), query, shortCode, req.URL)
	if err != nil {
		log.Printf("Ошибка при сохранении в бд: %v", err)
		http.Error(w, "Ошибка сервера", http.StatusInternalServerError)
		return
	}

	// отправляем ответ
	resp := ShortenResponse{
		ShortCode: shortCode,
		ShortUrl:  "http://localhost:8080/" + shortCode,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}

func getLinks(w http.ResponseWriter, r *http.Request) {
	// Проверяем метод
	if !checkMethod(w, r, http.MethodGet) {
		return
	}

	// 2. Запрашиваем все ссылки из БД
	query := `SELECT id, short_code, original_url, created_at FROM links ORDER BY id DESC`
	rows, err := db.Query(context.Background(), query)
	if err != nil {
		log.Printf("Ошибка при запросе из бд: %v", err)
		http.Error(w, "Ошибка сервера", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	// 3. Собираем результаты в слайс
	var links []Link
	for rows.Next() {
		var link Link
		err := rows.Scan(&link.ID, &link.ShortCode, &link.OriginalURL, &link.CreatedAt)
		if err != nil {
			log.Printf("Ошибка при чтении строки: %v", err)
			http.Error(w, "Ошибка сервера", http.StatusInternalServerError)
			return
		}
		links = append(links, link)
	}

	// 4. Проверяем ошибки после итерации
	if err = rows.Err(); err != nil {
		log.Printf("Ошибка после итерации: %v", err)
		http.Error(w, "Ошибка сервера", http.StatusInternalServerError)
		return
	}

	// 5. Если слайс пустой, инициализируем пустой массив (чтобы в JSON было [], а не null)
	if links == nil {
		links = []Link{}
	}

	// 6. Отправляем JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(links)
}

func generateSymbols() (string, error) {
	// генерим 4 байта (достаточно для 6 символов в base64 без padding)
	b := make([]byte, 4)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	// используем raw URL-safe base64 (без '+' и '/'), обрезаем до 6 символов
	code := base64.RawURLEncoding.EncodeToString(b)
	return code[:6], nil

}
