package app

import (
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"wallet/config"
	"wallet/internal/handler"
	"wallet/internal/repositiory"
	"wallet/internal/router"
	"wallet/internal/service"
)

func Run() error {

	err := godotenv.Load("config/.env")
	if err != nil {
		log.Println(".env  файл не найден")
	}

	db, err := config.InitDB()
	if err != nil {
		return err
	}

	repo := repositiory.NewRepositoryDB(db)
	serv := service.NewWalletService(repo)
	h := handler.NewHandler(serv)
	r := router.NewRouter(h)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Сервер запущен на порту: %s\n", port)
	return http.ListenAndServe(":"+port, r)
}
