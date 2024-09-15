package main

import "testing"

// Test for CalculatePoints function
func TestCalculatePoints(t *testing.T) {
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

	points := CalculatePoints(receipt)

	expectedPoints := 109
	if points != expectedPoints {
		t.Errorf("Expected %v points but got %v", expectedPoints, points)
	}
}
