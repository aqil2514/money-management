package handler

import (
	"encoding/json"
	"fmt"
	"money-backend/internal/model"
	"net/http"
)

func CreateTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000") // Izinkan Nuxt
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Hanya menerima POST method", http.StatusMethodNotAllowed)
		return
	}

	var input model.Transaction

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, "Format json tidak diketahui : "+err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Printf("Data masuk ! Kategory %s, Nominal %.2f\n", input.Category, input.Nominal)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	response := map[string]string{"message": "Transaksi berhasil diterima di server GO"}
	json.NewEncoder(w).Encode(response)
}
