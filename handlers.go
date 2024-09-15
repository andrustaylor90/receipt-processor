package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

// ProcessReceipt handles the /receipts/process endpoint.
func ProcessReceipt(w http.ResponseWriter, r *http.Request) {
	var receipt Receipt
	if err := json.NewDecoder(r.Body).Decode(&receipt); err != nil {
		http.Error(w, "Invalid receipt format", http.StatusBadRequest)
		return
	}

	// Generate a unique ID for the receipt.
	id := uuid.New().String()
	receipts[id] = receipt
	points := CalculatePoints(receipt)
	receiptPoints[id] = points

	// Return the receipt ID.
	response := map[string]string{"id": id}
	w.Header().Set("Content-Type", "application/json")
	fmt.Println(response)
	json.NewEncoder(w).Encode(response)
}

// GetPoints handles the /receipts/{id}/points endpoint.
func GetPoints(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	points, exists := receiptPoints[id]
	if !exists {
		http.Error(w, "Receipt not found", http.StatusNotFound)
		return
	}

	// Return the points awarded to the receipt.
	response := map[string]int{"points": points}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
