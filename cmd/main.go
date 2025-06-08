package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"wallet-service/internal/service"
)

var walletService service.WalletService

func main() {

	walletService = service.NewWalletService()

	http.HandleFunc("/wallet", handleCreateWallet)
	http.HandleFunc("/wallet/get", handleGetWallet)
	http.HandleFunc("/wallet/deposit", handleDeposit)
	http.HandleFunc("/wallet/withdraw", handleWithdraw)

	fmt.Println("Server is running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleCreateWallet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	ownerIDStr := r.URL.Query().Get("owner_id")
	ownerID, err := strconv.ParseInt(ownerIDStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid owner_id", http.StatusBadRequest)
		return
	}

	wallet := walletService.CreateWallet(ownerID)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(wallet)
}

func handleGetWallet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.URL.Query().Get("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return
	}

	wallet, err := walletService.GetWallet(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(wallet)
}

func handleDeposit(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.URL.Query().Get("id")
	amountStr := r.URL.Query().Get("amount")

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return
	}

	amount, err := strconv.ParseInt(amountStr, 10, 64)
	if err != nil || amount <= 0 {
		http.Error(w, "Invalid amount", http.StatusBadRequest)
		return
	}

	wallet, err := walletService.Deposit(id, amount)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(wallet)
}

func handleWithdraw(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.URL.Query().Get("id")
	amountStr := r.URL.Query().Get("amount")

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return
	}

	amount, err := strconv.ParseInt(amountStr, 10, 64)
	if err != nil || amount <= 0 {
		http.Error(w, "Invalid amount", http.StatusBadRequest)
		return
	}

	wallet, err := walletService.Withdraw(id, amount)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(wallet)

}
