package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

// Helper function to create a request and execute the handler
func executeRequest(req *http.Request, handlerFunc http.HandlerFunc) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlerFunc)
	handler.ServeHTTP(rr, req)
	return rr
}

// Test for the ProcessReceipt handler
func TestProcessReceipt(t *testing.T) {
	receipt := Receipt{
		Retailer:     "Target",
		PurchaseDate: "2022-01-01",
		PurchaseTime: "13:01",
		Items: []Item{
			{ShortDescription: "Mountain Dew 12PK", Price: "6.49"},
			{ShortDescription: "Emils Cheese Pizza", Price: "12.25"},
		},
		Total: "18.74",
	}
	payload, _ := json.Marshal(receipt)

	req, err := http.NewRequest("POST", "/receipts/process", bytes.NewBuffer(payload))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	rr := executeRequest(req, ProcessReceipt)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var response map[string]string
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("Unable to parse response: %v", err)
	}

	if _, exists := response["id"]; !exists {
		t.Errorf("Expected response to contain 'id' but got: %v", response)
	}
}

// Test for the GetPoints handler
func TestGetPoints(t *testing.T) {
	// Add a receipt manually to simulate processing
	receipt := Receipt{
		Retailer:     "M&M Corner Market",
		PurchaseDate: "2022-03-20",
		PurchaseTime: "14:33",
		Items: []Item{
			{ShortDescription: "Gatorade", Price: "2.25"},
			{ShortDescription: "Gatorade", Price: "2.25"},
			{ShortDescription: "Gatorade", Price: "2.25"},
			{ShortDescription: "Gatorade", Price: "2.25"},
		},
		Total: "9.00",
	}

	// Generate an ID and calculate points for the manual entry
	id := "test-receipt-id"
	receipts[id] = receipt
	receiptPoints[id] = CalculatePoints(receipt)

	// Create a new router and register the GetPoints handler with path variables
	router := mux.NewRouter()
	router.HandleFunc("/receipts/{id}/points", GetPoints).Methods("GET")

	// Create a request to get points using the generated ID
	req, err := http.NewRequest("GET", "/receipts/"+id+"/points", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Execute the request using the router
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	// Check if the response status code is 200 OK
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Parse the response body to verify the points
	var response map[string]int
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("Unable to parse response: %v", err)
	}

	// Expected points based on the receipt data
	expectedPoints := 109
	if response["points"] != expectedPoints {
		t.Errorf("Expected points %v but got %v", expectedPoints, response["points"])
	}
}
