package main

import (
	"fmt"
	"net/http"
)

// Adding random comment to see if the tests still get triggred


// --- MOCK DATABASE INFRA ---
// Change this value to trigger Test C (E2E)
const DB_VERSION = "PostgreSQL 15.4"

// --- BUSINESS LOGIC: TAX CALCULATOR ---
// Change the logic here to trigger Test A (Logic) or Test Unit
func CalculateTax(amount float64) float64 {
	if amount < 0 {
		return 0
	}
	return amount * 0.25 // 
}
// --- API HANDLER ---
// Change the JSON response to trigger Test D (API Validator)
func paymentStatusHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{"status": "active", "db_status": "connected", "version": "1.2.0"}`)
}

func main() {
	http.HandleFunc("/payment/status", paymentStatusHandler)
	fmt.Println("Payment Service running on :8080...")
}
