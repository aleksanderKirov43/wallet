package main

import (
	"log"
	"net/http"

	"wallet/internal/router"
)

func main() {

	r := router.NewRouter()

	log.Println("Сервер запущен на порту: 8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("Сервер не удалось запустить: %v", err)
	}

}
