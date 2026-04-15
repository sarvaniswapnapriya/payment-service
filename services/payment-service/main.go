package main

import (
	"fmt"
	"net/http"
	"os"
)

// --- SECTION: Infrastructure & Constants ---
// Changes here (like DB_URL) should trigger Test C (E2E)
const (
	ServiceVersion = "1.2.0"
	DatabaseURL    = "postgres://payment-db:5432"
)

// --- SECTION: Core Business Logic ---
// Changes here should trigger Test A (Logic Check)
func ValidateTransaction(amount float64) bool {
	if amount > 10000 {
		fmt.Println("Flagging for Fraud Review")
		return false
	}
	return amount > 0
}

// --- SECTION: API Handlers ---
// Changes here should trigger Test D (API Validator)
func statusHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{"status": "active", "version": "%s"}`, ServiceVersion)
}

func main() {
	http.HandleFunc("/status", statusHandler)
	port := os.Getenv("PORT")
	if port == "" { port = "8080" }
	http.ListenAndServe(":"+port, nil)
}
