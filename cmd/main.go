package main

import (
	"log"
	"wallet/internal/app"
)

func main() {

	if err := app.Run(); err != nil {
		log.Fatalf("Сервер не удалось запустить: %v", err)
	}

}
