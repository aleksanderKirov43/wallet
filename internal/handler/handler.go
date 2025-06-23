package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type WalletRequest struct {
	WalletID      int    `json:"walletId"`
	OperationType string `json:"operationType"`
	Amount        int64  `json:"amount,omitempty"`
}

func CreateWallet(w http.ResponseWriter, r *http.Request) {

	var req WalletRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Некорректный запрос", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Операция обработана: " + req.OperationType))

}

func PostBalance(w http.ResponseWriter, r *http.Request) {

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

func GetBalance(w http.ResponseWriter, r *http.Request) {

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

func UpdateBalance(w http.ResponseWriter, r *http.Request) {
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

func DeleteWallet(w http.ResponseWriter, r *http.Request) {
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
