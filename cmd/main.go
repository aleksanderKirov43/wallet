package main

import (
	"log"
	"wallet/internal/app"
)

func main() {

	log.Println("Сервер запущен на порту: 8080")
	if err := app.Run(); err != nil {
		log.Fatalf("Сервер не удалось запустить: %v", err)
	}

}
