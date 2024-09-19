package main

import (
	"html/template"
	"log"
	"net/http"
)

// Обработчик для главной страницы
func home_page(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/home_page.html")
	if err != nil {
		log.Printf("Ошибка при парсинге шаблона: %v", err)
		http.Error(w, "Ошибка сервера", http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Printf("Ошибка при выполнении шаблона: %v", err)
		http.Error(w, "Ошибка сервера", http.StatusInternalServerError)
		return
	}
}

// Обработчик для страницы sasal.html
func sasal_page(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/sasal.html")
	if err != nil {
		log.Printf("Ошибка при парсинге шаблона: %v", err)
		http.Error(w, "Ошибка сервера", http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Printf("Ошибка при выполнении шаблона: %v", err)
		http.Error(w, "Ошибка сервера", http.StatusInternalServerError)
		return
	}
}

// Обработчик для обработки запросов
func handleRequest() {
	http.HandleFunc("/", home_page)
	http.HandleFunc("/sasal", sasal_page) // Добавлено маршрутизирование для sasal_page
	log.Println("Сервер запущен на порту 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Ошибка при запуске сервера: %v", err)
	}
}

func main() {
	log.Println("Запуск сервера...")
	handleRequest()
}
