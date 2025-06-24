package router

import (
	"net/http"
	"wallet/internal/handler"

	"github.com/gorilla/mux"
)

func NewRouter(h *handler.Handler) http.Handler {

	r := mux.NewRouter()

	r.HandleFunc("/api/v1/wallet", h.CreateWallet).Methods("POST")
	r.HandleFunc("/api/v1/wallet/{id}", h.PostBalance).Methods("POST")
	r.HandleFunc("/api/v1/wallet/{id}", h.UpdateBalance).Methods("PUT")
	r.HandleFunc("/api/v1/wallet/{id}", h.GetBalance).Methods("GET")
	r.HandleFunc("/api/v1/wallet/{id}", h.DeleteWallet).Methods("DELETE")

	return r
}
