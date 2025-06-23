package router

import (
	"net/http"

	"wallet/internal/handler"

	"github.com/gorilla/mux"
)

func NewRouter() http.Handler {

	r := mux.NewRouter()

	r.HandleFunc("/api/v1/wallet", handler.CreateWallet).Methods("POST")
	r.HandleFunc("/api/v1/wallet/{id}", handler.PostBalance).Methods("GET")
	r.HandleFunc("/api/v1/wallet/{id}", handler.UpdateBalance).Methods("PUT")
	r.HandleFunc("/api/v1/wallet/{id}", handler.GetBalance).Methods("GET")
	r.HandleFunc("/api/v1/wallet/{id}", handler.DeleteWallet).Methods("DELETE")

	return r
}
