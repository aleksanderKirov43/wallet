package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"wallet/internal/model"
	"wallet/internal/service"

	"github.com/gorilla/mux"
)

type Handler struct {
	service service.WalletService
}

func NewHandler(s service.WalletService) *Handler {
	return &Handler{
		service: s,
	}
}

func (h *Handler) CreateWallet(w http.ResponseWriter, r *http.Request) {

	var req model.Wallet

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Некорректный запрос", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Операция обработана: " + req.OperationType))

}

func (h *Handler) PostBalance(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	walletIdStr := vars["id"]

	walletId, err := strconv.Atoi(walletIdStr)
	if err != nil {
		http.Error(w, "id должен быть числом", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("id кошелька: %d", walletId)))
}

func (h *Handler) GetBalance(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	walletIdStr := vars["id"]

	walletId, err := strconv.Atoi(walletIdStr)
	if err != nil {
		http.Error(w, "id должен быть числом", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("id кошелька: %d", walletId)))

}

func (h *Handler) UpdateBalance(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	walletIdStr := vars["id"]

	walletId, err := strconv.Atoi(walletIdStr)
	if err != nil {
		http.Error(w, "id должен быть числом", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("id кошелька: %d", walletId)))
}

func (h *Handler) DeleteWallet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	walletIdStr := vars["id"]

	walletId, err := strconv.Atoi(walletIdStr)
	if err != nil {
		http.Error(w, "id должен быть числом", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("id кошелька: %d", walletId)))
}
