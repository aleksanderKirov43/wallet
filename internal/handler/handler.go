package handler

import (
	"encoding/json"
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

	wallet, err := h.service.CreateWallet(r.Context(), &req)
	if err != nil {
		http.Error(w, "Ошибка созадния кошелька: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(wallet)
}

func (h *Handler) PostBalance(w http.ResponseWriter, r *http.Request) {

	var req model.Wallet

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Некорректный JSON", http.StatusBadRequest)
		return
	}

	vars := mux.Vars(r)
	walletIdStr := vars["id"]

	walletId, err := strconv.Atoi(walletIdStr)
	if err != nil {
		http.Error(w, "id должен быть числом", http.StatusBadRequest)
		return
	}

	req.WalletID = walletId

	result, err := h.service.PostBalance(r.Context(), &req)
	if err != nil {
		http.Error(w, "Ошибка операции: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}

func (h *Handler) GetBalance(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	walletIdStr := vars["id"]

	walletId, err := strconv.Atoi(walletIdStr)
	if err != nil {
		http.Error(w, "id должен быть числом", http.StatusBadRequest)
		return
	}

	balance, err := h.service.GetBalance(r.Context(), walletId)
	if err != nil {
		http.Error(w, "Ошибка операции: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(balance)

}

func (h *Handler) UpdateBalance(w http.ResponseWriter, r *http.Request) {
	var req model.Wallet

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Некорректный JSON", http.StatusBadRequest)
		return
	}

	vars := mux.Vars(r)
	walletIdStr := vars["id"]

	walletId, err := strconv.Atoi(walletIdStr)
	if err != nil {
		http.Error(w, "id должен быть числом", http.StatusBadRequest)
		return
	}

	req.WalletID = walletId

	balanceUp, err := h.service.UpdateBalance(r.Context(), &req)
	if err != nil {
		http.Error(w, "Ошибка операции: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(balanceUp)
}

func (h *Handler) DeleteWallet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	walletIdStr := vars["id"]

	walletId, err := strconv.Atoi(walletIdStr)
	if err != nil {
		http.Error(w, "id должен быть числом", http.StatusBadRequest)
		return
	}

	if err = h.service.DeleteWallet(r.Context(), walletId); err != nil {
		http.Error(w, "Ошибка Удаления: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Кошелек успешно удален"))
}
